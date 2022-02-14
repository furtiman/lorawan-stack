// Code generated by protoc-gen-go-flags. DO NOT EDIT.
// versions:
// - protoc-gen-go-flags v0.0.0-dev
// - protoc              v3.17.3
// source: lorawan-stack/api/qrcodegenerator.proto

package ttnpb

import (
	flagsplugin "github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin"
	pflag "github.com/spf13/pflag"
)

// AddSelectFlagsForGenerateEndDeviceQRCodeRequest adds flags to select fields in GenerateEndDeviceQRCodeRequest.
func AddSelectFlagsForGenerateEndDeviceQRCodeRequest(flags *pflag.FlagSet, prefix string) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("format-id", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("format-id", prefix), false)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("end-device", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("end-device", prefix), true)))
	AddSelectFlagsForEndDevice(flags, flagsplugin.Prefix("end-device", prefix))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("image", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("image", prefix), true)))
	// NOTE: image (GenerateEndDeviceQRCodeRequest_Image) does not seem to have select flags.
}

// SelectFromFlags outputs the fieldmask paths forGenerateEndDeviceQRCodeRequest message from select flags.
func PathsFromSelectFlagsForGenerateEndDeviceQRCodeRequest(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("format_id", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("format_id", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("end_device", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("end_device", prefix))
	}
	if selectPaths, err := PathsFromSelectFlagsForEndDevice(flags, flagsplugin.Prefix("end_device", prefix)); err != nil {
		return nil, err
	} else {
		paths = append(paths, selectPaths...)
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("image", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("image", prefix))
	}
	// NOTE: image (GenerateEndDeviceQRCodeRequest_Image) does not seem to have select flags.
	return paths, nil
}

// AddSetFlagsForGenerateEndDeviceQRCodeRequest adds flags to select fields in GenerateEndDeviceQRCodeRequest.
func AddSetFlagsForGenerateEndDeviceQRCodeRequest(flags *pflag.FlagSet, prefix string) {
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("format-id", prefix), ""))
	AddSetFlagsForEndDevice(flags, flagsplugin.Prefix("end-device", prefix))
	// FIXME: Skipping Image because it does not seem to implement AddSetFlags.
}

// SetFromFlags sets the GenerateEndDeviceQRCodeRequest message from flags.
func (m *GenerateEndDeviceQRCodeRequest) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("format_id", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.FormatId = val
		paths = append(paths, flagsplugin.Prefix("format_id", prefix))
	}
	if selected := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("end_device", prefix)); selected {
		m.EndDevice = &EndDevice{}
		if setPaths, err := m.EndDevice.SetFromFlags(flags, flagsplugin.Prefix("end_device", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	// FIXME: Skipping Image because it does not seem to implement AddSetFlags.
	return paths, nil
}
