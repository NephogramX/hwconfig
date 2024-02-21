package integration

import "github.com/NephogramX/hwconfig/configfile"

type IntegrationType int32

const (
	Buildin            IntegrationType = iota
	UdpPacketForwarder IntegrationType = iota
	BasicsStation      IntegrationType = iota
)

type Integration interface {
	GetType() IntegrationType
	HandleUdpPacketForwarder(configfile.PFSettings) *configfile.UdpPacketForwarder
	HandleGatewayBridge(configfile.GBSettings) *configfile.GatewayBridge
	HandleNetworkServer(configfile.NSSettings) *configfile.NetworkServer
}
