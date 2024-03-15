package integration

import (
	"github.com/NephogramX/hwconfig/band"
	cf "github.com/NephogramX/hwconfig/configfile"
)

type PacketForwarderSettings struct {
	Band band.Band
	ServerSettings
	CommSettings
}

type PacketForwarderIntegration struct {
	PacketForwarderSettings
}

func NewPacketForwarderIntegration(s *PacketForwarderSettings) (*PacketForwarderIntegration, error) {
	return &PacketForwarderIntegration{
		PacketForwarderSettings: *s,
	}, nil
}

func (i *PacketForwarderIntegration) Type() IntegrationType {
	return Builtin
}

func (i *PacketForwarderIntegration) HandleBasicsStationUri() *cf.BasicsStation {
	return nil
}
func (i *PacketForwarderIntegration) HandleBasicsStationKey() *cf.BasicsStation {
	return nil
}
func (i *PacketForwarderIntegration) HandleBasicsStationCrt() *cf.BasicsStation {
	return nil
}

func (i *PacketForwarderIntegration) HandleBasicsStationTrust() *cf.BasicsStation {
	return nil
}

func (i *PacketForwarderIntegration) HandleUdpPacketForwarder() *cf.UdpPacketForwarder {
	return cf.NewUdpPacketForwarderCF(&cf.PFSettings{
		Channel: *i.Band.GetChannelSettings(),
		Server:  i.ServerSettings,
		Comm:    i.CommSettings,
		File: cf.File{
			Name: PFName,
			Path: PFPath,
		},
	})
}

func (i *PacketForwarderIntegration) HandleGatewayBridge() *cf.GatewayBridge {
	return nil
}

func (i *PacketForwarderIntegration) HandleNetworkServer() *cf.NetworkServer {
	return nil
}
