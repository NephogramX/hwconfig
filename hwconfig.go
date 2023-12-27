package hwconfig

type Configs struct {
	PacketForwarder *PacketForwarderConfig
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
)

var (
	backendList = []string{
		SemtechUDP,
		BStation,
	}
)

func CheckBackend(b string) bool {
	for _, v := range backendList {
		if v == b {
			return true
		}
	}
	return false
}
