package configfile

import (
	"bytes"
	"errors"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

const (
	SemtechUDP = "semtech_udp"
	BStation   = "basic_station"
)

type SemtechUdp struct {
	UdpBind string `toml:"udp_bind"`
}

type MultiSF struct {
	Frequencies []int32 `toml:"frequencies"`
}

type LoraStd struct {
	Frequency       int32 `toml:"frequency"`
	Bandwidth       int32 `toml:"bandwidth"`
	SpreadingFactor int32 `toml:"spreading_factor"`
}

type Fsk struct {
	Frequency int32 `toml:"frequency"`
}

type Concentrators struct {
	MultiSF `toml:"multi_sf"`
	LoraStd *LoraStd `toml:"lora_std"`
	Fsk     *Fsk     `toml:"fsk"`
}

type GbBasicStation struct {
	Bind          string `toml:"bind"`
	Region        string `toml:"region"`
	FrequencyMin  int32  `toml:"frequency_min"`
	FrequencyMax  int32  `toml:"frequency_max"`
	Concentrators `toml:"concentrators"`
}

type Backend struct {
	Type         string          `toml:"type"`
	SemtechUdp   *SemtechUdp     `toml:"semtech_udp, omitempty"`
	BasicStation *GbBasicStation `toml:"basic_station, omitempty"`
}

type Generic struct {
	Server string `toml:"server"`
}

type Auth struct {
	Type    string `toml:"type"`
	Generic `toml:"generic"`
}

type Mqtt struct {
	EventTopicTemplate   string `toml:"event_topic_template"`
	CommandTopicTemplate string `toml:"command_topic_template"`
	Auth                 `toml:"auth"`
}

type Intergration struct {
	Marshaler string `toml:"marshaler"`
	Mqtt      `toml:"gb_mqtt"`
}

type Filter struct {
	NetIds   []string    `toml:"net_ids"`
	JoinEuis [][2]string `mapstructure:toml:"join_euis"`
}

type GatewayBridge struct {
	File         `toml:"-"`
	Backend      `toml:"backend"`
	Filter       `toml:"filter"`
	Intergration `toml:"intergration"`
}

type GBSettings struct {
	Backend
	Filter
	File
}

func NewGatewayBridge(s *GBSettings) *GatewayBridge {
	return &GatewayBridge{
		File:    s.File,
		Backend: s.Backend,
		Filter:  s.Filter,
		Intergration: Intergration{
			Marshaler: "protobuf",
			Mqtt: Mqtt{
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

func (c *GatewayBridge) Write() error {
	if c == nil {
		return nil
	}
	var buf bytes.Buffer
	encoder := toml.NewEncoder(&buf)
	if err := encoder.Encode(c); err != nil {
		return err
	}
	return writeFile(c.File.String(), buf.Bytes())
}

func (c *GatewayBridge) ReadFrom(p string) error {
	if c == nil {
		return errors.New("nil interface")
	}
	c.File = File{
		Name: filepath.Base(p),
		Path: filepath.Dir(p),
	}
	_, err := toml.DecodeFile(p, c)
	return err
}
