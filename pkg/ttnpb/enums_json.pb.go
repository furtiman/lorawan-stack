// Code generated by protoc-gen-go-json. DO NOT EDIT.
// versions:
// - protoc-gen-go-json v1.1.0
// - protoc             v3.9.1
// source: lorawan-stack/api/enums.proto

package ttnpb

import (
	jsonplugin "github.com/TheThingsIndustries/protoc-gen-go-json/jsonplugin"
)

// MarshalProtoJSON marshals the DownlinkPathConstraint to JSON.
func (x DownlinkPathConstraint) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	s.WriteEnumString(int32(x), DownlinkPathConstraint_name)
}

// DownlinkPathConstraint_customvalue contains custom string values that extend DownlinkPathConstraint_value.
var DownlinkPathConstraint_customvalue = map[string]int32{
	"NONE":         0,
	"PREFER_OTHER": 1,
	"NEVER":        2,
}

// UnmarshalProtoJSON unmarshals the DownlinkPathConstraint from JSON.
func (x *DownlinkPathConstraint) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	v := s.ReadEnum(DownlinkPathConstraint_value, DownlinkPathConstraint_customvalue)
	if err := s.Err(); err != nil {
		s.SetErrorf("could not read DownlinkPathConstraint enum: %v", err)
		return
	}
	*x = DownlinkPathConstraint(v)
}

// MarshalProtoJSON marshals the State to JSON.
func (x State) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	s.WriteEnumString(int32(x), State_name)
}

// State_customvalue contains custom string values that extend State_value.
var State_customvalue = map[string]int32{
	"REQUESTED": 0,
	"APPROVED":  1,
	"REJECTED":  2,
	"FLAGGED":   3,
	"SUSPENDED": 4,
}

// UnmarshalProtoJSON unmarshals the State from JSON.
func (x *State) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	v := s.ReadEnum(State_value, State_customvalue)
	if err := s.Err(); err != nil {
		s.SetErrorf("could not read State enum: %v", err)
		return
	}
	*x = State(v)
}
