package configfile

import (
	"bytes"

	"github.com/BurntSushi/toml"
)

// configfile
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

type ExtraChannels struct {
	Frequency int32 `toml:"frequency"`
	MinDr     int32 `toml:"min_dr"`
	MaxDr     int32 `toml:"max_dr"`
}

type NetworkSettings struct {
	InstallationMargin int32 `toml:"installation_margin"`
	RxWindow           int32 `toml:"rx_window"`
	Rx1Delay           int32 `toml:"rx1_delay"`
	Rx1DROffset        int32 `toml:"rx1_dr_offset"`
	Rx2DR              int32 `toml:"rx2_dr"`
	Rx2Frequency       int32 `toml:"rx2_frequency"`
	DownlinkTxPower    int32 `toml:"downlink_tx_power"`
	DisableADR         bool  `toml:"disable_adr"`

	ExtraChannels         *[]ExtraChannels `toml:"extra_channels,omitempty"`
	EnabledUplinkChannels *[]int32         `toml:"enabled_uplink_channels,omitempty"`
}

type NsNetworkServer struct {
	NetId           string `toml:"net_id"`
	Api             `toml:"api"`
	Band            `toml:"band"`
	Gateway         `toml:"gateway"`
	NetworkSettings `toml:"network_settings"`
}

type Default struct {
	Server string `toml:"server"`
}

type JoinServer struct {
	Default `toml:"default"`
}

type NetworkServer struct {
	Postgresql      `toml:"postgresql"`
	Redis           `toml:"redis"`
	NsNetworkServer `toml:"network_server"`
	JoinServer      `toml:"join_server"`
	File            `toml:"-"`
}

// NSSettings
type NSSettings struct {
	NetID                 string
	BandName              string
	DisableADR            bool
	ADRMargin             int32
	Rx1Delay              int32
	Rx1DROffset           int32
	Rx2Frequency          int32
	Rx2DR                 int32
	DownlinkTXPower       int32
	ExtraChannels         *[]ExtraChannels
	EnabledUplinkChannels *[]int32
	File
}

func NewNetworkServer(s *NSSettings) *NetworkServer {
	return &NetworkServer{
		File: s.File,
		Postgresql: Postgresql{
			Dsn: "postgres://chirpstack_ns:dfrobot@localhost/chirpstack_ns?sslmode=disable",
		},
		Redis: Redis{
			Url: "redis://localhost:6379",
		},
		NsNetworkServer: NsNetworkServer{
			NetId: s.NetID,
			Api: Api{
				Bind: "0.0.0.0:8000",
			},
			Band: Band{
				Name: s.BandName,
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
			NetworkSettings: NetworkSettings{
				InstallationMargin: s.ADRMargin,
				// RxWindow          :,
				Rx1Delay:        s.Rx1Delay,
				Rx1DROffset:     s.Rx1DROffset,
				Rx2DR:           s.Rx2DR,
				Rx2Frequency:    s.Rx2Frequency,
				DownlinkTxPower: s.DownlinkTXPower,
				DisableADR:      s.DisableADR,

				ExtraChannels:         s.ExtraChannels,
				EnabledUplinkChannels: s.EnabledUplinkChannels,
			},
		},
		JoinServer: JoinServer{
			Default: Default{
				Server: "http://localhost:8003",
			},
		},
	}
}

func (c *NetworkServer) Marshal() ([]byte, error) {
	var buf bytes.Buffer
	encoder := toml.NewEncoder(&buf)
	err := encoder.Encode(c)
	return buf.Bytes(), err
}

func (c *NetworkServer) GetFile() string {
	return c.File.String()
}

func (c *NetworkServer) IsNil() bool {
	return c == nil
}
