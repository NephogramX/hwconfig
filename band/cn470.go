package band

import (
	"fmt"

	cf "github.com/NephogramX/hwconfig/configfile"
	"github.com/pkg/errors"
)

type CN470Band struct {
	subbandIndex int32
}

func NewBandCN470(subbandIndex int32) (*CN470Band, error) {
	if subbandIndex < 0 || subbandIndex > 12 {
		return nil, errors.New(fmt.Sprint("unsupported subband index:", subbandIndex))
	}
	return &CN470Band{
		subbandIndex: subbandIndex,
	}, nil
}

func (b *CN470Band) String() string {
	return "CN470"
}

func (b CN470Band) GetChannelSettings() *cf.Channel {
	return &cf.Channel{
		RaidoCneterFrequency: [2]int32{
			470300000 + 1600000*(b.subbandIndex-1) + 1100000,
			470300000 + 1600000*(b.subbandIndex-1) + 300000,
		},
		MinTxFrequency: 470000000,
		MaxTxFrequency: 510000000,
		RssiOffset:     -207.0,
		ChanMultiSF: [8]cf.ChanMultiSF{
			{Enable: true, Radio: 0, IF: -300000},
			{Enable: true, Radio: 0, IF: -100000},
			{Enable: true, Radio: 0, IF: 100000},
			{Enable: true, Radio: 0, IF: 300000},
			{Enable: true, Radio: 1, IF: -300000},
			{Enable: true, Radio: 1, IF: -100000},
			{Enable: true, Radio: 1, IF: 100000},
			{Enable: true, Radio: 1, IF: 300000},
		},
		ChanLoRaStd: cf.ChanLoRaStd{
			ChanMultiSF: cf.ChanMultiSF{Enable: true, Radio: 1, IF: -200000},
			Bandwidth:   250000, SpreadFactor: 7,
		},
		ChanLoRaFsk: cf.ChanLoRaFSK{
			ChanMultiSF: cf.ChanMultiSF{Enable: false, Radio: 1},
		},
		TxGainLutItem: []cf.TxGainLutItem{
			{RFPower: -6, PaGain: 0, PwrIdx: 0},
			{RFPower: -3, PaGain: 0, PwrIdx: 1},
			{RFPower: 0, PaGain: 0, PwrIdx: 2},
			{RFPower: 3, PaGain: 1, PwrIdx: 3},
			{RFPower: 6, PaGain: 1, PwrIdx: 4},
			{RFPower: 10, PaGain: 1, PwrIdx: 5},
			{RFPower: 11, PaGain: 1, PwrIdx: 6},
			{RFPower: 12, PaGain: 1, PwrIdx: 7},
			{RFPower: 13, PaGain: 1, PwrIdx: 8},
			{RFPower: 14, PaGain: 1, PwrIdx: 9},
			{RFPower: 16, PaGain: 1, PwrIdx: 10},
			{RFPower: 20, PaGain: 1, PwrIdx: 11},
			{RFPower: 23, PaGain: 1, PwrIdx: 12},
			{RFPower: 25, PaGain: 1, PwrIdx: 13},
			{RFPower: 26, PaGain: 1, PwrIdx: 14},
			{RFPower: 27, PaGain: 1, PwrIdx: 15},
		},
	}
}

func (b CN470Band) GetExtraChannels() *[]cf.ExtraChannels {
	return nil
}

func (b CN470Band) GetUplinkChannels() *[]int32 {
	var ch int32 = (b.subbandIndex - 1) * 8
	return &[]int32{ch, ch + 1, ch + 2, ch + 3, ch + 4, ch + 5, ch + 6, ch + 7}
}
