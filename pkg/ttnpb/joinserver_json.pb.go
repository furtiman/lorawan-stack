// Code generated by protoc-gen-go-json. DO NOT EDIT.
// versions:
// - protoc-gen-go-json v1.3.1
// - protoc             v3.9.1
// source: lorawan-stack/api/joinserver.proto

package ttnpb

import (
	gogo "github.com/TheThingsIndustries/protoc-gen-go-json/gogo"
	jsonplugin "github.com/TheThingsIndustries/protoc-gen-go-json/jsonplugin"
)

// MarshalProtoJSON marshals the CryptoServicePayloadRequest message to JSON.
func (x *CryptoServicePayloadRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Ids != nil || s.HasField("ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("ids")
		// NOTE: EndDeviceIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.Ids)
	}
	if x.LorawanVersion != 0 || s.HasField("lorawan_version") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("lorawan_version")
		x.LorawanVersion.MarshalProtoJSON(s)
	}
	if len(x.Payload) > 0 || s.HasField("payload") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("payload")
		s.WriteBytes(x.Payload)
	}
	if x.ProvisionerId != "" || s.HasField("provisioner_id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("provisioner_id")
		s.WriteString(x.ProvisionerId)
	}
	if x.ProvisioningData != nil || s.HasField("provisioning_data") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("provisioning_data")
		if x.ProvisioningData == nil {
			s.WriteNil()
		} else {
			gogo.MarshalStruct(s, x.ProvisioningData)
		}
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the CryptoServicePayloadRequest to JSON.
func (x CryptoServicePayloadRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(&x)
}

// UnmarshalProtoJSON unmarshals the CryptoServicePayloadRequest message from JSON.
func (x *CryptoServicePayloadRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "ids":
			s.AddField("ids")
			if s.ReadNil() {
				x.Ids = nil
				return
			}
			// NOTE: EndDeviceIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v EndDeviceIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.Ids = &v
		case "lorawan_version", "lorawanVersion":
			s.AddField("lorawan_version")
			x.LorawanVersion.UnmarshalProtoJSON(s)
		case "payload":
			s.AddField("payload")
			x.Payload = s.ReadBytes()
		case "provisioner_id", "provisionerId":
			s.AddField("provisioner_id")
			x.ProvisionerId = s.ReadString()
		case "provisioning_data", "provisioningData":
			s.AddField("provisioning_data")
			if s.ReadNil() {
				x.ProvisioningData = nil
				return
			}
			v := gogo.UnmarshalStruct(s)
			if s.Err() != nil {
				return
			}
			x.ProvisioningData = v
		}
	})
}

// UnmarshalJSON unmarshals the CryptoServicePayloadRequest from JSON.
func (x *CryptoServicePayloadRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the JoinAcceptMICRequest message to JSON.
func (x *JoinAcceptMICRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.PayloadRequest != nil || s.HasField("payload_request") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("payload_request")
		x.PayloadRequest.MarshalProtoJSON(s.WithField("payload_request"))
	}
	if x.JoinRequestType != 0 || s.HasField("join_request_type") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("join_request_type")
		x.JoinRequestType.MarshalProtoJSON(s)
	}
	if len(x.DevNonce) > 0 || s.HasField("dev_nonce") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("dev_nonce")
		x.DevNonce.MarshalProtoJSON(s.WithField("dev_nonce"))
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the JoinAcceptMICRequest to JSON.
func (x JoinAcceptMICRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(&x)
}

// UnmarshalProtoJSON unmarshals the JoinAcceptMICRequest message from JSON.
func (x *JoinAcceptMICRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "payload_request", "payloadRequest":
			if s.ReadNil() {
				x.PayloadRequest = nil
				return
			}
			x.PayloadRequest = &CryptoServicePayloadRequest{}
			x.PayloadRequest.UnmarshalProtoJSON(s.WithField("payload_request", true))
		case "join_request_type", "joinRequestType":
			s.AddField("join_request_type")
			x.JoinRequestType.UnmarshalProtoJSON(s)
		case "dev_nonce", "devNonce":
			s.AddField("dev_nonce")
			x.DevNonce.UnmarshalProtoJSON(s.WithField("dev_nonce", false))
		}
	})
}

// UnmarshalJSON unmarshals the JoinAcceptMICRequest from JSON.
func (x *JoinAcceptMICRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the DeriveSessionKeysRequest message to JSON.
func (x *DeriveSessionKeysRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Ids != nil || s.HasField("ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("ids")
		// NOTE: EndDeviceIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.Ids)
	}
	if x.LorawanVersion != 0 || s.HasField("lorawan_version") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("lorawan_version")
		x.LorawanVersion.MarshalProtoJSON(s)
	}
	if len(x.JoinNonce) > 0 || s.HasField("join_nonce") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("join_nonce")
		x.JoinNonce.MarshalProtoJSON(s.WithField("join_nonce"))
	}
	if len(x.DevNonce) > 0 || s.HasField("dev_nonce") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("dev_nonce")
		x.DevNonce.MarshalProtoJSON(s.WithField("dev_nonce"))
	}
	if len(x.NetId) > 0 || s.HasField("net_id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("net_id")
		x.NetId.MarshalProtoJSON(s.WithField("net_id"))
	}
	if x.ProvisionerId != "" || s.HasField("provisioner_id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("provisioner_id")
		s.WriteString(x.ProvisionerId)
	}
	if x.ProvisioningData != nil || s.HasField("provisioning_data") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("provisioning_data")
		if x.ProvisioningData == nil {
			s.WriteNil()
		} else {
			gogo.MarshalStruct(s, x.ProvisioningData)
		}
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the DeriveSessionKeysRequest to JSON.
func (x DeriveSessionKeysRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(&x)
}

// UnmarshalProtoJSON unmarshals the DeriveSessionKeysRequest message from JSON.
func (x *DeriveSessionKeysRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "ids":
			s.AddField("ids")
			if s.ReadNil() {
				x.Ids = nil
				return
			}
			// NOTE: EndDeviceIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v EndDeviceIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.Ids = &v
		case "lorawan_version", "lorawanVersion":
			s.AddField("lorawan_version")
			x.LorawanVersion.UnmarshalProtoJSON(s)
		case "join_nonce", "joinNonce":
			s.AddField("join_nonce")
			x.JoinNonce.UnmarshalProtoJSON(s.WithField("join_nonce", false))
		case "dev_nonce", "devNonce":
			s.AddField("dev_nonce")
			x.DevNonce.UnmarshalProtoJSON(s.WithField("dev_nonce", false))
		case "net_id", "netId":
			s.AddField("net_id")
			x.NetId.UnmarshalProtoJSON(s.WithField("net_id", false))
		case "provisioner_id", "provisionerId":
			s.AddField("provisioner_id")
			x.ProvisionerId = s.ReadString()
		case "provisioning_data", "provisioningData":
			s.AddField("provisioning_data")
			if s.ReadNil() {
				x.ProvisioningData = nil
				return
			}
			v := gogo.UnmarshalStruct(s)
			if s.Err() != nil {
				return
			}
			x.ProvisioningData = v
		}
	})
}

// UnmarshalJSON unmarshals the DeriveSessionKeysRequest from JSON.
func (x *DeriveSessionKeysRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}
