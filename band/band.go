package band

import (
	"fmt"

	cf "github.com/NephogramX/hwconfig/configfile"
)

type Band interface {
	String() string
	GetChannelSettings() *cf.Channel
	GetExtraChannels() *[]cf.ExtraChannels
	GetUplinkChannels() *[]int32

	// GetDefaultChannels() int32
	// GetFSKDr() int32
	// GetStdDr() int32
}

func GetDefaultChannelSettings(r string) (*cf.Channel, error) {
	switch r {
	case "CN470":
		return GetCN470DefaultChannelSettings(), nil
	case "RU864":
		// return GetRU864DefaultChannelSettings(), nil
	case "IN865":
		// return GetIN865DefaultChannelSettings(), nil
	case "EU868":
		return GetEU868DefaultChannelSettings(), nil
	case "AU915":
		// return GetAU915DefaultChannelSettings(), nil
	case "US915":
		return GetUS915DefaultChannelSettings(), nil
	case "KR920":
		// return GetKR920DefaultChannelSettings(), nil
	}
	return nil, fmt.Errorf("unsupported region: %s", r)
}

func isDefaultChannel(ch int32, dc []int32) bool {
	for _, dch := range dc {
		if dch == ch {
			return true
		}
	}
	return false
}

func getExtraChannels(multiSFMaxDr int32, StdDr int32, FskDr int32, dc []int32, c *cf.Channel) *[]cf.ExtraChannels {
	chs := []cf.ExtraChannels{}
	radio := [2]int32{c.Radio0.Freq, c.Radio1.Freq}
	for _, v := range c.ChanMultiSF {
		if v.Enable {
			chf := radio[v.Radio] + v.IF
			if isDefaultChannel(chf, dc) {
				continue
			}
			chs = append(chs, cf.ExtraChannels{
				Frequency: chf,
				MinDr:     0,
				MaxDr:     multiSFMaxDr,
			})
		}
	}

	if c.ChanLoraStd.Enable {
		chs = append(chs, cf.ExtraChannels{
			Frequency: radio[c.ChanLoraStd.Radio] + c.ChanLoraStd.IF,
			MinDr:     StdDr,
			MaxDr:     StdDr,
		})
	}

	if c.ChanLoraFSK.Enable {
		chs = append(chs, cf.ExtraChannels{
			Frequency: radio[c.ChanLoraFSK.Radio] + c.ChanLoraFSK.IF,
			MinDr:     FskDr,
			MaxDr:     FskDr,
		})
	}

	return &chs
}

func rssiOffset900MHz() float32 {
	return -215.4
}

func rssiOffset500MHz() float32 {
	return -207.0
}

func rssiTcomp() *cf.RssiTcomp {
	return &cf.RssiTcomp{
		CoeffA: 0,
		CoeffB: 0,
		CoeffC: 20.41,
		CoeffD: 2162.56,
		CoeffE: 0,
	}
}

func txGainLut900MHz() []*cf.TxGainLutItem {
	return []*cf.TxGainLutItem{
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

func txGainLut500MHz() []*cf.TxGainLutItem {
	return []*cf.TxGainLutItem{
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
}

func ChanMultiSFAll() *cf.ChanMultiSFAll {
	return &cf.ChanMultiSFAll{
		SpreadingFactorEnable: []int32{5, 6, 7, 8, 9, 10, 11, 12},
		Radio:                 0,
		IF:                    0,
	}
}
