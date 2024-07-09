package conf

const Eu868NsDefault string = `
[postgresql]
dsn = "postgres://chirpstack_ns:dfrobot@localhost/chirpstack_ns?sslmode=disable"

[redis]
url = "redis://localhost:6379"

[network_server]
net_id = "000000"

[network_server.api]
bind = "0.0.0.0:8000"

[network_server.band]
name = "EU868"

[network_server.gateway]
[network_server.gateway.backend]
type = "mqtt"

[network_server.gateway.backend.mqtt]
command_topic_template = "gateway/{{ .GatewayID }}/command/{{ .CommandType }}"
event_topic = "gateway/+/event/+"
server = "tcp://localhost:1884"

[network_server.network_settings]
installation_margin = 10
rx_window = 0
rx1_delay = 1
rx1_dr_offset = 0
rx2_dr = 0
rx2_frequency = 869525000
downlink_tx_power = -1
disable_adr = false

[[network_server.network_settings.extra_channels]]
frequency = 867100000
min_dr = 0
max_dr = 5

[[network_server.network_settings.extra_channels]]
frequency = 867300000
min_dr = 0
max_dr = 5

[[network_server.network_settings.extra_channels]]
frequency = 867500000
min_dr = 0
max_dr = 5

[[network_server.network_settings.extra_channels]]
frequency = 867700000
min_dr = 0
max_dr = 5

[[network_server.network_settings.extra_channels]]
frequency = 867900000
min_dr = 0
max_dr = 5 

[join_server]
[join_server.default]
server = "http://localhost:8003"
`

const Eu868PfDefault = `
{
    "SX130x_conf": {
        "com_type": "SPI",
        "com_path": "/dev/spidev2.0",
        "lorawan_public": true,
        "clksrc": 0,
        "antenna_gain": 0,
        "full_duplex": false,
        "fine_timestamp": {
            "enable": false,
            "mode": "all_sf"
        },
        "radio_0": {
            "enable": true,
            "type": "SX1250",
            "freq": 867500000,
            "rssi_offset": -215.4,
            "rssi_tcomp": {"coeff_a": 0, "coeff_b": 0, "coeff_c": 20.41, "coeff_d": 2162.56, "coeff_e": 0},
            "tx_enable": true,
            "tx_freq_min": 863000000,
            "tx_freq_max": 870000000,
            "tx_gain_lut":[
                {"rf_power": 12, "pa_gain": 1, "pwr_idx": 4},
                {"rf_power": 13, "pa_gain": 1, "pwr_idx": 5},
                {"rf_power": 14, "pa_gain": 1, "pwr_idx": 6},
                {"rf_power": 15, "pa_gain": 1, "pwr_idx": 7},
                {"rf_power": 16, "pa_gain": 1, "pwr_idx": 8},
                {"rf_power": 17, "pa_gain": 1, "pwr_idx": 9},
                {"rf_power": 18, "pa_gain": 1, "pwr_idx": 10},
                {"rf_power": 19, "pa_gain": 1, "pwr_idx": 11},
                {"rf_power": 20, "pa_gain": 1, "pwr_idx": 12},
                {"rf_power": 21, "pa_gain": 1, "pwr_idx": 13},
                {"rf_power": 22, "pa_gain": 1, "pwr_idx": 14},
                {"rf_power": 23, "pa_gain": 1, "pwr_idx": 16},
                {"rf_power": 24, "pa_gain": 1, "pwr_idx": 17},
                {"rf_power": 25, "pa_gain": 1, "pwr_idx": 18},
                {"rf_power": 26, "pa_gain": 1, "pwr_idx": 19},
                {"rf_power": 27, "pa_gain": 1, "pwr_idx": 22}
            ]
        },
        "radio_1": {
            "enable": true,
            "type": "SX1250",
            "freq": 868500000,
            "rssi_offset": -215.4,
            "rssi_tcomp": {"coeff_a": 0, "coeff_b": 0, "coeff_c": 20.41, "coeff_d": 2162.56, "coeff_e": 0},
            "tx_enable": false
        },
        "chan_multiSF_All": {"spreading_factor_enable": [ 5, 6, 7, 8, 9, 10, 11, 12 ]},
        "chan_multiSF_0": {"enable": true, "radio": 1, "if": -400000},
        "chan_multiSF_1": {"enable": true, "radio": 1, "if": -200000},
        "chan_multiSF_2": {"enable": true, "radio": 1, "if":  0},
        "chan_multiSF_3": {"enable": true, "radio": 0, "if": -400000},
        "chan_multiSF_4": {"enable": true, "radio": 0, "if": -200000},
        "chan_multiSF_5": {"enable": true, "radio": 0, "if":  0},
        "chan_multiSF_6": {"enable": true, "radio": 0, "if":  200000},
        "chan_multiSF_7": {"enable": true, "radio": 0, "if":  400000},
        "chan_Lora_std":  {"enable": true, "radio": 1, "if": -200000, "bandwidth": 250000, "spread_factor": 7,
                           "implicit_hdr": false, "implicit_payload_length": 17, "implicit_crc_en": false, "implicit_coderate": 1},
        "chan_FSK":       {"enable": true, "radio": 1, "if":  300000, "bandwidth": 125000, "datarate": 50000}
    },
    "gateway_conf": {
        "gateway_ID": "AA55DFDF00000000",
        "server_address": "localhost",
        "serv_port_up": 1700,
        "serv_port_down": 1700,
        "keepalive_interval": 10,
        "stat_interval": 30,
        "push_timeout_ms": 100,
        "forward_crc_valid": true,
        "forward_crc_error": false,
        "forward_crc_disabled": false,
        "ref_latitude": 0.0,
        "ref_longitude": 0.0,
        "ref_altitude": 0,
        "beacon_period": 0,
        "beacon_freq_hz": 869525000, 
        "beacon_freq_nb": 1, 
        "beacon_freq_step": 0, 
        "beacon_datarate": 9, 
        "beacon_bw_hz": 125000, 
        "beacon_power": 27
    },
    "debug_conf": {
        "ref_payload":[
            {"id": "0xCAFE1234"},
            {"id": "0xCAFE2345"}
        ],
        "log_file": "loragw_hal.log"
    }
}
`
