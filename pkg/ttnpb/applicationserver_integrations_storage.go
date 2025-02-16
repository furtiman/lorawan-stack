// Copyright © 2020 The Things Network Foundation, The Things Industries B.V.
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

package ttnpb

var (
	_ interface {
		IDStringer
		ExtractRequestFields(dst map[string]interface{})
	} = (*GetStoredApplicationUpRequest)(nil)
	_ interface {
		IDStringer
		ExtractRequestFields(dst map[string]interface{})
	} = (*GetStoredApplicationUpCountRequest)(nil)
)

// StoredApplicationUpTypes is a list of available ApplicationUp message types.
var StoredApplicationUpTypes = map[string]struct{}{
	"":                           {},
	"uplink_message":             {},
	"join_accept":                {},
	"downlink_ack":               {},
	"downlink_nack":              {},
	"downlink_sent":              {},
	"downlink_failed":            {},
	"downlink_queued":            {},
	"downlink_queue_invalidated": {},
	"location_solved":            {},
	"service_data":               {},
}

// WithEndDeviceIds returns the request with set EndDeviceIdentifiers.
func (m *GetStoredApplicationUpRequest) WithEndDeviceIds(ids *EndDeviceIdentifiers) *GetStoredApplicationUpRequest {
	m.EndDeviceIds = ids
	return m
}

// WithApplicationIds returns the request with set ApplicationIdentifiers.
func (m *GetStoredApplicationUpRequest) WithApplicationIds(ids *ApplicationIdentifiers) *GetStoredApplicationUpRequest {
	m.ApplicationIds = ids
	return m
}

// EntityType implements IDStringer.
func (m *GetStoredApplicationUpRequest) EntityType() string {
	if ids := m.GetEndDeviceIds(); !ids.IsZero() {
		return ids.EntityType()
	}
	return m.GetApplicationIds().EntityType()
}

// IDString implements IDStringer.
func (m *GetStoredApplicationUpRequest) IDString() string {
	if ids := m.GetEndDeviceIds(); !ids.IsZero() {
		return ids.IDString()
	}
	return m.GetApplicationIds().IDString()
}

// ExtractRequestFields is used by github.com/grpc-ecosystem/go-grpc-middleware/tags.
func (m *GetStoredApplicationUpRequest) ExtractRequestFields(dst map[string]interface{}) {
	if ids := m.GetEndDeviceIds(); !ids.IsZero() {
		ids.ExtractRequestFields(dst)
		return
	}
	m.GetApplicationIds().ExtractRequestFields(dst)
}

// WithEndDeviceIds returns the request with set EndDeviceIdentifiers.
func (m *GetStoredApplicationUpCountRequest) WithEndDeviceIds(ids *EndDeviceIdentifiers) *GetStoredApplicationUpCountRequest {
	m.EndDeviceIds = ids
	return m
}

// WithApplicationIds returns the request with set ApplicationIdentifiers.
func (m *GetStoredApplicationUpCountRequest) WithApplicationIds(ids *ApplicationIdentifiers) *GetStoredApplicationUpCountRequest {
	m.ApplicationIds = ids
	return m
}

// EntityType implements IDStringer.
func (m *GetStoredApplicationUpCountRequest) EntityType() string {
	if ids := m.GetEndDeviceIds(); !ids.IsZero() {
		return ids.EntityType()
	}
	return m.GetApplicationIds().EntityType()
}

// IDString implements IDStringer.
func (m *GetStoredApplicationUpCountRequest) IDString() string {
	if ids := m.GetEndDeviceIds(); !ids.IsZero() {
		return ids.IDString()
	}
	return m.GetApplicationIds().IDString()
}

// ExtractRequestFields is used by github.com/grpc-ecosystem/go-grpc-middleware/tags.
func (m *GetStoredApplicationUpCountRequest) ExtractRequestFields(dst map[string]interface{}) {
	if ids := m.GetEndDeviceIds(); !ids.IsZero() {
		ids.ExtractRequestFields(dst)
		return
	}
	m.GetApplicationIds().ExtractRequestFields(dst)
}
