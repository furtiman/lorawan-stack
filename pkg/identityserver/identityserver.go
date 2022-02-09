// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package identityserver

import (
	"context"
	"fmt"

	pbtypes "github.com/gogo/protobuf/types"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Postgres database driver.
	"go.thethings.network/lorawan-stack/v3/pkg/account"
	account_store "go.thethings.network/lorawan-stack/v3/pkg/account/store"
	"go.thethings.network/lorawan-stack/v3/pkg/auth/rights"
	"go.thethings.network/lorawan-stack/v3/pkg/cluster"
	"go.thethings.network/lorawan-stack/v3/pkg/component"
	"go.thethings.network/lorawan-stack/v3/pkg/email"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	gormstore "go.thethings.network/lorawan-stack/v3/pkg/identityserver/gormstore"
	"go.thethings.network/lorawan-stack/v3/pkg/identityserver/store"
	"go.thethings.network/lorawan-stack/v3/pkg/interop"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/oauth"
	"go.thethings.network/lorawan-stack/v3/pkg/redis"
	"go.thethings.network/lorawan-stack/v3/pkg/rpcmiddleware/hooks"
	"go.thethings.network/lorawan-stack/v3/pkg/rpcmiddleware/rpclog"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/webui"
	"google.golang.org/grpc"
)

// IdentityServer implements the Identity Server component.
//
// The Identity Server exposes the Registry and Access services for Applications,
// OAuth clients, Gateways, Organizations and Users.
type IdentityServer struct {
	*component.Component
	ctx            context.Context
	config         *Config
	db             *gorm.DB
	redis          *redis.Client
	emailTemplates *email.TemplateRegistry
	account        account.Server
	oauth          oauth.Server
}

// Context returns the context of the Identity Server.
func (is *IdentityServer) Context() context.Context {
	return is.ctx
}

// SetRedisCache configures the given redis instance for caching.
func (is *IdentityServer) SetRedisCache(redis *redis.Client) {
	is.redis = redis
}

type ctxKeyType struct{}

var ctxKey ctxKeyType

func (is *IdentityServer) configFromContext(ctx context.Context) *Config {
	if config, ok := ctx.Value(ctxKey).(*Config); ok {
		return config
	}
	return is.config
}

// GenerateCSPString returns a Content-Security-Policy header value
// for OAuth and Account app template.
func GenerateCSPString(config *oauth.Config, nonce string) string {
	cspMap := webui.CleanCSP(map[string][]string{
		"connect-src": {
			"'self'",
			config.UI.StackConfig.IS.BaseURL,
			config.UI.SentryDSN,
			"gravatar.com",
			"www.gravatar.com",
		},
		"style-src": {
			"'self'",
			config.UI.AssetsBaseURL,
			config.UI.BrandingBaseURL,
		},
		"script-src": {
			"'self'",
			config.UI.AssetsBaseURL,
			config.UI.BrandingBaseURL,
			"'unsafe-eval'",
			"'strict-dynamic'",
			fmt.Sprintf("'nonce-%s'", nonce),
		},
		"base-uri": {
			"'self'",
		},
		"frame-ancestors": {
			"'none'",
		},
	})
	return webui.GenerateCSPString(cspMap)
}

type accountAppStore struct {
	db *gorm.DB

	store.UserStore
	store.LoginTokenStore
	store.UserSessionStore
}

// WithSoftDeleted implements account_store.Interface.
func (as *accountAppStore) WithSoftDeleted(ctx context.Context, onlyDeleted bool) context.Context {
	return store.WithSoftDeleted(ctx, onlyDeleted)
}

// Transact implements account_store.Interface.
func (as *accountAppStore) Transact(ctx context.Context, f func(context.Context, account_store.Interface) error) error {
	return gormstore.Transact(ctx, as.db, func(db *gorm.DB) error {
		return f(ctx, createAccountAppStore(db))
	})
}

func createAccountAppStore(db *gorm.DB) account_store.Interface {
	return &accountAppStore{
		db: db,

		UserStore:        gormstore.GetUserStore(db),
		LoginTokenStore:  gormstore.GetLoginTokenStore(db),
		UserSessionStore: gormstore.GetUserSessionStore(db),
	}
}

var (
	errDBNeedsMigration = errors.Define("db_needs_migration", "the database needs to be migrated")
	errNoDevEUIPrefix   = errors.Define("dev_eui_block_prefix_missing", "issuing DevEUI's from a block enabled but no prefix configured")
)

// New returns new *IdentityServer.
func New(c *component.Component, config *Config) (is *IdentityServer, err error) {
	is = &IdentityServer{
		Component: c,
		ctx:       log.NewContextWithField(c.Context(), "namespace", "identityserver"),
		config:    config,
	}
	is.db, err = gormstore.Open(is.Context(), is.config.DatabaseURI)
	if err != nil {
		return nil, err
	}
	if c.LogDebug() {
		is.db = is.db.Debug()
	}
	if err = gormstore.Check(is.db); err != nil {
		return nil, errDBNeedsMigration.WithCause(err)
	}
	go func() {
		<-is.Context().Done()
		is.db.Close()
	}()

	is.emailTemplates, err = is.initEmailTemplates(is.Context())
	if err != nil {
		return nil, err
	}

	if is.config.DevEUIBlock.Enabled {
		if is.config.DevEUIBlock.Prefix.IsZero() {
			return nil, errNoDevEUIPrefix.New()
		}
		err := is.withDatabase(is.Context(), func(db *gorm.DB) error {
			return gormstore.GetEUIStore(db).InitializeDevEUIBlock(is.Context(), is.config.DevEUIBlock.Prefix, is.config.DevEUIBlock.InitCounter)
		})
		if err != nil {
			return nil, err
		}
	}

	is.config.OAuth.CSRFAuthKey = is.GetBaseConfig(is.Context()).HTTP.Cookie.HashKey
	is.config.OAuth.UI.FrontendConfig.EnableUserRegistration = is.config.UserRegistration.Enabled
	is.oauth, err = oauth.NewServer(c, struct {
		store.UserStore
		store.UserSessionStore
		store.ClientStore
		store.OAuthStore
	}{
		UserStore:        gormstore.GetUserStore(is.db),
		UserSessionStore: gormstore.GetUserSessionStore(is.db),
		ClientStore:      gormstore.GetClientStore(is.db),
		OAuthStore:       gormstore.GetOAuthStore(is.db),
	}, is.config.OAuth, GenerateCSPString)

	is.account, err = account.NewServer(c, createAccountAppStore(is.db), is.config.OAuth, GenerateCSPString)
	if err != nil {
		return nil, err
	}

	c.AddContextFiller(func(ctx context.Context) context.Context {
		ctx = is.withRequestAccessCache(ctx)
		ctx = rights.NewContextWithFetcher(ctx, is)
		ctx = rights.NewContextWithCache(ctx)
		return ctx
	})

	for _, hook := range []struct {
		name       string
		middleware hooks.UnaryHandlerMiddleware
	}{
		{rpclog.NamespaceHook, rpclog.UnaryNamespaceHook("identityserver")},
		{cluster.HookName, c.ClusterAuthUnaryHook()},
	} {
		hooks.RegisterUnaryHook("/ttn.lorawan.v3.Is", hook.name, hook.middleware)
		hooks.RegisterUnaryHook("/ttn.lorawan.v3.ApplicationRegistry", hook.name, hook.middleware)
		hooks.RegisterUnaryHook("/ttn.lorawan.v3.ApplicationAccess", hook.name, hook.middleware)
		hooks.RegisterUnaryHook("/ttn.lorawan.v3.ClientRegistry", hook.name, hook.middleware)
		hooks.RegisterUnaryHook("/ttn.lorawan.v3.ClientAccess", hook.name, hook.middleware)
		hooks.RegisterUnaryHook("/ttn.lorawan.v3.EndDeviceRegistry", hook.name, hook.middleware)
		hooks.RegisterUnaryHook("/ttn.lorawan.v3.GatewayRegistry", hook.name, hook.middleware)
		hooks.RegisterUnaryHook("/ttn.lorawan.v3.GatewayAccess", hook.name, hook.middleware)
		hooks.RegisterUnaryHook("/ttn.lorawan.v3.OrganizationRegistry", hook.name, hook.middleware)
		hooks.RegisterUnaryHook("/ttn.lorawan.v3.OrganizationAccess", hook.name, hook.middleware)
		hooks.RegisterUnaryHook("/ttn.lorawan.v3.UserRegistry", hook.name, hook.middleware)
		hooks.RegisterUnaryHook("/ttn.lorawan.v3.UserAccess", hook.name, hook.middleware)
		hooks.RegisterUnaryHook("/ttn.lorawan.v3.UserSessionRegistry", hook.name, hook.middleware)
	}
	hooks.RegisterUnaryHook("/ttn.lorawan.v3.EntityAccess", rpclog.NamespaceHook, rpclog.UnaryNamespaceHook("identityserver"))
	hooks.RegisterUnaryHook("/ttn.lorawan.v3.EntityAccess", cluster.HookName, c.ClusterAuthUnaryHook())
	hooks.RegisterUnaryHook("/ttn.lorawan.v3.OAuthAuthorizationRegistry", rpclog.NamespaceHook, rpclog.UnaryNamespaceHook("identityserver"))

	c.RegisterGRPC(is)
	c.RegisterWeb(is.oauth)
	c.RegisterWeb(is.account)
	c.RegisterInterop(is)

	return is, nil
}

func (is *IdentityServer) withDatabase(ctx context.Context, f func(*gorm.DB) error) error {
	return gormstore.Transact(ctx, is.db, f)
}

// RegisterServices registers services provided by is at s.
func (is *IdentityServer) RegisterServices(s *grpc.Server) {
	ttnpb.RegisterIsServer(s, is)
	ttnpb.RegisterEntityAccessServer(s, &entityAccess{IdentityServer: is})
	ttnpb.RegisterApplicationRegistryServer(s, &applicationRegistry{IdentityServer: is})
	ttnpb.RegisterApplicationAccessServer(s, &applicationAccess{IdentityServer: is})
	ttnpb.RegisterClientRegistryServer(s, &clientRegistry{IdentityServer: is})
	ttnpb.RegisterClientAccessServer(s, &clientAccess{IdentityServer: is})
	ttnpb.RegisterEndDeviceRegistryServer(s, &endDeviceRegistry{IdentityServer: is})
	ttnpb.RegisterGatewayRegistryServer(s, &gatewayRegistry{IdentityServer: is})
	ttnpb.RegisterGatewayAccessServer(s, &gatewayAccess{IdentityServer: is})
	ttnpb.RegisterOrganizationRegistryServer(s, &organizationRegistry{IdentityServer: is})
	ttnpb.RegisterOrganizationAccessServer(s, &organizationAccess{IdentityServer: is})
	ttnpb.RegisterUserRegistryServer(s, &userRegistry{IdentityServer: is})
	ttnpb.RegisterUserAccessServer(s, &userAccess{IdentityServer: is})
	ttnpb.RegisterUserSessionRegistryServer(s, &userSessionRegistry{IdentityServer: is})
	ttnpb.RegisterUserInvitationRegistryServer(s, &invitationRegistry{IdentityServer: is})
	ttnpb.RegisterEntityRegistrySearchServer(s, &registrySearch{IdentityServer: is})
	ttnpb.RegisterEndDeviceRegistrySearchServer(s, &registrySearch{IdentityServer: is})
	ttnpb.RegisterOAuthAuthorizationRegistryServer(s, &oauthRegistry{IdentityServer: is})
	ttnpb.RegisterContactInfoRegistryServer(s, &contactInfoRegistry{IdentityServer: is})
}

// RegisterHandlers registers gRPC handlers.
func (is *IdentityServer) RegisterHandlers(s *runtime.ServeMux, conn *grpc.ClientConn) {
	ttnpb.RegisterIsHandler(is.Context(), s, conn)
	ttnpb.RegisterEntityAccessHandler(is.Context(), s, conn)
	ttnpb.RegisterApplicationRegistryHandler(is.Context(), s, conn)
	ttnpb.RegisterApplicationAccessHandler(is.Context(), s, conn)
	ttnpb.RegisterClientRegistryHandler(is.Context(), s, conn)
	ttnpb.RegisterClientAccessHandler(is.Context(), s, conn)
	ttnpb.RegisterEndDeviceRegistryHandler(is.Context(), s, conn)
	ttnpb.RegisterGatewayRegistryHandler(is.Context(), s, conn)
	ttnpb.RegisterGatewayAccessHandler(is.Context(), s, conn)
	ttnpb.RegisterOrganizationRegistryHandler(is.Context(), s, conn)
	ttnpb.RegisterOrganizationAccessHandler(is.Context(), s, conn)
	ttnpb.RegisterUserRegistryHandler(is.Context(), s, conn)
	ttnpb.RegisterUserAccessHandler(is.Context(), s, conn)
	ttnpb.RegisterUserSessionRegistryHandler(is.Context(), s, conn)
	ttnpb.RegisterUserInvitationRegistryHandler(is.Context(), s, conn)
	ttnpb.RegisterEntityRegistrySearchHandler(is.Context(), s, conn)
	ttnpb.RegisterEndDeviceRegistrySearchHandler(is.Context(), s, conn)
	ttnpb.RegisterOAuthAuthorizationRegistryHandler(is.Context(), s, conn)
	ttnpb.RegisterContactInfoRegistryHandler(is.Context(), s, conn)
}

// RegisterInterop registers the LoRaWAN Backend Interfaces interoperability services.
func (is *IdentityServer) RegisterInterop(srv *interop.Server) {
	srv.RegisterIS(&interopServer{IdentityServer: is})
}

// Roles returns the roles that the Identity Server fulfills.
func (is *IdentityServer) Roles() []ttnpb.ClusterRole {
	return []ttnpb.ClusterRole{ttnpb.ClusterRole_ACCESS, ttnpb.ClusterRole_ENTITY_REGISTRY}
}

func (is *IdentityServer) getMembershipStore(ctx context.Context, db *gorm.DB) store.MembershipStore {
	s := gormstore.GetMembershipStore(db)
	if is.redis != nil {
		if membershipTTL := is.configFromContext(ctx).AuthCache.MembershipTTL; membershipTTL > 0 {
			s = gormstore.GetMembershipCache(s, is.redis, membershipTTL)
		}
	}
	return s
}

var softDeleteFieldMask = &pbtypes.FieldMask{Paths: []string{"deleted_at"}}

var errRestoreWindowExpired = errors.DefineFailedPrecondition("restore_window_expired", "this entity can no longer be restored")
