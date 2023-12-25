package hwconfig

const (
	PacketForwarderBackend = iota
	BasicStationBackend
)

/* **************************************************
 * PacketForwarder Conifg
 */

type FineTimeStamp struct {
	Enable bool   `json:"enable,omitempty"`
	Mode   string `json:"mode,omitempty"`
}

type RssiTcomp struct {
	CoeffA float32 `json:"coeff_a"`
	CoeffB float32 `json:"coeff_b"`
	CoeffC float32 `json:"coeff_c"`
	CoeffD float32 `json:"coeff_d"`
	CoeffE float32 `json:"coeff_e"`
}

type TxGainLutItem struct {
	RFPower int `json:"rf_power"`
	PaGain  int `json:"pa_gain"`
	PwrIdx  int `json:"pwr_idx"`
}

type ChanMultiSF struct {
	Enable bool  `json:"enable"`
	Radio  int32 `json:"radio"`
	IF     int32 `json:"if"`
}

type ChanMultiSFAll struct {
	SpreadingFactorEnable []int32 `json:"spreading_factor_enable"`
	Radio                 int32   `json:"radio"`
	IF                    int32   `json:"if"`
}
type ChanLoraStd struct {
	Enable                bool  `json:"enable"`
	Radio                 int32 `json:"radio"`
	IF                    int32 `json:"if"`
	Bandwidth             int32 `json:"bandwidth"`
	SpreadFactor          int32 `json:"spread_factor"`
	ImplicitHdr           bool  `json:"implicit_hdr"`
	Implicitpayloadlength int32 `json:"implicit_payload_length"`
	ImplicitcrcEn         bool  `json:"implicit_crc_en"`
	Implicitcoderate      int32 `json:"implicit_coderate"`
}
type ChanLoraFSK struct {
	Enable    bool  `json:"enable"`
	Radio     int32 `json:"radio"`
	IF        int32 `json:"if"`
	Bandwidth int32 `json:"bandwidth"`
	Datarate  int32 `json:"datarate"`
}

type Radio0 struct {
	Enable          bool    `json:"enable"`
	Type            string  `json:"type"`
	SingleInputMode bool    `json:"single_input_mode"`
	Freq            int32   `json:"freq"`
	RssiOffset      float32 `json:"rssi_offset"`
	RssiTcomp       `json:"rssi_tcomp"`
	TxEnable        bool            `json:"tx_enable"`
	TxFreqMin       int32           `json:"tx_freq_min"`
	TxFreqMax       int32           `json:"tx_freq_max"`
	TxGainLut       []TxGainLutItem `json:"tx_gain_lut"`
}

type Radio1 struct {
	Enable          bool    `json:"enable"`
	Type            string  `json:"type"`
	SingleInputMode bool    `json:"single_input_mode"`
	Freq            int32   `json:"freq"`
	RssiOffset      float32 `json:"rssi_offset"`
	RssiTcomp       `json:"rssi_tcomp"`
	TxEnable        bool `json:"tx_enable"`
}

type SX130xConfig struct {
	ComType        string `json:"com_type"`
	ComPath        string `json:"com_path"`
	LorawanPublic  bool   `json:"lorawan_public"`
	ClkSrc         int32  `json:"clksrc"`
	AntennaGain    int32  `json:"antenna_gain"`
	FullDuplex     bool   `json:"full_duplex"`
	FineTimeStamp  `json:"fine_timestamp"`
	Radio0         `json:"radio_0"`
	Radio1         `json:"radio_1"`
	ChanMultiSFAll ChanMultiSFAll `json:"chan_multiSF_All"`
	ChanMultiSF0   ChanMultiSF    `json:"chan_multiSF_0"`
	ChanMultiSF1   ChanMultiSF    `json:"chan_multiSF_1"`
	ChanMultiSF2   ChanMultiSF    `json:"chan_multiSF_2"`
	ChanMultiSF3   ChanMultiSF    `json:"chan_multiSF_3"`
	ChanMultiSF4   ChanMultiSF    `json:"chan_multiSF_4"`
	ChanMultiSF5   ChanMultiSF    `json:"chan_multiSF_5"`
	ChanMultiSF6   ChanMultiSF    `json:"chan_multiSF_6"`
	ChanMultiSF7   ChanMultiSF    `json:"chan_multiSF_7"`
	ChanLoraStd    `json:"chan_Lora_std"`

	ChanLoraFSK `json:"chan_FSK"`
}

type GateWayConfig struct {
	GatewayID          string  `json:"gateway_ID"`
	ServerAddress      string  `json:"server_address"`
	ServPortUp         int32   `json:"serv_port_up"`
	ServPortDown       int32   `json:"serv_port_down"`
	KeepaliveInterval  int32   `json:"keepalive_interval"`
	StatInterval       int32   `json:"stat_interval"`
	PushTimeoutMs      int32   `json:"push_timeout_ms"`
	ForwardCrcValid    bool    `json:"forward_crc_valid"`
	ForwardCrcError    bool    `json:"forward_crc_error"`
	ForwardCrcDisabled bool    `json:"forward_crc_disabled"`
	GPSTTYPath         string  `json:"gps_tty_path"`
	RefLatitude        float32 `json:"ref_latitude"`
	RefLongitude       float32 `json:"ref_longitude"`
	RefAltitude        int32   `json:"ref_altitude"`
	AutoquitThreshold  int32   `json:"autoquit_threshold"`
	BeaconPeriod       int32   `json:"beacon_period"`
	BeaconFreqHZ       int32   `json:"beacon_freq_hz"`
	BeaconFreqNB       int32   `json:"beacon_freq_nb"`
	BeaconFreqStep     int32   `json:"beacon_freq_step"`
	BeaconDatarate     int32   `json:"beacon_datarate"`
	BeaconBwHZ         int32   `json:"beacon_bw_hz"`
	BeaconPower        int32   `json:"beacon_power"`
	BeaconInfodesc     int32   `json:"beacon_infodesc"`
}

type RefPayloadItem struct {
	ID string `json:"id"`
}

type DebugConf struct {
	RefPayload []RefPayloadItem `json:"ref_payload"`
	LogFile    string           `json:"log_file"`
}

type PacketForwarderConfig struct {
	SX130xConfig  `json:"SX130x_conf"`
	GateWayConfig `json:"gateway_conf"`
	DebugConf     `json:"debug_conf"`
}

/* **************************************************
 * GatewayBridge Conifg
 */

type SemtechUdp struct {
	UdpBind string `toml:"udp_bind"`
}

type BasicStation struct {
	Bind string `toml:"bind"`
}

type GbBackend struct {
	Type         string `toml:"type"`
	SemtechUdp   `toml:"semtech_udp"`
	BasicStation `toml:"basicStation"`
}

type Generic struct {
	Server string `toml:"server"`
}

type Auth struct {
	Type    string `toml:"type"`
	Generic `toml:"generic"`
}

type GbMqtt struct {
	EventTopicTemplate   string `toml:"event_topic_template"`
	CommandTopicTemplate string `toml:"command_topic_template"`
	Auth                 `toml:"auth"`
}

type Intergration struct {
	Marshaler string `toml:"marshaler"`
	GbMqtt    `toml:"gb_mqtt"`
}

type GatewayBridgeConfig struct {
	GbBackend    `toml:"backend"`
	Intergration `toml:"intergration"`
}

/* **************************************************
 * NetworkServer Conifg
 */

type Postgresql struct {
	Dsn string `toml:"dsn"`
}

type Redis struct {
	Url string `toml:"url"`
}

type Api struct {
	Bind string `toml:"bind"`
}

type NsMqtt struct {
	CommandTopicTemplate string `toml:"command_topic_template"`
	EventTopic           string `toml:"event_topic"`
	Server               string `toml:"server"`
}

type Band struct {
	Name string `toml:"name"`
}

type NsBackend struct {
	Type   string `toml:"type"`
	NsMqtt `toml:"mqtt"`
}

type Gateway struct {
	NsBackend `toml:"backend"`
}

type NetworkServer struct {
	NetId   string `toml:"net_id"`
	Api     `toml:"api"`
	Band    `toml:"band"`
	Gateway `toml:"gateway"`
}

type NetworkServerConfig struct {
	Postgresql    `toml:"postgresql"`
	Redis         `toml:"redis"`
	NetworkServer `toml:"NetworkServer"`
}

/* **************************************************
 * Region Interface
 */

type Configs struct {
	PacketForwarder *PacketForwarderConfig
	GatewayBridge   *GatewayBridgeConfig
	NetworkServer   *NetworkServerConfig
}

type CustomBand struct {
	CenterFrequency int
	FrequencyShift  [5]int
}

type Builder interface {
	// Build create a new config based on region & backend，the config will be nil if no change
	Build() (*Configs, error)

	// SetBackend sets the gateway backend (PacketForwarderBackend/BasicStationBackend), default is PacketForwarder
	SetBackend(b int)

	// SetCustomBand sets the EU868 custom band
	SetCustomBand(c CustomBand)

	// SetSubband sets CN470 & US915's subband, default is 0
	SetSubband(fsb int)
}

func FillPacketForwarder(pfc *PacketForwarderConfig) *PacketForwarderConfig {
	return &PacketForwarderConfig{
		SX130xConfig: SX130xConfig{
			ComType:       "SPI",
			ComPath:       "/dev/spidev2.0",
			LorawanPublic: true,
			ClkSrc:        0,
			AntennaGain:   0,
			FullDuplex:    false,
			FineTimeStamp: FineTimeStamp{
				Enable: false,
				Mode:   "all_sf",
			},
			Radio0: Radio0{
				Enable:     true,
				Type:       "SX1250",
				Freq:       pfc.Radio0.Freq,
				RssiOffset: -215.4,
				RssiTcomp: RssiTcomp{
					CoeffA: 0,
					CoeffB: 0,
					CoeffC: 20.41,
					CoeffD: 2162.56,
					CoeffE: 0,
				},
				TxEnable:  true,
				TxFreqMin: pfc.Radio0.TxFreqMin,
				TxFreqMax: pfc.Radio0.TxFreqMin,
				TxGainLut: pfc.Radio0.TxGainLut,
			},
			Radio1: Radio1{
				Enable:     true,
				Type:       "SX1250",
				Freq:       pfc.Radio1.Freq,
				RssiOffset: -215.4,
				RssiTcomp: RssiTcomp{
					CoeffA: 0,
					CoeffB: 0,
					CoeffC: 20.41,
					CoeffD: 2162.56,
					CoeffE: 0,
				},
				TxEnable: false,
			},
			ChanMultiSFAll: ChanMultiSFAll{
				SpreadingFactorEnable: []int32{5, 6, 7, 8, 9, 10, 11, 12},
			},
			ChanMultiSF0: pfc.ChanMultiSF0,
			ChanMultiSF1: pfc.ChanMultiSF1,
			ChanMultiSF2: pfc.ChanMultiSF2,
			ChanMultiSF3: pfc.ChanMultiSF3,
			ChanMultiSF4: pfc.ChanMultiSF4,
			ChanMultiSF5: pfc.ChanMultiSF5,
			ChanMultiSF6: pfc.ChanMultiSF6,
			ChanMultiSF7: pfc.ChanMultiSF7,
			ChanLoraStd:  pfc.ChanLoraStd,
			ChanLoraFSK:  pfc.ChanLoraFSK,
		},
		GateWayConfig: GateWayConfig{
			GatewayID:          "0000000000000000",
			ServerAddress:      "localhost",
			ServPortUp:         1700,
			ServPortDown:       1700,
			KeepaliveInterval:  10,
			StatInterval:       30,
			PushTimeoutMs:      100,
			ForwardCrcValid:    true,
			ForwardCrcError:    false,
			ForwardCrcDisabled: false,
			RefLatitude:        0.0,
			RefLongitude:       0.0,
			RefAltitude:        0,
			BeaconPeriod:       pfc.BeaconPeriod,
			BeaconFreqHZ:       pfc.BeaconFreqHZ,
			BeaconFreqNB:       pfc.BeaconFreqNB,
			BeaconFreqStep:     pfc.BeaconFreqStep,
			BeaconDatarate:     pfc.BeaconDatarate,
			BeaconBwHZ:         pfc.BeaconBwHZ,
			BeaconPower:        pfc.BeaconPower,
			BeaconInfodesc:     pfc.BeaconInfodesc,
		},
		DebugConf: DebugConf{
			RefPayload: []RefPayloadItem{{ID: "0xCAFE1234"}, {ID: "0xCAFE2345"}},
			LogFile:    "loragw_hal.log",
		},
	}
}

func NewGatewayBridge(b *GbBackend) *GatewayBridgeConfig {
	return &GatewayBridgeConfig{
		GbBackend: *b,
		Intergration: Intergration{
			Marshaler: "protobuf",
			GbMqtt: GbMqtt{
				EventTopicTemplate:   "gateway/{{ .GatewayID }}/event/{{ .EventType }}",
				CommandTopicTemplate: "gateway/{{ .GatewayID }}/command/#",
				Auth: Auth{
					Type: "generic",
					Generic: Generic{
						Server: "tcp://127.0.0.1:1883",
					},
				},
			},
		},
	}
}

func NewNetworkServerConfig(region string) *NetworkServerConfig {
	return &NetworkServerConfig{
		Postgresql: Postgresql{
			Dsn: "postgres://chirpstack_ns:dfrobot@localhost/chirpstack_ns?sslmode=disable",
		},
		Redis: Redis{
			Url: "redis://localhost:6379",
		},
		NetworkServer: NetworkServer{
			NetId: "000000",
			Api: Api{
				Bind: "0.0.0.0:8000",
			},
			Band: Band{
				Name: region,
			},
			Gateway: Gateway{
				NsBackend: NsBackend{
					Type: "mqtt",
					NsMqtt: NsMqtt{
						CommandTopicTemplate: "gateway/{{ .GatewayID }}/command/{{ .CommandType }}",
						EventTopic:           "gateway/+/event/+",
						Server:               "tcp://localhost:1883",
					},
				},
			},
		},
	}
}
