// Code generated by protoc-gen-go-flags. DO NOT EDIT.
// versions:
// - protoc-gen-go-flags v0.0.0-dev
// - protoc              v3.17.3
// source: lorawan-stack/api/events.proto

package ttnpb

import (
	flagsplugin "github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin"
	gogo "github.com/TheThingsIndustries/protoc-gen-go-flags/gogo"
	pflag "github.com/spf13/pflag"
)

// AddSelectFlagsForEvent adds flags to select fields in Event.
func AddSelectFlagsForEvent(flags *pflag.FlagSet, prefix string) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("name", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("name", prefix), false)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("time", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("time", prefix), false)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("identifiers", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("identifiers", prefix), false)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("data", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("data", prefix), false)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("correlation-ids", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("correlation-ids", prefix), false)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("origin", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("origin", prefix), false)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("context", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("context", prefix), false)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("visibility", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("visibility", prefix), true)))
	// NOTE: visibility (Rights) does not seem to have select flags.
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("authentication", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("authentication", prefix), true)))
	// NOTE: authentication (Event_Authentication) does not seem to have select flags.
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("remote-ip", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("remote-ip", prefix), false)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("user-agent", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("user-agent", prefix), false)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("unique-id", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("unique-id", prefix), false)))
}

// SelectFromFlags outputs the fieldmask paths forEvent message from select flags.
func PathsFromSelectFlagsForEvent(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("name", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("name", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("time", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("time", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("identifiers", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("identifiers", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("data", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("data", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("correlation_ids", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("correlation_ids", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("origin", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("origin", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("context", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("context", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("visibility", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("visibility", prefix))
	}
	// NOTE: visibility (Rights) does not seem to have select flags.
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("authentication", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("authentication", prefix))
	}
	// NOTE: authentication (Event_Authentication) does not seem to have select flags.
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("remote_ip", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("remote_ip", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("user_agent", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("user_agent", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("unique_id", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("unique_id", prefix))
	}
	return paths, nil
}

// AddSetFlagsForEvent adds flags to select fields in Event.
func AddSetFlagsForEvent(flags *pflag.FlagSet, prefix string) {
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("name", prefix), ""))
	flags.AddFlag(flagsplugin.NewTimestampFlag(flagsplugin.Prefix("time", prefix), ""))
	// FIXME: Skipping Identifiers because repeated messages are currently not supported.
	// FIXME: Skipping Data because this WKT is currently not supported.
	flags.AddFlag(flagsplugin.NewStringSliceFlag(flagsplugin.Prefix("correlation-ids", prefix), ""))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("origin", prefix), ""))
	flags.AddFlag(flagsplugin.NewStringBytesMapFlag(flagsplugin.Prefix("context", prefix), ""))
	// FIXME: Skipping Visibility because it does not seem to implement AddSetFlags.
	// FIXME: Skipping Authentication because it does not seem to implement AddSetFlags.
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("remote-ip", prefix), ""))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("user-agent", prefix), ""))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("unique-id", prefix), ""))
}

// SetFromFlags sets the Event message from flags.
func (m *Event) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("name", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Name = val
		paths = append(paths, flagsplugin.Prefix("name", prefix))
	}
	if val, selected, err := flagsplugin.GetTimestamp(flags, flagsplugin.Prefix("time", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Time = gogo.SetTimestamp(val)
		paths = append(paths, flagsplugin.Prefix("time", prefix))
	}
	// FIXME: Skipping Identifiers because it does not seem to implement AddSetFlags.
	// FIXME: Skipping Data because this WKT is not supported.
	if val, selected, err := flagsplugin.GetStringSlice(flags, flagsplugin.Prefix("correlation_ids", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.CorrelationIds = val
		paths = append(paths, flagsplugin.Prefix("correlation_ids", prefix))
	}
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("origin", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Origin = val
		paths = append(paths, flagsplugin.Prefix("origin", prefix))
	}
	if val, selected, err := flagsplugin.GetStringBytesMap(flags, flagsplugin.Prefix("context", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Context = val
		paths = append(paths, flagsplugin.Prefix("context", prefix))
	}
	// FIXME: Skipping Visibility because it does not seem to implement AddSetFlags.
	// FIXME: Skipping Authentication because it does not seem to implement AddSetFlags.
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("remote_ip", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.RemoteIp = val
		paths = append(paths, flagsplugin.Prefix("remote_ip", prefix))
	}
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("user_agent", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.UserAgent = val
		paths = append(paths, flagsplugin.Prefix("user_agent", prefix))
	}
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("unique_id", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.UniqueId = val
		paths = append(paths, flagsplugin.Prefix("unique_id", prefix))
	}
	return paths, nil
}

// AddSelectFlagsForStreamEventsRequest adds flags to select fields in StreamEventsRequest.
func AddSelectFlagsForStreamEventsRequest(flags *pflag.FlagSet, prefix string) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("identifiers", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("identifiers", prefix), false)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("tail", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("tail", prefix), false)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("after", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("after", prefix), false)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("names", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("names", prefix), false)))
}

// SelectFromFlags outputs the fieldmask paths forStreamEventsRequest message from select flags.
func PathsFromSelectFlagsForStreamEventsRequest(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("identifiers", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("identifiers", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("tail", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("tail", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("after", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("after", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("names", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("names", prefix))
	}
	return paths, nil
}

// AddSetFlagsForStreamEventsRequest adds flags to select fields in StreamEventsRequest.
func AddSetFlagsForStreamEventsRequest(flags *pflag.FlagSet, prefix string) {
	// FIXME: Skipping Identifiers because repeated messages are currently not supported.
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("tail", prefix), ""))
	flags.AddFlag(flagsplugin.NewTimestampFlag(flagsplugin.Prefix("after", prefix), ""))
	flags.AddFlag(flagsplugin.NewStringSliceFlag(flagsplugin.Prefix("names", prefix), ""))
}

// SetFromFlags sets the StreamEventsRequest message from flags.
func (m *StreamEventsRequest) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	// FIXME: Skipping Identifiers because it does not seem to implement AddSetFlags.
	if val, selected, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("tail", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Tail = val
		paths = append(paths, flagsplugin.Prefix("tail", prefix))
	}
	if val, selected, err := flagsplugin.GetTimestamp(flags, flagsplugin.Prefix("after", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.After = gogo.SetTimestamp(val)
		paths = append(paths, flagsplugin.Prefix("after", prefix))
	}
	if val, selected, err := flagsplugin.GetStringSlice(flags, flagsplugin.Prefix("names", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Names = val
		paths = append(paths, flagsplugin.Prefix("names", prefix))
	}
	return paths, nil
}

// AddSelectFlagsForFindRelatedEventsResponse adds flags to select fields in FindRelatedEventsResponse.
func AddSelectFlagsForFindRelatedEventsResponse(flags *pflag.FlagSet, prefix string) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("events", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("events", prefix), false)))
}

// SelectFromFlags outputs the fieldmask paths forFindRelatedEventsResponse message from select flags.
func PathsFromSelectFlagsForFindRelatedEventsResponse(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("events", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("events", prefix))
	}
	return paths, nil
}

// AddSetFlagsForFindRelatedEventsResponse adds flags to select fields in FindRelatedEventsResponse.
func AddSetFlagsForFindRelatedEventsResponse(flags *pflag.FlagSet, prefix string) {
	// FIXME: Skipping Events because repeated messages are currently not supported.
}

// SetFromFlags sets the FindRelatedEventsResponse message from flags.
func (m *FindRelatedEventsResponse) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	// FIXME: Skipping Events because it does not seem to implement AddSetFlags.
	return paths, nil
}
