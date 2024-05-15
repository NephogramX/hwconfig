package band

import (
	cf "github.com/NephogramX/hwconfig/configfile"
)

type EU868Band struct {
	cf.Channel
}

func GetEU868DefaultChannelSettings() *cf.Channel {
	return &cf.Channel{
		Radio0: cf.Radio0{
			Enable:     true,
			Type:       "SX1250",
			Freq:       867500000,
			RssiOffset: rssiOffset900MHz(),
			RssiTcomp:  *rssiTcomp(),
			TxEnable:   true,
			TxFreqMin:  863000000,
			TxFreqMax:  870000000,
			TxGainLut:  txGainLut900MHz(),
		},
		Radio1: cf.Radio1{
			Enable:     true,
			Type:       "SX1250",
			Freq:       868500000,
			RssiOffset: rssiOffset900MHz(),
			RssiTcomp:  *rssiTcomp(),
			TxEnable:   true,
		},
		ChanMultiSFAll: *ChanMultiSFAll(),
		ChanMultiSF: [8]*cf.ChanMultiSF{
			{Enable: true, Radio: 1, IF: -400000},
			{Enable: true, Radio: 1, IF: -200000},
			{Enable: true, Radio: 1, IF: 0},
			{Enable: true, Radio: 0, IF: -400000},
			{Enable: true, Radio: 0, IF: -200000},
			{Enable: true, Radio: 0, IF: 0},
			{Enable: true, Radio: 0, IF: 200000},
			{Enable: true, Radio: 0, IF: 400000},
		},
		ChanLoraStd: cf.ChanLoraStd{
			Enable:                true,
			Radio:                 1,
			IF:                    -200000,
			Bandwidth:             250000,
			SpreadFactor:          7,
			ImplicitHdr:           false,
			Implicitpayloadlength: 17,
			ImplicitcrcEn:         false,
			Implicitcoderate:      1,
		},
		ChanLoraFSK: cf.ChanLoraFSK{
			Enable:    true,
			Radio:     1,
			IF:        300000,
			Bandwidth: 125000,
			Datarate:  50000,
		},
	}
}

func NewBandEU868(c *cf.Channel) (*EU868Band, error) {
	return &EU868Band{
		Channel: *c,
	}, nil
}

func (b *EU868Band) String() string {
	return "EU868"
}

func (b *EU868Band) GetChannelSettings() *cf.Channel {
	return &b.Channel
}

func (b *EU868Band) GetExtraChannels() *[]cf.ExtraChannels {
	return getExtraChannels(5, 6, 7, []int32{868100000, 868300000, 868500000}, &b.Channel)
}

func (b *EU868Band) GetUplinkChannels() *[]int32 {
	return nil
}
