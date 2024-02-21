package configfile

const (
	UdpPacketForwarderCFPath = "/gw/opt/lora_pkt_fwd/global_config"
	GatewayBridgeCFPath      = "/gw/etc/chirpstack-gateway-bridge/chirpstack-gateway-bridge.toml"
	NetworkServerCFPath      = "/gw/etc/chirpstack-network-server/chirpstack-network-server.toml"
)

type Marshaller interface {
	// Backend configuration marshal
	Marshal() ([]byte, error)

	// Check if the struct value is nil
	IsNil() bool
}
