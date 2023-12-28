package hwconfig

import (
	"errors"
)

type BandEu868 struct {
	backend    string
	customBand CustomBand
	Builder    Builder
}

func (r *BandEu868) SetBackend(b string) {
	r.backend = b
}

func (r *BandEu868) SetCustomBand(c CustomBand) {
	r.customBand = c
}

func (r *BandEu868) SetSubband(fsb int32) {
}

func (r *BandEu868) Build() (*Configs, error) {
	if !checkBackend(r.backend) {
		return nil, errors.New("unkown gateway backend")
	}

	return &Configs{
		PacketForwarder: r.buildPacketForwarder(),
		GatewayBridge:   r.buildGatewayBridge(),
		NetworkServer:   r.buildNetworkServer(),
	}, nil
}

func (r *BandEu868) buildPacketForwarder() *SemtechUdpConfig {
	if r.backend != SemtechUDP {
		return nil
	}

	return fillPacketForwarder(&SemtechUdpConfig{
		SX130xConfig: SX130xConfig{
			Radio0: Radio0{
				SingleInputMode: false,
				Freq:            r.customBand.CenterFrequency,
				RssiOffset:      -215.4,
				TxFreqMin:       863000000,
				TxFreqMax:       870000000,
				TxGainLut: []TxGainLutItem{
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
			Radio1: Radio1{
				SingleInputMode: false,
				Freq:            868500000,
				RssiOffset:      -215.4,
			},
			ChanMultiSF0: ChanMultiSF{
				Enable: true,
				Radio:  1,
				IF:     -400000,
			},
			ChanMultiSF1: ChanMultiSF{
				Enable: true,
				Radio:  1,
				IF:     -200000,
			},
			ChanMultiSF2: ChanMultiSF{
				Enable: true,
				Radio:  1,
				IF:     0,
			},
			ChanMultiSF3: ChanMultiSF{
				Enable: true,
				Radio:  0,
				IF:     r.customBand.FrequencyShift[0],
			},
			ChanMultiSF4: ChanMultiSF{
				Enable: true,
				Radio:  0,
				IF:     r.customBand.FrequencyShift[1],
			},
			ChanMultiSF5: ChanMultiSF{
				Enable: true,
				Radio:  0,
				IF:     r.customBand.FrequencyShift[2],
			},
			ChanMultiSF6: ChanMultiSF{
				Enable: true,
				Radio:  0,
				IF:     r.customBand.FrequencyShift[3],
			},
			ChanMultiSF7: ChanMultiSF{
				Enable: true,
				Radio:  0,
				IF:     r.customBand.FrequencyShift[4],
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
			BeaconFreqNB:   1,
			BeaconFreqStep: 0,
			BeaconDatarate: 9,
			BeaconBwHZ:     125000,
			BeaconPower:    27,
			BeaconInfodesc: 0,
		},
	})
}

func (r *BandEu868) buildGatewayBridge() *GatewayBridgeConfig {
	b := &GbBackend{
		Type: r.backend,
	}

	switch r.backend {
	case BStation:
		b.SemtechUdp = nil
		b.BasicStation = &BasicStation{
			Bind:         "0.0.0.0:3001",
			Region:       "EU868",
			FrequencyMin: 863000000,
			FrequencyMax: 870000000,
			Concentrators: Concentrators{
				MultiSF: MultiSF{
					Frequencies: []int32{
						868100000, 868300000, 868500000,
						r.customBand.CenterFrequency + r.customBand.FrequencyShift[0],
						r.customBand.CenterFrequency + r.customBand.FrequencyShift[1],
						r.customBand.CenterFrequency + r.customBand.FrequencyShift[2],
						r.customBand.CenterFrequency + r.customBand.FrequencyShift[3],
						r.customBand.CenterFrequency + r.customBand.FrequencyShift[4],
					},
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

func (r *BandEu868) buildNetworkServer() *NetworkServerConfig {
	ec := [5]ExtraChannels{}

	for i := range ec {
		ec[i].Frequency = r.customBand.CenterFrequency + r.customBand.FrequencyShift[i]
		ec[i].MinDr = 0
		ec[i].MaxDr = 5
	}

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
				Name: "EU868",
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
			NetworkSettings: NetworkSettings{
				ExtraChannels: ec[:],
			},
		},
		JoinServer: JoinServer{
			Default: Default{
				Server: "http://localhost:8003",
			},
		},
	}
}
