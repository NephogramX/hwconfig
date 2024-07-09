package conf

import (
	"bytes"
	"strconv"
	"strings"

	"gitee.com/dfrobotcd/chirpstack-api/go/as/external/api"
	"github.com/pkg/errors"
)

var ErrorBsConfigFileNotFound = errors.New("bs config file not found")

type BsConfig struct {
	BsAuth
}

type BsAuth interface {
	GetAuth(dir string) *Auth
	ApiMode() *api.GateWayMode
}

type NoAuth struct {
	UseCups    bool
	ServerAddr string
	ServerPort int32
}

type TlsServerAuth struct {
	NoAuth
	Trust []byte
}

type TlsServerAndClientAuth struct {
	NoAuth
	Key   []byte
	Trust []byte
	Crt   []byte
}

type TlsServerAuthAndClientToken struct {
	NoAuth
	Trust []byte
	Key   []byte
}

func GenerateUri(cups bool, tls bool, addr string, port int) string {
	uri := bytes.Buffer{}
	if cups {
		if tls {
			uri.WriteString("https://")
		} else {
			uri.WriteString("http://")
		}
	} else {
		if tls {
			uri.WriteString("wss://")
		} else {
			uri.WriteString("ws://")
		}
	}
	uri.WriteString(addr)
	uri.WriteString(":")
	uri.WriteString(strconv.Itoa(int(port)))
	return uri.String()
}

func GenerateDir(cups bool, path string, suffix string) string {
	dir := strings.Builder{}
	dir.WriteString(path)
	if cups {
		dir.WriteString("cups")
	} else {
		dir.WriteString("tc")
	}
	dir.WriteString(suffix)
	return dir.String()
}

func (c *NoAuth) GetAuth(dir string) *Auth {
	return &Auth{
		Uri: &File{
			Dir:     GenerateDir(c.UseCups, dir, ".uri"),
			Content: []byte(GenerateUri(c.UseCups, false, c.ServerAddr, int(c.ServerPort))),
		},
		Key:   nil,
		Crt:   nil,
		Trust: nil,
	}
}

func (c *NoAuth) ApiMode() *api.GateWayMode {
	a := &api.GateWayMode{
		Mode: "BS",
		ModeConfig: &api.GateWayMode_Bs{
			Bs: &api.BasicsStation{
				Type:   "LNS",
				Server: c.ServerAddr,
				Port:   int64(c.ServerPort),
				Auth: &api.BSAuth{
					Mode:    "NO_AUTH",
					CaCert:  "",
					CliCert: "",
					CliKey:  "",
					Token:   "",
				},
			},
		},
	}
	if c.UseCups {
		a.GetBs().Type = "CUPS"
	}
	return a
}

func (c *TlsServerAuth) GetAuth(dir string) *Auth {
	return &Auth{
		Uri: &File{
			Dir:     GenerateDir(c.UseCups, dir, ".uri"),
			Content: []byte(GenerateUri(c.UseCups, true, c.ServerAddr, int(c.ServerPort))),
		},
		Trust: &File{
			Dir:     GenerateDir(c.UseCups, dir, ".trust"),
			Content: c.Trust,
		},
		Key: nil,
		Crt: nil,
	}
}

func (c *TlsServerAuth) ApiMode() *api.GateWayMode {
	a := c.NoAuth.ApiMode()
	a.GetBs().Auth.Mode = "TLS_Server"
	a.GetBs().Auth.CaCert = string(c.Trust)
	return a
}

func (c *TlsServerAndClientAuth) GetAuth(dir string) *Auth {
	return &Auth{
		Uri: &File{
			Dir:     GenerateDir(c.UseCups, dir, ".uri"),
			Content: []byte(GenerateUri(c.UseCups, true, c.ServerAddr, int(c.ServerPort))),
		},
		Key: &File{
			Dir:     GenerateDir(c.UseCups, dir, ".key"),
			Content: c.Key,
		},
		Crt: &File{
			Dir:     GenerateDir(c.UseCups, dir, ".crt"),
			Content: c.Crt,
		},
		Trust: &File{
			Dir:     GenerateDir(c.UseCups, dir, ".trust"),
			Content: c.Trust,
		},
	}
}

func (c *TlsServerAndClientAuth) ApiMode() *api.GateWayMode {
	a := c.NoAuth.ApiMode()
	a.GetBs().Auth.Mode = "TLS_Server_Client"
	a.GetBs().Auth.CliKey = string(c.Key)
	a.GetBs().Auth.CaCert = string(c.Trust)
	a.GetBs().Auth.CliCert = string(c.Crt)
	return a
}

func (c *TlsServerAuthAndClientToken) GetAuth(dir string) *Auth {
	return &Auth{
		Uri: &File{
			Dir:     GenerateDir(c.UseCups, dir, ".uri"),
			Content: []byte(GenerateUri(c.UseCups, true, c.ServerAddr, int(c.ServerPort))),
		},
		Key: &File{
			Dir:     GenerateDir(c.UseCups, dir, ".key"),
			Content: c.Key,
		},
		Trust: &File{
			Dir:     GenerateDir(c.UseCups, dir, ".trust"),
			Content: c.Trust,
		},
		Crt: nil,
	}
}

func (c *TlsServerAuthAndClientToken) ApiMode() *api.GateWayMode {
	a := c.NoAuth.ApiMode()
	a.GetBs().Auth.Mode = "TLS_Server_Client_Token"
	a.GetBs().Auth.Token = string(c.Key)
	a.GetBs().Auth.CaCert = string(c.Trust)
	return a
}

type File struct {
	Dir     string
	Content []byte
}

type Auth struct {
	Uri   *File
	Key   *File
	Crt   *File
	Trust *File
}
