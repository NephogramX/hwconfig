package main

import "gitee.com/dfrobotcd/chirpstack-api/go/as/external/api"

var (
	CN470 *api.GateWayRegion = &api.GateWayRegion{
		RegionId: "CN470",
		RegionConfig: &api.GateWayRegion_Cn470{
			Cn470: &api.CN470Config{
				SubBandId: 2,
			},
		},
	}

	EU868 *api.GateWayRegion = &api.GateWayRegion{
		RegionId: "EU868",
		RegionConfig: &api.GateWayRegion_Eu868{
			Eu868: &api.EU868Config{
				Radio_0: &api.EU868Radio0{
					Freq: 867500000,
				},
				ChanMultiSF_3: &api.EU868ChannelMultiSF{
					Offset: -400000,
				},
				ChanMultiSF_4: &api.EU868ChannelMultiSF{
					Offset: -200000,
				},
				ChanMultiSF_5: &api.EU868ChannelMultiSF{
					Offset: 0,
				},
				ChanMultiSF_6: &api.EU868ChannelMultiSF{
					Offset: 200000,
				},
				ChanMultiSF_7: &api.EU868ChannelMultiSF{
					Offset: 400000,
				},
			},
		},
	}

	IN865 *api.GateWayRegion = &api.GateWayRegion{
		RegionId: "IN865",
	}

	RU864 *api.GateWayRegion = &api.GateWayRegion{
		RegionId: "RU864",
	}

	US915 *api.GateWayRegion = &api.GateWayRegion{
		RegionId: "US915",
		RegionConfig: &api.GateWayRegion_Us915{
			Us915: &api.US915Config{
				SubBandId: 2,
			},
		},
	}

	AU915 *api.GateWayRegion = &api.GateWayRegion{
		RegionId: "AU915",
	}

	KR920 *api.GateWayRegion = &api.GateWayRegion{
		RegionId: "KR920",
	}

	AS923 *api.GateWayRegion = &api.GateWayRegion{
		RegionId: "AS923",
	}
)
