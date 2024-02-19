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

type ExtraChannels struct {
	Frequency int32 `toml:"frequency"`
	MinDr     int32 `toml:"min_dr"`
	MaxDr     int32 `toml:"max_dr"`
}

type NetworkSettings struct {
	ExtraChannels         []ExtraChannels `toml:"extra_channels,omitempty"`
	EnabledUplinkChannels []int32         `toml:"enabled_uplink_channels,omitempty"`
}

type NetworkServer struct {
	NetId           string `toml:"net_id"`
	Api             `toml:"api"`
	Band            `toml:"band"`
	Gateway         `toml:"gateway"`
	NetworkSettings NetworkSettings `toml:"network_settings"`
}

type Default struct {
	Server string `toml:"server"`
}

type JoinServer struct {
	Default `toml:"default"`
}

type NetworkServerConfig struct {
	Postgresql    `toml:"postgresql"`
	Redis         `toml:"redis"`
	NetworkServer `toml:"network_server"`
	JoinServer    `toml:"join_server"`
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
