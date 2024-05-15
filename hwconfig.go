package hwconfig

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	cf "github.com/NephogramX/hwconfig/configfile"

	"gitee.com/dfrobotcd/chirpstack-api/go/as/external/api"
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

func Setup(region string) error {
	if _, err := os.Stat(HWPath + HWName); os.IsNotExist(err) {
		restoreDefault(region)
		if err != nil {
			return err
		}
		c := api.ConfigGateWayModeRegionRequest{
			Mode:   hwconfig.Mode,
			Region: hwconfig.Region,
			Filter: hwconfig.Filter,
		}
		save(&c)
	} else if err != nil {
		return err
	}
	return load(&hwconfig)
}

func SetupDebug() error {
	HWPath = "./build/"
	integration.BSPath = "./build/"
	integration.PFPath = "./build/"
	integration.GBPath = "./build/"
	integration.NSPath = "./build/"
	if _, err := os.Stat(HWPath + HWName); os.IsNotExist(err) {
		err = restoreDefault("EU868")
		if err != nil {
			return err
		}
		c := api.ConfigGateWayModeRegionRequest{
			Mode:   hwconfig.Mode,
			Region: hwconfig.Region,
			Filter: hwconfig.Filter,
		}
		return save(&c)
	} else if err != nil {
		return err
	} else {
		return load(&hwconfig)
	}
}

func GetMode() *api.GateWayMode {
	return hwconfig.Mode
}

func GetRegion() *api.GateWayRegion {
	return hwconfig.Region
}

func GetFilter() *api.Filter {
	return hwconfig.Filter
}

func SetMode(c *api.GateWayMode) error {
	if c == nil {
		return fmt.Errorf("illegal parameters: config is nil")
	}

	tmp := &api.ConfigGateWayModeRegionRequest{
		Mode:   c,
		Region: hwconfig.Region,
		Filter: hwconfig.Filter,
	}

	if err := applySettings(tmp); err != nil {
		return err
	}

	hwconfig.Mode = c
	return save(tmp)
}

func SetRegion(c *api.GateWayRegion) error {
	if c == nil {
		return fmt.Errorf("illegal parameters: config is nil")
	}

	tmp := &api.ConfigGateWayModeRegionRequest{
		Mode:   hwconfig.Mode,
		Region: c,
		Filter: hwconfig.Filter,
	}

	if err := applySettings(tmp); err != nil {
		return err
	}

	hwconfig.Region = c
	return save(tmp)
}

func SetFilter(c *api.Filter) error {
	if c == nil {
		return fmt.Errorf("illegal parameters: config is nil")
	}

	tmp := &api.ConfigGateWayModeRegionRequest{
		Mode:   hwconfig.Mode,
		Region: hwconfig.Region,
		Filter: c,
	}

	if err := applySettings(tmp); err != nil {
		return err
	}

	hwconfig.Filter = c
	return save(tmp)
}

// private tools

func save(c *api.ConfigGateWayModeRegionRequest) error {
	b, err := proto.Marshal(c)
	if err != nil {
		return err
	}

	// b, err := json.MarshalIndent(c, "  ", "   ")
	// if err != nil {
	// 	return err
	// }

	return cf.WriteFile(HWPath+HWName, b)
}

func load(c *api.GetGateWayModeRegionResponse) error {
	b, err := cf.ReadFile(HWPath + HWName)
	if err != nil {
		return err
	}
	// file, err := os.Open(HWPath + HWName)
	// if err != nil {
	// 	return fmt.Errorf("error opening file: %v, %v", HWPath+HWName, err)
	// }
	// defer file.Close()

	return proto.Unmarshal(b, c)
}

func transCfTxgainlut(c []*cf.TxGainLutItem) []*api.TxGainLutItem {
	out := []*api.TxGainLutItem{}

	for _, v := range c {
		out = append(out, &api.TxGainLutItem{
			Rfpower: int64(v.RFPower),
			Pagain:  int64(v.PaGain),
			Pwridx:  int64(v.PwrIdx),
		})
	}

	return out
}

func transCfChanMultiSF(c *cf.ChanMultiSF) *api.EU868ChannelMultiSF {
	return &api.EU868ChannelMultiSF{
		Enable: c.Enable,
		Radio:  c.Radio,
		Offset: c.IF,
	}
}

func transCfChannel(c *cf.Channel) *api.EU868Config {
	return &api.EU868Config{
		Radio_0: &api.EU868Radio0{
			Enable:          c.Radio0.Enable,
			Type:            c.Radio0.Type,
			SingleInputMode: c.Radio0.SingleInputMode,
			Freq:            c.Radio0.Freq,
			RssiOffset:      c.Radio0.RssiOffset,
			Rssicomp: &api.RssiTcomp{
				Coeffa: c.Radio0.RssiTcomp.CoeffA,
				Coeffb: c.Radio0.RssiTcomp.CoeffB,
				Coeffc: c.Radio0.RssiTcomp.CoeffC,
				Coeffd: c.Radio0.RssiTcomp.CoeffD,
				Coeffe: c.Radio0.RssiTcomp.CoeffE,
			},
			TxEnable:  c.Radio0.TxEnable,
			TxFreqMin: c.Radio0.TxFreqMin,
			TxFreqMax: c.Radio0.TxFreqMax,
			Txgainlut: transCfTxgainlut(c.Radio0.TxGainLut),
		},
		Radio_1: &api.EU868Radio1{
			Enable:          c.Radio1.Enable,
			Type:            c.Radio1.Type,
			SingleInputMode: c.Radio1.SingleInputMode,
			Freq:            c.Radio1.Freq,
			RssiOffset:      c.Radio1.RssiOffset,
			Rssicomp: &api.RssiTcomp{
				Coeffa: c.Radio1.RssiTcomp.CoeffA,
				Coeffb: c.Radio1.RssiTcomp.CoeffB,
				Coeffc: c.Radio1.RssiTcomp.CoeffC,
				Coeffd: c.Radio1.RssiTcomp.CoeffD,
				Coeffe: c.Radio1.RssiTcomp.CoeffE,
			},
			TxEnable: c.Radio1.TxEnable,
		},
		Chan_LoraStd: &api.EU868ChannelLoraStandard{
			Enable:                c.ChanLoraStd.Enable,
			Radio:                 c.ChanLoraStd.Radio,
			If:                    c.ChanLoraStd.IF,
			Bandwidth:             c.ChanLoraStd.Bandwidth,
			SpreadFactor:          c.ChanLoraStd.SpreadFactor,
			ImplicitHdr:           c.ChanLoraStd.ImplicitHdr,
			Implicitpayloadlength: c.ChanLoraStd.Implicitpayloadlength,
			ImplicitcrcEn:         c.ChanLoraStd.ImplicitcrcEn,
			Implicitcoderate:      c.ChanLoraStd.Implicitcoderate,
		},
		ChanMultiSF_0: transCfChanMultiSF(c.ChanMultiSF[0]),
		ChanMultiSF_1: transCfChanMultiSF(c.ChanMultiSF[1]),
		ChanMultiSF_2: transCfChanMultiSF(c.ChanMultiSF[2]),
		ChanMultiSF_3: transCfChanMultiSF(c.ChanMultiSF[3]),
		ChanMultiSF_4: transCfChanMultiSF(c.ChanMultiSF[4]),
		ChanMultiSF_5: transCfChanMultiSF(c.ChanMultiSF[5]),
		ChanMultiSF_6: transCfChanMultiSF(c.ChanMultiSF[6]),
		ChanMultiSF_7: transCfChanMultiSF(c.ChanMultiSF[7]),
	}
}

func transApiTxgainlut(c []*api.TxGainLutItem) []*cf.TxGainLutItem {
	out := []*cf.TxGainLutItem{}

	for _, v := range c {
		out = append(out, &cf.TxGainLutItem{
			RFPower: int(v.Rfpower),
			PaGain:  int(v.Pagain),
			PwrIdx:  int(v.Pwridx),
		})
	}

	return out
}

func transApiChanMultiSF(c *api.EU868ChannelMultiSF) *cf.ChanMultiSF {
	return &cf.ChanMultiSF{
		Enable: c.Enable,
		Radio:  c.Radio,
		IF:     c.Offset,
	}
}

func transApiChannel(c *api.EU868Config) *cf.Channel {
	return &cf.Channel{
		Radio0: cf.Radio0{
			Enable:          c.Radio_0.Enable,
			Type:            c.Radio_0.Type,
			SingleInputMode: c.Radio_0.SingleInputMode,
			Freq:            c.Radio_0.Freq,
			RssiOffset:      c.Radio_0.RssiOffset,
			RssiTcomp: cf.RssiTcomp{
				CoeffA: c.Radio_0.Rssicomp.Coeffa,
				CoeffB: c.Radio_0.Rssicomp.Coeffb,
				CoeffC: c.Radio_0.Rssicomp.Coeffc,
				CoeffD: c.Radio_0.Rssicomp.Coeffd,
				CoeffE: c.Radio_0.Rssicomp.Coeffe,
			},
			TxEnable:  c.Radio_0.TxEnable,
			TxFreqMin: c.Radio_0.TxFreqMin,
			TxFreqMax: c.Radio_0.TxFreqMax,
			TxGainLut: transApiTxgainlut(c.Radio_0.Txgainlut),
		},
		Radio1: cf.Radio1{
			Enable:          c.Radio_1.Enable,
			Type:            c.Radio_1.Type,
			SingleInputMode: c.Radio_1.SingleInputMode,
			Freq:            c.Radio_1.Freq,
			RssiOffset:      c.Radio_1.RssiOffset,
			RssiTcomp: cf.RssiTcomp{
				CoeffA: c.Radio_1.Rssicomp.Coeffa,
				CoeffB: c.Radio_1.Rssicomp.Coeffb,
				CoeffC: c.Radio_1.Rssicomp.Coeffc,
				CoeffD: c.Radio_1.Rssicomp.Coeffd,
				CoeffE: c.Radio_1.Rssicomp.Coeffe,
			},
			TxEnable: c.Radio_1.TxEnable,
		},
		ChanLoraStd: cf.ChanLoraStd{
			Enable:                c.Chan_LoraStd.Enable,
			Radio:                 c.Chan_LoraStd.Radio,
			IF:                    c.Chan_LoraStd.If,
			Bandwidth:             c.Chan_LoraStd.Bandwidth,
			SpreadFactor:          c.Chan_LoraStd.SpreadFactor,
			ImplicitHdr:           c.Chan_LoraStd.ImplicitHdr,
			Implicitpayloadlength: c.Chan_LoraStd.Implicitpayloadlength,
			ImplicitcrcEn:         c.Chan_LoraStd.ImplicitcrcEn,
			Implicitcoderate:      c.Chan_LoraStd.Implicitcoderate,
		},
		ChanLoraFSK: cf.ChanLoraFSK{
			Enable: false,
		},
		ChanMultiSFAll: *band.ChanMultiSFAll(),
		ChanMultiSF: [8]*cf.ChanMultiSF{
			transApiChanMultiSF(c.ChanMultiSF_0),
			transApiChanMultiSF(c.ChanMultiSF_1),
			transApiChanMultiSF(c.ChanMultiSF_2),
			transApiChanMultiSF(c.ChanMultiSF_3),
			transApiChanMultiSF(c.ChanMultiSF_4),
			transApiChanMultiSF(c.ChanMultiSF_5),
			transApiChanMultiSF(c.ChanMultiSF_6),
			transApiChanMultiSF(c.ChanMultiSF_7),
		},
	}
}

func transRequest(c *api.ConfigGateWayModeRegionRequest) *api.GetGateWayModeRegionResponse {
	out := api.GetGateWayModeRegionResponse{
		Mode:   &api.GateWayMode{},
		Region: &api.GateWayRegion{},
		Filter: &api.Filter{}}
	deepCopy(c.Mode, out.Mode)
	deepCopy(c.Region, out.Region)
	deepCopy(c.Filter, out.Filter)
	return &out
}

func transResponse(c *api.GetGateWayModeRegionResponse) *api.ConfigGateWayModeRegionRequest {
	out := api.ConfigGateWayModeRegionRequest{
		Mode:   &api.GateWayMode{},
		Region: &api.GateWayRegion{},
		Filter: &api.Filter{},
	}
	deepCopy(c.Mode, out.Mode)
	deepCopy(c.Region, out.Region)
	deepCopy(c.Filter, out.Filter)
	return &out
}

func isAdrUpdated(c *api.NSADR) bool {
	if !c.GetEnable() {
		return false
	}

	if hwconfig.GetMode().GetMode() != "NS" {
		return true
	}

	if c.GetDrIdMax() != hwconfig.GetMode().GetNs().GetAdr().GetDrIdMax() ||
		c.GetDrIdMin() != hwconfig.GetMode().GetNs().GetAdr().GetDrIdMin() {
		return true
	}

	return false
}

func restoreDefault(bd string) error {
	hwconfig = api.GetGateWayModeRegionResponse{
		Mode: &api.GateWayMode{
			Mode: "NS",
			ModeConfig: &api.GateWayMode_Ns{
				Ns: &api.BuiltInNetworkServer{
					NetworkId: "000000",
					Adr: &api.NSADR{
						Enable:  true,
						DrIdMin: 0,
						DrIdMax: 5,
						Margin:  10,
					},
					Rx1: &api.NSRX1{
						DrOffset: 0,
						Delay:    1,
					},
					Rx2: &api.NSRX2{
						Freq:    -1,
						DrIndex: -1,
					},
					DwellTimeLimit:  &api.NSDwellTimeLimit{}, // Avoid frontend parse failure
					DownlinkTxPower: -1,
				},
			},
		},
		Filter: &api.Filter{
			AutoFilter: &api.AutoFilter{
				Enable: false,
			},
			WhiteList: &api.WhiteList{
				Enable: false,
			},
		},
	}
	switch bd {
	case "CN470":
		hwconfig.Region = &api.GateWayRegion{
			RegionId: "CN470",
			RegionConfig: &api.GateWayRegion_Cn470{
				Cn470: &api.CN470Config{
					SubBandId: 2,
				},
			},
		}
	case "EU868":
		hwconfig.Region = &api.GateWayRegion{
			RegionId: "EU868",
			RegionConfig: &api.GateWayRegion_Eu868{
				Eu868: transCfChannel(band.GetEU868DefaultChannelSettings()),
			},
		}
	case "US915":
		hwconfig.Region = &api.GateWayRegion{
			RegionId: "US915",
			RegionConfig: &api.GateWayRegion_Us915{
				Us915: &api.US915Config{
					SubBandId: 2,
				},
			},
		}
	default:
		return fmt.Errorf("unsupported region: %v", bd)
	}
	return applySettings(transResponse(&hwconfig))
}

func applySettings(c *api.ConfigGateWayModeRegionRequest) error {
	// gateway region
	var (
		b   band.Band
		i   integration.Integration
		err error
	)

	switch c.GetRegion().GetRegionId() {
	case "EU868":
		r := c.GetRegion().GetEu868()
		if r == nil {
			return errors.New("fail to parse region parameters")
		}
		b, err = band.NewBandEU868(transApiChannel(r))
	case "US915":
		r := c.GetRegion().GetUs915()
		if r == nil {
			return errors.New("fail to parse region parameters")
		}
		b, err = band.NewBandUS915(r.GetSubBandId())
	default:
		err = fmt.Errorf("unsupport region: %v", c.Region.RegionId)
	}

	if err != nil {
		return err
	}

	// gateway mode
	switch c.GetMode().GetMode() {
	case "NS":
		m := c.GetMode().GetNs()
		ni, je := []string{}, [][2]string{}
		if m == nil {
			return errors.New("fail to parse mode parameters")
		}
		if c.GetFilter().GetWhiteList().GetEnable() {
			ni = c.GetFilter().GetWhiteList().GetOuiList()
			for _, jl := range c.GetFilter().GetWhiteList().GetJoinList() {
				je = append(je, [2]string{jl.GetFrom(), jl.GetTo()})
			}
		}

		if isAdrUpdated(m.GetAdr()) {

		}

		i, err = integration.NewBuiltinIntegration(&integration.BuiltinNSSettings{
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
			return errors.New("fail to parse mode parameters")
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
			return errors.New("fail to parse mode parameters")
		}

		a := integration.Authentication{}
		p, c, k, t := []byte(m.GetAuth().GetCaCert()), []byte(m.GetAuth().GetCliCert()), []byte(m.GetAuth().GetCliKey()), []byte(m.GetAuth().GetToken())

		switch m.GetAuth().GetMode() {
		case "NO_AUTH":
			m.GetAuth().CaCert = ""
			m.GetAuth().CliCert = ""
			m.GetAuth().CliKey = ""
			m.GetAuth().Token = ""
		case "TLS_Server":
			a.ServerCert = &p
			m.GetAuth().CliCert = ""
			m.GetAuth().CliKey = ""
			m.GetAuth().Token = ""
		case "TLS_Server_Client":
			a.ServerCert = &p
			a.ClientCert = &c
			a.Key = &k
			m.GetAuth().Token = ""
		case "TLS_Server_Client_Token":
			a.ServerCert = &p
			a.Key = &t
			m.GetAuth().CliCert = ""
			m.GetAuth().CliKey = ""
		}

		i, err = integration.NewBasicsStationIntegration(&integration.BasicsStationSetting{
			Protocol:       m.GetType(),
			ServerAddress:  m.GetServer(),
			ServerPort:     int32(m.GetPort()),
			Authentication: a,
		})
	default:
		err = fmt.Errorf("unsupported mode: %v", c.GetMode().GetMode())
	}
	if err != nil {
		return err
	}

	return integration.ApplySettings(i)
}

func deepCopy(cin interface{}, cout interface{}) error {
	switch cin.(type) {
	case *api.GateWayMode:
		in, _ := cin.(*api.GateWayMode)
		b, err := proto.Marshal(in)
		if err != nil {
			return err
		}
		out, ok := cout.(*api.GateWayMode)
		if !ok {
			return fmt.Errorf("i/o type mismatch")
		}
		return proto.Unmarshal(b, out)
	case *api.GateWayRegion:
		in, _ := cin.(*api.GateWayRegion)
		b, err := proto.Marshal(in)
		if err != nil {
			return err
		}
		out, ok := cout.(*api.GateWayRegion)
		if !ok {
			return fmt.Errorf("i/o type mismatch")
		}
		return proto.Unmarshal(b, out)
	case *api.Filter:
		in, _ := cin.(*api.Filter)
		b, err := proto.Marshal(in)
		if err != nil {
			return err
		}
		out, ok := cout.(*api.Filter)
		if !ok {
			return fmt.Errorf("i/o type mismatch")
		}
		return proto.Unmarshal(b, out)
	default:
		return fmt.Errorf("unsupported type %T", cin)
	}
}

func Print(a any) {
	b, err := json.MarshalIndent(a, " ", "	  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", string(b))
}
