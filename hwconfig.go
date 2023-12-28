package hwconfig

import "errors"

type Configs struct {
	PacketForwarder *SemtechUdpConfig
	GatewayBridge   *GatewayBridgeConfig
	NetworkServer   *NetworkServerConfig
}

type CustomBand struct {
	CenterFrequency int32
	FrequencyShift  [5]int32
}

type Builder interface {
	// Build create a new config based on region & backend，the config will be nil if no change
	Build() (*Configs, error)

	// SetBackend sets the gateway backend (PacketForwarderBackend/BasicStationBackend), default is PacketForwarder
	SetBackend(b string)

	// SetCustomBand sets the EU868 custom band
	SetCustomBand(c CustomBand)

	// SetSubband sets CN470 & US915's subband, default is 0
	SetSubband(fsb int32)
}

type Marshaler interface {
	// Backend configuration marshal
	Marshal() ([]byte, error)

	// Check if the struct value is nil
	IsNil() bool
}

const (
	SemtechUDP = "semtech_udp"
	BStation   = "basic_station"

	CN470 = "CN470"
	EU868 = "EU868"
	US915 = "US915"
)

var (
	backendList = []string{
		SemtechUDP,
		BStation,
	}

	bandList = []struct {
		Name   string
		Struct Builder
	}{
		{CN470, &BandCn470{}},
		{EU868, &BandEu868{}},
		{US915, &BandUs915{}},
	}
)

func checkBackend(b string) bool {
	for _, v := range backendList {
		if v == b {
			return true
		}
	}
	return false
}

func NewBuilder(r string) (Builder, error) {
	for _, v := range bandList {
		if v.Name == r {
			return v.Struct, nil
		}
	}
	return nil, errors.New("unsupported region")
}
