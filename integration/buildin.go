package integration

import (
	"github.com/NephogramX/hwconfig/band"
	cf "github.com/NephogramX/hwconfig/configfile"
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

type BuildinNSSettings struct {
	Band band.Band
	LoRaSettings
	CommSettings
}

type BuildinIntegration struct {
	BuildinNSSettings
}

func NewBuildinIntegration(s *BuildinNSSettings) (*BuildinIntegration, error) {
	return &BuildinIntegration{
		BuildinNSSettings: *s,
	}, nil
}

func (i *BuildinIntegration) HandleBasicsStationUri() *cf.BasicsStation {
	return nil
}

func (i *BuildinIntegration) HandleBasicsStationKey() *cf.BasicsStation {
	return nil
}

func (i *BuildinIntegration) HandleBasicsStationCrt() *cf.BasicsStation {
	return nil
}

func (i *BuildinIntegration) HandleBasicsStationTrust() *cf.BasicsStation {
	return nil
}

func (i *BuildinIntegration) HandleUdpPacketForwarder() *cf.UdpPacketForwarder {
	return cf.NewUdpPacketForwarderCF(&cf.PFSettings{
		Channel: *i.Band.GetChannelSettings(),
		Server: cf.Server{
			Address:  "localhost",
			PortUp:   1700,
			PortDown: 1700,
		},
		Comm: i.CommSettings,
		File: cf.File{
			Name: PFName,
			Path: PFPath,
		},
	})
}

func (i *BuildinIntegration) HandleGatewayBridge() *cf.GatewayBridge {
	return cf.NewGatewayBridge(&cf.GBSettings{
		Backend: cf.Backend{
			Type: cf.SemtechUDP,
			SemtechUdp: &cf.SemtechUdp{
				UdpBind: "0.0.0.0:1700",
			},
			BasicStation: nil,
		},
		File: cf.File{
			Name: GBName,
			Path: GBPath,
		},
	})
}

func (i *BuildinIntegration) HandleNetworkServer() *cf.NetworkServer {
	return cf.NewNetworkServer(&cf.NSSettings{
		NetID:                 i.NetID,
		BandName:              i.Band.String(),
		DisableADR:            !i.Enable,
		ADRMargin:             i.AdrMargin,
		Rx1Delay:              i.Rx1Delay,
		Rx1DROffset:           i.Rx1DROffset,
		Rx2Frequency:          i.Rx2Frequency,
		Rx2DR:                 i.Rx2DR,
		DownlinkTXPower:       i.DownlinkTXPower,
		ExtraChannels:         i.Band.GetExtraChannels(),
		EnabledUplinkChannels: i.Band.GetUplinkChannels(),
		File: cf.File{
			Name: NSName,
			Path: NSPath,
		},
	})
}
