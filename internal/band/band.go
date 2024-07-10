package band

import (
	"gitee.com/dfrobotcd/chirpstack-api/go/as/external/api"
	"github.com/pkg/errors"
)

type ExtraChannel struct {
	Frequency int `mapstructure:"frequency" toml:"frequency"`
	MinDR     int `mapstructure:"min_dr" toml:"min_dr"`
	MaxDR     int `mapstructure:"max_dr" toml:"max_dr"`
}

type SX130xConfig struct {
	ComType        string `json:"com_type"`
	ComPath        string `json:"com_path"`
	LorawanPublic  bool   `json:"lorawan_public"`
	ClkSrc         int32  `json:"clksrc"`
	AntennaGain    int32  `json:"antenna_gain"`
	FullDuplex     bool   `json:"full_duplex"`
	FineTimeStamp  `json:"fine_timestamp"`
	Radio0         `json:"radio_0" mapstructure:"radio_0"`
	Radio1         `json:"radio_1" mapstructure:"radio_1"`
	ChanMultiSFAll ChanMultiSFAll `json:"chan_multiSF_All" mapstructure:"chan_multiSF_All"`
	ChanMultiSF0   ChanMultiSF    `json:"chan_multiSF_0" mapstructure:"chan_multiSF_0"`
	ChanMultiSF1   ChanMultiSF    `json:"chan_multiSF_1" mapstructure:"chan_multiSF_1"`
	ChanMultiSF2   ChanMultiSF    `json:"chan_multiSF_2" mapstructure:"chan_multiSF_2"`
	ChanMultiSF3   ChanMultiSF    `json:"chan_multiSF_3" mapstructure:"chan_multiSF_3"`
	ChanMultiSF4   ChanMultiSF    `json:"chan_multiSF_4" mapstructure:"chan_multiSF_4"`
	ChanMultiSF5   ChanMultiSF    `json:"chan_multiSF_5" mapstructure:"chan_multiSF_5"`
	ChanMultiSF6   ChanMultiSF    `json:"chan_multiSF_6" mapstructure:"chan_multiSF_6"`
	ChanMultiSF7   ChanMultiSF    `json:"chan_multiSF_7" mapstructure:"chan_multiSF_7"`
	ChanLoraStd    `json:"chan_Lora_std" mapstructure:"chan_Lora_std"`
	ChanLoraFSK    `json:"chan_FSK" mapstructure:"chan_FSK"`
}

type FineTimeStamp struct {
	Enable bool   `json:"enable"`
	Mode   string `json:"mode"`
}

type Radio0 struct {
	Enable          bool    `json:"enable" mapstructure:"enable"`
	Type            string  `json:"type" mapstructure:"type"`
	SingleInputMode bool    `json:"single_input_mode" mapstructure:"single_input_mode"`
	Freq            int32   `json:"freq" mapstructure:"freq"`
	RssiOffset      float32 `json:"rssi_offset" mapstructure:"rssi_offset"`
	RssiTcomp       `json:"rssi_tcomp" mapstructure:"rssi_tcomp"`
	TxEnable        bool            `json:"tx_enable" mapstructure:"tx_enable"`
	TxFreqMin       int32           `json:"tx_freq_min" mapstructure:"tx_freq_min"`
	TxFreqMax       int32           `json:"tx_freq_max" mapstructure:"tx_freq_max"`
	TxGainLut       []TxGainLutItem `json:"tx_gain_lut" mapstructure:"tx_gain_lut"`
}

type Radio1 struct {
	Enable          bool    `json:"enable" mapstructure:"enable"`
	Type            string  `json:"type" mapstructure:"type"`
	SingleInputMode bool    `json:"single_input_mode" mapstructure:"single_input_mode"`
	Freq            int32   `json:"freq" mapstructure:"freq"`
	RssiOffset      float32 `json:"rssi_offset" mapstructure:"rssi_offset"`
	RssiTcomp       `json:"rssi_tcomp" mapstructure:"rssi_tcomp"`
	TxEnable        bool `json:"tx_enable" mapstructure:"tx_enable"`
}

type ChanMultiSFAll struct {
	SpreadingFactorEnable []int32 `json:"spreading_factor_enable" mapstructure:"spreading_factor_enable"`
}

type ChanMultiSF struct {
	Enable bool  `json:"enable" mapstructure:"enable"`
	Radio  int32 `json:"radio" mapstructure:"radio"`
	IF     int32 `json:"if" mapstructure:"if"`
}

type ChanLoraStd struct {
	Enable                bool  `json:"enable" mapstructure:"enable"`
	Radio                 int32 `json:"radio" mapstructure:"radio"`
	IF                    int32 `json:"if" mapstructure:"if"`
	Bandwidth             int32 `json:"bandwidth" mapstructure:"bandwidth"`
	SpreadFactor          int32 `json:"spread_factor" mapstructure:"spread_factor"`
	ImplicitHdr           bool  `json:"implicit_hdr" mapstructure:"implicit_hdr"`
	Implicitpayloadlength int32 `json:"implicit_payload_length" mapstructure:"implicit_payload_length"`
	ImplicitcrcEn         bool  `json:"implicit_crc_en" mapstructure:"implicit_crc_en"`
	Implicitcoderate      int32 `json:"implicit_coderate" mapstructure:"implicit_coderate"`
}

type ChanLoraFSK struct {
	Enable    bool  `json:"enable" mapstructure:"enable"`
	Radio     int32 `json:"radio" mapstructure:"radio"`
	IF        int32 `json:"if" mapstructure:"if"`
	Bandwidth int32 `json:"bandwidth" mapstructure:"bandwidth"`
	Datarate  int32 `json:"datarate" mapstructure:"datarate"`
}

type RssiTcomp struct {
	CoeffA float32 `json:"coeff_a" mapstructure:"coeff_a"`
	CoeffB float32 `json:"coeff_b" mapstructure:"coeff_b"`
	CoeffC float32 `json:"coeff_c" mapstructure:"coeff_c"`
	CoeffD float32 `json:"coeff_d" mapstructure:"coeff_d"`
	CoeffE float32 `json:"coeff_e" mapstructure:"coeff_e"`
}

type TxGainLutItem struct {
	RFPower int `json:"rf_power" mapstructure:"rf_power"`
	PaGain  int `json:"pa_gain" mapstructure:"pa_gain"`
	PwrIdx  int `json:"pwr_idx" mapstructure:"pwr_idx"`
}

type Band interface {
	String() string

	GetSX130xConfig() *SX130xConfig

	GetExtraChannels() []ExtraChannel

	GetUplinkChannels() []int

	GetSubband() (int, error)

	SetSubband(int) error

	GetApiRegion() *api.GateWayRegion

	GetAdrRange() (int, int)

	// GetDefaultRssiTComp() *RssiTcomp

	// GetDefaultTxGainLut() *[]TxGainLutItem
}

func errRegionHasNoSubband(region string) error {
	return errors.Errorf("region %s has no subband", region)
}

func errFailToFindSubband(chs []int) error {
	return errors.Errorf("fail to find subband with channel: %v", chs)
}

func HasSubband(region string) (bool, error) {
	switch region {
	case "AU915":
		fallthrough
	case "AS923":
		fallthrough
	case "US915":
		fallthrough
	case "KR920":
		fallthrough
	case "CN470":
		return true, nil
	case "EU868":
		fallthrough
	case "IN865":
		fallthrough
	case "RU864":
		return false, nil
	}

	return false, errors.Errorf("unsupported region %s", region)
}

func TryFindSubband(region string, enableChannels []int) (int, error) {
	switch region {
	case "AU915":
	case "AS923":
	case "US915":
	case "KR920":
	case "CN470":
	}

	return 0, errors.Errorf("region %s has no subband", region)
}

func NewWithSubband(region string, config *SX130xConfig, subband int) (Band, error) {
	b := Band(nil)
	switch region {
	case "AU915":
	case "AS923":
	case "US915":
		b = &US915{
			SX130xConfig: config,
		}
	case "KR920":
	case "CN470":
	}
	if b == nil {
		return nil, errors.Errorf("unsupported region %s", region)
	}

	if err := b.SetSubband(subband); err != nil {
		return nil, err
	}
	return b, nil
}

func NewWithChannel(region string, config *SX130xConfig) (Band, error) {
	b := Band(nil)
	switch region {
	case "AU915":
	case "AS923":
	case "US915":
		b = &US915{
			SX130xConfig: config,
		}
	case "KR920":
	case "CN470":
	case "EU868":
		b = &EU868{
			SX130xConfig: config,
		}
	case "IN865":
	case "RU864":
	}

	if b == nil {
		return nil, errors.Errorf("unsupported region %s", region)
	}
	return b, nil
}

func SetChannel(c *SX130xConfig, a *api.EU868Config) *SX130xConfig {
	if a == nil {
		return &SX130xConfig{}
	}

	c.LorawanPublic = a.LoraWanPublic

	// c.Radio0.Enable = a.Radio_0.Enable
	// c.Radio0.Type = "SX1250" // a.Radio_0.Type
	// c.Radio0.SingleInputMode = a.Radio_0.SingleInputMode
	c.Radio0.Freq = a.Radio_0.Freq
	// c.Radio0.RssiOffset = a.Radio_0.RssiOffset
	// c.Radio0.TxEnable = a.Radio_0.TxEnable
	// c.Radio0.TxFreqMin = a.Radio_0.TxFreqMin
	// c.Radio0.TxFreqMax = a.Radio_0.TxFreqMax

	// c.Radio1.Enable = a.Radio_1.Enable
	// c.Radio1.Type = a.Radio_1.Type
	// c.Radio1.SingleInputMode = a.Radio_1.SingleInputMode
	c.Radio1.Freq = a.Radio_1.Freq
	// c.Radio1.RssiOffset = a.Radio_1.RssiOffset
	// c.Radio1.TxEnable = a.Radio_1.TxEnable

	c.ChanLoraStd.Enable = a.Chan_LoraStd.Enable
	c.ChanLoraStd.Radio = a.Chan_LoraStd.Radio
	c.ChanLoraStd.IF = a.Chan_LoraStd.If
	c.ChanLoraStd.Bandwidth = a.Chan_LoraStd.Bandwidth
	c.ChanLoraStd.SpreadFactor = a.Chan_LoraStd.SpreadFactor
	c.ChanLoraStd.ImplicitHdr = a.Chan_LoraStd.ImplicitHdr
	c.ChanLoraStd.Implicitpayloadlength = a.Chan_LoraStd.Implicitpayloadlength
	c.ChanLoraStd.ImplicitcrcEn = a.Chan_LoraStd.ImplicitcrcEn
	c.ChanLoraStd.Implicitcoderate = a.Chan_LoraStd.Implicitcoderate

	cms := []struct {
		c *ChanMultiSF
		a *api.EU868ChannelMultiSF
	}{
		{&c.ChanMultiSF0, a.ChanMultiSF_0},
		{&c.ChanMultiSF1, a.ChanMultiSF_1},
		{&c.ChanMultiSF2, a.ChanMultiSF_2},
		{&c.ChanMultiSF3, a.ChanMultiSF_3},
		{&c.ChanMultiSF4, a.ChanMultiSF_4},
		{&c.ChanMultiSF5, a.ChanMultiSF_5},
		{&c.ChanMultiSF6, a.ChanMultiSF_6},
		{&c.ChanMultiSF7, a.ChanMultiSF_7},
	}

	for _, ch := range cms {
		ch.c.Enable = ch.a.Enable
		ch.c.Radio = ch.a.Radio
		ch.c.IF = ch.a.Offset
	}

	return c
}

func (c *SX130xConfig) ApiRegion() *api.EU868Config {
	a := &api.EU868Config{
		LoraWanPublic: c.LorawanPublic,
		SyncWord:      52,
		Radio_0: &api.EU868Radio0{
			Enable:          c.Radio0.Enable,
			Type:            "SX1250", // c.Radio0.Type,
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
		},
		Radio_1: &api.EU868Radio1{
			Enable:          c.Radio1.Enable,
			Type:            "SX1250", // c.Radio1.Type,
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
		ChanMultiSF_0: &api.EU868ChannelMultiSF{},
		ChanMultiSF_1: &api.EU868ChannelMultiSF{},
		ChanMultiSF_2: &api.EU868ChannelMultiSF{},
		ChanMultiSF_3: &api.EU868ChannelMultiSF{},
		ChanMultiSF_4: &api.EU868ChannelMultiSF{},
		ChanMultiSF_5: &api.EU868ChannelMultiSF{},
		ChanMultiSF_6: &api.EU868ChannelMultiSF{},
		ChanMultiSF_7: &api.EU868ChannelMultiSF{},
	}

	for _, v := range c.TxGainLut {
		a.Radio_0.Txgainlut = append(a.Radio_0.Txgainlut, &api.TxGainLutItem{
			Rfpower: int64(v.RFPower),
			Pagain:  int64(v.PaGain),
			Pwridx:  int64(v.PwrIdx),
		})
	}

	cms := []struct {
		c *ChanMultiSF
		a *api.EU868ChannelMultiSF
	}{
		{&c.ChanMultiSF0, a.ChanMultiSF_0},
		{&c.ChanMultiSF1, a.ChanMultiSF_1},
		{&c.ChanMultiSF2, a.ChanMultiSF_2},
		{&c.ChanMultiSF3, a.ChanMultiSF_3},
		{&c.ChanMultiSF4, a.ChanMultiSF_4},
		{&c.ChanMultiSF5, a.ChanMultiSF_5},
		{&c.ChanMultiSF6, a.ChanMultiSF_6},
		{&c.ChanMultiSF7, a.ChanMultiSF_7},
	}

	for _, ch := range cms {
		ch.a.Enable = ch.c.Enable
		ch.a.Radio = ch.c.Radio
		ch.a.Offset = ch.c.IF
	}

	return a
}

func NewTxGainLutItem(a *api.TxGainLutItem) *TxGainLutItem {
	return &TxGainLutItem{
		RFPower: int(a.Rfpower),
		PaGain:  int(a.Pagain),
		PwrIdx:  int(a.Pwridx),
	}
}

func extraChannel(normalDr int, stdDr int, normalChan []*ChanMultiSF, stdChan *ChanLoraStd, defaultChanFreqs []int32, centerFreqs [2]int32) []ExtraChannel {
	ec := []ExtraChannel{}

	for _, ch := range normalChan {
		if ch.Enable {
			chfreq := centerFreqs[ch.Radio] + ch.IF
			for _, freq := range defaultChanFreqs {
				if freq == chfreq {
					goto next_loop
				}
			}
			ec = append(ec, ExtraChannel{
				Frequency: int(chfreq),
				MinDR:     0,
				MaxDR:     normalDr,
			})
		next_loop:
		}
	}

	if stdChan.Enable {
		ec = append(ec, ExtraChannel{
			Frequency: int(centerFreqs[stdChan.Radio] + stdChan.IF),
			MinDR:     int(stdDr),
			MaxDR:     int(stdDr),
		})
	}

	return ec
}
