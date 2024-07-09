package conf

import (
	"gitee.com/dfrobotcd/chirpstack-api/go/as/external/api"
	"github.com/NephogramX/hwconfig/internal/band"
)

type PfConfig struct {
	band.SX130xConfig `json:"SX130x_conf" mapstructure:"SX130x_conf"`
	GatewayConfig     `json:"gateway_conf" mapstructure:"gateway_conf"`
	DebugConf         `json:"debug_conf" mapstructure:"debug_conf"`
}

type GatewayConfig struct {
	ServerAddress string `json:"server_address" mapstructure:"server_address"`
	ServPortUp    int32  `json:"serv_port_up" mapstructure:"serv_port_up"`
	ServPortDown  int32  `json:"serv_port_down" mapstructure:"serv_port_down"`

	GatewayID          string  `json:"gateway_ID"`
	KeepaliveInterval  int32   `json:"keepalive_interval"`
	StatInterval       int32   `json:"stat_interval"`
	PushTimeoutMs      int32   `json:"push_timeout_ms"`
	ForwardCrcValid    bool    `json:"forward_crc_valid"`
	ForwardCrcError    bool    `json:"forward_crc_error"`
	ForwardCrcDisabled bool    `json:"forward_crc_disabled"`
	GPSTTYPath         string  `json:"gps_tty_path,omitempty"`
	RefLatitude        float32 `json:"ref_latitude"`
	RefLongitude       float32 `json:"ref_longitude"`
	RefAltitude        int32   `json:"ref_altitude"`
	AutoquitThreshold  int32   `json:"autoquit_threshold,omitempty"`
	BeaconPeriod       int32   `json:"beacon_period"`
	BeaconFreqHZ       int32   `json:"beacon_freq_hz"`
	BeaconFreqNB       int32   `json:"beacon_freq_nb"`
	BeaconFreqStep     int32   `json:"beacon_freq_step"`
	BeaconDatarate     int32   `json:"beacon_datarate"`
	BeaconBwHZ         int32   `json:"beacon_bw_hz"`
	BeaconPower        int32   `json:"beacon_power"`
	BeaconInfodesc     int32   `json:"beacon_infodesc,omitempty"`
}

type DebugConf struct {
	RefPayload []struct {
		Id string `json:"id"`
	} `json:"ref_payload"`

	LogFile string `json:"log_file"`
}

func (c *PfConfig) ApiMode() *api.GateWayMode {
	return &api.GateWayMode{
		Mode: "PF",
		ModeConfig: &api.GateWayMode_Pf{
			Pf: &api.PacketForwarder{
				Protocol: &api.PFProtocol{
					Settings: &api.PFProtocol_Gwmp{
						Gwmp: &api.GWMPSSettings{
							Port: &api.GWMPPort{
								Uplink:   c.ServPortUp,
								Downlink: c.ServPortDown,
							},
							Server: c.ServerAddress,
							// not used
							PushTimeout:          0,
							AutoRestartThreshold: 0,
							KeepaliveInterval:    0,
						},
					},
					// not used
					Type:          "GWMP",
					StatsInterval: 0,
				},
			},
		},
	}
}
