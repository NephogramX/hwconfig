package builder

import (
	"github.com/NephogramX/hwconfig/band"
	"github.com/NephogramX/hwconfig/configfile"
	"github.com/NephogramX/hwconfig/integration"
)

type Region = band.Region

const (
	EU868 Region = band.EU868
	US915 Region = band.US915
	CN470 Region = band.CN470
)

type IntegrationType = integration.IntegrationType

const (
	Buildin            IntegrationType = integration.Buildin
	UdpPacketForwarder IntegrationType = integration.UdpPacketForwarder
	BasicsStation      IntegrationType = integration.BasicsStation
)

type Builder struct {
	band        band.Band
	integration integration.Integration
}

type ConfigFile struct {
	PacketForwarder *configfile.UdpPacketForwarder
	GatewayBridge   *configfile.GatewayBridge
	NetworkServer   *configfile.NetworkServer
}

func NewBuilder(b band.Band, i integration.Integration) *Builder {
	return &Builder{b, i}
}
