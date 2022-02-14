// Code generated by protoc-gen-go-flags. DO NOT EDIT.
// versions:
// - protoc-gen-go-flags v0.0.0-dev
// - protoc              v3.17.3
// source: lorawan-stack/api/identityserver.proto

package ttnpb

import (
	flagsplugin "github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin"
	pflag "github.com/spf13/pflag"
)

// AddSelectFlagsForAuthInfoResponse_APIKeyAccess adds flags to select fields in AuthInfoResponse_APIKeyAccess.
func AddSelectFlagsForAuthInfoResponse_APIKeyAccess(flags *pflag.FlagSet, prefix string) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("api-key", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("api-key", prefix), true)))
	AddSelectFlagsForAPIKey(flags, flagsplugin.Prefix("api-key", prefix))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("entity-ids", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("entity-ids", prefix), true)))
	AddSelectFlagsForEntityIdentifiers(flags, flagsplugin.Prefix("entity-ids", prefix))
}

// SelectFromFlags outputs the fieldmask paths forAuthInfoResponse_APIKeyAccess message from select flags.
func PathsFromSelectFlagsForAuthInfoResponse_APIKeyAccess(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("api_key", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("api_key", prefix))
	}
	if selectPaths, err := PathsFromSelectFlagsForAPIKey(flags, flagsplugin.Prefix("api_key", prefix)); err != nil {
		return nil, err
	} else {
		paths = append(paths, selectPaths...)
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("entity_ids", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("entity_ids", prefix))
	}
	if selectPaths, err := PathsFromSelectFlagsForEntityIdentifiers(flags, flagsplugin.Prefix("entity_ids", prefix)); err != nil {
		return nil, err
	} else {
		paths = append(paths, selectPaths...)
	}
	return paths, nil
}

// AddSetFlagsForAuthInfoResponse_APIKeyAccess adds flags to select fields in AuthInfoResponse_APIKeyAccess.
func AddSetFlagsForAuthInfoResponse_APIKeyAccess(flags *pflag.FlagSet, prefix string) {
	AddSetFlagsForAPIKey(flags, flagsplugin.Prefix("api-key", prefix))
	AddSetFlagsForEntityIdentifiers(flags, flagsplugin.Prefix("entity-ids", prefix))
}

// SetFromFlags sets the AuthInfoResponse_APIKeyAccess message from flags.
func (m *AuthInfoResponse_APIKeyAccess) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if selected := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("api_key", prefix)); selected {
		m.ApiKey = &APIKey{}
		if setPaths, err := m.ApiKey.SetFromFlags(flags, flagsplugin.Prefix("api_key", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	if selected := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("entity_ids", prefix)); selected {
		m.EntityIds = &EntityIdentifiers{}
		if setPaths, err := m.EntityIds.SetFromFlags(flags, flagsplugin.Prefix("entity_ids", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	return paths, nil
}

// AddSelectFlagsForAuthInfoResponse adds flags to select fields in AuthInfoResponse.
func AddSelectFlagsForAuthInfoResponse(flags *pflag.FlagSet, prefix string) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("access-method.api-key", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("access-method.api-key", prefix), true)))
	AddSelectFlagsForAuthInfoResponse_APIKeyAccess(flags, flagsplugin.Prefix("access-method.api-key", prefix))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("access-method.oauth-access-token", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("access-method.oauth-access-token", prefix), true)))
	AddSelectFlagsForOAuthAccessToken(flags, flagsplugin.Prefix("access-method.oauth-access-token", prefix))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("access-method.user-session", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("access-method.user-session", prefix), true)))
	AddSelectFlagsForUserSession(flags, flagsplugin.Prefix("access-method.user-session", prefix))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("universal-rights", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("universal-rights", prefix), true)))
	// NOTE: universal_rights (Rights) does not seem to have select flags.
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("is-admin", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("is-admin", prefix), false)))
}

// SelectFromFlags outputs the fieldmask paths forAuthInfoResponse message from select flags.
func PathsFromSelectFlagsForAuthInfoResponse(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("access_method.api_key", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("access_method.api_key", prefix))
	}
	if selectPaths, err := PathsFromSelectFlagsForAuthInfoResponse_APIKeyAccess(flags, flagsplugin.Prefix("access_method.api_key", prefix)); err != nil {
		return nil, err
	} else {
		paths = append(paths, selectPaths...)
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("access_method.oauth_access_token", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("access_method.oauth_access_token", prefix))
	}
	if selectPaths, err := PathsFromSelectFlagsForOAuthAccessToken(flags, flagsplugin.Prefix("access_method.oauth_access_token", prefix)); err != nil {
		return nil, err
	} else {
		paths = append(paths, selectPaths...)
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("access_method.user_session", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("access_method.user_session", prefix))
	}
	if selectPaths, err := PathsFromSelectFlagsForUserSession(flags, flagsplugin.Prefix("access_method.user_session", prefix)); err != nil {
		return nil, err
	} else {
		paths = append(paths, selectPaths...)
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("universal_rights", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("universal_rights", prefix))
	}
	// NOTE: universal_rights (Rights) does not seem to have select flags.
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("is_admin", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("is_admin", prefix))
	}
	return paths, nil
}

// AddSetFlagsForAuthInfoResponse adds flags to select fields in AuthInfoResponse.
func AddSetFlagsForAuthInfoResponse(flags *pflag.FlagSet, prefix string) {
	AddSetFlagsForAuthInfoResponse_APIKeyAccess(flags, flagsplugin.Prefix("access-method.api-key", prefix))
	AddSetFlagsForOAuthAccessToken(flags, flagsplugin.Prefix("access-method.oauth-access-token", prefix))
	AddSetFlagsForUserSession(flags, flagsplugin.Prefix("access-method.user-session", prefix))
	// FIXME: Skipping UniversalRights because it does not seem to implement AddSetFlags.
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("is-admin", prefix), ""))
}

// SetFromFlags sets the AuthInfoResponse message from flags.
func (m *AuthInfoResponse) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if selected := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("access_method.api_key", prefix)); selected {
		ov := &AuthInfoResponse_ApiKey{}
		ov.ApiKey = &AuthInfoResponse_APIKeyAccess{}
		if setPaths, err := ov.ApiKey.SetFromFlags(flags, flagsplugin.Prefix("access_method.api_key", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
		m.AccessMethod = ov
	}
	if selected := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("access_method.oauth_access_token", prefix)); selected {
		ov := &AuthInfoResponse_OauthAccessToken{}
		ov.OauthAccessToken = &OAuthAccessToken{}
		if setPaths, err := ov.OauthAccessToken.SetFromFlags(flags, flagsplugin.Prefix("access_method.oauth_access_token", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
		m.AccessMethod = ov
	}
	if selected := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("access_method.user_session", prefix)); selected {
		ov := &AuthInfoResponse_UserSession{}
		ov.UserSession = &UserSession{}
		if setPaths, err := ov.UserSession.SetFromFlags(flags, flagsplugin.Prefix("access_method.user_session", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
		m.AccessMethod = ov
	}
	// FIXME: Skipping UniversalRights because it does not seem to implement AddSetFlags.
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("is_admin", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.IsAdmin = val
		paths = append(paths, flagsplugin.Prefix("is_admin", prefix))
	}
	return paths, nil
}
