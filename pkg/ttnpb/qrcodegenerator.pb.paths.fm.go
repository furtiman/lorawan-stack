// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

var QRCodeFormatFieldPathsNested = []string{
	"description",
	"field_mask",
	"name",
}

var QRCodeFormatFieldPathsTopLevel = []string{
	"description",
	"field_mask",
	"name",
}
var QRCodeFormatsFieldPathsNested = []string{
	"formats",
}

var QRCodeFormatsFieldPathsTopLevel = []string{
	"formats",
}
var GetQRCodeFormatRequestFieldPathsNested = []string{
	"format_id",
}

var GetQRCodeFormatRequestFieldPathsTopLevel = []string{
	"format_id",
}
var GenerateEndDeviceQRCodeRequestFieldPathsNested = []string{
	"end_device",
	"end_device.activated_at",
	"end_device.application_server_address",
	"end_device.application_server_id",
	"end_device.application_server_kek_label",
	"end_device.attributes",
	"end_device.battery_percentage",
	"end_device.claim_authentication_code",
	"end_device.claim_authentication_code.valid_from",
	"end_device.claim_authentication_code.valid_to",
	"end_device.claim_authentication_code.value",
	"end_device.created_at",
	"end_device.description",
	"end_device.downlink_margin",
	"end_device.formatters",
	"end_device.formatters.down_formatter",
	"end_device.formatters.down_formatter_parameter",
	"end_device.formatters.up_formatter",
	"end_device.formatters.up_formatter_parameter",
	"end_device.frequency_plan_id",
	"end_device.ids",
	"end_device.ids.application_ids",
	"end_device.ids.application_ids.application_id",
	"end_device.ids.dev_addr",
	"end_device.ids.dev_eui",
	"end_device.ids.device_id",
	"end_device.ids.join_eui",
	"end_device.join_server_address",
	"end_device.last_dev_nonce",
	"end_device.last_dev_status_received_at",
	"end_device.last_join_nonce",
	"end_device.last_rj_count_0",
	"end_device.last_rj_count_1",
	"end_device.locations",
	"end_device.lorawan_phy_version",
	"end_device.lorawan_version",
	"end_device.mac_settings",
	"end_device.mac_settings.adr_margin",
	"end_device.mac_settings.beacon_frequency",
	"end_device.mac_settings.beacon_frequency.value",
	"end_device.mac_settings.class_b_c_downlink_interval",
	"end_device.mac_settings.class_b_timeout",
	"end_device.mac_settings.class_c_timeout",
	"end_device.mac_settings.desired_adr_ack_delay_exponent",
	"end_device.mac_settings.desired_adr_ack_delay_exponent.value",
	"end_device.mac_settings.desired_adr_ack_limit_exponent",
	"end_device.mac_settings.desired_adr_ack_limit_exponent.value",
	"end_device.mac_settings.desired_beacon_frequency",
	"end_device.mac_settings.desired_beacon_frequency.value",
	"end_device.mac_settings.desired_max_duty_cycle",
	"end_device.mac_settings.desired_max_duty_cycle.value",
	"end_device.mac_settings.desired_max_eirp",
	"end_device.mac_settings.desired_max_eirp.value",
	"end_device.mac_settings.desired_ping_slot_data_rate_index",
	"end_device.mac_settings.desired_ping_slot_data_rate_index.value",
	"end_device.mac_settings.desired_ping_slot_frequency",
	"end_device.mac_settings.desired_ping_slot_frequency.value",
	"end_device.mac_settings.desired_rx1_data_rate_offset",
	"end_device.mac_settings.desired_rx1_data_rate_offset.value",
	"end_device.mac_settings.desired_rx1_delay",
	"end_device.mac_settings.desired_rx1_delay.value",
	"end_device.mac_settings.desired_rx2_data_rate_index",
	"end_device.mac_settings.desired_rx2_data_rate_index.value",
	"end_device.mac_settings.desired_rx2_frequency",
	"end_device.mac_settings.desired_rx2_frequency.value",
	"end_device.mac_settings.factory_preset_frequencies",
	"end_device.mac_settings.max_duty_cycle",
	"end_device.mac_settings.max_duty_cycle.value",
	"end_device.mac_settings.ping_slot_data_rate_index",
	"end_device.mac_settings.ping_slot_data_rate_index.value",
	"end_device.mac_settings.ping_slot_frequency",
	"end_device.mac_settings.ping_slot_frequency.value",
	"end_device.mac_settings.ping_slot_periodicity",
	"end_device.mac_settings.ping_slot_periodicity.value",
	"end_device.mac_settings.resets_f_cnt",
	"end_device.mac_settings.resets_f_cnt.value",
	"end_device.mac_settings.rx1_data_rate_offset",
	"end_device.mac_settings.rx1_data_rate_offset.value",
	"end_device.mac_settings.rx1_delay",
	"end_device.mac_settings.rx1_delay.value",
	"end_device.mac_settings.rx2_data_rate_index",
	"end_device.mac_settings.rx2_data_rate_index.value",
	"end_device.mac_settings.rx2_frequency",
	"end_device.mac_settings.rx2_frequency.value",
	"end_device.mac_settings.status_count_periodicity",
	"end_device.mac_settings.status_time_periodicity",
	"end_device.mac_settings.supports_32_bit_f_cnt",
	"end_device.mac_settings.supports_32_bit_f_cnt.value",
	"end_device.mac_settings.use_adr",
	"end_device.mac_settings.use_adr.value",
	"end_device.mac_state",
	"end_device.mac_state.current_parameters",
	"end_device.mac_state.current_parameters.adr_ack_delay",
	"end_device.mac_state.current_parameters.adr_ack_delay_exponent",
	"end_device.mac_state.current_parameters.adr_ack_delay_exponent.value",
	"end_device.mac_state.current_parameters.adr_ack_limit",
	"end_device.mac_state.current_parameters.adr_ack_limit_exponent",
	"end_device.mac_state.current_parameters.adr_ack_limit_exponent.value",
	"end_device.mac_state.current_parameters.adr_data_rate_index",
	"end_device.mac_state.current_parameters.adr_nb_trans",
	"end_device.mac_state.current_parameters.adr_tx_power_index",
	"end_device.mac_state.current_parameters.beacon_frequency",
	"end_device.mac_state.current_parameters.channels",
	"end_device.mac_state.current_parameters.downlink_dwell_time",
	"end_device.mac_state.current_parameters.downlink_dwell_time.value",
	"end_device.mac_state.current_parameters.max_duty_cycle",
	"end_device.mac_state.current_parameters.max_eirp",
	"end_device.mac_state.current_parameters.ping_slot_data_rate_index",
	"end_device.mac_state.current_parameters.ping_slot_data_rate_index_value",
	"end_device.mac_state.current_parameters.ping_slot_data_rate_index_value.value",
	"end_device.mac_state.current_parameters.ping_slot_frequency",
	"end_device.mac_state.current_parameters.rejoin_count_periodicity",
	"end_device.mac_state.current_parameters.rejoin_time_periodicity",
	"end_device.mac_state.current_parameters.rx1_data_rate_offset",
	"end_device.mac_state.current_parameters.rx1_delay",
	"end_device.mac_state.current_parameters.rx2_data_rate_index",
	"end_device.mac_state.current_parameters.rx2_frequency",
	"end_device.mac_state.current_parameters.uplink_dwell_time",
	"end_device.mac_state.current_parameters.uplink_dwell_time.value",
	"end_device.mac_state.desired_parameters",
	"end_device.mac_state.desired_parameters.adr_ack_delay",
	"end_device.mac_state.desired_parameters.adr_ack_delay_exponent",
	"end_device.mac_state.desired_parameters.adr_ack_delay_exponent.value",
	"end_device.mac_state.desired_parameters.adr_ack_limit",
	"end_device.mac_state.desired_parameters.adr_ack_limit_exponent",
	"end_device.mac_state.desired_parameters.adr_ack_limit_exponent.value",
	"end_device.mac_state.desired_parameters.adr_data_rate_index",
	"end_device.mac_state.desired_parameters.adr_nb_trans",
	"end_device.mac_state.desired_parameters.adr_tx_power_index",
	"end_device.mac_state.desired_parameters.beacon_frequency",
	"end_device.mac_state.desired_parameters.channels",
	"end_device.mac_state.desired_parameters.downlink_dwell_time",
	"end_device.mac_state.desired_parameters.downlink_dwell_time.value",
	"end_device.mac_state.desired_parameters.max_duty_cycle",
	"end_device.mac_state.desired_parameters.max_eirp",
	"end_device.mac_state.desired_parameters.ping_slot_data_rate_index",
	"end_device.mac_state.desired_parameters.ping_slot_data_rate_index_value",
	"end_device.mac_state.desired_parameters.ping_slot_data_rate_index_value.value",
	"end_device.mac_state.desired_parameters.ping_slot_frequency",
	"end_device.mac_state.desired_parameters.rejoin_count_periodicity",
	"end_device.mac_state.desired_parameters.rejoin_time_periodicity",
	"end_device.mac_state.desired_parameters.rx1_data_rate_offset",
	"end_device.mac_state.desired_parameters.rx1_delay",
	"end_device.mac_state.desired_parameters.rx2_data_rate_index",
	"end_device.mac_state.desired_parameters.rx2_frequency",
	"end_device.mac_state.desired_parameters.uplink_dwell_time",
	"end_device.mac_state.desired_parameters.uplink_dwell_time.value",
	"end_device.mac_state.device_class",
	"end_device.mac_state.last_adr_change_f_cnt_up",
	"end_device.mac_state.last_confirmed_downlink_at",
	"end_device.mac_state.last_dev_status_f_cnt_up",
	"end_device.mac_state.last_downlink_at",
	"end_device.mac_state.last_network_initiated_downlink_at",
	"end_device.mac_state.lorawan_version",
	"end_device.mac_state.pending_application_downlink",
	"end_device.mac_state.pending_application_downlink.class_b_c",
	"end_device.mac_state.pending_application_downlink.class_b_c.absolute_time",
	"end_device.mac_state.pending_application_downlink.class_b_c.gateways",
	"end_device.mac_state.pending_application_downlink.confirmed",
	"end_device.mac_state.pending_application_downlink.correlation_ids",
	"end_device.mac_state.pending_application_downlink.decoded_payload",
	"end_device.mac_state.pending_application_downlink.decoded_payload_warnings",
	"end_device.mac_state.pending_application_downlink.f_cnt",
	"end_device.mac_state.pending_application_downlink.f_port",
	"end_device.mac_state.pending_application_downlink.frm_payload",
	"end_device.mac_state.pending_application_downlink.priority",
	"end_device.mac_state.pending_application_downlink.session_key_id",
	"end_device.mac_state.pending_join_request",
	"end_device.mac_state.pending_join_request.cf_list",
	"end_device.mac_state.pending_join_request.cf_list.ch_masks",
	"end_device.mac_state.pending_join_request.cf_list.freq",
	"end_device.mac_state.pending_join_request.cf_list.type",
	"end_device.mac_state.pending_join_request.downlink_settings",
	"end_device.mac_state.pending_join_request.downlink_settings.opt_neg",
	"end_device.mac_state.pending_join_request.downlink_settings.rx1_dr_offset",
	"end_device.mac_state.pending_join_request.downlink_settings.rx2_dr",
	"end_device.mac_state.pending_join_request.rx_delay",
	"end_device.mac_state.pending_requests",
	"end_device.mac_state.ping_slot_periodicity",
	"end_device.mac_state.ping_slot_periodicity.value",
	"end_device.mac_state.queued_join_accept",
	"end_device.mac_state.queued_join_accept.correlation_ids",
	"end_device.mac_state.queued_join_accept.dev_addr",
	"end_device.mac_state.queued_join_accept.keys",
	"end_device.mac_state.queued_join_accept.keys.app_s_key",
	"end_device.mac_state.queued_join_accept.keys.app_s_key.encrypted_key",
	"end_device.mac_state.queued_join_accept.keys.app_s_key.kek_label",
	"end_device.mac_state.queued_join_accept.keys.app_s_key.key",
	"end_device.mac_state.queued_join_accept.keys.f_nwk_s_int_key",
	"end_device.mac_state.queued_join_accept.keys.f_nwk_s_int_key.encrypted_key",
	"end_device.mac_state.queued_join_accept.keys.f_nwk_s_int_key.kek_label",
	"end_device.mac_state.queued_join_accept.keys.f_nwk_s_int_key.key",
	"end_device.mac_state.queued_join_accept.keys.nwk_s_enc_key",
	"end_device.mac_state.queued_join_accept.keys.nwk_s_enc_key.encrypted_key",
	"end_device.mac_state.queued_join_accept.keys.nwk_s_enc_key.kek_label",
	"end_device.mac_state.queued_join_accept.keys.nwk_s_enc_key.key",
	"end_device.mac_state.queued_join_accept.keys.s_nwk_s_int_key",
	"end_device.mac_state.queued_join_accept.keys.s_nwk_s_int_key.encrypted_key",
	"end_device.mac_state.queued_join_accept.keys.s_nwk_s_int_key.kek_label",
	"end_device.mac_state.queued_join_accept.keys.s_nwk_s_int_key.key",
	"end_device.mac_state.queued_join_accept.keys.session_key_id",
	"end_device.mac_state.queued_join_accept.net_id",
	"end_device.mac_state.queued_join_accept.payload",
	"end_device.mac_state.queued_join_accept.request",
	"end_device.mac_state.queued_join_accept.request.cf_list",
	"end_device.mac_state.queued_join_accept.request.cf_list.ch_masks",
	"end_device.mac_state.queued_join_accept.request.cf_list.freq",
	"end_device.mac_state.queued_join_accept.request.cf_list.type",
	"end_device.mac_state.queued_join_accept.request.downlink_settings",
	"end_device.mac_state.queued_join_accept.request.downlink_settings.opt_neg",
	"end_device.mac_state.queued_join_accept.request.downlink_settings.rx1_dr_offset",
	"end_device.mac_state.queued_join_accept.request.downlink_settings.rx2_dr",
	"end_device.mac_state.queued_join_accept.request.rx_delay",
	"end_device.mac_state.queued_responses",
	"end_device.mac_state.recent_downlinks",
	"end_device.mac_state.recent_uplinks",
	"end_device.mac_state.rejected_adr_data_rate_indexes",
	"end_device.mac_state.rejected_adr_tx_power_indexes",
	"end_device.mac_state.rejected_data_rate_ranges",
	"end_device.mac_state.rejected_frequencies",
	"end_device.mac_state.rx_windows_available",
	"end_device.max_frequency",
	"end_device.min_frequency",
	"end_device.multicast",
	"end_device.name",
	"end_device.net_id",
	"end_device.network_server_address",
	"end_device.network_server_kek_label",
	"end_device.pending_mac_state",
	"end_device.pending_mac_state.current_parameters",
	"end_device.pending_mac_state.current_parameters.adr_ack_delay",
	"end_device.pending_mac_state.current_parameters.adr_ack_delay_exponent",
	"end_device.pending_mac_state.current_parameters.adr_ack_delay_exponent.value",
	"end_device.pending_mac_state.current_parameters.adr_ack_limit",
	"end_device.pending_mac_state.current_parameters.adr_ack_limit_exponent",
	"end_device.pending_mac_state.current_parameters.adr_ack_limit_exponent.value",
	"end_device.pending_mac_state.current_parameters.adr_data_rate_index",
	"end_device.pending_mac_state.current_parameters.adr_nb_trans",
	"end_device.pending_mac_state.current_parameters.adr_tx_power_index",
	"end_device.pending_mac_state.current_parameters.beacon_frequency",
	"end_device.pending_mac_state.current_parameters.channels",
	"end_device.pending_mac_state.current_parameters.downlink_dwell_time",
	"end_device.pending_mac_state.current_parameters.downlink_dwell_time.value",
	"end_device.pending_mac_state.current_parameters.max_duty_cycle",
	"end_device.pending_mac_state.current_parameters.max_eirp",
	"end_device.pending_mac_state.current_parameters.ping_slot_data_rate_index",
	"end_device.pending_mac_state.current_parameters.ping_slot_data_rate_index_value",
	"end_device.pending_mac_state.current_parameters.ping_slot_data_rate_index_value.value",
	"end_device.pending_mac_state.current_parameters.ping_slot_frequency",
	"end_device.pending_mac_state.current_parameters.rejoin_count_periodicity",
	"end_device.pending_mac_state.current_parameters.rejoin_time_periodicity",
	"end_device.pending_mac_state.current_parameters.rx1_data_rate_offset",
	"end_device.pending_mac_state.current_parameters.rx1_delay",
	"end_device.pending_mac_state.current_parameters.rx2_data_rate_index",
	"end_device.pending_mac_state.current_parameters.rx2_frequency",
	"end_device.pending_mac_state.current_parameters.uplink_dwell_time",
	"end_device.pending_mac_state.current_parameters.uplink_dwell_time.value",
	"end_device.pending_mac_state.desired_parameters",
	"end_device.pending_mac_state.desired_parameters.adr_ack_delay",
	"end_device.pending_mac_state.desired_parameters.adr_ack_delay_exponent",
	"end_device.pending_mac_state.desired_parameters.adr_ack_delay_exponent.value",
	"end_device.pending_mac_state.desired_parameters.adr_ack_limit",
	"end_device.pending_mac_state.desired_parameters.adr_ack_limit_exponent",
	"end_device.pending_mac_state.desired_parameters.adr_ack_limit_exponent.value",
	"end_device.pending_mac_state.desired_parameters.adr_data_rate_index",
	"end_device.pending_mac_state.desired_parameters.adr_nb_trans",
	"end_device.pending_mac_state.desired_parameters.adr_tx_power_index",
	"end_device.pending_mac_state.desired_parameters.beacon_frequency",
	"end_device.pending_mac_state.desired_parameters.channels",
	"end_device.pending_mac_state.desired_parameters.downlink_dwell_time",
	"end_device.pending_mac_state.desired_parameters.downlink_dwell_time.value",
	"end_device.pending_mac_state.desired_parameters.max_duty_cycle",
	"end_device.pending_mac_state.desired_parameters.max_eirp",
	"end_device.pending_mac_state.desired_parameters.ping_slot_data_rate_index",
	"end_device.pending_mac_state.desired_parameters.ping_slot_data_rate_index_value",
	"end_device.pending_mac_state.desired_parameters.ping_slot_data_rate_index_value.value",
	"end_device.pending_mac_state.desired_parameters.ping_slot_frequency",
	"end_device.pending_mac_state.desired_parameters.rejoin_count_periodicity",
	"end_device.pending_mac_state.desired_parameters.rejoin_time_periodicity",
	"end_device.pending_mac_state.desired_parameters.rx1_data_rate_offset",
	"end_device.pending_mac_state.desired_parameters.rx1_delay",
	"end_device.pending_mac_state.desired_parameters.rx2_data_rate_index",
	"end_device.pending_mac_state.desired_parameters.rx2_frequency",
	"end_device.pending_mac_state.desired_parameters.uplink_dwell_time",
	"end_device.pending_mac_state.desired_parameters.uplink_dwell_time.value",
	"end_device.pending_mac_state.device_class",
	"end_device.pending_mac_state.last_adr_change_f_cnt_up",
	"end_device.pending_mac_state.last_confirmed_downlink_at",
	"end_device.pending_mac_state.last_dev_status_f_cnt_up",
	"end_device.pending_mac_state.last_downlink_at",
	"end_device.pending_mac_state.last_network_initiated_downlink_at",
	"end_device.pending_mac_state.lorawan_version",
	"end_device.pending_mac_state.pending_application_downlink",
	"end_device.pending_mac_state.pending_application_downlink.class_b_c",
	"end_device.pending_mac_state.pending_application_downlink.class_b_c.absolute_time",
	"end_device.pending_mac_state.pending_application_downlink.class_b_c.gateways",
	"end_device.pending_mac_state.pending_application_downlink.confirmed",
	"end_device.pending_mac_state.pending_application_downlink.correlation_ids",
	"end_device.pending_mac_state.pending_application_downlink.decoded_payload",
	"end_device.pending_mac_state.pending_application_downlink.decoded_payload_warnings",
	"end_device.pending_mac_state.pending_application_downlink.f_cnt",
	"end_device.pending_mac_state.pending_application_downlink.f_port",
	"end_device.pending_mac_state.pending_application_downlink.frm_payload",
	"end_device.pending_mac_state.pending_application_downlink.priority",
	"end_device.pending_mac_state.pending_application_downlink.session_key_id",
	"end_device.pending_mac_state.pending_join_request",
	"end_device.pending_mac_state.pending_join_request.cf_list",
	"end_device.pending_mac_state.pending_join_request.cf_list.ch_masks",
	"end_device.pending_mac_state.pending_join_request.cf_list.freq",
	"end_device.pending_mac_state.pending_join_request.cf_list.type",
	"end_device.pending_mac_state.pending_join_request.downlink_settings",
	"end_device.pending_mac_state.pending_join_request.downlink_settings.opt_neg",
	"end_device.pending_mac_state.pending_join_request.downlink_settings.rx1_dr_offset",
	"end_device.pending_mac_state.pending_join_request.downlink_settings.rx2_dr",
	"end_device.pending_mac_state.pending_join_request.rx_delay",
	"end_device.pending_mac_state.pending_requests",
	"end_device.pending_mac_state.ping_slot_periodicity",
	"end_device.pending_mac_state.ping_slot_periodicity.value",
	"end_device.pending_mac_state.queued_join_accept",
	"end_device.pending_mac_state.queued_join_accept.correlation_ids",
	"end_device.pending_mac_state.queued_join_accept.dev_addr",
	"end_device.pending_mac_state.queued_join_accept.keys",
	"end_device.pending_mac_state.queued_join_accept.keys.app_s_key",
	"end_device.pending_mac_state.queued_join_accept.keys.app_s_key.encrypted_key",
	"end_device.pending_mac_state.queued_join_accept.keys.app_s_key.kek_label",
	"end_device.pending_mac_state.queued_join_accept.keys.app_s_key.key",
	"end_device.pending_mac_state.queued_join_accept.keys.f_nwk_s_int_key",
	"end_device.pending_mac_state.queued_join_accept.keys.f_nwk_s_int_key.encrypted_key",
	"end_device.pending_mac_state.queued_join_accept.keys.f_nwk_s_int_key.kek_label",
	"end_device.pending_mac_state.queued_join_accept.keys.f_nwk_s_int_key.key",
	"end_device.pending_mac_state.queued_join_accept.keys.nwk_s_enc_key",
	"end_device.pending_mac_state.queued_join_accept.keys.nwk_s_enc_key.encrypted_key",
	"end_device.pending_mac_state.queued_join_accept.keys.nwk_s_enc_key.kek_label",
	"end_device.pending_mac_state.queued_join_accept.keys.nwk_s_enc_key.key",
	"end_device.pending_mac_state.queued_join_accept.keys.s_nwk_s_int_key",
	"end_device.pending_mac_state.queued_join_accept.keys.s_nwk_s_int_key.encrypted_key",
	"end_device.pending_mac_state.queued_join_accept.keys.s_nwk_s_int_key.kek_label",
	"end_device.pending_mac_state.queued_join_accept.keys.s_nwk_s_int_key.key",
	"end_device.pending_mac_state.queued_join_accept.keys.session_key_id",
	"end_device.pending_mac_state.queued_join_accept.net_id",
	"end_device.pending_mac_state.queued_join_accept.payload",
	"end_device.pending_mac_state.queued_join_accept.request",
	"end_device.pending_mac_state.queued_join_accept.request.cf_list",
	"end_device.pending_mac_state.queued_join_accept.request.cf_list.ch_masks",
	"end_device.pending_mac_state.queued_join_accept.request.cf_list.freq",
	"end_device.pending_mac_state.queued_join_accept.request.cf_list.type",
	"end_device.pending_mac_state.queued_join_accept.request.downlink_settings",
	"end_device.pending_mac_state.queued_join_accept.request.downlink_settings.opt_neg",
	"end_device.pending_mac_state.queued_join_accept.request.downlink_settings.rx1_dr_offset",
	"end_device.pending_mac_state.queued_join_accept.request.downlink_settings.rx2_dr",
	"end_device.pending_mac_state.queued_join_accept.request.rx_delay",
	"end_device.pending_mac_state.queued_responses",
	"end_device.pending_mac_state.recent_downlinks",
	"end_device.pending_mac_state.recent_uplinks",
	"end_device.pending_mac_state.rejected_adr_data_rate_indexes",
	"end_device.pending_mac_state.rejected_adr_tx_power_indexes",
	"end_device.pending_mac_state.rejected_data_rate_ranges",
	"end_device.pending_mac_state.rejected_frequencies",
	"end_device.pending_mac_state.rx_windows_available",
	"end_device.pending_session",
	"end_device.pending_session.dev_addr",
	"end_device.pending_session.keys",
	"end_device.pending_session.keys.app_s_key",
	"end_device.pending_session.keys.app_s_key.encrypted_key",
	"end_device.pending_session.keys.app_s_key.kek_label",
	"end_device.pending_session.keys.app_s_key.key",
	"end_device.pending_session.keys.f_nwk_s_int_key",
	"end_device.pending_session.keys.f_nwk_s_int_key.encrypted_key",
	"end_device.pending_session.keys.f_nwk_s_int_key.kek_label",
	"end_device.pending_session.keys.f_nwk_s_int_key.key",
	"end_device.pending_session.keys.nwk_s_enc_key",
	"end_device.pending_session.keys.nwk_s_enc_key.encrypted_key",
	"end_device.pending_session.keys.nwk_s_enc_key.kek_label",
	"end_device.pending_session.keys.nwk_s_enc_key.key",
	"end_device.pending_session.keys.s_nwk_s_int_key",
	"end_device.pending_session.keys.s_nwk_s_int_key.encrypted_key",
	"end_device.pending_session.keys.s_nwk_s_int_key.kek_label",
	"end_device.pending_session.keys.s_nwk_s_int_key.key",
	"end_device.pending_session.keys.session_key_id",
	"end_device.pending_session.last_a_f_cnt_down",
	"end_device.pending_session.last_conf_f_cnt_down",
	"end_device.pending_session.last_f_cnt_up",
	"end_device.pending_session.last_n_f_cnt_down",
	"end_device.pending_session.queued_application_downlinks",
	"end_device.pending_session.started_at",
	"end_device.picture",
	"end_device.picture.embedded",
	"end_device.picture.embedded.data",
	"end_device.picture.embedded.mime_type",
	"end_device.picture.sizes",
	"end_device.power_state",
	"end_device.provisioner_id",
	"end_device.provisioning_data",
	"end_device.queued_application_downlinks",
	"end_device.resets_join_nonces",
	"end_device.root_keys",
	"end_device.root_keys.app_key",
	"end_device.root_keys.app_key.encrypted_key",
	"end_device.root_keys.app_key.kek_label",
	"end_device.root_keys.app_key.key",
	"end_device.root_keys.nwk_key",
	"end_device.root_keys.nwk_key.encrypted_key",
	"end_device.root_keys.nwk_key.kek_label",
	"end_device.root_keys.nwk_key.key",
	"end_device.root_keys.root_key_id",
	"end_device.service_profile_id",
	"end_device.session",
	"end_device.session.dev_addr",
	"end_device.session.keys",
	"end_device.session.keys.app_s_key",
	"end_device.session.keys.app_s_key.encrypted_key",
	"end_device.session.keys.app_s_key.kek_label",
	"end_device.session.keys.app_s_key.key",
	"end_device.session.keys.f_nwk_s_int_key",
	"end_device.session.keys.f_nwk_s_int_key.encrypted_key",
	"end_device.session.keys.f_nwk_s_int_key.kek_label",
	"end_device.session.keys.f_nwk_s_int_key.key",
	"end_device.session.keys.nwk_s_enc_key",
	"end_device.session.keys.nwk_s_enc_key.encrypted_key",
	"end_device.session.keys.nwk_s_enc_key.kek_label",
	"end_device.session.keys.nwk_s_enc_key.key",
	"end_device.session.keys.s_nwk_s_int_key",
	"end_device.session.keys.s_nwk_s_int_key.encrypted_key",
	"end_device.session.keys.s_nwk_s_int_key.kek_label",
	"end_device.session.keys.s_nwk_s_int_key.key",
	"end_device.session.keys.session_key_id",
	"end_device.session.last_a_f_cnt_down",
	"end_device.session.last_conf_f_cnt_down",
	"end_device.session.last_f_cnt_up",
	"end_device.session.last_n_f_cnt_down",
	"end_device.session.queued_application_downlinks",
	"end_device.session.started_at",
	"end_device.skip_payload_crypto",
	"end_device.skip_payload_crypto_override",
	"end_device.supports_class_b",
	"end_device.supports_class_c",
	"end_device.supports_join",
	"end_device.updated_at",
	"end_device.used_dev_nonces",
	"end_device.version_ids",
	"end_device.version_ids.band_id",
	"end_device.version_ids.brand_id",
	"end_device.version_ids.firmware_version",
	"end_device.version_ids.hardware_version",
	"end_device.version_ids.model_id",
	"format_id",
	"image",
	"image.image_size",
}

var GenerateEndDeviceQRCodeRequestFieldPathsTopLevel = []string{
	"end_device",
	"format_id",
	"image",
}
var GenerateQRCodeResponseFieldPathsNested = []string{
	"image",
	"image.embedded",
	"image.embedded.data",
	"image.embedded.mime_type",
	"image.sizes",
	"text",
}

var GenerateQRCodeResponseFieldPathsTopLevel = []string{
	"image",
	"text",
}
var EndDeviceOnboardingDataFieldPathsNested = []string{
	"checksum",
	"claim_authentication_code",
	"dev_eui",
	"join_eui",
	"model_id",
	"proprietary",
	"serial_number",
	"vendor_id",
}

var EndDeviceOnboardingDataFieldPathsTopLevel = []string{
	"checksum",
	"claim_authentication_code",
	"dev_eui",
	"join_eui",
	"model_id",
	"proprietary",
	"serial_number",
	"vendor_id",
}
var EntityOnboardingDataFieldPathsNested = []string{
	"data",
	"data.end_device_onboarding_data",
	"data.end_device_onboarding_data.checksum",
	"data.end_device_onboarding_data.claim_authentication_code",
	"data.end_device_onboarding_data.dev_eui",
	"data.end_device_onboarding_data.join_eui",
	"data.end_device_onboarding_data.model_id",
	"data.end_device_onboarding_data.proprietary",
	"data.end_device_onboarding_data.serial_number",
	"data.end_device_onboarding_data.vendor_id",
	"format_id",
}

var EntityOnboardingDataFieldPathsTopLevel = []string{
	"data",
	"format_id",
}
var ParseQRCodeRequestFieldPathsNested = []string{
	"format_id",
	"qr_code",
}

var ParseQRCodeRequestFieldPathsTopLevel = []string{
	"format_id",
	"qr_code",
}
var ParseQRCodeResponseFieldPathsNested = []string{
	"entity_onboarding_data",
	"entity_onboarding_data.data",
	"entity_onboarding_data.data.end_device_onboarding_data",
	"entity_onboarding_data.data.end_device_onboarding_data.checksum",
	"entity_onboarding_data.data.end_device_onboarding_data.claim_authentication_code",
	"entity_onboarding_data.data.end_device_onboarding_data.dev_eui",
	"entity_onboarding_data.data.end_device_onboarding_data.join_eui",
	"entity_onboarding_data.data.end_device_onboarding_data.model_id",
	"entity_onboarding_data.data.end_device_onboarding_data.proprietary",
	"entity_onboarding_data.data.end_device_onboarding_data.serial_number",
	"entity_onboarding_data.data.end_device_onboarding_data.vendor_id",
	"entity_onboarding_data.format_id",
}

var ParseQRCodeResponseFieldPathsTopLevel = []string{
	"entity_onboarding_data",
}
var GenerateEndDeviceQRCodeRequest_ImageFieldPathsNested = []string{
	"image_size",
}

var GenerateEndDeviceQRCodeRequest_ImageFieldPathsTopLevel = []string{
	"image_size",
}
