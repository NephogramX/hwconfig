package main

import (
	"encoding/json"
	"fmt"

	"gitee.com/arya123/chirpstack-api/go/as/external/api"
	"github.com/NephogramX/hwconfig"
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
	err := hwconfig.SetupDebug()
	// if err != nil {
	// 	panic(err)
	// }

	// c, err := hwconfig.Get(true)
	// if err != nil {
	// 	panic(err)
	// }
	ns()
	c, err := hwconfig.LoadFromOriginFile()
	if err != nil {
		panic(err)
	}
	b, err := json.MarshalIndent(&c, " ", "	  ")
	fmt.Printf("%+v\n", string(b))

	// bs()
	// bs()
	// ns()
	// pf()
}

func pf() {
	_, err := hwconfig.Set(&api.ConfigGateWayModeRegionRequest{
		Mode: &api.GateWayMode{
			Mode: "PF",
			ModeConfig: &api.GateWayMode_Pf{
				Pf: &api.PacketForwarder{
					Protocol: &api.PFProtocol{
						Settings: &api.PFProtocol_Gwmp{
							Gwmp: &api.GWMPSSettings{
								Port: &api.GWMPPort{
									Uplink:   1700,
									Downlink: 1700,
								},
								Server: "localhost",
							},
						},
					},
				},
			},
		},
		Region: &api.GateWayRegion{
			RegionId: "US915",
			RegionConfig: &api.GateWayRegion_Us915{
				Us915: &api.US915Config{
					SubBandId: 2,
				},
			},
		},
		Filter: &api.Filter{
			WhiteList: &api.WhiteList{
				Enable:  true,
				OuiList: []string{"000001", "000002"},
				JoinList: []*api.JoinEUIs{
					{From: "aaaabbbbccccdddd", To: "aaaabbbbcccceeee"},
					{From: "aaaabbbbccccffff", To: "aaaabbbbcccfffff"},
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}
}

func ns() {
	adr, err := hwconfig.Set(&api.ConfigGateWayModeRegionRequest{
		Mode: &api.GateWayMode{
			Mode: "NS",
			ModeConfig: &api.GateWayMode_Ns{
				Ns: &api.BuiltInNetworkServer{
					NetworkId: "000000",
					Adr: &api.NSADR{
						Enable:  true,
						DrIdMax: 4,
						DrIdMin: 0,
						Margin:  10,
					},
					Rx1: &api.NSRX1{
						DrOffset: -1,
						Delay:    -1,
					},
					Rx2: &api.NSRX2{
						Freq:    -1,
						DrIndex: -1,
					},
					DownlinkTxPower: 23,
				},
			},
		},
		// Region: &api.GateWayRegion{
		// 	RegionId: "US915",
		// 	RegionConfig: &api.GateWayRegion_Us915{
		// 		Us915: &api.US915Config{
		// 			SubBandId: 2,
		// 		},
		// 	},
		// },
		// Region: &api.GateWayRegion{
		// 	RegionId: "CN470",
		// 	RegionConfig: &api.GateWayRegion_Cn470{
		// 		Cn470: &api.CN470Config{
		// 			SubBandId: 2,
		// 		},
		// 	},
		// },
		Region: &api.GateWayRegion{
			RegionId: "EU868",
			RegionConfig: &api.GateWayRegion_Eu868{
				Eu868: &api.EU868Config{
					Radio_1: &api.EU868Radio1{
						Freq: 867500000,
					},
					ChanMultiSF_3: &api.EU868ChannelMultiSF{
						Offset: -400000,
					},
					ChanMultiSF_4: &api.EU868ChannelMultiSF{
						Offset: -200000,
					},
					ChanMultiSF_5: &api.EU868ChannelMultiSF{
						Offset: 0,
					},
					ChanMultiSF_6: &api.EU868ChannelMultiSF{
						Offset: 200000,
					},
					ChanMultiSF_7: &api.EU868ChannelMultiSF{
						Offset: 400000,
					},
				},
			},
		},
		Filter: &api.Filter{
			WhiteList: &api.WhiteList{
				Enable:  true,
				OuiList: []string{"000001", "000002"},
				JoinList: []*api.JoinEUIs{
					{From: "aaaabbbbccccdddd", To: "aaaabbbbcccceeee"},
					{From: "aaaabbbbccccffff", To: "aaaabbbbcccfffff"},
				},
			},
		},
	})
	fmt.Print("----------")
	fmt.Println(adr)
	if err != nil {
		panic(err)
	}
}

func bs() {
	_, err := hwconfig.Set(&api.ConfigGateWayModeRegionRequest{
		Mode: &api.GateWayMode{
			Mode: "BS",
			ModeConfig: &api.GateWayMode_Bs{
				Bs: &api.BasicsStation{
					Type:   "CUPS",
					Server: "ALYE5ZL8JJGTT.cups.lorawan.ap-southeast-2.amazonaws.com",
					Port:   443,
					Auth: &api.BSAuth{
						Mode:    "TLS_Server_Client",
						CaCert:  string(LnsAWSCert.Pem),
						CliCert: string(LnsAWSCert.Crt),
						CliKey:  string(LnsAWSCert.Key),
					},
					// Auth: &api.BSAuth{
					// 	Mode:   "TLS_Server_Client_Token",
					// 	CaCert: string(LnsTTNCert.Pem),
					// 	CliKey: string(LnsTTNCert.Key),
					// 	Cli
					// },
				},
			},
		},
		Region: &api.GateWayRegion{
			RegionId: "EU868",
			RegionConfig: &api.GateWayRegion_Eu868{
				Eu868: &api.EU868Config{
					Radio_1: &api.EU868Radio1{
						Freq: 867500000,
					},
					ChanMultiSF_3: &api.EU868ChannelMultiSF{
						Offset: -400000,
					},
					ChanMultiSF_4: &api.EU868ChannelMultiSF{
						Offset: -200000,
					},
					ChanMultiSF_5: &api.EU868ChannelMultiSF{
						Offset: 0,
					},
					ChanMultiSF_6: &api.EU868ChannelMultiSF{
						Offset: 200000,
					},
					ChanMultiSF_7: &api.EU868ChannelMultiSF{
						Offset: 400000,
					},
				},
			},
		},
		Filter: &api.Filter{
			WhiteList: &api.WhiteList{
				Enable:  true,
				OuiList: []string{"000001", "000002"},
				JoinList: []*api.JoinEUIs{
					{From: "aaaabbbbccccdddd", To: "aaaabbbbcccceeee"},
					{From: "aaaabbbbccccffff", To: "aaaabbbbcccfffff"},
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}
}

/*
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
*/
