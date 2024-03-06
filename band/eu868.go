package band

import (
	cf "github.com/NephogramX/hwconfig/configfile"
)

type EU868Band struct {
	CenterFrequency int32
	FrequencyShift  []int32
}

func NewBandEU868(centerFrequency int32, frequencyShift []int32) (*EU868Band, error) {
	return &EU868Band{
		CenterFrequency: centerFrequency,
		FrequencyShift:  frequencyShift,
	}, nil
}

func (b *EU868Band) String() string {
	return "EU868"
}

func (b *EU868Band) GetChannelSettings() *cf.Channel {
	fs := [5]struct {
		Enable         bool
		FrequencyShift int32
	}{}

	for i, f := range b.FrequencyShift {
		if i >= 5 {
			break
		}
		fs[i].Enable = true
		fs[i].FrequencyShift = f
	}

	return &cf.Channel{
		RaidoCneterFrequency: [2]int32{b.CenterFrequency, 868500000},
		MinTxFrequency:       863000000,
		MaxTxFrequency:       870000000,
		RssiOffset:           -215.4,
		ChanMultiSF: [8]cf.ChanMultiSF{
			{Enable: true, Radio: 1, IF: -400000},
			{Enable: true, Radio: 1, IF: -200000},
			{Enable: true, Radio: 1, IF: 0},
			{Enable: fs[0].Enable, Radio: 0, IF: fs[0].FrequencyShift},
			{Enable: fs[1].Enable, Radio: 0, IF: fs[1].FrequencyShift},
			{Enable: fs[2].Enable, Radio: 0, IF: fs[2].FrequencyShift},
			{Enable: fs[3].Enable, Radio: 0, IF: fs[3].FrequencyShift},
			{Enable: fs[4].Enable, Radio: 0, IF: fs[4].FrequencyShift},
		},
		ChanLoRaStd: cf.ChanLoRaStd{
			ChanMultiSF:  cf.ChanMultiSF{Enable: true, Radio: 1},
			Bandwidth:    250000,
			SpreadFactor: 7,
		},
		ChanLoRaFsk: cf.ChanLoRaFSK{
			ChanMultiSF: cf.ChanMultiSF{Enable: false, Radio: 1},
		},
		TxGainLutItem: []cf.TxGainLutItem{
			{RFPower: 12, PaGain: 1, PwrIdx: 4},
			{RFPower: 13, PaGain: 1, PwrIdx: 5},
			{RFPower: 14, PaGain: 1, PwrIdx: 6},
			{RFPower: 15, PaGain: 1, PwrIdx: 7},
			{RFPower: 16, PaGain: 1, PwrIdx: 8},
			{RFPower: 17, PaGain: 1, PwrIdx: 9},
			{RFPower: 18, PaGain: 1, PwrIdx: 10},
			{RFPower: 19, PaGain: 1, PwrIdx: 11},
			{RFPower: 20, PaGain: 1, PwrIdx: 12},
			{RFPower: 21, PaGain: 1, PwrIdx: 13},
			{RFPower: 22, PaGain: 1, PwrIdx: 14},
			{RFPower: 23, PaGain: 1, PwrIdx: 16},
			{RFPower: 24, PaGain: 1, PwrIdx: 17},
			{RFPower: 25, PaGain: 1, PwrIdx: 18},
			{RFPower: 26, PaGain: 1, PwrIdx: 19},
			{RFPower: 27, PaGain: 1, PwrIdx: 22},
		},
	}
}

func (b *EU868Band) GetExtraChannels() *[]cf.ExtraChannels {
	ec := make([]cf.ExtraChannels, 5)

	for i := range b.FrequencyShift {
		ec[i].Frequency = b.CenterFrequency + b.FrequencyShift[i]
		ec[i].MinDr = 0
		ec[i].MaxDr = 5
	}

	return &ec
}

func (b *EU868Band) GetUplinkChannels() *[]int32 {
	return nil
}
