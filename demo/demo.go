package main

import (
	"encoding/json"
	"fmt"

	"gitee.com/dfrobotcd/chirpstack-api/go/as/external/api"
)

func main() {
	err := hwconfig.SetupDebug()
	if err != nil {
		panic(err)
	}

	hwconfig.Print(hwconfig.GetMode())
	hwconfig.Print(hwconfig.GetRegion())
	hwconfig.Print(hwconfig.GetFilter())

	// hwconfig.SetMode(hwconfig.GetMode())

	// band.GetEU868DefaultChannelSettings()
	// if err := hwconfig.SetRegion(&api.GateWayRegion{
	// 	RegionId: "EU868",
	// 	RegionConfig: &api.GateWayRegion_Eu868{
	// 		Eu868:,
	// 	},
	// }); err != nil {
	// 	panic(err)
	// }
}
