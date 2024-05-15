package band

import (
	"fmt"

	"github.com/pkg/errors"

	cf "github.com/NephogramX/hwconfig/configfile"
)

type US915Band struct {
	cf.Channel
	uplinkChannels []int32
}

func GetUS915DefaultChannelSettings() *cf.Channel {
	b, _ := NewBandUS915(2)
	return b.GetChannelSettings()
}

func NewBandUS915(subbandIndex int32) (*US915Band, error) {
	if subbandIndex < 1 || subbandIndex > 8 {
		return nil, errors.New(fmt.Sprint("unsupported subband index:", subbandIndex))
	}
	ch := [9]int32{}
	for i := range ch {
		ch[i] = (subbandIndex-1)*8 + int32(i)
	}
	ch[8] = 64 + (subbandIndex - 1)
	return &US915Band{
		Channel: cf.Channel{
			Radio0: cf.Radio0{
				Enable:     true,
				Type:       "SX1250",
				Freq:       902300000 + 1600000*(subbandIndex-1) + 400000,
				RssiOffset: rssiOffset900MHz(),
				RssiTcomp:  *rssiTcomp(),
				TxEnable:   true,
				TxFreqMin:  902000000,
				TxFreqMax:  928000000,
				TxGainLut:  txGainLut900MHz(),
			},
			Radio1: cf.Radio1{
				Enable:     true,
				Type:       "SX1250",
				Freq:       902300000 + 1600000*(subbandIndex-1) + 1100000,
				RssiOffset: rssiOffset900MHz(),
				RssiTcomp:  *rssiTcomp(),
				TxEnable:   true,
			},
			ChanMultiSFAll: *ChanMultiSFAll(),
			ChanMultiSF: [8]*cf.ChanMultiSF{
				{Enable: true, Radio: 0, IF: -400000},
				{Enable: true, Radio: 0, IF: -200000},
				{Enable: true, Radio: 0, IF: 0},
				{Enable: true, Radio: 0, IF: 200000},
				{Enable: true, Radio: 1, IF: -300000},
				{Enable: true, Radio: 1, IF: 0},
				{Enable: true, Radio: 1, IF: 100000},
				{Enable: true, Radio: 1, IF: 300000},
			},
			ChanLoraStd: cf.ChanLoraStd{
				Enable:                true,
				Radio:                 0,
				IF:                    300000,
				Bandwidth:             500000,
				SpreadFactor:          8,
				ImplicitHdr:           false,
				Implicitpayloadlength: 17,
				ImplicitcrcEn:         false,
				Implicitcoderate:      1,
			},
			ChanLoraFSK: cf.ChanLoraFSK{
				Enable:    false,
				Radio:     1,
				IF:        300000,
				Bandwidth: 125000,
				Datarate:  50000,
			},
		},
		uplinkChannels: ch[:],
	}, nil
}

func (b *US915Band) String() string {
	return "US915"
}

func (b US915Band) GetChannelSettings() *cf.Channel {
	return &b.Channel
}

func (b US915Band) GetExtraChannels() *[]cf.ExtraChannels {
	return nil
}

func (b US915Band) GetUplinkChannels() *[]int32 {
	return &b.uplinkChannels
}
