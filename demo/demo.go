package main

import (
	"encoding/json"
	"fmt"

	"gitee.com/arya123/chirpstack-api/go/as/external/api"
	"github.com/NephogramX/hwconfig"
)

func main() {
	setup()
	get()

	// build("BS", EU868)
}

func setup() {
	err := hwconfig.SetupDebug()
	if err != nil {
		panic(err)
	}
}

func build(m string, r *api.GateWayRegion) {
	switch m {
	case "NS":
		_, err := hwconfig.Set(&api.ConfigGateWayModeRegionRequest{
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
							DrOffset: 0,
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
			Region: r,
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
	case "PF":
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
			Region: r,
		})
		if err != nil {
			panic(err)
		}
	case "BS":
		_, err := hwconfig.Set(&api.ConfigGateWayModeRegionRequest{
			Mode: &api.GateWayMode{
				Mode: "BS",
				ModeConfig: &api.GateWayMode_Bs{
					Bs: &api.BasicsStation{
						Type: "CUPS",
						// Server: "ALYE5ZL8JJGTT.lns.lorawan.eu-central-1.amazonaws.com",
						Server: "au1.cloud.thethings.network",
						Port:   443,
						// Auth: &api.BSAuth{
						// 	Mode:    "TLS_Server_Client",
						// 	CaCert:  string(LnsAWSCert.Pem),
						// 	CliCert: string(LnsAWSCert.Crt),
						// 	CliKey:  string(LnsAWSCert.Key),
						// },
						Auth: &api.BSAuth{
							Mode:   "TLS_Server_Client_Token",
							CaCert: string(CupsTTNCert.Pem),
							CliKey: string(CupsTTNCert.Key),
						},
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
}

func get() {
	c, err := hwconfig.Get(true)
	if err != nil {
		panic(err)
	}
	b, err := json.MarshalIndent(&c, " ", "	  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", string(b))
}

// ERRO[0121] uplink: processing uplink frame error         ctx_id=a4fe4538-37ed-4610-a623-efb0c3b68fe4 error="join-request to join-server error: json marshal error: json: error calling MarshalText for type lorawan.DLSettings: lorawan: max value of RX1DROffset is 7"
