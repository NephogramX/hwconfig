package hwconfig

import (
	"bytes"

	"github.com/BurntSushi/toml"
)

type SemtechUdp struct {
	UdpBind string `toml:"udp_bind"`
}

type MultiSF struct {
	Frequencies []int32 `toml:"frequencies"`
}

type Concentrators struct {
	MultiSF `toml:"multi_sf"`
}

type BasicStation struct {
	Bind          string `toml:"bind"`
	Region        string `toml:"region"`
	FrequencyMin  int32  `toml:"frequency_min"`
	FrequencyMax  int32  `toml:"frequency_max"`
	Concentrators `toml:"concentrators"`
}

type GbBackend struct {
	Type         string        `toml:"type"`
	SemtechUdp   *SemtechUdp   `toml:"semtech_udp, omitempty"`
	BasicStation *BasicStation `toml:"basic_station, omitempty"`
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

func (c *GatewayBridgeConfig) Marshal() ([]byte, error) {
	var buf bytes.Buffer
	encoder := toml.NewEncoder(&buf)
	err := encoder.Encode(c)
	return buf.Bytes(), err
}

func (c *GatewayBridgeConfig) IsNil() bool {
	return c == nil
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
