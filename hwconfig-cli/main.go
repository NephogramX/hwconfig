package main

import (
	"fmt"
	"io/ioutil"

	"github.com/NephogramX/hwconfig"
	"github.com/NephogramX/hwconfig/band"
	"github.com/NephogramX/hwconfig/integration"
)

func main() {

	packetForwarderPath := "D:\\Software\\Go\\gopath\\pkg\\mod\\github.com\\!nephogram!x\\hwconfig@v0.0.0-20231225023832-3f182f8142c2\\hwconfig-cli\\build\\global_conf.json"
	//gatewayBridgePath := path + "/chirpstack-gateway-bridge.toml"
	//networkServerPath := path + "/chirpstack-network-server.toml"

	b := hwconfig.NewBuilder()
	if err := b.SetRegion(band.CN470); err != nil {
		panic(err)
	}
	if err := b.SetSubband(2); err != nil {
		panic(err)
	}
	if err := b.SetBackend(hwconfig.UdpPacketForwarder); err != nil {
		panic(err)
	}
	if err := b.SetIntegration(integration.NewBuildinIntegration()); err != nil {
		panic(err)
	}
	c, err := b.Build()
	if err != nil {
		panic(err)
	}
	buf, err := c.PacketForwarder.Marshal()
	if err != nil {
		panic(err)
	}
	fmt.Println(packetForwarderPath)
	if err := ioutil.WriteFile(packetForwarderPath, buf, 0644); err != nil {
		panic(err)
	}
}
