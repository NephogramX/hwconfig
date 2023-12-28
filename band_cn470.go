package hwconfig

import (
	"errors"
	"fmt"
)

type BandCn470 struct {
	backend string
	subband int32
	Builder Builder
}

func (r *BandCn470) SetBackend(b string) {
	r.backend = b
}

func (r *BandCn470) SetCustomBand(c CustomBand) {
}

func (r *BandCn470) SetSubband(fsb int32) {
	r.subband = fsb
}

func (r *BandCn470) Build() (*Configs, error) {
	if r.subband < 0 || r.subband > 9 {
		return nil, errors.New(fmt.Sprint("unknow subband:", r.subband, " in CN470"))
	}

	if !checkBackend(r.backend) {
		return nil, errors.New("unkown gateway backend")
	}

	return &Configs{
		PacketForwarder: r.buildPacketForwarder(),
		GatewayBridge:   r.buildGatewayBridge(),
		NetworkServer:   r.buildNetworkServer(),
	}, nil
}

func (r *BandCn470) buildPacketForwarder() *SemtechUdpConfig {
	if r.backend != SemtechUDP {
		return nil
	}

	var (
		Radio0Freq, Radio1Freq int32
		RadioFreqOffsets       [8]int32
	)

	Radio0Freq = 470300000 + 1600000*(r.subband-1) + 1100000
	Radio1Freq = 470300000 + 1600000*(r.subband-1) + 300000
	RadioFreqOffsets = [8]int32{-300000, -100000, 100000, 300000, -300000, -100000, 100000, 300000}

	return fillPacketForwarder(&SemtechUdpConfig{
		SX130xConfig: SX130xConfig{
			Radio0: Radio0{
				SingleInputMode: true,
				Freq:            Radio0Freq,
				RssiOffset:      -207.0,
				TxFreqMin:       500000000,
				TxFreqMax:       510000000,
				TxGainLut: []TxGainLutItem{
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
			},
			Radio1: Radio1{
				SingleInputMode: true,
				Freq:            Radio1Freq,
				RssiOffset:      -207.0,
			},
			ChanMultiSF0: ChanMultiSF{
				Enable: true,
				Radio:  0,
				IF:     RadioFreqOffsets[0],
			},
			ChanMultiSF1: ChanMultiSF{
				Enable: true,
				Radio:  0,
				IF:     RadioFreqOffsets[1],
			},
			ChanMultiSF2: ChanMultiSF{
				Enable: true,
				Radio:  0,
				IF:     RadioFreqOffsets[2],
			},
			ChanMultiSF3: ChanMultiSF{
				Enable: true,
				Radio:  0,
				IF:     RadioFreqOffsets[3],
			},
			ChanMultiSF4: ChanMultiSF{
				Enable: true,
				Radio:  1,
				IF:     RadioFreqOffsets[4],
			},
			ChanMultiSF5: ChanMultiSF{
				Enable: true,
				Radio:  1,
				IF:     RadioFreqOffsets[5],
			},
			ChanMultiSF6: ChanMultiSF{
				Enable: true,
				Radio:  1,
				IF:     RadioFreqOffsets[6],
			},
			ChanMultiSF7: ChanMultiSF{
				Enable: true,
				Radio:  1,
				IF:     RadioFreqOffsets[7],
			},
			ChanLoraStd: ChanLoraStd{
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
			ChanLoraFSK: ChanLoraFSK{
				Enable:    true,
				Radio:     1,
				IF:        300000,
				Bandwidth: 125000,
				Datarate:  50000,
			},
		},
		GateWayConfig: GateWayConfig{
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

func (r *BandCn470) buildGatewayBridge() *GatewayBridgeConfig {
	b := &GbBackend{
		Type: r.backend,
	}

	var frequencies [8]int32
	for i := range frequencies {
		frequencies[i] = 470300000 + (r.subband-1)*1600000 + int32(i)*200000
	}

	switch r.backend {
	case BStation:
		b.SemtechUdp = nil
		b.BasicStation = &BasicStation{
			Bind:         "0.0.0.0:3001",
			Region:       "CN470",
			FrequencyMin: 470000000,
			FrequencyMax: 510000000,
			Concentrators: Concentrators{
				MultiSF: MultiSF{
					Frequencies: frequencies[:],
				},
			},
		}
	case SemtechUDP:
		b.BasicStation = nil
		b.SemtechUdp = &SemtechUdp{
			UdpBind: "0.0.0.0:1700",
		}
	}

	return &GatewayBridgeConfig{
		GbBackend: *b,
		Intergration: Intergration{
			Marshaler: "protobuf",
			GbMqtt: GbMqtt{
				EventTopicTemplate:   "gateway/{{ .GatewayID }}/event/{{ .EventType }}",
				CommandTopicTemplate: "gateway/{{ .GatewayID }}/command/#",
				Auth: Auth{
					Type: "generic",
					Generic: Generic{
						Server: "tcp://127.0.0.1:1883",
					},
				},
			},
		},
	}
}

func (r *BandCn470) buildNetworkServer() *NetworkServerConfig {
	return &NetworkServerConfig{
		Postgresql: Postgresql{
			Dsn: "postgres://chirpstack_ns:dfrobot@localhost/chirpstack_ns?sslmode=disable",
		},
		Redis: Redis{
			Url: "redis://localhost:6379",
		},
		NetworkServer: NetworkServer{
			NetId: "000000",
			Api: Api{
				Bind: "0.0.0.0:8000",
			},
			Band: Band{
				Name: "CN470",
			},
			Gateway: Gateway{
				NsBackend: NsBackend{
					Type: "mqtt",
					NsMqtt: NsMqtt{
						CommandTopicTemplate: "gateway/{{ .GatewayID }}/command/{{ .CommandType }}",
						EventTopic:           "gateway/+/event/+",
						Server:               "tcp://localhost:1883",
					},
				},
			},
			NetworkSettings: NetworkSettings{},
		},
		JoinServer: JoinServer{
			Default: Default{
				Server: "http://localhost:8003",
			},
		},
	}
}
