package integration

import (
	cf "github.com/NephogramX/hwconfig/configfile"
)

type IntegrationType int32
type CommSettings = cf.Comm
type ServerSettings = cf.Server

const (
	Builtin            IntegrationType = iota
	UdpPacketForwarder IntegrationType = iota
	BasicsStation      IntegrationType = iota

	PFName = "global_conf.json"
	GBName = "chirpstack-gateway-bridge.toml"
	NSName = "chirpstack-network-server.toml"
)

var (
	BSPath string = "/gw/opt/lora_pkt_fwd/"
	PFPath string = "/gw/opt/lora_pkt_fwd/"
	GBPath string = "/gw/etc/chirpstack-gateway-bridge/"
	NSPath string = "/gw/etc/chirpstack-network-server/"
)

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
