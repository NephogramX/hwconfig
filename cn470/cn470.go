package cn470

// import (
// 	"errors"
// 	"fmt"

// 	"github.com/NephogramX/hwconfig"
// )

// type CN470 struct {
// 	backend int32
// 	subband int32
// 	Builder hwconfig.Builder
// }

// func NewBuilder() *CN470 {
// 	return &CN470{}
// }

// func (r *CN470) SetBackend(b int32) {
// 	r.backend = b
// }

// func (r *CN470) SetCustomBand(c hwconfig.CustomBand) {
// }

// func (r *CN470) SetSubband(fsb int32) {
// 	r.subband = fsb
// }

// func (r *CN470) Build() (*hwconfig.Configs, error) {
// 	if r.subband < 0 || r.subband > 9 {
// 		return nil, errors.New(fmt.Sprint("unknow subband:", r.subband, " in CN470"))
// 	}

// 	if r.backend != hwconfig.BasicStationBackend && r.backend != hwconfig.PacketForwarderBackend {
// 		return nil, errors.New("unkown gateway backend")
// 	}

// 	return &hwconfig.Configs{
// 		PacketForwarder: r.buildPacketForwarder(),
// 		GatewayBridge:   r.buildGatewayBridge(),
// 		NetworkServer:   r.buildNetworkServer(),
// 	}, nil
// }

// func (r *CN470) buildPacketForwarder() *hwconfig.PacketForwarderConfig {
// 	if r.backend != hwconfig.PacketForwarderBackend {
// 		return nil
// 	}

// 	var (
// 		Radio0Freq, Radio1Freq int32
// 		RadioFreqOffsets       [8]int32
// 	)

// 	Radio0Freq = 470300000 + 1600000*(r.subband-1) + 1100000
// 	Radio1Freq = 470300000 + 1600000*(r.subband-1) + 300000
// 	RadioFreqOffsets = [8]int32{-300000, -100000, 100000, 300000, -300000, -100000, 100000, 300000}

// 	return hwconfig.FillPacketForwarder(&hwconfig.PacketForwarderConfig{
// 		SX130xConfig: hwconfig.SX130xConfig{
// 			Radio0: hwconfig.Radio0{
// 				SingleInputMode: true,
// 				Freq:            Radio0Freq,
// 				RssiOffset:      -207.0,
// 				TxFreqMin:       500000000,
// 				TxFreqMax:       510000000,
// 				TxGainLut: []hwconfig.TxGainLutItem{
// 					{RFPower: -6, PaGain: 0, PwrIdx: 0},
// 					{RFPower: -3, PaGain: 0, PwrIdx: 1},
// 					{RFPower: 0, PaGain: 0, PwrIdx: 2},
// 					{RFPower: 3, PaGain: 1, PwrIdx: 3},
// 					{RFPower: 6, PaGain: 1, PwrIdx: 4},
// 					{RFPower: 10, PaGain: 1, PwrIdx: 5},
// 					{RFPower: 11, PaGain: 1, PwrIdx: 6},
// 					{RFPower: 12, PaGain: 1, PwrIdx: 7},
// 					{RFPower: 13, PaGain: 1, PwrIdx: 8},
// 					{RFPower: 14, PaGain: 1, PwrIdx: 9},
// 					{RFPower: 16, PaGain: 1, PwrIdx: 10},
// 					{RFPower: 20, PaGain: 1, PwrIdx: 11},
// 					{RFPower: 23, PaGain: 1, PwrIdx: 12},
// 					{RFPower: 25, PaGain: 1, PwrIdx: 13},
// 					{RFPower: 26, PaGain: 1, PwrIdx: 14},
// 					{RFPower: 27, PaGain: 1, PwrIdx: 15},
// 				},
// 			},
// 			Radio1: hwconfig.Radio1{
// 				SingleInputMode: true,
// 				Freq:            Radio1Freq,
// 				RssiOffset:      -207.0,
// 			},
// 			ChanMultiSF0: hwconfig.ChanMultiSF{
// 				Enable: true,
// 				Radio:  0,
// 				IF:     RadioFreqOffsets[0],
// 			},
// 			ChanMultiSF1: hwconfig.ChanMultiSF{
// 				Enable: true,
// 				Radio:  0,
// 				IF:     RadioFreqOffsets[1],
// 			},
// 			ChanMultiSF2: hwconfig.ChanMultiSF{
// 				Enable: true,
// 				Radio:  0,
// 				IF:     RadioFreqOffsets[2],
// 			},
// 			ChanMultiSF3: hwconfig.ChanMultiSF{
// 				Enable: true,
// 				Radio:  0,
// 				IF:     RadioFreqOffsets[3],
// 			},
// 			ChanMultiSF4: hwconfig.ChanMultiSF{
// 				Enable: true,
// 				Radio:  1,
// 				IF:     RadioFreqOffsets[4],
// 			},
// 			ChanMultiSF5: hwconfig.ChanMultiSF{
// 				Enable: true,
// 				Radio:  1,
// 				IF:     RadioFreqOffsets[5],
// 			},
// 			ChanMultiSF6: hwconfig.ChanMultiSF{
// 				Enable: true,
// 				Radio:  1,
// 				IF:     RadioFreqOffsets[6],
// 			},
// 			ChanMultiSF7: hwconfig.ChanMultiSF{
// 				Enable: true,
// 				Radio:  1,
// 				IF:     RadioFreqOffsets[7],
// 			},
// 			ChanLoraStd: hwconfig.ChanLoraStd{
// 				Enable:                true,
// 				Radio:                 1,
// 				IF:                    -200000,
// 				Bandwidth:             250000,
// 				SpreadFactor:          7,
// 				ImplicitHdr:           false,
// 				Implicitpayloadlength: 17,
// 				ImplicitcrcEn:         false,
// 				Implicitcoderate:      1,
// 			},
// 			ChanLoraFSK: hwconfig.ChanLoraFSK{
// 				Enable:    true,
// 				Radio:     1,
// 				IF:        300000,
// 				Bandwidth: 125000,
// 				Datarate:  50000,
// 			},
// 		},
// 		GateWayConfig: hwconfig.GateWayConfig{
// 			BeaconPeriod:   0,
// 			BeaconFreqHZ:   869525000,
// 			BeaconFreqNB:   0,
// 			BeaconFreqStep: 0,
// 			BeaconDatarate: 9,
// 			BeaconBwHZ:     125000,
// 			BeaconPower:    14,
// 			BeaconInfodesc: 0,
// 		},
// 	})
// }

// func (r *CN470) buildGatewayBridge() *hwconfig.GatewayBridgeConfig {
// 	var frequencies [8]int32
// 	for i := range frequencies {
// 		frequencies[i] = 470300000 + (r.subband-1)*1600000 + int32(i)*200000
// 	}
// 	return hwconfig.NewGatewayBridge(r.backend, frequencies[:])
// }

// func (r *CN470) buildNetworkServer() *hwconfig.NetworkServerConfig {
// 	return hwconfig.NewNetworkServerConfig("CN470")
// }
