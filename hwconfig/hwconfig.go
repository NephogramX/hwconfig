package hwconfig

import (
	"context"
	"os"

	"gitee.com/dfrobotcd/chirpstack-api/go/as/external/api"
	"github.com/NephogramX/hwconfig/internal/band"
	"github.com/NephogramX/hwconfig/internal/conf"
	"github.com/NephogramX/hwconfig/internal/file"
	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
)

var getAdrRange func(ctx context.Context) (int, int, error) = nil
var setAdrRange func(ctx context.Context, min int, max int) error = nil

func Setup(region string, nsdir string, gbdir string, pfdir string, bsdir string, GetAdrRange func(context.Context) (int, int, error), SetAdrRange func(context.Context, int, int) error) error {
	if err := file.Setup(
		region,
		nsdir,
		gbdir,
		pfdir,
		bsdir,
	); err != nil {
		return err
	}
	if GetAdrRange == nil || SetAdrRange == nil {
		return errors.New("nil get/setAdrRange func")
	}
	getAdrRange = GetAdrRange
	setAdrRange = SetAdrRange
	return nil
}

func Get(ctx context.Context, hasPermission bool) (*api.GetGateWayModeRegionResponse, error) {
	var (
		notBsMode bool = false
		notPfMode bool = false
		conf           = struct {
			Ns *conf.NsConfig
			Gb *conf.GbConfig
			Pf *conf.PfConfig
			Bs *conf.BsConfig
		}{&conf.NsConfig{}, &conf.GbConfig{}, &conf.PfConfig{}, &conf.BsConfig{}}
	)

	// read config file
	if err := file.Get(conf.Ns); err != nil {
		return nil, errors.Wrap(err, "failed to get ns config")
	}

	if err := file.Get(conf.Gb); err != nil {
		return nil, errors.Wrap(err, "failed to get gb config")
	}

	if err := file.Get(conf.Pf); err != nil {
		return nil, errors.Wrap(err, "failed to get pf config")
	}
	if conf.Pf.ServerAddress == "localhost" || conf.Pf.ServerAddress == "127.0.0.1" {
		notPfMode = true
	}

	if err := file.Get(conf.Bs); err == os.ErrNotExist {
		notBsMode = true
	} else if err != nil {
		return nil, errors.Wrap(err, "failed to get bs config")
	}

	b, err := band.NewWithChannel(conf.Ns.NetworkServer.Band.Name, &conf.Pf.SX130xConfig)

	if err != nil {
		return nil, errors.Wrap(err, "bad region config:")
	}
	res := &api.GetGateWayModeRegionResponse{}

	// recognize gateway mode
	if !notBsMode { // BsMode
		spew.Dump(conf.Bs)

		res.Mode = conf.Bs.ApiMode()
		res.Region = b.GetApiRegion()
		res.Filter = conf.Gb.ApiFilter()

		if !hasPermission {
			res.Mode.GetBs().Auth.CaCert = "******"
			res.Mode.GetBs().Auth.CliCert = "******"
			res.Mode.GetBs().Auth.CliKey = "******"
			res.Mode.GetBs().Auth.Token = "******"
		}

	} else if notPfMode { // NsMode
		min, max, err := getAdrRange(ctx)
		if err != nil {
			return nil, err
		}
		res.Mode = conf.Ns.ApiMode(min, max)
		res.Region = b.GetApiRegion()
		res.Filter = conf.Gb.ApiFilter()

	} else { // PfMode
		res.Mode = conf.Pf.ApiMode()
		res.Region = b.GetApiRegion()
		res.Filter = conf.Gb.ApiFilter()
	}

	return res, nil
}

func Set(ctx context.Context, req *api.ConfigGateWayModeRegionRequest) error {
	var b band.Band
	pfConf := conf.PfConfig{}
	if err := file.Get(&pfConf); err != nil {
		return errors.Wrap(err, "failed to get pf config")
	}
	// fmt.Println("old band settings:")
	// j, _ := json.MarshalIndent(pfConf.SX130xConfig, "", "  ")
	// fmt.Println(string(j))

	var (
		ok  bool  = false
		err error = nil
	)
	if ok, err = band.HasSubband(req.GetRegion().GetRegionId()); err != nil {
		return err
	} else if ok {
		b, err = band.NewWithSubband(req.GetRegion().GetRegionId(), &pfConf.SX130xConfig, int(req.GetRegion().GetUs915().GetSubBandId()))
	} else {
		b, err = band.NewWithChannel(req.GetRegion().GetRegionId(), band.SetChannel(&pfConf.SX130xConfig, req.GetRegion().GetEu868()))
	}
	if err != nil {
		return errors.Wrap(err, "bad region config:")
	}

	switch req.GetMode().GetMode() {
	case "NS":
		if err := setAdrRange(ctx, int(req.GetMode().GetNs().GetAdr().DrIdMin), int(req.GetMode().GetNs().GetAdr().DrIdMax)); err != nil {
			return err
		}
		if err := file.Set(conf.NewNsConfig(b, req.GetMode().GetNs())); err != nil {
			return errors.Wrap(err, "set ns config fail:")
		}
		if err := file.Set(conf.NewGbConfig(req.GetFilter())); err != nil {
			return errors.Wrap(err, "set gb config fail:")
		}

		pfConf.GatewayConfig.ServerAddress = "127.0.0.1"
		pfConf.GatewayConfig.ServPortUp = 1700
		pfConf.GatewayConfig.ServPortDown = 1700
		if err := file.Set(&pfConf); err != nil {
			return errors.Wrap(err, "set pf config fail")
		}
		file.CleanAuthFile()

	case "PF":
		pfConf.GatewayConfig.ServerAddress = req.GetMode().GetPf().GetProtocol().GetGwmp().Server
		pfConf.GatewayConfig.ServPortUp = req.GetMode().GetPf().GetProtocol().GetGwmp().Port.Uplink
		pfConf.GatewayConfig.ServPortDown = req.GetMode().GetPf().GetProtocol().GetGwmp().Port.Downlink
		if err := file.Set(&pfConf); err != nil {
			return errors.Wrap(err, "set pf config fail")
		}
		file.CleanAuthFile()

	case "BS":
		useCups := false
		switch req.GetMode().GetBs().GetType() {
		case "CUPS":
			useCups = true
		case "LNS":
			useCups = false
		}

		server := &conf.NoAuth{
			UseCups:    useCups,
			ServerAddr: req.GetMode().GetBs().GetServer(),
			ServerPort: int32(req.GetMode().GetBs().GetPort()),
		}

		switch req.GetMode().GetBs().GetAuth().GetMode() {
		case "NO_AUTH":
			err = file.Set(&conf.BsConfig{BsAuth: server})

		case "TLS_Server":
			err = file.Set(&conf.BsConfig{
				BsAuth: &conf.TlsServerAuth{
					NoAuth: *server,
					Trust:  []byte(req.GetMode().GetBs().GetAuth().GetCaCert()),
				}})

		case "TLS_Server_Client":
			err = file.Set(&conf.BsConfig{
				BsAuth: &conf.TlsServerAndClientAuth{
					NoAuth: *server,
					Key:    []byte(req.GetMode().GetBs().GetAuth().GetCliKey()),
					Trust:  []byte(req.GetMode().GetBs().GetAuth().GetCaCert()),
					Crt:    []byte(req.GetMode().GetBs().GetAuth().GetCliCert()),
				}})

		case "TLS_Server_Client_Token":
			err = file.Set(&conf.BsConfig{
				BsAuth: &conf.TlsServerAuthAndClientToken{
					NoAuth: *server,
					Key:    []byte(req.GetMode().GetBs().GetAuth().GetToken()),
					Trust:  []byte(req.GetMode().GetBs().GetAuth().GetCaCert()),
				}})
		}

	default:
		err = errors.Errorf("unsupported authorization mode: %s", req.GetMode().GetMode())
	}

	// fmt.Println("new band settings:")
	// j, _ = json.MarshalIndent(b.GetSX130xConfig(), "", "  ")
	// fmt.Println(string(j))

	return err

}

func GetAdrRange(region string) (int, int, error) {
	b, err := band.NewWithChannel(region, nil)
	if err != nil {
		return -1, -1, err
	}
	min, max := b.GetAdrRange()
	return min, max, err
}
