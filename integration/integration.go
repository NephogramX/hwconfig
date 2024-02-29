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

	BSPath = "/gw/opt/lora_pkt_fwd/"
	PFPath = "/gw/opt/lora_pkt_fwd/"
	GBPath = "/gw/etc/chirpstack-ttn-gateway/"
	NSPath = "/gw/etc/chirpstack-ttn-gateway/"

	// BSPath = "./build/"
	// PFPath = "./build/"
	// GBPath = "./build/"
	// NSPath = "./build/"

	PFName = "global_conf.json"
	GBName = "chirpstack-gateway-bridge.toml"
	NSName = "chirpstack-network-server.toml"
)

type Integration interface {
	HandleBasicsStationUri() *cf.BasicsStation
	HandleBasicsStationKey() *cf.BasicsStation
	HandleBasicsStationCrt() *cf.BasicsStation
	HandleBasicsStationTrust() *cf.BasicsStation
	HandleUdpPacketForwarder() *cf.UdpPacketForwarder
	HandleGatewayBridge() *cf.GatewayBridge
	HandleNetworkServer() *cf.NetworkServer
}

func ApplySettings(i Integration) error {
	if err := cf.CreateConfigFile(i.HandleBasicsStationUri()); err != nil {
		return err
	}
	if err := cf.CreateConfigFile(i.HandleBasicsStationTrust()); err != nil {
		return err
	}
	if err := cf.CreateConfigFile(i.HandleBasicsStationKey()); err != nil {
		return err
	}
	if err := cf.CreateConfigFile(i.HandleBasicsStationCrt()); err != nil {
		return err
	}
	if err := cf.CreateConfigFile(i.HandleUdpPacketForwarder()); err != nil {
		return err
	}
	if err := cf.CreateConfigFile(i.HandleGatewayBridge()); err != nil {
		return err
	}
	if err := cf.CreateConfigFile(i.HandleNetworkServer()); err != nil {
		return err
	}
	return nil
}
