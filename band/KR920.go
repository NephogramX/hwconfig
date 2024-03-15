package band

import (
	"fmt"

	"github.com/pkg/errors"

	cf "github.com/NephogramX/hwconfig/configfile"
)

type KR920Band struct {
	subbandIndex int32
}

func NewBandKR920(subbandIndex int32) (*KR920Band, error) {
	if subbandIndex < 0 || subbandIndex > 7 {
		return nil, errors.New(fmt.Sprint("unsupported subband index:", subbandIndex))
	}
	return &KR920Band{
		subbandIndex: subbandIndex,
	}, nil
}

func (b *KR920Band) String() string {
	return "KR920"
}

func (b KR920Band) GetChannelSettings() *cf.Channel {
	return &cf.Channel{
		RaidoCneterFrequency: [2]int32{
			922100000 + 1600000*(b.subbandIndex-1) + 400000,
			922100000 + 1600000*(b.subbandIndex-1) + 1100000,
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
			{Enable: false},
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

func (b KR920Band) GetExtraChannels() *[]cf.ExtraChannels {
	ec := make([]cf.ExtraChannels, 5)

	for i := range ec {
		ec[i].Frequency = int32(922700000 + i*200000)
		ec[i].MinDr = 0
		ec[i].MaxDr = 5
	}

	return &ec
}

func (b KR920Band) GetUplinkChannels() *[]int32 {
	return nil
}
