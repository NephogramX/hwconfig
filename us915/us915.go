package us915

import (
	"errors"
	"fmt"

	"github.com/NephogramX/hwconfig"
)

type US915 struct {
	backend string
	subband int32
}

func NewBuilder() *US915 {
	return &US915{}
}

func (r *US915) SetBackend(b string) {
	r.backend = b
}

func (r *US915) SetCustomBand(c hwconfig.CustomBand) {
}

func (r *US915) SetSubband(fsb int32) {
	r.subband = fsb
}

func (r *US915) Build() (*hwconfig.Configs, error) {
	if r.subband < 0 || r.subband > 9 {
		return nil, errors.New(fmt.Sprint("unknow subband:", r.subband, " in US915"))
	}

	if !hwconfig.CheckBackend(r.backend) {
		return nil, errors.New(fmt.Sprint("unsupported backend: ", r.backend))
	}

	return &hwconfig.Configs{
		PacketForwarder: r.buildPacketForwarder(),
		GatewayBridge:   r.buildGatewayBridge(),
		NetworkServer:   r.buildNetworkServer(),
	}, nil
}

func (r *US915) buildPacketForwarder() *hwconfig.PacketForwarderConfig {
	if r.backend != hwconfig.SemtechUDP {
		return nil
	}

	var (
		Radio0Freq, Radio1Freq int32
		RadioFreqOffsets       [8]int32
	)

	if r.subband == 9 {
		Radio0Freq = 903000000 + 1600000*2
		Radio1Freq = 903000000 + 1600000*6
		RadioFreqOffsets = [8]int32{-3200000, -1600000, 0, 1600000, -3200000, -1600000, 0, 1600000}
	} else {
		Radio0Freq = 902300000 + 1600000*(r.subband-1) + 400000
		Radio1Freq = 902300000 + 1600000*(r.subband-1) + 1000000
		RadioFreqOffsets = [8]int32{-400000, -2000000, 0, 200000, -400000, -2000000, 0, 200000}
	}

	return hwconfig.FillPacketForwarder(&hwconfig.PacketForwarderConfig{
		SX130xConfig: hwconfig.SX130xConfig{
			Radio0: hwconfig.Radio0{
				SingleInputMode: false,
				Freq:            Radio0Freq,
				RssiOffset:      -215.4,
				TxFreqMin:       923000000,
				TxFreqMax:       928000000,
				TxGainLut: []hwconfig.TxGainLutItem{
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
			},
			Radio1: hwconfig.Radio1{
				SingleInputMode: false,
				Freq:            Radio1Freq,
				RssiOffset:      -215.4,
			},
			ChanMultiSF0: hwconfig.ChanMultiSF{
				Enable: true,
				Radio:  0,
				IF:     RadioFreqOffsets[0],
			},
			ChanMultiSF1: hwconfig.ChanMultiSF{
				Enable: true,
				Radio:  0,
				IF:     RadioFreqOffsets[1],
			},
			ChanMultiSF2: hwconfig.ChanMultiSF{
				Enable: true,
				Radio:  0,
				IF:     RadioFreqOffsets[2],
			},
			ChanMultiSF3: hwconfig.ChanMultiSF{
				Enable: true,
				Radio:  0,
				IF:     RadioFreqOffsets[3],
			},
			ChanMultiSF4: hwconfig.ChanMultiSF{
				Enable: true,
				Radio:  1,
				IF:     RadioFreqOffsets[4],
			},
			ChanMultiSF5: hwconfig.ChanMultiSF{
				Enable: true,
				Radio:  1,
				IF:     RadioFreqOffsets[5],
			},
			ChanMultiSF6: hwconfig.ChanMultiSF{
				Enable: true,
				Radio:  1,
				IF:     RadioFreqOffsets[6],
			},
			ChanMultiSF7: hwconfig.ChanMultiSF{
				Enable: true,
				Radio:  1,
				IF:     RadioFreqOffsets[7],
			},
			ChanLoraStd: hwconfig.ChanLoraStd{
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
			ChanLoraFSK: hwconfig.ChanLoraFSK{
				Enable:    false,
				Radio:     1,
				IF:        300000,
				Bandwidth: 125000,
				Datarate:  50000,
			},
		},
		GateWayConfig: hwconfig.GateWayConfig{
			BeaconPeriod:   0,
			BeaconFreqHZ:   869525000,
			BeaconFreqNB:   1,
			BeaconFreqStep: 0,
			BeaconDatarate: 9,
			BeaconBwHZ:     125000,
			BeaconPower:    14,
			BeaconInfodesc: 0,
		},
	})
}

func (r *US915) buildGatewayBridge() *hwconfig.GatewayBridgeConfig {
	b := &hwconfig.GbBackend{
		Type: r.backend,
	}

	var frequencies [8]int32

	if r.subband == 9 {
		for i := range frequencies {
			frequencies[i] = int32(903000000 + i*1600000)
		}
	} else {
		for i := range frequencies {
			frequencies[i] = 902300000 + (r.subband-1)*1600000 + int32(i)*200000
		}
	}

	switch r.backend {
	case hwconfig.BStation:
		b.SemtechUdp = nil
		b.BasicStation = &hwconfig.BasicStation{
			Bind:         "0.0.0.0:3001",
			Region:       "US915",
			FrequencyMin: 902300000,
			FrequencyMax: 927500000,
			Concentrators: hwconfig.Concentrators{
				MultiSF: hwconfig.MultiSF{
					Frequencies: frequencies[:],
				},
			},
		}
	case hwconfig.SemtechUDP:
		b.BasicStation = nil
		b.SemtechUdp = &hwconfig.SemtechUdp{
			UdpBind: "0.0.0.0:1700",
		}
	}
	return hwconfig.NewGatewayBridge(b)
}

func (r *US915) buildNetworkServer() *hwconfig.NetworkServerConfig {
	return hwconfig.NewNetworkServerConfig("US915")
}
