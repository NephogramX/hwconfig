package band

import (
	"fmt"

	"github.com/pkg/errors"

	cf "github.com/NephogramX/hwconfig/configfile"
)

type AU915Band struct {
	subbandIndex int32
}

func NewBandAU915(subbandIndex int32) (*AU915Band, error) {
	if subbandIndex < 0 || subbandIndex > 8 {
		return nil, errors.New(fmt.Sprint("unsupported subband index:", subbandIndex))
	}
	return &AU915Band{
		subbandIndex: subbandIndex,
	}, nil
}

func (b *AU915Band) String() string {
	return "AU915"
}

func (b AU915Band) GetChannelSettings() *cf.Channel {
	return &cf.Channel{
		RaidoCneterFrequency: [2]int32{
			915200000 + 1600000*(b.subbandIndex-1) + 400000,
			915200000 + 1600000*(b.subbandIndex-1) + 1100000,
		},
		MinTxFrequency: 915000000,
		MaxTxFrequency: 928000000,
		RssiOffset:     -215.4,
		ChanMultiSF: [8]cf.ChanMultiSF{
			{Enable: true, Radio: 0, IF: -400000},
			{Enable: true, Radio: 0, IF: -200000},
			{Enable: true, Radio: 0, IF: 0},
			{Enable: true, Radio: 0, IF: 200000},
			{Enable: true, Radio: 1, IF: -300000},
			{Enable: true, Radio: 1, IF: -100000},
			{Enable: true, Radio: 1, IF: 100000},
			{Enable: true, Radio: 1, IF: 300000},
		},
		ChanLoRaStd: cf.ChanLoRaStd{
			ChanMultiSF:  cf.ChanMultiSF{Enable: true, Radio: 0, IF: 300000},
			Bandwidth:    500000,
			SpreadFactor: 8,
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

func (b AU915Band) GetExtraChannels() *[]cf.ExtraChannels {
	return nil
}

func (b AU915Band) GetUplinkChannels() *[]int32 {
	var ch int32 = (b.subbandIndex - 1) * 8
	return &[]int32{ch, ch + 1, ch + 2, ch + 3, ch + 4, ch + 5, ch + 6, ch + 7, 64 + b.subbandIndex}
}
