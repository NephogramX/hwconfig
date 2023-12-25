package cn470

import (
	"errors"
	"fmt"

	"github.com/NephogramX/hwconfig"
)

type CN470 struct {
	backend int32
	subband int32
}

func NewCN470() *CN470 {
	return &CN470{}
}

func (r *CN470) SetSubband(fsb int32) {
	r.subband = fsb
}

func (r *CN470) SetCustomBand(c hwconfig.CustomBand) {
}

func (r *CN470) Build() (*hwconfig.Configs, error) {
	if r.subband < 0 || r.subband > 9 {
		return nil, errors.New(fmt.Sprint("unknow subband:", r.subband, " in CN470"))
	}

	if r.backend != hwconfig.BasicStationBackend && r.backend != hwconfig.PacketForwarderBackend {
		return nil, errors.New("unkown gateway backend")
	}

	return &hwconfig.Configs{
		PacketForwarder: r.buildPacketForwarder(),
		GatewayBridge:   r.buildGatewayBridge(),
		NetworkServer:   r.buildNetworkServer(),
	}, nil
}

func (r *CN470) buildPacketForwarder() *hwconfig.PacketForwarderConfig {
	if r.backend != hwconfig.PacketForwarderBackend {
		return nil
	}

	var (
		Radio0Freq, Radio1Freq int32
		RadioFreqOffsets       [8]int32
	)

	// TODO: Config
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
				Freq:      Radio0Freq,
				TxFreqMin: 923000000,
				TxFreqMax: 928000000,
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
				Freq: Radio1Freq,
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
				IF:     RadioFreqOffsets[0],
			},
			ChanMultiSF5: hwconfig.ChanMultiSF{
				Enable: true,
				Radio:  1,
				IF:     RadioFreqOffsets[1],
			},
			ChanMultiSF6: hwconfig.ChanMultiSF{
				Enable: true,
				Radio:  1,
				IF:     RadioFreqOffsets[2],
			},
			ChanMultiSF7: hwconfig.ChanMultiSF{
				Enable: true,
				Radio:  1,
				IF:     RadioFreqOffsets[3],
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
			BeaconFreqNB:   0,
			BeaconFreqStep: 0,
			BeaconDatarate: 9,
			BeaconBwHZ:     125000,
			BeaconPower:    14,
			BeaconInfodesc: 0,
		},
	})
}

func (r *CN470) buildGatewayBridge() *hwconfig.GatewayBridgeConfig {
	var b *hwconfig.GbBackend

	switch r.backend {
	case hwconfig.BasicStationBackend:
		b = &hwconfig.GbBackend{
			Type: "basic_station",
			BasicStation: hwconfig.BasicStation{
				Bind: "0.0.0.0:3001",
			},
		}
	case hwconfig.PacketForwarderBackend:
		b = &hwconfig.GbBackend{
			Type: "semtech_udp",
			SemtechUdp: hwconfig.SemtechUdp{
				UdpBind: "0.0.0.0:1700",
			},
		}
	}

	return hwconfig.NewGatewayBridge(b)
}

func (r *CN470) buildNetworkServer() *hwconfig.NetworkServerConfig {
	return hwconfig.NewNetworkServerConfig("CN470")
}
