package hwconfig

import (
	"errors"
	"fmt"
)

type BandUs915 struct {
	backend string
	subband int32
}

func (r *BandUs915) SetBackend(b string) {
	r.backend = b
}

func (r *BandUs915) SetCustomBand(c CustomBand) {
}

func (r *BandUs915) SetSubband(fsb int32) {
	r.subband = fsb
}

func (r *BandUs915) Build() (*Configs, error) {
	if r.subband < 0 || r.subband > 9 {
		return nil, errors.New(fmt.Sprint("unknow subband:", r.subband, " in US915"))
	}

	if r.subband == 9 {
		return nil, errors.New("unsupported subband: 9")
	}

	if !checkBackend(r.backend) {
		return nil, errors.New(fmt.Sprint("unsupported backend: ", r.backend))
	}

	return &Configs{
		PacketForwarder: r.buildPacketForwarder(),
		GatewayBridge:   r.buildGatewayBridge(),
		NetworkServer:   r.buildNetworkServer(),
	}, nil
}

func (r *BandUs915) buildPacketForwarder() *SemtechUdpConfig {
	if r.backend != SemtechUDP {
		return nil
	}

	Radio0Freq := 902300000 + 1600000*(r.subband-1) + 400000
	Radio1Freq := 902300000 + 1600000*(r.subband-1) + 1200000
	RadioFreqOffsets := [8]int32{-400000, -200000, 0, 200000, -400000, -200000, 0, 200000}
	LoRaStdFreqOffset := int32(300000)

	return fillPacketForwarder(&SemtechUdpConfig{
		SX130xConfig: SX130xConfig{
			Radio0: Radio0{
				SingleInputMode: false,
				Freq:            Radio0Freq,
				RssiOffset:      -215.4,
				TxFreqMin:       923000000,
				TxFreqMax:       928000000,
				TxGainLut: []TxGainLutItem{
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
			Radio1: Radio1{
				SingleInputMode: false,
				Freq:            Radio1Freq,
				RssiOffset:      -215.4,
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
				Enable: true,
				Radio:  0,
				IF:     LoRaStdFreqOffset,
				// IF:                    300000,
				Bandwidth:             500000,
				SpreadFactor:          8,
				ImplicitHdr:           false,
				Implicitpayloadlength: 17,
				ImplicitcrcEn:         false,
				Implicitcoderate:      1,
			},
			ChanLoraFSK: ChanLoraFSK{
				Enable:    false,
				Radio:     1,
				IF:        300000,
				Bandwidth: 125000,
				Datarate:  50000,
			},
		},
		GateWayConfig: GateWayConfig{
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

func (r *BandUs915) buildGatewayBridge() *GatewayBridgeConfig {
	b := &GbBackend{
		Type: r.backend,
	}

	switch r.backend {
	case BStation:
		b.SemtechUdp = nil
		b.BasicStation = &BasicStation{
			Bind:         "0.0.0.0:3001",
			Region:       "US915",
			FrequencyMin: 902300000,
			FrequencyMax: 927500000,
			Concentrators: Concentrators{
				MultiSF: MultiSF{
					Frequencies: []int32{
						902300000 + (r.subband-1)*1600000 + 0*200000,
						902300000 + (r.subband-1)*1600000 + 1*200000,
						902300000 + (r.subband-1)*1600000 + 2*200000,
						902300000 + (r.subband-1)*1600000 + 3*200000,
						902300000 + (r.subband-1)*1600000 + 4*200000,
						902300000 + (r.subband-1)*1600000 + 5*200000,
						902300000 + (r.subband-1)*1600000 + 6*200000,
						902300000 + (r.subband-1)*1600000 + 7*200000,
					},
				},
				LoraStd: &LoraStd{
					Frequency:       903000000 + (r.subband-1)*1600000,
					Bandwidth:       500000,
					SpreadingFactor: 8,
				},
				Fsk: nil,
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

func (r *BandUs915) buildNetworkServer() *NetworkServerConfig {
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
				Name: "US915",
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
