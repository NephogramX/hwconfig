package integration

import (
	cf "github.com/NephogramX/hwconfig/configfile"
)

type IntegrationType int32
type CommSettings = cf.Comm
type ServerSettings = cf.Server

const (
	Buildin            IntegrationType = iota
	UdpPacketForwarder IntegrationType = iota
	BasicsStation      IntegrationType = iota

	// BSPath = "/gw/opt/lora_pkt_fwd/"
	// PFPath = "/gw/opt/lora_pkt_fwd/"
	// GBPath = "/gw/etc/chirpstack-ttn-gateway/"
	// NSPath = "/gw/etc/chirpstack-ttn-gateway/"

	BSPath = "./build/"
	PFPath = "./build/"
	GBPath = "./build/"
	NSPath = "./build/"

	PFName = "global_conf.json"
	GBName = "chirpstack-gateway-bridge.toml"
	NSName = "chirpstack-network-server.toml"
)

// type Settings struct {
// 	Integration
// }

// func (s *Settings) GobMarshal() ([]byte, error) {
// 	var buf bytes.Buffer
// 	encoder := gob.NewEncoder(&buf)
// 	err := encoder.Encode(s)
// 	return buf.Bytes(), err
// }

// func (s *Settings) GobUnmarshal(b []byte) error {
// 	buf := bytes.NewBuffer(b)
// 	decoder := gob.NewDecoder(buf)
// 	if err := decoder.Decode(s); err != nil {
// 		return err
// 	}
// 	return nil
// }

type Integration interface {
	Type() IntegrationType
	HandleBasicsStationUri() *cf.BasicsStation
	HandleBasicsStationKey() *cf.BasicsStation
	HandleBasicsStationCrt() *cf.BasicsStation
	HandleBasicsStationTrust() *cf.BasicsStation
	HandleUdpPacketForwarder() *cf.UdpPacketForwarder
	HandleGatewayBridge() *cf.GatewayBridge
	HandleNetworkServer() *cf.NetworkServer
}

func ApplySettings(i Integration) error {

	if err := i.HandleBasicsStationUri().Write(); err != nil {
		return err
	}
	if err := i.HandleBasicsStationTrust().Write(); err != nil {
		return err
	}
	if err := i.HandleBasicsStationKey().Write(); err != nil {
		return err
	}
	if err := i.HandleBasicsStationCrt().Write(); err != nil {
		return err
	}
	if err := i.HandleUdpPacketForwarder().Write(); err != nil {
		return err
	}
	if err := i.HandleGatewayBridge().Write(); err != nil {
		return err
	}
	if err := i.HandleNetworkServer().Write(); err != nil {
		return err
	}
	return nil
}

// func GetSettings() Integration {
// 	return &BuildinIntegration{}
// }

// func saveSettings(path string) Integration {

// }

// func loadSettings(path string) Integration {

// }
