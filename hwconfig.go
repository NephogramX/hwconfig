package hwconfig

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/NephogramX/hwconfig/band"
	"github.com/NephogramX/hwconfig/configfile/gb"
	"github.com/NephogramX/hwconfig/configfile/ns"
	"github.com/NephogramX/hwconfig/configfile/pf"
	"github.com/NephogramX/hwconfig/integration"
)

type Region = band.Region
type Integration = integration.Integration

type Backend int32

const (
	UdpPacketForwarder Backend = iota
	BasicStation       Backend = iota
)

type Builder struct {
	band        band.Band
	region      Region
	backend     Backend
	integration integration.Integration
}

type ConfigFile struct {
	PacketForwarder *pf.UdpPacketForwarder
	GatewayBridge   *gb.GatewayBridge
	NetworkServer   *ns.NetworkServer
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) SetRegion(r band.Region) error {
	var err error = nil
	switch r {
	case band.EU868:
	case band.US915:
	case band.CN470:
	default:
		err = errors.New(fmt.Sprint("unsupported region:", r))
	}
	b.region = r
	return err
}

func (b *Builder) SetSubband(subbandIndex int32) error {
	var err error = nil
	switch b.region {
	case band.EU868:
	case band.US915:
	case band.CN470:
		b.band, err = band.NewBandCN470(subbandIndex)
	default:
		err = errors.New(fmt.Sprint("unsupported region:", b.region))
	}
	return err
}

func (b *Builder) SetCustomBand(centerFrequency int32, frequencyShift [5]int32) error {
	return nil
}

func (b *Builder) SetBackend(be Backend) error {
	switch be {
	case UdpPacketForwarder:
		b.backend = be
	case BasicStation:
		b.backend = be
	default:
		return errors.New(fmt.Sprint("unsupported backend:", be))
	}
	return nil
}

func (b *Builder) SetIntegration(i integration.Integration) error {
	b.integration = i
	return nil
}

func (b *Builder) Build() (*ConfigFile, error) {
	var cf ConfigFile

	if b.backend == UdpPacketForwarder {
		cf.PacketForwarder = pf.NewUdpPacketForwarder()
		if err := b.band.HandleUdpPacketForwarder(cf.PacketForwarder); err != nil {
			return nil, err
		}
		if err := b.integration.HandleUdpPacketForwarder(cf.PacketForwarder); err != nil {
			return nil, err
		}
	} else if b.backend == BasicStation {
		if err := b.integration.HandleBasicStation(); err != nil {
			return nil, err
		}
	}

	if b.integration.Get() == integration.Buildin {
		cf.GatewayBridge = gb.NewGatewayBridge()
		if err := b.band.HandleGatewayBridge(cf.GatewayBridge); err != nil {
			return nil, err
		}

		cf.NetworkServer = ns.NewNetworkServer()
		if err := b.band.HandleNetworkServer(cf.NetworkServer); err != nil {
			return nil, err
		}
	}

	return &cf, nil
}
