package band

import cf "github.com/NephogramX/hwconfig/configfile"

type IN865Band struct {
	cf.Channel
}

func GetIN865DefaultChannelSettings() *cf.Channel {
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

func NewBandIN865() (*IN865Band, error) {
	return &IN865Band{}, nil
}

func (b *IN865Band) String() string {
	return "IN865"
}

func (b *IN865Band) GetChannelSettings() *cf.Channel {
	// return &cf.Channel{
	// 	RaidoCneterFrequency: [2]int32{865200000, 866385000},
	// 	MinTxFrequency:       865000000,
	// 	MaxTxFrequency:       867000000,
	// 	RssiOffset:           -215.4,
	// 	ChanMultiSF: [8]cf.ChanMultiSF{
	// 		{Enable: true, Radio: 0, IF: -137500},
	// 		{Enable: true, Radio: 0, IF: 202500},
	// 		{Enable: true, Radio: 1, IF: -400000},
	// 		{Enable: true, Radio: 0, IF: 32500},
	// 		{Enable: true, Radio: 1, IF: -200000},
	// 		{Enable: true, Radio: 1, IF: 0},
	// 		{Enable: true, Radio: 1, IF: 200000},
	// 		{Enable: true, Radio: 1, IF: 400000},
	// 	},
	// 	ChanLoRaStd: cf.ChanLoRaStd{
	// 		ChanMultiSF: cf.ChanMultiSF{Enable: false},
	// 	},
	// 	ChanLoRaFsk: cf.ChanLoRaFSK{
	// 		ChanMultiSF: cf.ChanMultiSF{Enable: false},
	// 	},
	// 	TxGainLutItem: []cf.TxGainLutItem{
	// 		{RFPower: 12, PaGain: 0, PwrIdx: 15},
	// 		{RFPower: 13, PaGain: 0, PwrIdx: 16},
	// 		{RFPower: 14, PaGain: 0, PwrIdx: 17},
	// 		{RFPower: 15, PaGain: 0, PwrIdx: 19},
	// 		{RFPower: 16, PaGain: 0, PwrIdx: 20},
	// 		{RFPower: 17, PaGain: 0, PwrIdx: 22},
	// 		{RFPower: 18, PaGain: 1, PwrIdx: 1},
	// 		{RFPower: 19, PaGain: 1, PwrIdx: 2},
	// 		{RFPower: 20, PaGain: 1, PwrIdx: 3},
	// 		{RFPower: 21, PaGain: 1, PwrIdx: 4},
	// 		{RFPower: 22, PaGain: 1, PwrIdx: 5},
	// 		{RFPower: 23, PaGain: 1, PwrIdx: 6},
	// 		{RFPower: 24, PaGain: 1, PwrIdx: 7},
	// 		{RFPower: 25, PaGain: 1, PwrIdx: 9},
	// 		{RFPower: 26, PaGain: 1, PwrIdx: 11},
	// 		{RFPower: 27, PaGain: 1, PwrIdx: 14},
	// 	},
	// }
	return nil
}

func (b *IN865Band) GetExtraChannels() *[]cf.ExtraChannels {
	ec := make([]cf.ExtraChannels, 5)
	ecList := [5]int32{865232500, 865572500, 865742500, 865912500, 866155000}

	for i := range ecList {
		ec[i].Frequency = ecList[i]
		ec[i].MinDr = 0
		ec[i].MaxDr = 5
	}

	return &ec
}

func (b *IN865Band) GetUplinkChannels() *[]int32 {
	return nil
}
