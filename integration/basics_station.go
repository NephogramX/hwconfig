package integration

import (
	"strconv"

	cf "github.com/NephogramX/hwconfig/configfile"
	"github.com/pkg/errors"
)

const LetsEncryptRootCA = `-----BEGIN CERTIFICATE-----
MIIFazCCA1OgAwIBAgIRAIIQz7DSQONZRGPgu2OCiwAwDQYJKoZIhvcNAQELBQAw
TzELMAkGA1UEBhMCVVMxKTAnBgNVBAoTIEludGVybmV0IFNlY3VyaXR5IFJlc2Vh
cmNoIEdyb3VwMRUwEwYDVQQDEwxJU1JHIFJvb3QgWDEwHhcNMTUwNjA0MTEwNDM4
WhcNMzUwNjA0MTEwNDM4WjBPMQswCQYDVQQGEwJVUzEpMCcGA1UEChMgSW50ZXJu
ZXQgU2VjdXJpdHkgUmVzZWFyY2ggR3JvdXAxFTATBgNVBAMTDElTUkcgUm9vdCBY
MTCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIBAK3oJHP0FDfzm54rVygc
h77ct984kIxuPOZXoHj3dcKi/vVqbvYATyjb3miGbESTtrFj/RQSa78f0uoxmyF+
0TM8ukj13Xnfs7j/EvEhmkvBioZxaUpmZmyPfjxwv60pIgbz5MDmgK7iS4+3mX6U
A5/TR5d8mUgjU+g4rk8Kb4Mu0UlXjIB0ttov0DiNewNwIRt18jA8+o+u3dpjq+sW
T8KOEUt+zwvo/7V3LvSye0rgTBIlDHCNAymg4VMk7BPZ7hm/ELNKjD+Jo2FR3qyH
B5T0Y3HsLuJvW5iB4YlcNHlsdu87kGJ55tukmi8mxdAQ4Q7e2RCOFvu396j3x+UC
B5iPNgiV5+I3lg02dZ77DnKxHZu8A/lJBdiB3QW0KtZB6awBdpUKD9jf1b0SHzUv
KBds0pjBqAlkd25HN7rOrFleaJ1/ctaJxQZBKT5ZPt0m9STJEadao0xAH0ahmbWn
OlFuhjuefXKnEgV4We0+UXgVCwOPjdAvBbI+e0ocS3MFEvzG6uBQE3xDk3SzynTn
jh8BCNAw1FtxNrQHusEwMFxIt4I7mKZ9YIqioymCzLq9gwQbooMDQaHWBfEbwrbw
qHyGO0aoSCqI3Haadr8faqU9GY/rOPNk3sgrDQoo//fb4hVC1CLQJ13hef4Y53CI
rU7m2Ys6xt0nUW7/vGT1M0NPAgMBAAGjQjBAMA4GA1UdDwEB/wQEAwIBBjAPBgNV
HRMBAf8EBTADAQH/MB0GA1UdDgQWBBR5tFnme7bl5AFzgAiIyBpY9umbbjANBgkq
hkiG9w0BAQsFAAOCAgEAVR9YqbyyqFDQDLHYGmkgJykIrGF1XIpu+ILlaS/V9lZL
ubhzEFnTIZd+50xx+7LSYK05qAvqFyFWhfFQDlnrzuBZ6brJFe+GnY+EgPbk6ZGQ
3BebYhtF8GaV0nxvwuo77x/Py9auJ/GpsMiu/X1+mvoiBOv/2X/qkSsisRcOj/KK
NFtY2PwByVS5uCbMiogziUwthDyC3+6WVwW6LLv3xLfHTjuCvjHIInNzktHCgKQ5
ORAzI4JMPJ+GslWYHb4phowim57iaztXOoJwTdwJx4nLCgdNbOhdjsnvzqvHu7Ur
TkXWStAmzOVyyghqpZXjFaH3pO3JLF+l+/+sKAIuvtd7u+Nxe5AW0wdeRlN8NwdC
jNPElpzVmbUq4JUagEiuTDkHzsxHpFKVK7q4+63SM1N95R1NbdWhscdCb+ZAJzVc
oyi3B43njTOQ5yOf+1CceWxG1bQVs5ZufpsMljq4Ui0/1lvh+wjChP4kqKOJ2qxq
4RgqsahDYVvTH9w7jXbyLeiNdd8XM2w9U/t7y0Ff/9yi0GE44Za4rF2LN9d11TPA
mRGunUHBcnWEvgJBQl9nJEiU0Zsnvgc/ubhPgXRR4Xq37Z0j4r7g1SgEEzwxA57d
emyPxgcYxn/eR44/KJ4EBs+lVDR3veyJm+kXQ99b21/+jh5Xos1AnX5iItreGCc=
-----END CERTIFICATE-----`

type Portocol = string

const (
	LNS  Portocol = "LNS"
	CUPS Portocol = "CUPS"
)

/*
 * Basics Station has 3 kinds of authentication ways:
 * 1. No authentication: *.uri
 * 2. TLS Server Authentication: *uri *.trust
 * 3. TLS Server and Client Authentication: *.uri *.key *.trust *.crt
 * 4. TLS Server Authentication and Client Token: *.uri *.trust *.key
 */

type Authentication struct {
	Key        *[]byte // *.key
	ServerCert *[]byte // *.trust
	ClientCert *[]byte // *.crt
}

type BasicsStationSetting struct {
	Protocol      Portocol
	ServerAddress string
	ServerPort    int32

	Authentication
}

// basic station integration
type BasicsStationIntegration struct {
	filename string
	Uri      *[]byte
	Authentication
}

func NewBasicsStationIntegration(s *BasicsStationSetting) (*BasicsStationIntegration, error) {
	i := &BasicsStationIntegration{}
	tlsDisable := false
	protocol := ""

	if s.ServerCert == nil && s.ClientCert == nil && s.Key == nil {
		tlsDisable = true
	} else if s.ServerCert != nil && s.Key == nil && s.ClientCert == nil {
	} else if s.ServerCert != nil && s.Key != nil && s.ClientCert == nil {
	} else if s.ServerCert != nil && s.Key != nil && s.ClientCert != nil {
	} else {
		return nil, errors.New("invalid authentication format")
	}

	switch s.Protocol {
	case LNS:
		i.filename = "tc"
		if tlsDisable {
			protocol = "ws"
		} else {
			protocol = "wss"
		}
	case CUPS:
		i.filename = "cups"
		if tlsDisable {
			protocol = "http"
		} else {
			protocol = "https"
		}
	default:
		return nil, errors.Errorf("unsupported protocol: %v", s.Protocol)
	}

	Uri := []byte(protocol + "://" + s.ServerAddress + ":" + strconv.Itoa(int(s.ServerPort)))
	i.Uri = &Uri
	i.Authentication = s.Authentication

	return i, nil
}

func (i *BasicsStationIntegration) Type() IntegrationType {
	return BasicsStation
}

func (i *BasicsStationIntegration) HandleBasicsStationUri() *cf.BasicsStation {
	if i.Uri == nil {
		return nil
	}
	return cf.NewAuthenticationFile(cf.File{Name: i.filename + ".uri", Path: BSPath}, i.Uri)
}

func (i *BasicsStationIntegration) HandleBasicsStationKey() *cf.BasicsStation {
	if i.Key == nil {
		return nil
	}
	return cf.NewAuthenticationFile(cf.File{Name: i.filename + ".key", Path: BSPath}, i.Key)
}

func (i *BasicsStationIntegration) HandleBasicsStationCrt() *cf.BasicsStation {
	if i.ClientCert == nil {
		return nil
	}
	return cf.NewAuthenticationFile(cf.File{Name: i.filename + ".crt", Path: BSPath}, i.ClientCert)
}

func (i *BasicsStationIntegration) HandleBasicsStationTrust() *cf.BasicsStation {
	if i.ServerCert == nil {
		return nil
	}
	return cf.NewAuthenticationFile(cf.File{Name: i.filename + ".trust", Path: BSPath}, i.ServerCert)
}

func (i *BasicsStationIntegration) HandleUdpPacketForwarder() *cf.UdpPacketForwarder {
	return nil
}

func (i *BasicsStationIntegration) HandleGatewayBridge() *cf.GatewayBridge {
	return nil
}

func (i *BasicsStationIntegration) HandleNetworkServer() *cf.NetworkServer {
	return nil
}
