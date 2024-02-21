package integration

import (
	"github.com/NephogramX/hwconfig/band"
	"github.com/NephogramX/hwconfig/configfile"
)

type ADRSettings struct {
	Enable    bool
	MinDR     int32
	MaxDR     int32
	AdrMargin int32
}

type LoRaSettings struct {
	ADRSettings
	NetID           string
	Rx1Delay        int32
	Rx1DROffset     int32
	Rx2Frequency    int32
	Rx2DR           int32
	DownlinkTXPower int32
}

type CommSettings = configfile.Comm

type BuildinNSSettings struct {
	Band band.Band
	LoRaSettings
	CommSettings
}

type BuildinIntegration struct {
	BuildinNSSettings
}

func NewBuildinIntegration(s *BuildinNSSettings) *BuildinIntegration {
	return &BuildinIntegration{
		BuildinNSSettings: *s,
	}
}

func (b *BuildinIntegration) GetType() IntegrationType {
	return Buildin
}

func (b *BuildinIntegration) HandleUdpPacketForwarder() *configfile.UdpPacketForwarder {
	return configfile.NewUdpPacketForwarderCF(&configfile.PFSettings{
		Channel: *b.Band.GetChannelSettings(),
		Server: configfile.Server{
			Address:  "localhost",
			PortUp:   1700,
			PortDown: 1700,
		},
		Comm: b.CommSettings,
	})
}

func (b *BuildinIntegration) HandleGatewayBridge() *configfile.GatewayBridge {
	return configfile.NewGatewayBridge(&configfile.GBSettings{
		Backend: configfile.Backend{
			Type: configfile.SemtechUDP,
			SemtechUdp: &configfile.SemtechUdp{
				UdpBind: "0.0.0.0:1700",
			},
			BasicStation: nil,
		},
	})
}

func (b *BuildinIntegration) HandleNetworkServer() *configfile.NetworkServer {
	return configfile.NewNetworkServer(&configfile.NSSettings{
		NetID:                 b.NetID,
		BandName:              b.Band.String(),
		DisableADR:            b.Enable,
		ADRMargin:             b.AdrMargin,
		Rx1Delay:              b.Rx1Delay,
		Rx1DROffset:           b.Rx1DROffset,
		Rx2Frequency:          b.Rx2Frequency,
		Rx2DR:                 b.Rx2DR,
		DownlinkTXPower:       b.DownlinkTXPower,
		ExtraChannels:         b.Band.GetExtraChannels(),
		EnabledUplinkChannels: b.Band.GetUplinkChannels(),
	})
}
