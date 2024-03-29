package band

import cf "github.com/NephogramX/hwconfig/configfile"

type IN865Band struct {
}

func NewBandIN865() (*IN865Band, error) {
	return &IN865Band{}, nil
}

func (b *IN865Band) String() string {
	return "IN865"
}

func (b *IN865Band) GetChannelSettings() *cf.Channel {
	return &cf.Channel{
		RaidoCneterFrequency: [2]int32{865200000, 866385000},
		MinTxFrequency:       865000000,
		MaxTxFrequency:       867000000,
		RssiOffset:           -215.4,
		ChanMultiSF: [8]cf.ChanMultiSF{
			{Enable: true, Radio: 0, IF: -137500},
			{Enable: true, Radio: 0, IF: 202500},
			{Enable: true, Radio: 1, IF: -400000},
			{Enable: true, Radio: 0, IF: 32500},
			{Enable: true, Radio: 1, IF: -200000},
			{Enable: true, Radio: 1, IF: 0},
			{Enable: true, Radio: 1, IF: 200000},
			{Enable: true, Radio: 1, IF: 400000},
		},
		ChanLoRaStd: cf.ChanLoRaStd{
			ChanMultiSF: cf.ChanMultiSF{Enable: false},
		},
		ChanLoRaFsk: cf.ChanLoRaFSK{
			ChanMultiSF: cf.ChanMultiSF{Enable: false},
		},
		TxGainLutItem: []cf.TxGainLutItem{
			{RFPower: 12, PaGain: 0, PwrIdx: 15},
			{RFPower: 13, PaGain: 0, PwrIdx: 16},
			{RFPower: 14, PaGain: 0, PwrIdx: 17},
			{RFPower: 15, PaGain: 0, PwrIdx: 19},
			{RFPower: 16, PaGain: 0, PwrIdx: 20},
			{RFPower: 17, PaGain: 0, PwrIdx: 22},
			{RFPower: 18, PaGain: 1, PwrIdx: 1},
			{RFPower: 19, PaGain: 1, PwrIdx: 2},
			{RFPower: 20, PaGain: 1, PwrIdx: 3},
			{RFPower: 21, PaGain: 1, PwrIdx: 4},
			{RFPower: 22, PaGain: 1, PwrIdx: 5},
			{RFPower: 23, PaGain: 1, PwrIdx: 6},
			{RFPower: 24, PaGain: 1, PwrIdx: 7},
			{RFPower: 25, PaGain: 1, PwrIdx: 9},
			{RFPower: 26, PaGain: 1, PwrIdx: 11},
			{RFPower: 27, PaGain: 1, PwrIdx: 14},
		},
	}
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
