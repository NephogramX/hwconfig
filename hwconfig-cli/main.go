package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/NephogramX/hwconfig"
	"github.com/NephogramX/hwconfig/us915"
)

type arr []int

type Args struct {
	Region     string
	Backend    string
	Subband    int
	CenterFreq int
	FreqShift  arr
}

type Param struct {
	Region     string
	Backend    string
	Subband    int32
	CustomBand hwconfig.CustomBand
}

func (a *arr) String() string {
	return fmt.Sprintf("%v", *a)
}

func (a *arr) Set(s string) error {
	as := strings.Split(s, ",")
	fmt.Println(s, as)
	for _, v := range as {
		i, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		*a = append(*a, i)
	}
	return nil
}

func (a *Args) Parse() (*Param, error) {
	var (
		shift [5]int32
	)

	if len(a.FreqShift) == 0 {
	} else if len(a.FreqShift) == 5 {
		for i := range a.FreqShift {
			shift[i] = int32(a.FreqShift[i])
		}
	} else {
		return nil, errors.New("invalid frequency shift, must len 5")
	}

	return &Param{
		Region:  a.Region,
		Backend: a.Backend,
		Subband: int32(a.Subband),
		CustomBand: hwconfig.CustomBand{
			CenterFrequency: int32(a.CenterFreq),
			FrequencyShift:  shift,
		},
	}, nil
}

const (
	packetForwarderPath = "./global_conf.json"
	gatewayBridgePath   = "./chirpstack-gateway-bridge.toml"
	networkServerPath   = "./chirpstack-network-server.toml"
)

var (
	args Args
)

func update(path string, m hwconfig.Marshaler) error {
	// can't use "m == nil", the type of interface is not nil while value is nil
	if m.IsNil() {
		return nil
	}

	b, err := m.Marshal()
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(b)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	flag.StringVar(&args.Region, "r", "EU868", "lorawan region")
	flag.StringVar(&args.Backend, "b", "pkt_fwd", "backend")
	flag.IntVar(&args.Subband, "s", 1, "subband")
	flag.IntVar(&args.CenterFreq, "c", 0, "custom band center frequency")
	flag.Var(&args.FreqShift, "f", "custom band frequency shift")
}

func main() {
	flag.Parse()
	p, err := args.Parse()
	if err != nil {
		panic(err)
	}

	var (
		c *hwconfig.Configs
		b hwconfig.Builder
	)

	switch p.Region {
	case "CN470":
		// b = cn470.NewBuilder()
	case "EU868":
		// b = eu868.NewBuilder()
	case "US915":
		b = us915.NewBuilder()
	default:
		panic("unsupported region")
	}

	b.SetSubband(p.Subband)
	b.SetBackend(p.Backend)
	b.SetCustomBand(p.CustomBand)

	c, err = b.Build()
	if err != nil {
		panic(err)
	}

	UpdateList := []struct {
		FilePath  string
		Marshaler hwconfig.Marshaler
	}{
		{packetForwarderPath, c.PacketForwarder},
		{gatewayBridgePath, c.GatewayBridge},
		{networkServerPath, c.NetworkServer},
	}

	for _, m := range UpdateList {
		if err := update(m.FilePath, m.Marshaler); err != nil {
			panic(err)
		}
	}

	fmt.Println("==========================")
	fmt.Println("region:      ", p.Region)
	fmt.Println("backend:     ", args.Backend)
	fmt.Println("==========================")

	if p.Region == "EU868" {
		fmt.Println("center freq: ", p.CustomBand.CenterFrequency)
		fmt.Println("freq shift:  \n  ", p.CustomBand.FrequencyShift)
	} else {
		fmt.Println("subband:     ", p.Subband)
	}
	fmt.Println("==========================")
}
