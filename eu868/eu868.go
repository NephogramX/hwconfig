package eu868

import (
	"errors"

	"github.com/NephogramX/hwconfig"
)

type EU868 struct {
	backend    string
	customBand hwconfig.CustomBand
	Builder    hwconfig.Builder
}

func NewBuilder() *EU868 {
	return &EU868{}
}

func (r *EU868) SetBackend(b string) {
	r.backend = b
}

func (r *EU868) SetCustomBand(c hwconfig.CustomBand) {
	r.customBand = c
}

func (r *EU868) SetSubband(fsb int32) {
}

func (r *EU868) Build() (*hwconfig.Configs, error) {
	if !hwconfig.CheckBackend(r.backend) {
		return nil, errors.New("unkown gateway backend")
	}

	return &hwconfig.Configs{
		PacketForwarder: r.buildPacketForwarder(),
		GatewayBridge:   r.buildGatewayBridge(),
		NetworkServer:   r.buildNetworkServer(),
	}, nil
}

func (r *EU868) buildPacketForwarder() *hwconfig.PacketForwarderConfig {
	if r.backend != hwconfig.SemtechUDP {
		return nil
	}

	return hwconfig.FillPacketForwarder(&hwconfig.PacketForwarderConfig{
		SX130xConfig: hwconfig.SX130xConfig{
			Radio0: hwconfig.Radio0{
				SingleInputMode: false,
				Freq:            r.customBand.CenterFrequency,
				RssiOffset:      -215.4,
				TxFreqMin:       863000000,
				TxFreqMax:       870000000,
				TxGainLut: []hwconfig.TxGainLutItem{
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
			},
			Radio1: hwconfig.Radio1{
				SingleInputMode: false,
				Freq:            868500000,
				RssiOffset:      -215.4,
			},
			ChanMultiSF0: hwconfig.ChanMultiSF{
				Enable: true,
				Radio:  1,
				IF:     -400000,
			},
			ChanMultiSF1: hwconfig.ChanMultiSF{
				Enable: true,
				Radio:  1,
				IF:     -200000,
			},
			ChanMultiSF2: hwconfig.ChanMultiSF{
				Enable: true,
				Radio:  1,
				IF:     0,
			},
			ChanMultiSF3: hwconfig.ChanMultiSF{
				Enable: true,
				Radio:  0,
				IF:     r.customBand.FrequencyShift[0],
			},
			ChanMultiSF4: hwconfig.ChanMultiSF{
				Enable: true,
				Radio:  0,
				IF:     r.customBand.FrequencyShift[1],
			},
			ChanMultiSF5: hwconfig.ChanMultiSF{
				Enable: true,
				Radio:  0,
				IF:     r.customBand.FrequencyShift[2],
			},
			ChanMultiSF6: hwconfig.ChanMultiSF{
				Enable: true,
				Radio:  0,
				IF:     r.customBand.FrequencyShift[3],
			},
			ChanMultiSF7: hwconfig.ChanMultiSF{
				Enable: true,
				Radio:  0,
				IF:     r.customBand.FrequencyShift[4],
			},
			ChanLoraStd: hwconfig.ChanLoraStd{
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
			ChanLoraFSK: hwconfig.ChanLoraFSK{
				Enable:    true,
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
			BeaconPower:    27,
			BeaconInfodesc: 0,
		},
	})
}

func (r *EU868) buildGatewayBridge() *hwconfig.GatewayBridgeConfig {
	return hwconfig.NewGatewayBridge(&hwconfig.GbBackend{})

	// i := []int32{
	// 	868100000, 868300000, 868500000,
	// 	r.customBand.CenterFrequency + r.customBand.FrequencyShift[0],
	// 	r.customBand.CenterFrequency + r.customBand.FrequencyShift[1],
	// 	r.customBand.CenterFrequency + r.customBand.FrequencyShift[2],
	// 	r.customBand.CenterFrequency + r.customBand.FrequencyShift[3],
	// 	r.customBand.CenterFrequency + r.customBand.FrequencyShift[4],
	// }
}

func (r *EU868) buildNetworkServer() *hwconfig.NetworkServerConfig {
	return hwconfig.NewNetworkServerConfig("EU868")
}
