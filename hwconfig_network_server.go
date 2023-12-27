package hwconfig

import (
	"bytes"

	"github.com/BurntSushi/toml"
)

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
	NetworkServer `toml:"network_server"`
}

func (c *NetworkServerConfig) Marshal() ([]byte, error) {
	var buf bytes.Buffer
	encoder := toml.NewEncoder(&buf)
	err := encoder.Encode(c)
	return buf.Bytes(), err
}

func (c *NetworkServerConfig) IsNil() bool {
	return c == nil
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
