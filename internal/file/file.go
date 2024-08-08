package file

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	"github.com/NephogramX/hwconfig/internal/conf"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type HwConfFile struct {
	Dir   string
	Viper *viper.Viper
}

type HwJsonFile struct {
	Dir     string
	Default string
}

type HwAuthFile struct {
	Dir string
}

type File conf.File

var (
	Ns *HwConfFile = nil
	Gb *HwConfFile = nil
	Pf *HwJsonFile = nil
	Bs *HwAuthFile = nil
)

func Setup(region string, nsdir string, gbdir string, pfdir string, bsdir string) error {
	switch region {
	case "EU868":
		Ns = &HwConfFile{nsdir, NewViper(nsdir, conf.Eu868NsDefault, "toml")}
		Pf = &HwJsonFile{pfdir, conf.Eu868PfDefault}
	case "US915":
		Ns = &HwConfFile{nsdir, NewViper(nsdir, conf.Us915NsDefault, "toml")}
		Pf = &HwJsonFile{pfdir, conf.Us915PfDefault}
	default:
		return errors.Errorf("unsupported region: %s", region)
	}

	Gb = &HwConfFile{gbdir, NewViper(gbdir, conf.GbDefault, "toml")}
	Bs = &HwAuthFile{bsdir} // Basics Station only has authentication file, the config will download from ns, so we don't use viper` to config`
	return nil
}

func NewViper(dir string, configdefault string, configtype string) *viper.Viper {
	v := viper.New()
	v.SetConfigFile(dir)
	v.SetConfigType(configtype)
	if err := v.ReadConfig(strings.NewReader(configdefault)); err != nil {
		fmt.Println("get default config fail:", err)
	}

	return v
}

func Get(c any) error {
	switch c := c.(type) {
	case *conf.NsConfig:
		return Ns.Get(c)
	case *conf.GbConfig:
		return Gb.Get(c)
	case *conf.PfConfig:
		return Pf.Get(c)
	case *conf.BsConfig:
		return Bs.Get(c)
	default:
		return errors.Errorf("get fail, unsupported config structure %s", reflect.TypeOf(c).String())
	}
}

func Set(c any) error {
	switch c := c.(type) {
	case *conf.NsConfig:
		return Ns.Set(c)
	case *conf.GbConfig:
		return Gb.Set(c)
	case *conf.PfConfig:
		return Pf.Set(c)
	case *conf.BsConfig:
		return Bs.Set(c)
	default:
		return errors.Errorf("set fail, unsupported config structure %s", reflect.TypeOf(c).String())
	}
}

func (c *HwConfFile) Get(_c any) error {
	if err := c.Viper.ReadInConfig(); os.IsNotExist(err) {
		fmt.Printf("config file: %s not exist, try generate\n", c.Dir)
		if err := c.Viper.WriteConfig(); err != nil {
			return errors.Wrapf(err, "config file: %s not exist, try generate fail:", c.Dir)
		}
	} else if err != nil {
		return errors.Wrapf(err, "failed to read config: %s", c.Dir)
	}
	if err := c.Viper.Unmarshal(_c); err != nil {
		return errors.Wrap(err, "failed to unmarshal config")
	}
	return nil
}

func (c *HwConfFile) Set(_c any) error {
	var mapConfig map[string]any
	c.Viper.ReadInConfig()
	if err := mapstructure.Decode(_c, &mapConfig); err != nil {
		return errors.Wrap(err, "failed to decode config to map")
	}

	for k, v := range mapConfig {
		c.Viper.Set(k, v)
	}

	c.Viper.WriteConfig()
	return nil
}

func (c *HwAuthFile) Set(_c *conf.BsConfig) error {
	auth := _c.GetAuth(c.Dir)

	if auth.Uri != nil {
		NewFile(auth.Uri.Dir, auth.Uri.Content).Wirte()
	}

	if auth.Key != nil {
		NewFile(auth.Key.Dir, auth.Key.Content).Wirte()
	}

	if auth.Crt != nil {
		NewFile(auth.Crt.Dir, auth.Crt.Content).Wirte()
	}

	if auth.Trust != nil {
		NewFile(auth.Trust.Dir, auth.Trust.Content).Wirte()
	}

	return nil
}

func (c *HwAuthFile) Get(_c *conf.BsConfig) error {
	if _c == nil {
		return errors.New("structure config.BsConfig uninitialized")
	}

	var (
		uri, key, crt, trust *File = nil, nil, nil, nil
		useCups              bool  = true
		err                  error = nil
	)

	uri = NewFile(Bs.Dir+"cups.uri", nil)
	if err := uri.Read(); err != nil {
		return errors.Wrapf(err, "failed to read %s", uri.Dir)
	}
	if uri.Content == nil {
		useCups = false
		uri.Dir = Bs.Dir + "tc.uri"
		if err := uri.Read(); err != nil {
			return errors.Wrapf(err, "failed to read %s", uri.Dir)
		}
		if uri.Content == nil {
			return os.ErrNotExist
		}
	}

	if useCups {
		key = NewFile(Bs.Dir+"cups.key", nil)
		if err := key.Read(); err != nil {
			return errors.Wrapf(err, "failed to read %s", key.Dir)
		}
		crt = NewFile(Bs.Dir+"cups.crt", nil)
		if err := crt.Read(); err != nil {
			return errors.Wrapf(err, "failed to read %s", crt.Dir)
		}
		trust = NewFile(Bs.Dir+"cups.key", nil)
		if err := trust.Read(); err != nil {
			return errors.Wrapf(err, "failed to read %s", trust.Dir)
		}
	} else {
		key = NewFile(Bs.Dir+"tc.key", nil)
		if err := key.Read(); err != nil {
			return errors.Wrapf(err, "failed to read %s", key.Dir)
		}
		crt = NewFile(Bs.Dir+"tc.crt", nil)
		if err := crt.Read(); err != nil {
			return errors.Wrapf(err, "failed to read %s", crt.Dir)
		}
		trust = NewFile(Bs.Dir+"tc.trust", nil)
		if err := trust.Read(); err != nil {
			return errors.Wrapf(err, "failed to read %s", trust.Dir)
		}
	}

	u, err := url.Parse(string(uri.Content))
	if err != nil {
		return errors.Wrapf(err, "failed to parse *.uri (%s)", uri)
	}
	port, err := strconv.Atoi(u.Port())
	if err != nil {
		return errors.Wrapf(err, "failed to parse *.uri port (%s)", u.Port())
	}

	server := conf.NoAuth{
		UseCups:    useCups,
		ServerAddr: u.Hostname(),
		ServerPort: int32(port),
	}
	if crt.Content == nil && trust.Content == nil && key.Content == nil {
		*_c = conf.BsConfig{
			BsAuth: &server,
		}

	} else if crt.Content == nil && trust.Content != nil && key.Content == nil {
		*_c = conf.BsConfig{
			BsAuth: &conf.TlsServerAuth{
				NoAuth: server,
				Trust:  trust.Content,
			},
		}
	} else if crt.Content != nil && trust.Content != nil && key.Content != nil {
		*_c = conf.BsConfig{
			BsAuth: &conf.TlsServerAndClientAuth{
				NoAuth: server,
				Crt:    crt.Content,
				Trust:  trust.Content,
				Key:    key.Content,
			},
		}

	} else if crt.Content == nil && trust.Content != nil && key.Content != nil {
		*_c = conf.BsConfig{
			BsAuth: &conf.TlsServerAuthAndClientToken{
				NoAuth: server,
				Trust:  trust.Content,
				Key:    key.Content,
			},
		}
	} else {
		return errors.New("incomplete authentication file")
	}

	return nil
}

func (c *HwJsonFile) Set(_c *conf.PfConfig) error {
	j, err := json.MarshalIndent(_c, "", "  ")
	if err != nil {
		return errors.Wrap(err, "failed to marshal pf config to json")
	}
	return os.WriteFile(c.Dir, j, os.ModePerm)
}

func (c *HwJsonFile) Get(_c *conf.PfConfig) error {
	b, err := os.ReadFile(c.Dir)
	if os.IsNotExist(err) {
		fmt.Printf("config file: %s not exist, try generate\n", c.Dir)
		b = []byte(c.Default)
		if err := os.WriteFile(c.Dir, b, os.ModePerm); err != nil {
			return errors.Wrap(err, "pf config file not exist: failed to generate defautl config file")
		}
	} else if err != nil {
		return errors.Wrapf(err, "failed to read config: %s", c.Dir)
	}

	if err := json.Unmarshal(b, &_c); err != nil {
		return errors.Wrap(err, "failed to unmarshal pf config")
	}

	return nil
}

func CleanAuthFile() error {
	pattern := filepath.Join(Bs.Dir, "tc.*")
	f1, err := filepath.Glob(pattern)
	if err != nil {
		return errors.Wrap(err, "failed to clean auth file")
	}
	pattern = filepath.Join(Bs.Dir, "cups.*")
	f2, err := filepath.Glob(pattern)
	if err != nil {
		return errors.Wrap(err, "failed to clean auth file")
	}

	f1 = append(f1, f2...)

	if f1 == nil {
		fmt.Println("no auth file found")
		return nil
	}
	for _, f := range f1 {
		err := os.Remove(f)
		if err != nil {
			return errors.Wrapf(err, "failed to remove file: %s", f)
		}
		fmt.Println("remove auth file:", f)
	}
	return nil
}

func BsConfigExist() (bool, error) {
	if ok, err := NewFile(Bs.Dir+"cups.uri", nil).Exist(); err != nil {
		return false, err
	} else if ok {
		return true, nil
	}

	if ok, err := NewFile(Bs.Dir+"tc.uri", nil).Exist(); err != nil {
		return false, err
	} else {
		return ok, nil
	}
}

func NewFile(dir string, content []byte) *File {
	return &File{
		Dir:     dir,
		Content: content,
	}
}

func (f *File) Exist() (bool, error) {
	if _, err := os.Stat(f.Dir); err == nil {
		return true, nil
	} else if err != os.ErrNotExist {
		return false, nil
	} else {
		return false, err
	}
}

func (f *File) Read() error {
	if _, err := os.Stat(f.Dir); os.IsNotExist(err) {
		fmt.Println("can not find ", f.Dir)
		f.Content = nil
		return nil
	} else if err != nil {
		fmt.Println("can not read ", f.Dir, err.Error())
		return err
	}

	var err error = nil
	f.Content, err = os.ReadFile(f.Dir)
	return err
}

func (f *File) Wirte() error {
	fd, err := os.Create(f.Dir)
	if err != nil {
		return errors.Wrapf(err, "failed to create %s", f.Dir)
	}
	defer fd.Close()

	if _, err = fd.Write(f.Content); err != nil {
		return errors.Wrap(err, "failed to write auth file")
	}
	return nil
}

func GetFromDefault(region string) (*conf.NsConfig, *conf.PfConfig, error) {
	var (
		nsd string
		pfd string
	)

	switch region {
	case "EU868":
		nsd = conf.Eu868NsDefault
		pfd = conf.Eu868PfDefault
	case "US915":
		nsd = conf.Us915NsDefault
		pfd = conf.Us915PfDefault
	default:
		return nil, nil, errors.Errorf("unsupported region: %s", region)
	}

	var (
		nsc conf.NsConfig
		pfc conf.PfConfig
	)

	

	nsv := viper.New()
	nsv.SetConfigType("toml")
	if err := nsv.ReadConfig(strings.NewReader(nsd)); err != nil {
		fmt.Println("get ns default config fail:", err)
	}
	if err := nsv.Unmarshal(&nsc); err != nil {
		return nil, nil, errors.Errorf("unmarshl ns default config fail: %s", err)
	}
	if err := json.Unmarshal([]byte(pfd), &pfc); err != nil {
		return nil, nil, errors.Errorf("unmarshl pf default config fail: %s", err)
	}

	return &nsc, &pfc, nil
}
