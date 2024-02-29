package main

import (
	"github.com/NephogramX/hwconfig/band"
	"github.com/NephogramX/hwconfig/integration"
)

/*
 * std. produce:
 *	1. create band			:	band.New{{.region}}(*)
 * 	2. create integration	: 	integration.New{{.IntegrationType}}(*)
 *  3. stop daemon
 *  4. clean old config files
 *  5. apply settings		:	integration.ApplySettings(*)
 */

func main() {
	// b, _ := band.NewBandEU868(867500000, [5]int32{-400000, -200000, 0, 200000, 400000})

	// BSLoriot()
	// PFLoriot(b)

	// b, _ := band.NewBandCN470(11)
	// PFTTN(b)

	// LNSTTN()
	CUPSTTN()

	// buildin(b)

	// PFTencentCloud()

}

func buildin(b band.Band) error {
	i, err := integration.NewBuildinIntegration(&integration.BuildinNSSettings{
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
	if err != nil {
		return err
	}
	return integration.ApplySettings(i)
}

func PFTencentCloud() error {
	b, err := band.NewBandCN470(11)
	if err != nil {
		return err
	}
	i, err := integration.NewPacketForwarderIntegration(&integration.PacketForwarderSettings{
		Band: b,
		ServerSettings: integration.ServerSettings{
			Address:  "loragw.things.qcloud.com",
			PortUp:   1700,
			PortDown: 1700,
		},
		CommSettings: integration.CommSettings{
			PushTimeoutMs:        100,
			StatIntervalSec:      600,
			KeepaliveIntervalSec: 60,
		},
	})
	if err != nil {
		return err
	}
	return integration.ApplySettings(i)
}

func PFTTN(b band.Band) error {
	i, err := integration.NewPacketForwarderIntegration(&integration.PacketForwarderSettings{
		Band: b,
		ServerSettings: integration.ServerSettings{
			Address:  "au1.cloud.thethings.network",
			PortUp:   1700,
			PortDown: 1700,
		},
		CommSettings: integration.CommSettings{
			PushTimeoutMs:        100,
			StatIntervalSec:      600,
			KeepaliveIntervalSec: 60,
		},
	})
	if err != nil {
		return err
	}
	return integration.ApplySettings(i)
}

func PFLoriot(b band.Band) error {
	i, err := integration.NewPacketForwarderIntegration(&integration.PacketForwarderSettings{
		Band: b,
		ServerSettings: integration.ServerSettings{
			Address:  "eu1.loriot.io",
			PortUp:   1780,
			PortDown: 1780,
		},
		CommSettings: integration.CommSettings{
			PushTimeoutMs:        100,
			StatIntervalSec:      30,
			KeepaliveIntervalSec: 10,
		},
	})
	if err != nil {
		return err
	}
	return integration.ApplySettings(i)
}

func LNSTTN() error {
	i, err := integration.NewBasicsStationIntegration(&integration.BasicsStationSetting{
		Protocol:      integration.LNS,
		ServerAddress: "au1.cloud.thethings.network",
		ServerPort:    8887,
		Authentication: integration.Authentication{
			ServerCert: &LNSTTNCert.Pem,
			Key:        &LNSTTNCert.Key,
			ClientCert: nil,
		},
	})
	if err != nil {
		return err
	}
	return integration.ApplySettings(i)
}

func CUPSTTN() error {
	i, err := integration.NewBasicsStationIntegration(&integration.BasicsStationSetting{
		Protocol:      integration.CUPS,
		ServerAddress: "au1.cloud.thethings.network",
		ServerPort:    443,
		Authentication: integration.Authentication{
			ServerCert: &CUPSTTNCert.Pem,
			Key:        &CUPSTTNCert.Key,
			ClientCert: nil,
		},
	})
	if err != nil {
		return err
	}
	return integration.ApplySettings(i)
}

func LNSLoriot() error {
	i, err := integration.NewBasicsStationIntegration(&integration.BasicsStationSetting{
		Protocol:      integration.LNS,
		ServerAddress: "eu1.loriot.io",
		ServerPort:    717,
		Authentication: integration.Authentication{
			ServerCert: &LNSLoriotCert.Pem,
			Key:        &LNSLoriotCert.Key,
			ClientCert: &LNSLoriotCert.Crt,
		},
	})
	if err != nil {
		return err
	}

	return integration.ApplySettings(i)
}
