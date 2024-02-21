package band

import (
	"github.com/NephogramX/hwconfig/configfile"
)

type EU868Band struct {
	centerFrequency int32
	frequencyShift  [5]int32
}

type CustomBand struct {
}

func NewBandEU868(centerFrequency int32, frequencyShift [5]int32) (*EU868Band, error) {
	return &EU868Band{
		centerFrequency: centerFrequency,
		frequencyShift:  frequencyShift,
	}, nil
}

func (b *EU868Band) String() string {
	return "EU868"
}

func (b EU868Band) GetChannelSettings() *configfile.Channel {
	return &configfile.Channel{
		RaidoCneterFrequency: [2]int32{b.centerFrequency, 868500000},
		MinTxFrequency:       863000000,
		MaxTxFrequency:       870000000,
		RssiOffset:           -215.4,
		ChanMultiSF: [8]configfile.ChanMultiSF{
			{Enable: true, Radio: 1, IF: -400000},
			{Enable: true, Radio: 1, IF: -200000},
			{Enable: true, Radio: 1, IF: 0},
			{Enable: true, Radio: 0, IF: b.frequencyShift[0]},
			{Enable: true, Radio: 0, IF: b.frequencyShift[1]},
			{Enable: true, Radio: 0, IF: b.frequencyShift[2]},
			{Enable: true, Radio: 0, IF: b.frequencyShift[3]},
			{Enable: true, Radio: 0, IF: b.frequencyShift[4]},
		},
		ChanLoRaStd: configfile.ChanLoRaStd{
			ChanMultiSF: configfile.ChanMultiSF{Enable: false, Radio: 1},
			// Bandwidth:   250000, SpreadFactor: 7,
		},
		ChanLoRaFsk: configfile.ChanLoRaFSK{
			ChanMultiSF: configfile.ChanMultiSF{Enable: false, Radio: 1},
		},
		TxGainLutItem: []configfile.TxGainLutItem{
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

func (b EU868Band) GetExtraChannels() *[]configfile.ExtraChannels {
	ec := make([]configfile.ExtraChannels, 5)

	for i := range b.frequencyShift {
		ec[i].Frequency = b.centerFrequency + b.frequencyShift[i]
		ec[i].MinDr = 0
		ec[i].MaxDr = 5
	}

	return &ec
}

func (b EU868Band) GetUplinkChannels() *[]int32 {
	return nil
}
