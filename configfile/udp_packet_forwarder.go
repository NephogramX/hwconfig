package configfile

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

// global_config.json

type FineTimeStamp struct {
	Enable bool   `json:"enable"`
	Mode   string `json:"mode"`
}

type RssiTcomp struct {
	CoeffA float32 `json:"coeff_a"`
	CoeffB float32 `json:"coeff_b"`
	CoeffC float32 `json:"coeff_c"`
	CoeffD float32 `json:"coeff_d"`
	CoeffE float32 `json:"coeff_e"`
}

type TxGainLutItem struct {
	RFPower int `json:"rf_power"`
	PaGain  int `json:"pa_gain"`
	PwrIdx  int `json:"pwr_idx"`
}

type ChanMultiSF struct {
	Enable bool  `json:"enable"`
	Radio  int32 `json:"radio"`
	IF     int32 `json:"if"`
}

type ChanMultiSFAll struct {
	SpreadingFactorEnable []int32 `json:"spreading_factor_enable"`
	Radio                 int32   `json:"radio"`
	IF                    int32   `json:"if"`
}

type ChanLoraStd struct {
	Enable                bool  `json:"enable"`
	Radio                 int32 `json:"radio"`
	IF                    int32 `json:"if"`
	Bandwidth             int32 `json:"bandwidth"`
	SpreadFactor          int32 `json:"spread_factor"`
	ImplicitHdr           bool  `json:"implicit_hdr"`
	Implicitpayloadlength int32 `json:"implicit_payload_length"`
	ImplicitcrcEn         bool  `json:"implicit_crc_en"`
	Implicitcoderate      int32 `json:"implicit_coderate"`
}

type ChanLoraFSK struct {
	Enable    bool  `json:"enable"`
	Radio     int32 `json:"radio"`
	IF        int32 `json:"if"`
	Bandwidth int32 `json:"bandwidth"`
	Datarate  int32 `json:"datarate"`
}

type Radio0 struct {
	Enable          bool    `json:"enable"`
	Type            string  `json:"type"`
	SingleInputMode bool    `json:"single_input_mode"`
	Freq            int32   `json:"freq"`
	RssiOffset      float32 `json:"rssi_offset"`
	RssiTcomp       `json:"rssi_tcomp"`
	TxEnable        bool             `json:"tx_enable"`
	TxFreqMin       int32            `json:"tx_freq_min"`
	TxFreqMax       int32            `json:"tx_freq_max"`
	TxGainLut       []*TxGainLutItem `json:"tx_gain_lut"`
}

type Radio1 struct {
	Enable          bool    `json:"enable"`
	Type            string  `json:"type"`
	SingleInputMode bool    `json:"single_input_mode"`
	Freq            int32   `json:"freq"`
	RssiOffset      float32 `json:"rssi_offset"`
	RssiTcomp       `json:"rssi_tcomp"`
	TxEnable        bool `json:"tx_enable"`
}

type SX130xConfig struct {
	ComType        string `json:"com_type"`
	ComPath        string `json:"com_path"`
	LorawanPublic  bool   `json:"lorawan_public"`
	ClkSrc         int32  `json:"clksrc"`
	AntennaGain    int32  `json:"antenna_gain"`
	FullDuplex     bool   `json:"full_duplex"`
	FineTimeStamp  `json:"fine_timestamp"`
	Radio0         `json:"radio_0"`
	Radio1         `json:"radio_1"`
	ChanMultiSFAll ChanMultiSFAll `json:"chan_multiSF_All"`
	ChanMultiSF0   ChanMultiSF    `json:"chan_multiSF_0"`
	ChanMultiSF1   ChanMultiSF    `json:"chan_multiSF_1"`
	ChanMultiSF2   ChanMultiSF    `json:"chan_multiSF_2"`
	ChanMultiSF3   ChanMultiSF    `json:"chan_multiSF_3"`
	ChanMultiSF4   ChanMultiSF    `json:"chan_multiSF_4"`
	ChanMultiSF5   ChanMultiSF    `json:"chan_multiSF_5"`
	ChanMultiSF6   ChanMultiSF    `json:"chan_multiSF_6"`
	ChanMultiSF7   ChanMultiSF    `json:"chan_multiSF_7"`
	ChanLoraStd    `json:"chan_Lora_std"`
	ChanLoraFSK    `json:"chan_FSK"`
}

type GateWayConfig struct {
	GatewayID          string `json:"gateway_ID"`
	ServerAddress      string `json:"server_address"`
	ServPortUp         int32  `json:"serv_port_up"`
	ServPortDown       int32  `json:"serv_port_down"`
	KeepaliveInterval  int32  `json:"keepalive_interval"`
	StatInterval       int32  `json:"stat_interval"`
	PushTimeoutMs      int32  `json:"push_timeout_ms"`
	ForwardCrcValid    bool   `json:"forward_crc_valid"`
	ForwardCrcError    bool   `json:"forward_crc_error"`
	ForwardCrcDisabled bool   `json:"forward_crc_disabled"`

	// Class B & GPS parameters, to do in future

	// GPSTTYPath        string  `json:"gps_tty_path"`
	// RefLatitude       float32 `json:"ref_latitude"`
	// RefLongitude      float32 `json:"ref_longitude"`
	// RefAltitude       int32   `json:"ref_altitude"`
	// AutoquitThreshold int32   `json:"autoquit_threshold"`
	// BeaconPeriod      int32   `json:"beacon_period"`
	// BeaconFreqHZ      int32   `json:"beacon_freq_hz"`
	// BeaconFreqNB      int32   `json:"beacon_freq_nb"`
	// BeaconFreqStep    int32   `json:"beacon_freq_step"`
	// BeaconDatarate    int32   `json:"beacon_datarate"`
	// BeaconBwHZ        int32   `json:"beacon_bw_hz"`
	// BeaconPower       int32   `json:"beacon_power"`
	// BeaconInfodesc    int32   `json:"beacon_infodesc"`
}

type RefPayloadItem struct {
	ID string `json:"id"`
}

// type DebugConf struct {
// 	RefPayload []RefPayloadItem `json:"ref_payload"`
// 	LogFile    string           `json:"log_file"`
// }

type UdpPacketForwarder struct {
	SX130xConfig  `json:"SX130x_conf"`
	GateWayConfig `json:"gateway_conf"`
	File          `json:"-"`
}

// Channel Settings
type Channel struct {
	Radio0         Radio0
	Radio1         Radio1
	ChanMultiSFAll ChanMultiSFAll
	ChanMultiSF    [8]*ChanMultiSF
	ChanLoraStd    ChanLoraStd
	ChanLoraFSK    ChanLoraFSK
}

// server settings
type Server struct {
	Address  string
	PortUp   int32
	PortDown int32
}

type Comm struct {
	PushTimeoutMs        int32
	StatIntervalSec      int32
	KeepaliveIntervalSec int32
}

// lora_pkt_fwd settings
type PFSettings struct {
	Channel
	Server
	Comm
	File
}

func NewUdpPacketForwarderCF(s *PFSettings) *UdpPacketForwarder {
	return &UdpPacketForwarder{
		File: s.File,
		SX130xConfig: SX130xConfig{
			ComType:       "SPI",
			ComPath:       "/dev/spidev2.0",
			LorawanPublic: true,
			ClkSrc:        0,
			AntennaGain:   0,
			FullDuplex:    false,
			FineTimeStamp: FineTimeStamp{
				Enable: false,
				Mode:   "all_sf",
			},
			Radio0:         s.Radio0,
			Radio1:         s.Radio1,
			ChanMultiSFAll: s.ChanMultiSFAll,
			ChanMultiSF0:   *s.ChanMultiSF[0],
			ChanMultiSF1:   *s.ChanMultiSF[1],
			ChanMultiSF2:   *s.ChanMultiSF[2],
			ChanMultiSF3:   *s.ChanMultiSF[3],
			ChanMultiSF4:   *s.ChanMultiSF[4],
			ChanMultiSF5:   *s.ChanMultiSF[5],
			ChanMultiSF6:   *s.ChanMultiSF[6],
			ChanMultiSF7:   *s.ChanMultiSF[7],
			ChanLoraStd:    s.ChanLoraStd,
			ChanLoraFSK:    s.ChanLoraFSK,
		},
		GateWayConfig: GateWayConfig{
			GatewayID:          "0000000000000000",
			ServerAddress:      s.Address,
			ServPortUp:         s.PortUp,
			ServPortDown:       s.PortDown,
			KeepaliveInterval:  s.KeepaliveIntervalSec,
			StatInterval:       s.StatIntervalSec,
			PushTimeoutMs:      s.PushTimeoutMs,
			ForwardCrcValid:    true,
			ForwardCrcError:    false,
			ForwardCrcDisabled: false,
		},
	}
}

func (c *UdpPacketForwarder) Write() error {
	if c == nil {
		return nil
	}
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	return WriteFile(c.File.String(), jsonData)
}

func (c *UdpPacketForwarder) ReadFrom(p string) error {
	if c == nil {
		return errors.New("nil interface")
	}
	c.File = File{
		Name: filepath.Base(p),
		Path: filepath.Dir(p),
	}
	file, err := os.Open(p)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(c)
}
