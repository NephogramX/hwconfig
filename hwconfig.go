package main

import (
	"os"

	"github.com/NephogramX/hwconfig/band"
	"github.com/NephogramX/hwconfig/configfile"
	"github.com/NephogramX/hwconfig/integration"
)

func main() {
	i := "EU868"
	switch i {
	case "CN470":
		testCN470()
	case "EU868":
		testEU868()
	case "US915":
		testUS915()
	}
}

func testCN470() {
	b, _ := band.NewBandCN470(2)
	i := integration.NewBuildinIntegration(&integration.BuildinNSSettings{
		Band: b,
		LoRaSettings: integration.LoRaSettings{
			ADRSettings: integration.ADRSettings{
				Enable:    true,
				MinDR:     0,
				MaxDR:     5,
				AdrMargin: 10,
			},
			NetID:           "000000",
			Rx1Delay:        1,
			Rx1DROffset:     0,
			Rx2Frequency:    -1,
			Rx2DR:           -1,
			DownlinkTXPower: 23,
		},
		CommSettings: integration.CommSettings{
			PushTimeoutMs:        100,
			StatIntervalSec:      600,
			KeepaliveIntervalSec: 60,
		},
	})

	generateCF(i.HandleUdpPacketForwarder(), "./build/global_config.json")
	generateCF(i.HandleGatewayBridge(), "./build/chirpstack-gateway-bridge.toml")
	generateCF(i.HandleNetworkServer(), "./build/chirpstack-network-server.toml")
}

func testEU868() {
	b, _ := band.NewBandEU868(86750000, [5]int32{-400000, -200000, 0, 200000, 400000})
	i := integration.NewBuildinIntegration(&integration.BuildinNSSettings{
		Band: b,
		LoRaSettings: integration.LoRaSettings{
			ADRSettings: integration.ADRSettings{
				Enable:    true,
				MinDR:     0,
				MaxDR:     5,
				AdrMargin: 10,
			},
			NetID:           "000000",
			Rx1Delay:        1,
			Rx1DROffset:     0,
			Rx2Frequency:    -1,
			Rx2DR:           -1,
			DownlinkTXPower: 23,
		},
		CommSettings: integration.CommSettings{
			PushTimeoutMs:        100,
			StatIntervalSec:      600,
			KeepaliveIntervalSec: 60,
		},
	})

	generateCF(i.HandleUdpPacketForwarder(), "./build/global_config.json")
	generateCF(i.HandleGatewayBridge(), "./build/chirpstack-gateway-bridge.toml")
	generateCF(i.HandleNetworkServer(), "./build/chirpstack-network-server.toml")
}

func testUS915() {
	b, _ := band.NewBandUS915(2)
	i := integration.NewBuildinIntegration(&integration.BuildinNSSettings{
		Band: b,
		LoRaSettings: integration.LoRaSettings{
			ADRSettings: integration.ADRSettings{
				Enable:    true,
				MinDR:     0,
				MaxDR:     5,
				AdrMargin: 10,
			},
			NetID:           "000000",
			Rx1Delay:        1,
			Rx1DROffset:     0,
			Rx2Frequency:    -1,
			Rx2DR:           -1,
			DownlinkTXPower: 23,
		},
		CommSettings: integration.CommSettings{
			PushTimeoutMs:        100,
			StatIntervalSec:      600,
			KeepaliveIntervalSec: 60,
		},
	})

	generateCF(i.HandleUdpPacketForwarder(), "./build/global_config.json")
	generateCF(i.HandleGatewayBridge(), "./build/chirpstack-gateway-bridge.toml")
	generateCF(i.HandleNetworkServer(), "./build/chirpstack-network-server.toml")
}

func generateCF(c configfile.Marshaller, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	b, err := c.Marshal()
	if err != nil {
		return err
	}
	_, err = file.Write(b)
	return err
}
