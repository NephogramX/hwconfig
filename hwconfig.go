package hwconfig

import (
	"errors"
	"fmt"
	"os"

	"gitee.com/arya123/chirpstack-api/go/as/external/api"
	"github.com/NephogramX/hwconfig/band"
	"github.com/NephogramX/hwconfig/integration"
	"github.com/golang/protobuf/proto"
)

const (
	// HWPath = "/gw/config/hwconfig/"
	HWPath = "./build/"
	HWName = ".hwconfig"
)

var hwconfig api.GetGateWayModeRegionResponse

func Setup() error {
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
		b, err = band.NewBandEU868(r.GetRadio_1().GetFreq(), []int32{
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

func Get() (*api.GetGateWayModeRegionResponse, error) {
	return &hwconfig, nil
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
