package band

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/NephogramX/hwconfig/configfile/gb"
	"github.com/NephogramX/hwconfig/configfile/ns"
	"github.com/NephogramX/hwconfig/configfile/pf"
)

type CN470Band struct {
	subbandIndex int32
}

func NewBandCN470(subbandIndex int32) (*CN470Band, error) {
	if subbandIndex < 0 || subbandIndex > 9 {
		return nil, errors.New(fmt.Sprint("unsupported subband index:", subbandIndex))
	}
	return &CN470Band{
		subbandIndex: subbandIndex,
	}, nil
}

func (b *CN470Band) HandleUdpPacketForwarder(cf *pf.UdpPacketForwarder) error {

	Radio0Freq := 470300000 + 1600000*(b.subbandIndex-1) + 1100000
	Radio1Freq := 470300000 + 1600000*(b.subbandIndex-1) + 300000
	RadioFreqOffsets := [8]int32{-300000, -100000, 100000, 300000, -300000, -100000, 100000, 300000}

	cf.SX130xConfig.Radio0.SingleInputMode = true
	cf.SX130xConfig.Radio0.Freq = Radio0Freq
	cf.SX130xConfig.Radio0.RssiOffset = -207.0
	cf.SX130xConfig.Radio0.TxFreqMin = 500000000
	cf.SX130xConfig.Radio0.TxFreqMax = 510000000
	cf.SX130xConfig.Radio0.TxGainLut = []pf.TxGainLutItem{
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
	}
	cf.SX130xConfig.Radio1.SingleInputMode = true
	cf.SX130xConfig.Radio1.Freq = Radio1Freq
	cf.SX130xConfig.Radio1.RssiOffset = -207.0
	cf.SX130xConfig.ChanMultiSF0 = pf.ChanMultiSF{
		Enable: true,
		Radio:  0,
		IF:     RadioFreqOffsets[0],
	}
	cf.SX130xConfig.ChanMultiSF1 = pf.ChanMultiSF{
		Enable: true,
		Radio:  0,
		IF:     RadioFreqOffsets[1],
	}
	cf.SX130xConfig.ChanMultiSF2 = pf.ChanMultiSF{
		Enable: true,
		Radio:  0,
		IF:     RadioFreqOffsets[2],
	}
	cf.SX130xConfig.ChanMultiSF3 = pf.ChanMultiSF{
		Enable: true,
		Radio:  0,
		IF:     RadioFreqOffsets[3],
	}
	cf.SX130xConfig.ChanMultiSF4 = pf.ChanMultiSF{
		Enable: true,
		Radio:  1,
		IF:     RadioFreqOffsets[4],
	}
	cf.SX130xConfig.ChanMultiSF5 = pf.ChanMultiSF{
		Enable: true,
		Radio:  1,
		IF:     RadioFreqOffsets[5],
	}
	cf.SX130xConfig.ChanMultiSF6 = pf.ChanMultiSF{
		Enable: true,
		Radio:  1,
		IF:     RadioFreqOffsets[6],
	}
	cf.SX130xConfig.ChanMultiSF7 = pf.ChanMultiSF{
		Enable: true,
		Radio:  1,
		IF:     RadioFreqOffsets[7],
	}
	cf.SX130xConfig.ChanLoraStd = pf.ChanLoraStd{
		Enable:                true,
		Radio:                 1,
		IF:                    -200000,
		Bandwidth:             250000,
		SpreadFactor:          7,
		ImplicitHdr:           false,
		Implicitpayloadlength: 17,
		ImplicitcrcEn:         false,
		Implicitcoderate:      1,
	}
	cf.SX130xConfig.ChanLoraFSK = pf.ChanLoraFSK{
		Enable:    false,
		Radio:     1,
		IF:        300000,
		Bandwidth: 125000,
		Datarate:  50000,
	}

	return nil
}

func (b *CN470Band) HandleGatewayBridge(cf *gb.GatewayBridge) error {
	return nil
}
func (b *CN470Band) HandleNetworkServer(cf *ns.NetworkServer) error {
	return nil
}
