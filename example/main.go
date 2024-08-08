package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"gitee.com/dfrobotcd/chirpstack-api/go/as/external/api"
	"github.com/NephogramX/hwconfig/hwconfig"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	spew.Config.DisableMethods = true
	os.Mkdir("./build", os.ModePerm)
	hwconfig.Setup(
		"EU868",
		"./build/chirpstack-network-server.toml",
		"./build/chirpstack-gateway-bridge.toml",
		"./build/global_conf.json",
		"./build/",
		func(ctx context.Context) (int, int, error) {
			return 0, 5, nil
		},
		func(ctx context.Context, min int, max int) error {
			return nil
		},
	)

	exampleDefault()
}

func exampleIO() {
	res, err := hwconfig.Get(context.Background(), true)
	if err != nil {
		panic(err)
	}

	req := &api.ConfigGateWayModeRegionRequest{
		Mode:   res.Mode,
		Region: res.Region,
		Filter: res.Filter,
	}

	if err := hwconfig.Set(context.Background(), req); err != nil {
		panic(err)
	}
}

func exampleBs() {
	res, err := hwconfig.Get(context.Background(), true)
	if err != nil {
		panic(err)
	}
	j, _ := json.MarshalIndent(res.Mode, "", "  ")
	fmt.Println(string(j))

	if err := hwconfig.Set(context.Background(), &api.ConfigGateWayModeRegionRequest{
		Mode: &api.GateWayMode{
			Mode: "BS",
			ModeConfig: &api.GateWayMode_Bs{
				Bs: &api.BasicsStation{
					Type:   "LNS",
					Server: "nam1.cloud.thethings.network",
					Port:   8887,
					Auth: &api.BSAuth{
						CaCert:  "123",
						CliCert: "123",
						CliKey:  "132",
						Mode:    "NO_AUTH",
						Token:   "Authorization: Bearer \n",
					},
				},
			},
		},
		Region: res.Region,
		Filter: res.Filter,
	}); err != nil {
		panic(err)
	}
}

func exampleAdrRange() {
	min, max, err := hwconfig.GetAdrRange("EU868")
	if err != nil {
		panic(err)
	}
	fmt.Println("EU868 ADR Range: ", min, max)
	min, max, err = hwconfig.GetAdrRange("US915")
	if err != nil {
		panic(err)
	}
	fmt.Println("US915 ADR Range: ", min, max)
	min, max, err = hwconfig.GetAdrRange("AB123")
	if err != nil {
		panic(err)
	}
	fmt.Println("should panic")
}

func exampleDefault() {
	res, err := hwconfig.GetFromDefault()
	if err != nil {
		panic(err)
	}
	b, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println("-------------------------------------\n", string(b))
}
