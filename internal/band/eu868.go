package band

import "gitee.com/dfrobotcd/chirpstack-api/go/as/external/api"

type EU868 struct {
	*SX130xConfig
}

func (b *EU868) String() string {
	return "EU868"
}

func (b *EU868) GetSubband() (int, error) {
	return -1, errRegionHasNoSubband(b.String())
}

func (b *EU868) SetSubband(id int) error {
	return errRegionHasNoSubband(b.String())
}

func (b *EU868) GetExtraChannels() []ExtraChannel {
	nc := []*ChanMultiSF{
		&b.ChanMultiSF0,
		&b.ChanMultiSF1,
		&b.ChanMultiSF2,
		&b.ChanMultiSF3,
		&b.ChanMultiSF4,
		&b.ChanMultiSF5,
		&b.ChanMultiSF6,
		&b.ChanMultiSF7,
	}

	return extraChannel(5, 6, nc, &b.ChanLoraStd, []int32{868100000, 868100000, 868100000}, [2]int32{b.Radio0.Freq, b.Radio1.Freq})
}

func (b *EU868) GetUplinkChannels() []int {
	return nil
}

func (b *EU868) GetSX130xConfig() *SX130xConfig {
	return b.SX130xConfig
}

func (b *EU868) GetApiRegion() *api.GateWayRegion {
	return &api.GateWayRegion{
		RegionId: b.String(),
		RegionConfig: &api.GateWayRegion_Eu868{
			Eu868: b.SX130xConfig.ApiRegion(),
		},
	}
}

func (b *EU868) GetDefaultRssiTComp() *RssiTcomp {
	return &RssiTcomp{
		CoeffA: 0,
		CoeffB: 0,
		CoeffC: 20.41,
		CoeffD: 2162.56,
		CoeffE: 0,
	}
}

func (b *EU868) GetDefaultTxGainLut() *[]TxGainLutItem {
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
