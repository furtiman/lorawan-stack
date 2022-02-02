// Copyright © 2022 The Things Network Foundation, The Things Industries B.V.
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

package gatewayconfigurationserver

import (
	"context"
	"encoding/json"

	types "github.com/gogo/protobuf/types"
	"go.thethings.network/lorawan-stack/v3/pkg/auth/rights"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/pfconfig/cpf"
	"go.thethings.network/lorawan-stack/v3/pkg/pfconfig/semtechudp"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

var errInvalidConfigurationFormat = errors.DefineInvalidArgument("invalid_configuration_format", "configuration format `{format}` is not supported")

func (s *Server) getGateway(ctx context.Context, gtwID *ttnpb.GatewayIdentifiers) (*ttnpb.Gateway, error) {
	cc, err := s.GetPeerConn(ctx, ttnpb.ClusterRole_ENTITY_REGISTRY, nil)
	if err != nil {
		return nil, err
	}
	client := ttnpb.NewGatewayRegistryClient(cc)
	gtw, err := client.Get(ctx, &ttnpb.GetGatewayRequest{
		GatewayIds: gtwID,
		FieldMask: &types.FieldMask{
			Paths: []string{
				"antennas",
				"frequency_plan_id",
				"gateway_server_address",
			},
		},
	}, s.WithClusterAuth())
	if err != nil {
		return nil, err
	}
	return gtw, nil
}

func (s *Server) GetGatewayConfiguration(ctx context.Context, req *ttnpb.GetGatewayConfigurationRequest) (*ttnpb.GetGatewayConfigurationResponse, error) {
	if s.config.RequireAuth {
		err := rights.RequireGateway(ctx, *req.GatewayIds, ttnpb.Right_RIGHT_GATEWAY_INFO)
		if err != nil {
			return nil, err
		}
	}

	gtw, err := s.getGateway(ctx, req.GetGatewayIds())
	if err != nil {
		return nil, err
	}

	fps, err := s.FrequencyPlansStore(ctx)
	if err != nil {
		return nil, err
	}

	var configContent []byte
	switch req.Format {
	case "semtechudp":
		config, err := semtechudp.Build(gtw, fps)
		if err != nil {
			return nil, err
		}
		configContent, err = json.Marshal(config)

	case "lorad":
		config, err := cpf.BuildLorad(gtw, fps)
		if err != nil {
			return nil, err
		}
		configContent, err = json.Marshal(config)

	case "lorafwd":
		config, err := cpf.BuildLorafwd(gtw)
		if err != nil {
			return nil, err
		}
		configContent, err = config.MarshalText()
	default:
		return nil, errInvalidConfigurationFormat.WithAttributes("format", req.Format)
	}

	return &ttnpb.GetGatewayConfigurationResponse{
		Contents: configContent,
	}, err
}
