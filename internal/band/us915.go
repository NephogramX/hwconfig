package band

import (
	"sort"

	"gitee.com/dfrobotcd/chirpstack-api/go/as/external/api"
	"github.com/pkg/errors"
)

type US915 struct {
	subband int
	*SX130xConfig
}

func (b *US915) String() string {
	return "US915"
}

func (b *US915) GetSubband() (int, error) {
	radiofreq := [2]int32{b.Radio0.Freq, b.Radio1.Freq}
	multisf := []int{
		int(radiofreq[b.ChanMultiSF0.Radio] + b.ChanMultiSF0.IF),
		int(radiofreq[b.ChanMultiSF1.Radio] + b.ChanMultiSF1.IF),
		int(radiofreq[b.ChanMultiSF2.Radio] + b.ChanMultiSF2.IF),
		int(radiofreq[b.ChanMultiSF3.Radio] + b.ChanMultiSF3.IF),
		int(radiofreq[b.ChanMultiSF4.Radio] + b.ChanMultiSF4.IF),
		int(radiofreq[b.ChanMultiSF5.Radio] + b.ChanMultiSF5.IF),
		int(radiofreq[b.ChanMultiSF6.Radio] + b.ChanMultiSF6.IF),
		int(radiofreq[b.ChanMultiSF7.Radio] + b.ChanMultiSF7.IF),
	}
	sort.Ints(multisf)
	// lorastd := radiofreq[b.ChanLoraStd.Radio] + b.ChanLoraStd.IF

	if (multisf[0]-902300000)%1600000 != 0 {
		return -1, errFailToFindSubband(multisf)
	}
	return (multisf[0]-902300000)/1600000 + 1, nil
}

func (b *US915) SetSubband(id int) error {
	if id < 1 || id > 8 {
		return errors.Errorf("unsupported subband for us915: %d", id)
	}
	b.Radio0.Freq = int32(902300000 + 1600000*(id-1) + 400000)
	b.Radio1.Freq = int32(902300000 + 1600000*(id-1) + 1100000)
	b.ChanMultiSF0 = ChanMultiSF{Enable: true, Radio: 0, IF: -400000}
	b.ChanMultiSF1 = ChanMultiSF{Enable: true, Radio: 0, IF: -200000}
	b.ChanMultiSF2 = ChanMultiSF{Enable: true, Radio: 0, IF: 0}
	b.ChanMultiSF3 = ChanMultiSF{Enable: true, Radio: 0, IF: 200000}
	b.ChanMultiSF4 = ChanMultiSF{Enable: true, Radio: 1, IF: -300000}
	b.ChanMultiSF5 = ChanMultiSF{Enable: true, Radio: 1, IF: -100000}
	b.ChanMultiSF6 = ChanMultiSF{Enable: true, Radio: 1, IF: 100000}
	b.ChanMultiSF7 = ChanMultiSF{Enable: true, Radio: 1, IF: 300000}
	b.ChanLoraFSK = ChanLoraFSK{Enable: false}
	b.ChanLoraStd = ChanLoraStd{
		Enable:                true,
		Radio:                 0,
		IF:                    300000,
		Bandwidth:             500000,
		SpreadFactor:          8,
		ImplicitHdr:           false,
		Implicitpayloadlength: 17,
		ImplicitcrcEn:         false,
		Implicitcoderate:      1,
	}
	b.subband = id
	return nil
}

func (b *US915) GetExtraChannels() []ExtraChannel {
	return nil
}

func (b *US915) GetUplinkChannels() []int {
	ch := (b.subband - 1) * 8
	return []int{ch, ch + 1, ch + 2, ch + 3, ch + 4, ch + 5, ch + 6, ch + 7}
}

func (b *US915) GetSX130xConfig() *SX130xConfig {
	return b.SX130xConfig
}

func (b *US915) GetApiRegion() *api.GateWayRegion {
	subbandId, err := b.GetSubband()
	if err != nil {
		subbandId = -1
	}
	return &api.GateWayRegion{
		RegionId: b.String(),
		RegionConfig: &api.GateWayRegion_Us915{
			Us915: &api.US915Config{
				LoraWanPublic: false,
				SubBandId:     int32(subbandId),
				TxFreq:        &api.TXFreqItem{Min: 0, Max: 0},
			},
		},
	}
}

func (b *US915) GetDefaultRssiTComp() *RssiTcomp {
	return &RssiTcomp{
		CoeffA: 0,
		CoeffB: 0,
		CoeffC: 20.41,
		CoeffD: 2162.56,
		CoeffE: 0,
	}
}

func (b *US915) GetDefaultTxGainLut() *[]TxGainLutItem {
	return &[]TxGainLutItem{
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
	}
}

func (b *US915) GetAdrRange() (int, int) {
	return 0, 4
}
