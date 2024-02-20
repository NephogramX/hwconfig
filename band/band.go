package band

import (
	"github.com/NephogramX/hwconfig/configfile/gb"
	"github.com/NephogramX/hwconfig/configfile/ns"
	"github.com/NephogramX/hwconfig/configfile/pf"
)

type Region int32

const (
	EU868 Region = iota
	US915 Region = iota
	CN470 Region = iota
)

type BandSettings struct {
	Region          Region
	SubbandIndex    *int32
	CenterFrequency *int32
	FrequencyShift  *[5]int32
}

type Band interface {
	HandleUdpPacketForwarder(cf *pf.UdpPacketForwarder) error
	HandleGatewayBridge(cf *gb.GatewayBridge) error
	HandleNetworkServer(cf *ns.NetworkServer) error
}

// func NewBand(r Region) Band {

// }
