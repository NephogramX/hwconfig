package band

import (
	"fmt"

	cf "github.com/NephogramX/hwconfig/configfile"
	"github.com/pkg/errors"
)

type CN470Band struct {
	cf.Channel
	uplinkChannels []int32
}

func GetCN470DefaultChannelSettings() *cf.Channel {
	b, _ := NewBandCN470(2)
	return b.GetChannelSettings()
}

func NewBandCN470(subbandIndex int32) (*CN470Band, error) {
	if subbandIndex < 1 || subbandIndex > 12 {
		return nil, errors.New(fmt.Sprint("unsupported subband index:", subbandIndex))
	}

	ch := [8]int32{}
	for i := range ch {
		ch[i] = (subbandIndex-1)*8 + int32(i)
	}

	return &CN470Band{
		Channel: cf.Channel{
			Radio0: cf.Radio0{
				Enable:     true,
				Type:       "SX1250",
				Freq:       470300000 + 1600000*(subbandIndex-1) + 1100000,
				RssiOffset: rssiOffset500MHz(),
				RssiTcomp:  *rssiTcomp(),
				TxEnable:   true,
				TxFreqMin:  470000000,
				TxFreqMax:  510000000,
				TxGainLut:  txGainLut500MHz(),
			},
			Radio1: cf.Radio1{
				Enable:     true,
				Type:       "SX1250",
				Freq:       470300000 + 1600000*(subbandIndex-1) + 300000,
				RssiOffset: rssiOffset500MHz(),
				RssiTcomp:  *rssiTcomp(),
				TxEnable:   true,
			},
			ChanMultiSFAll: *ChanMultiSFAll(),
			ChanMultiSF: [8]*cf.ChanMultiSF{
				{Enable: true, Radio: 0, IF: -300000},
				{Enable: true, Radio: 0, IF: -100000},
				{Enable: true, Radio: 0, IF: 100000},
				{Enable: true, Radio: 0, IF: 300000},
				{Enable: true, Radio: 1, IF: -300000},
				{Enable: true, Radio: 1, IF: -100000},
				{Enable: true, Radio: 1, IF: 100000},
				{Enable: true, Radio: 1, IF: 300000},
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
		},
		uplinkChannels: ch[:],
	}, nil
}

func (b *CN470Band) String() string {
	return "CN470"
}

func (b CN470Band) GetChannelSettings() *cf.Channel {
	return &b.Channel
}

func (b CN470Band) GetExtraChannels() *[]cf.ExtraChannels {
	return nil
}

func (b CN470Band) GetUplinkChannels() *[]int32 {
	return &b.uplinkChannels
}
