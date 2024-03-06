package hwconfig

import (
	"errors"
	"fmt"
	"os"

	cf "github.com/NephogramX/hwconfig/configfile"

	"gitee.com/arya123/chirpstack-api/go/as/external/api"
	"github.com/NephogramX/hwconfig/band"
	"github.com/NephogramX/hwconfig/integration"
	"github.com/golang/protobuf/proto"
)

const (
	HWName = ".hwconfig"
)

var (
	HWPath string = "/gw/config/hwconfig/"
)

var hwconfig api.GetGateWayModeRegionResponse

func Setup() error {
	if _, err := os.Stat(HWPath + HWName); os.IsNotExist(err) {
		hwconfig, err := LoadFromOriginFile()
		if err != nil {
			return err
		}
		c := api.ConfigGateWayModeRegionRequest{
			Mode:   hwconfig.Mode,
			Region: hwconfig.Region,
			Filter: hwconfig.Filter,
		}
		Save(&c)
	} else if err != nil {
		return err
	}
	return Load(&hwconfig)
}

func SetupDebug() error {
	HWPath = "./build/"
	integration.BSPath = "./build/"
	integration.PFPath = "./build/"
	integration.GBPath = "./build/"
	integration.NSPath = "./build/"
	if _, err := os.Stat(HWPath + HWName); os.IsNotExist(err) {
		hwconfig, err := LoadFromOriginFile()
		if err != nil {
			return err
		}
		c := api.ConfigGateWayModeRegionRequest{
			Mode:   hwconfig.Mode,
			Region: hwconfig.Region,
			Filter: hwconfig.Filter,
		}
		Save(&c)
	} else if err != nil {
		return err
	}
	return Load(&hwconfig)
}

func Set(c *api.ConfigGateWayModeRegionRequest) ([2]int32, error) {
	// gateway region
	var (
		b   band.Band
		i   integration.Integration
		err error
		adr [2]int32 = [2]int32{-1, -1}
	)

	switch c.GetRegion().GetRegionId() {
	case "CN470":
		r := c.GetRegion().GetCn470()
		if r == nil {
			return adr, errors.New("fail to parse region parameters")
		}
		b, err = band.NewBandCN470(r.GetSubBandId())
	case "EU868":
		r := c.GetRegion().GetEu868()
		if r == nil {
			return adr, errors.New("fail to parse region parameters")
		}
		b, err = band.NewBandEU868(r.GetRadio_0().GetFreq(), []int32{
			r.GetChanMultiSF_3().GetOffset(),
			r.GetChanMultiSF_4().GetOffset(),
			r.GetChanMultiSF_5().GetOffset(),
			r.GetChanMultiSF_6().GetOffset(),
			r.GetChanMultiSF_7().GetOffset(),
		})
	case "US915":
		r := c.GetRegion().GetUs915()
		if r == nil {
			return adr, errors.New("fail to parse region parameters")
		}
		b, err = band.NewBandUS915(r.GetSubBandId())
	default:
		err = fmt.Errorf("unsupport region: %v", c.Region.RegionId)
	}

	if err != nil {
		return adr, err
	}

	// gateway mode
	switch c.GetMode().GetMode() {
	case "NS":
		m := c.GetMode().GetNs()
		ni, je := []string{}, [][2]string{}
		if m == nil {
			return adr, errors.New("fail to parse mode parameters")
		}
		if c.GetFilter().GetWhiteList().GetEnable() {
			ni = c.GetFilter().GetWhiteList().GetOuiList()
			for _, jl := range c.GetFilter().GetWhiteList().GetJoinList() {
				je = append(je, [2]string{jl.GetFrom(), jl.GetTo()})
			}
		}
		if m.GetAdr().GetEnable() && (hwconfig.GetMode().GetMode() != "NS" ||
			m.GetAdr().GetDrIdMax() != hwconfig.GetMode().GetNs().GetAdr().GetDrIdMax() ||
			m.GetAdr().GetDrIdMin() != hwconfig.GetMode().GetNs().GetAdr().GetDrIdMin()) {
			adr = [2]int32{m.GetAdr().GetDrIdMin(), m.GetAdr().GetDrIdMax()}
		}

		i, err = integration.NewBuildinIntegration(&integration.BuildinNSSettings{
			Band: b,
			LoRaSettings: integration.LoRaSettings{
				ADRSettings: integration.ADRSettings{
					Enable:    m.GetAdr().GetEnable(),
					AdrMargin: m.GetAdr().GetMargin(),
				},
				FilterSettings: integration.FilterSettings{
					NetIds:   ni,
					JoinEuis: je,
				},
				NetID:           m.GetNetworkId(),
				Rx1Delay:        m.GetRx1().GetDelay(),
				Rx1DROffset:     m.GetRx1().GetDrOffset(),
				Rx2Frequency:    m.GetRx2().GetFreq(),
				Rx2DR:           m.GetRx2().GetDrIndex(),
				DownlinkTXPower: m.GetDownlinkTxPower(),
			},
			CommSettings: integration.CommSettings{
				PushTimeoutMs:        100,
				StatIntervalSec:      30,
				KeepaliveIntervalSec: 10,
			},
		})
	case "PF":
		m := c.GetMode().GetPf()
		if m == nil {
			return adr, errors.New("fail to parse mode parameters")
		}
		i, err = integration.NewPacketForwarderIntegration(&integration.PacketForwarderSettings{
			Band: b,
			ServerSettings: integration.ServerSettings{
				Address:  m.GetProtocol().GetGwmp().GetServer(),
				PortUp:   m.GetProtocol().GetGwmp().GetPort().GetUplink(),
				PortDown: m.GetProtocol().GetGwmp().GetPort().GetDownlink(),
			},
			CommSettings: integration.CommSettings{
				PushTimeoutMs:        100,
				StatIntervalSec:      30,
				KeepaliveIntervalSec: 10,
			},
		})
	case "BS":
		m := c.GetMode().GetBs()
		if m == nil {
			return adr, errors.New("fail to parse mode parameters")
		}

		a := integration.Authentication{}
		p, c, k := []byte(m.GetAuth().GetCaCert()), []byte(m.GetAuth().GetCliCert()), []byte(m.GetAuth().GetCliKey())

		switch m.GetAuth().GetMode() {
		case "NO_AUTH":
		case "TLS_Server":
			a.ServerCert = &p
		case "TLS_Server_Client":
			a.ServerCert = &p
			a.ClientCert = &c
			a.Key = &k
		case "TLS_Server_Client_Token":
			a.ServerCert = &p
			a.Key = &k
		}

		i, err = integration.NewBasicsStationIntegration(&integration.BasicsStationSetting{
			Protocol:       m.GetType(),
			ServerAddress:  m.GetServer(),
			ServerPort:     int32(m.GetPort()),
			Authentication: a,
		})
	default:
		err = fmt.Errorf("unsupport mode: %v", c.GetMode().GetBs())
	}
	if err != nil {
		return adr, err
	}

	if err = integration.ApplySettings(i); err != nil {
		return adr, err
	}

	hwconfig.Mode = c.Mode
	hwconfig.Region = c.Region
	hwconfig.Filter = c.Filter

	return adr, Save(c)
}

func Get(isAdmin bool) (*api.GetGateWayModeRegionResponse, error) {
	// proto: in order to achieve deep copy function
	b, err := proto.Marshal(&hwconfig)
	if err != nil {
		return nil, err
	}

	c := api.GetGateWayModeRegionResponse{}
	err = proto.Unmarshal(b, &c)
	if err != nil {
		return nil, err
	}

	if c.GetMode().GetMode() == "PF" {
		c.Mode.GetPf().GetProtocol().Type = "GWMP"
	}

	if (!isAdmin) && (c.GetMode().GetMode() == "BS") {
		c.GetMode().GetBs().GetAuth().CaCert = "······"
		c.GetMode().GetBs().GetAuth().CliCert = "······"
		c.GetMode().GetBs().GetAuth().CliKey = "······"
	}

	return &c, nil
}

func Save(c *api.ConfigGateWayModeRegionRequest) error {
	b, err := proto.Marshal(c)
	if err != nil {
		return err
	}

	return writeFile(HWPath+HWName, b)
}

func Load(c *api.GetGateWayModeRegionResponse) error {
	b, err := readFile(HWPath + HWName)
	if err != nil {
		return err
	}

	return proto.Unmarshal(b, c)
}

func LoadFromOriginFile() (*api.GetGateWayModeRegionResponse, error) {
	pf := cf.UdpPacketForwarder{}
	pf.ReadFrom(integration.PFPath + integration.PFName)

	ns := cf.NetworkServer{}
	ns.ReadFrom(integration.NSPath + integration.NSName)

	// gb := cf.GatewayBridge{}
	// gb.ReadFrom(integration.GBPath + integration.GBName)

	var r *api.GateWayRegion

	switch ns.Band.Name {
	case "CN470":
		sb := int32((pf.SX130xConfig.Radio0.Freq-470300000-1100000)/1600000 + 1)
		r = &api.GateWayRegion{
			RegionId: "CN470",
			RegionConfig: &api.GateWayRegion_Cn470{
				Cn470: &api.CN470Config{
					SubBandId: sb,
				},
			},
		}
	case "EU868":
		t := [](*api.TxGainLutItem){}
		for _, v := range pf.SX130xConfig.Radio0.TxGainLut {
			t = append(t, &api.TxGainLutItem{
				Rfpower: int64(v.RFPower),
				Pagain:  int64(v.PaGain),
				Pwridx:  int64(v.PwrIdx),
			})
		}
		r = &api.GateWayRegion{
			RegionId: "EU868",
			RegionConfig: &api.GateWayRegion_Eu868{
				Eu868: &api.EU868Config{
					TxFreq: &api.TXFreqItem{
						Min: pf.SX130xConfig.Radio0.TxFreqMin,
						Max: pf.SX130xConfig.Radio0.TxFreqMax,
					},
					Radio_0: &api.EU868Radio0{
						Enable:          pf.SX130xConfig.Radio0.Enable,
						Type:            pf.SX130xConfig.Radio0.Type,
						SingleInputMode: pf.SX130xConfig.Radio0.SingleInputMode,
						Freq:            pf.SX130xConfig.Radio0.Freq,
						RssiOffset:      pf.SX130xConfig.Radio0.RssiOffset,
						Rssicomp: &api.RssiTcomp{
							Coeffa: pf.SX130xConfig.Radio0.RssiTcomp.CoeffA,
							Coeffb: pf.SX130xConfig.Radio0.RssiTcomp.CoeffB,
							Coeffc: pf.SX130xConfig.Radio0.RssiTcomp.CoeffC,
							Coeffd: pf.SX130xConfig.Radio0.RssiTcomp.CoeffD,
							Coeffe: pf.SX130xConfig.Radio0.RssiTcomp.CoeffE,
						},
						TxEnable:  pf.SX130xConfig.Radio0.TxEnable,
						TxFreqMin: pf.SX130xConfig.Radio0.TxFreqMin,
						TxFreqMax: pf.SX130xConfig.Radio0.TxFreqMax,
						Txgainlut: t,
					},
					Radio_1: &api.EU868Radio1{
						Enable:          pf.SX130xConfig.Radio1.Enable,
						Type:            pf.SX130xConfig.Radio1.Type,
						SingleInputMode: pf.SX130xConfig.Radio1.SingleInputMode,
						Freq:            pf.SX130xConfig.Radio1.Freq,
						RssiOffset:      pf.SX130xConfig.Radio1.RssiOffset,
						Rssicomp: &api.RssiTcomp{
							Coeffa: pf.SX130xConfig.Radio1.RssiTcomp.CoeffA,
							Coeffb: pf.SX130xConfig.Radio1.RssiTcomp.CoeffB,
							Coeffc: pf.SX130xConfig.Radio1.RssiTcomp.CoeffC,
							Coeffd: pf.SX130xConfig.Radio1.RssiTcomp.CoeffD,
							Coeffe: pf.SX130xConfig.Radio1.RssiTcomp.CoeffE,
						},
						TxEnable: pf.SX130xConfig.Radio1.TxEnable,
					},
					Chan_LoraStd: &api.EU868ChannelLoraStandard{
						Enable:                pf.SX130xConfig.ChanLoraStd.Enable,
						Radio:                 pf.SX130xConfig.ChanLoraStd.Radio,
						If:                    pf.SX130xConfig.ChanLoraStd.IF,
						Bandwidth:             pf.SX130xConfig.ChanLoraStd.Bandwidth,
						SpreadFactor:          pf.SX130xConfig.ChanLoraStd.SpreadFactor,
						ImplicitHdr:           pf.SX130xConfig.ChanLoraStd.ImplicitHdr,
						Implicitpayloadlength: pf.SX130xConfig.ChanLoraStd.Implicitpayloadlength,
						ImplicitcrcEn:         pf.SX130xConfig.ChanLoraStd.ImplicitcrcEn,
						Implicitcoderate:      pf.SX130xConfig.ChanLoraStd.Implicitcoderate,
					},
					ChanMultiSF_0: &api.EU868ChannelMultiSF{
						Enable: pf.SX130xConfig.ChanMultiSF0.Enable,
						Radio:  pf.SX130xConfig.ChanMultiSF0.Radio,
						Offset: pf.SX130xConfig.ChanMultiSF0.IF,
					},
					ChanMultiSF_1: &api.EU868ChannelMultiSF{
						Enable: pf.SX130xConfig.ChanMultiSF1.Enable,
						Radio:  pf.SX130xConfig.ChanMultiSF1.Radio,
						Offset: pf.SX130xConfig.ChanMultiSF1.IF,
					},
					ChanMultiSF_2: &api.EU868ChannelMultiSF{
						Enable: pf.SX130xConfig.ChanMultiSF2.Enable,
						Radio:  pf.SX130xConfig.ChanMultiSF2.Radio,
						Offset: pf.SX130xConfig.ChanMultiSF2.IF,
					},
					ChanMultiSF_3: &api.EU868ChannelMultiSF{
						Enable: pf.SX130xConfig.ChanMultiSF3.Enable,
						Radio:  pf.SX130xConfig.ChanMultiSF3.Radio,
						Offset: pf.SX130xConfig.ChanMultiSF3.IF,
					},
					ChanMultiSF_4: &api.EU868ChannelMultiSF{
						Enable: pf.SX130xConfig.ChanMultiSF4.Enable,
						Radio:  pf.SX130xConfig.ChanMultiSF4.Radio,
						Offset: pf.SX130xConfig.ChanMultiSF4.IF,
					},
					ChanMultiSF_5: &api.EU868ChannelMultiSF{
						Enable: pf.SX130xConfig.ChanMultiSF5.Enable,
						Radio:  pf.SX130xConfig.ChanMultiSF5.Radio,
						Offset: pf.SX130xConfig.ChanMultiSF5.IF,
					},
					ChanMultiSF_6: &api.EU868ChannelMultiSF{
						Enable: pf.SX130xConfig.ChanMultiSF6.Enable,
						Radio:  pf.SX130xConfig.ChanMultiSF6.Radio,
						Offset: pf.SX130xConfig.ChanMultiSF6.IF,
					},
					ChanMultiSF_7: &api.EU868ChannelMultiSF{
						Enable: pf.SX130xConfig.ChanMultiSF7.Enable,
						Radio:  pf.SX130xConfig.ChanMultiSF7.Radio,
						Offset: pf.SX130xConfig.ChanMultiSF7.IF,
					},
				},
			},
		}
	case "US915":
		sb := int32((pf.SX130xConfig.Radio0.Freq-470300000-1100000)/1600000 + 1)
		r = &api.GateWayRegion{
			RegionId: "US915",
			RegionConfig: &api.GateWayRegion_Us915{
				Us915: &api.US915Config{
					SubBandId: sb,
				},
			},
		}
	}
	return &api.GetGateWayModeRegionResponse{
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
		Filter: &api.Filter{
			WhiteList: &api.WhiteList{},
		},
	}, nil
}

func readFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("error getting file information: %v", err)
	}
	fileSize := fileInfo.Size()

	data := make([]byte, fileSize)
	_, err = file.Read(data)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return data, nil
}

func writeFile(path string, data []byte) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}
	return nil
}
