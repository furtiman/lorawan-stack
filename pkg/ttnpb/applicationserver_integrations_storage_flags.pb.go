// Code generated by protoc-gen-go-flags. DO NOT EDIT.
// versions:
// - protoc-gen-go-flags v0.0.0-dev
// - protoc              v3.17.3
// source: lorawan-stack/api/applicationserver_integrations_storage.proto

package ttnpb

import (
	flagsplugin "github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin"
	gogo "github.com/TheThingsIndustries/protoc-gen-go-flags/gogo"
	types "github.com/gogo/protobuf/types"
	pflag "github.com/spf13/pflag"
)

// AddSelectFlagsForGetStoredApplicationUpRequest adds flags to select fields in GetStoredApplicationUpRequest.
func AddSelectFlagsForGetStoredApplicationUpRequest(flags *pflag.FlagSet, prefix string) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("application-ids", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("application-ids", prefix), true)))
	AddSelectFlagsForApplicationIdentifiers(flags, flagsplugin.Prefix("application-ids", prefix))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("end-device-ids", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("end-device-ids", prefix), true)))
	AddSelectFlagsForEndDeviceIdentifiers(flags, flagsplugin.Prefix("end-device-ids", prefix))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("type", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("type", prefix), false)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("limit", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("limit", prefix), false)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("after", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("after", prefix), false)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("before", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("before", prefix), false)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("f-port", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("f-port", prefix), false)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("order", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("order", prefix), false)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("field-mask", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("field-mask", prefix), false)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("last", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("last", prefix), false)))
}

// SelectFromFlags outputs the fieldmask paths forGetStoredApplicationUpRequest message from select flags.
func PathsFromSelectFlagsForGetStoredApplicationUpRequest(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("application_ids", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("application_ids", prefix))
	}
	if selectPaths, err := PathsFromSelectFlagsForApplicationIdentifiers(flags, flagsplugin.Prefix("application_ids", prefix)); err != nil {
		return nil, err
	} else {
		paths = append(paths, selectPaths...)
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("end_device_ids", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("end_device_ids", prefix))
	}
	if selectPaths, err := PathsFromSelectFlagsForEndDeviceIdentifiers(flags, flagsplugin.Prefix("end_device_ids", prefix)); err != nil {
		return nil, err
	} else {
		paths = append(paths, selectPaths...)
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("type", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("type", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("limit", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("limit", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("after", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("after", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("before", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("before", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("f_port", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("f_port", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("order", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("order", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("field_mask", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("field_mask", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("last", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("last", prefix))
	}
	return paths, nil
}

// AddSetFlagsForGetStoredApplicationUpRequest adds flags to select fields in GetStoredApplicationUpRequest.
func AddSetFlagsForGetStoredApplicationUpRequest(flags *pflag.FlagSet, prefix string) {
	AddSetFlagsForApplicationIdentifiers(flags, flagsplugin.Prefix("application-ids", prefix))
	AddSetFlagsForEndDeviceIdentifiers(flags, flagsplugin.Prefix("end-device-ids", prefix))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("type", prefix), ""))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("limit.value", prefix), ""))
	flagsplugin.AddAlias(flags, flagsplugin.Prefix("limit.value", prefix), flagsplugin.Prefix("limit", prefix))
	flags.AddFlag(flagsplugin.NewTimestampFlag(flagsplugin.Prefix("after", prefix), ""))
	flags.AddFlag(flagsplugin.NewTimestampFlag(flagsplugin.Prefix("before", prefix), ""))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("f-port.value", prefix), ""))
	flagsplugin.AddAlias(flags, flagsplugin.Prefix("f-port.value", prefix), flagsplugin.Prefix("f-port", prefix))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("order", prefix), ""))
	flags.AddFlag(flagsplugin.NewStringSliceFlag(flagsplugin.Prefix("field-mask", prefix), ""))
	flags.AddFlag(flagsplugin.NewDurationFlag(flagsplugin.Prefix("last", prefix), ""))
}

// SetFromFlags sets the GetStoredApplicationUpRequest message from flags.
func (m *GetStoredApplicationUpRequest) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if selected := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("application_ids", prefix)); selected {
		m.ApplicationIds = &ApplicationIdentifiers{}
		if setPaths, err := m.ApplicationIds.SetFromFlags(flags, flagsplugin.Prefix("application_ids", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	if selected := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("end_device_ids", prefix)); selected {
		m.EndDeviceIds = &EndDeviceIdentifiers{}
		if setPaths, err := m.EndDeviceIds.SetFromFlags(flags, flagsplugin.Prefix("end_device_ids", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("type", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Type = val
		paths = append(paths, flagsplugin.Prefix("type", prefix))
	}
	if val, selected, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("limit.value", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Limit = &types.UInt32Value{Value: val}
		paths = append(paths, flagsplugin.Prefix("limit", prefix))
	}
	if val, selected, err := flagsplugin.GetTimestamp(flags, flagsplugin.Prefix("after", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.After = gogo.SetTimestamp(val)
		paths = append(paths, flagsplugin.Prefix("after", prefix))
	}
	if val, selected, err := flagsplugin.GetTimestamp(flags, flagsplugin.Prefix("before", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Before = gogo.SetTimestamp(val)
		paths = append(paths, flagsplugin.Prefix("before", prefix))
	}
	if val, selected, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("f_port.value", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.FPort = &types.UInt32Value{Value: val}
		paths = append(paths, flagsplugin.Prefix("f_port", prefix))
	}
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("order", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Order = val
		paths = append(paths, flagsplugin.Prefix("order", prefix))
	}
	if val, selected, err := flagsplugin.GetStringSlice(flags, flagsplugin.Prefix("field_mask", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.FieldMask = gogo.SetFieldMask(val)
		paths = append(paths, flagsplugin.Prefix("field_mask", prefix))
	}
	if val, selected, err := flagsplugin.GetDuration(flags, flagsplugin.Prefix("last", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Last = gogo.SetDuration(val)
		paths = append(paths, flagsplugin.Prefix("last", prefix))
	}
	return paths, nil
}

// AddSelectFlagsForGetStoredApplicationUpCountRequest adds flags to select fields in GetStoredApplicationUpCountRequest.
func AddSelectFlagsForGetStoredApplicationUpCountRequest(flags *pflag.FlagSet, prefix string) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("application-ids", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("application-ids", prefix), true)))
	AddSelectFlagsForApplicationIdentifiers(flags, flagsplugin.Prefix("application-ids", prefix))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("end-device-ids", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("end-device-ids", prefix), true)))
	AddSelectFlagsForEndDeviceIdentifiers(flags, flagsplugin.Prefix("end-device-ids", prefix))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("type", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("type", prefix), false)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("after", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("after", prefix), false)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("before", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("before", prefix), false)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("f-port", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("f-port", prefix), false)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("last", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("last", prefix), false)))
}

// SelectFromFlags outputs the fieldmask paths forGetStoredApplicationUpCountRequest message from select flags.
func PathsFromSelectFlagsForGetStoredApplicationUpCountRequest(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("application_ids", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("application_ids", prefix))
	}
	if selectPaths, err := PathsFromSelectFlagsForApplicationIdentifiers(flags, flagsplugin.Prefix("application_ids", prefix)); err != nil {
		return nil, err
	} else {
		paths = append(paths, selectPaths...)
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("end_device_ids", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("end_device_ids", prefix))
	}
	if selectPaths, err := PathsFromSelectFlagsForEndDeviceIdentifiers(flags, flagsplugin.Prefix("end_device_ids", prefix)); err != nil {
		return nil, err
	} else {
		paths = append(paths, selectPaths...)
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("type", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("type", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("after", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("after", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("before", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("before", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("f_port", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("f_port", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("last", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("last", prefix))
	}
	return paths, nil
}

// AddSetFlagsForGetStoredApplicationUpCountRequest adds flags to select fields in GetStoredApplicationUpCountRequest.
func AddSetFlagsForGetStoredApplicationUpCountRequest(flags *pflag.FlagSet, prefix string) {
	AddSetFlagsForApplicationIdentifiers(flags, flagsplugin.Prefix("application-ids", prefix))
	AddSetFlagsForEndDeviceIdentifiers(flags, flagsplugin.Prefix("end-device-ids", prefix))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("type", prefix), ""))
	flags.AddFlag(flagsplugin.NewTimestampFlag(flagsplugin.Prefix("after", prefix), ""))
	flags.AddFlag(flagsplugin.NewTimestampFlag(flagsplugin.Prefix("before", prefix), ""))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("f-port.value", prefix), ""))
	flagsplugin.AddAlias(flags, flagsplugin.Prefix("f-port.value", prefix), flagsplugin.Prefix("f-port", prefix))
	flags.AddFlag(flagsplugin.NewDurationFlag(flagsplugin.Prefix("last", prefix), ""))
}

// SetFromFlags sets the GetStoredApplicationUpCountRequest message from flags.
func (m *GetStoredApplicationUpCountRequest) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if selected := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("application_ids", prefix)); selected {
		m.ApplicationIds = &ApplicationIdentifiers{}
		if setPaths, err := m.ApplicationIds.SetFromFlags(flags, flagsplugin.Prefix("application_ids", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	if selected := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("end_device_ids", prefix)); selected {
		m.EndDeviceIds = &EndDeviceIdentifiers{}
		if setPaths, err := m.EndDeviceIds.SetFromFlags(flags, flagsplugin.Prefix("end_device_ids", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("type", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Type = val
		paths = append(paths, flagsplugin.Prefix("type", prefix))
	}
	if val, selected, err := flagsplugin.GetTimestamp(flags, flagsplugin.Prefix("after", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.After = gogo.SetTimestamp(val)
		paths = append(paths, flagsplugin.Prefix("after", prefix))
	}
	if val, selected, err := flagsplugin.GetTimestamp(flags, flagsplugin.Prefix("before", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Before = gogo.SetTimestamp(val)
		paths = append(paths, flagsplugin.Prefix("before", prefix))
	}
	if val, selected, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("f_port.value", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.FPort = &types.UInt32Value{Value: val}
		paths = append(paths, flagsplugin.Prefix("f_port", prefix))
	}
	if val, selected, err := flagsplugin.GetDuration(flags, flagsplugin.Prefix("last", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Last = gogo.SetDuration(val)
		paths = append(paths, flagsplugin.Prefix("last", prefix))
	}
	return paths, nil
}
