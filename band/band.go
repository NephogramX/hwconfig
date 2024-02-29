package band

import cf "github.com/NephogramX/hwconfig/configfile"

type Region int32

const (
	EU868 Region = iota
	US915 Region = iota
	CN470 Region = iota
)

type BandSettings struct {
	Region          Region
	SubbandIndex    *int32
	CenterFrequency *int32
	FrequencyShift  *[5]int32
}

type Band interface {
	String() string
	GetChannelSettings() *cf.Channel
	GetExtraChannels() *[]cf.ExtraChannels
	GetUplinkChannels() *[]int32
}
