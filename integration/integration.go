package integration

import "github.com/NephogramX/hwconfig/configfile/pf"

type IntegrationType int32

const (
	Buildin            IntegrationType = iota
	UdpPacketForwarder IntegrationType = iota
	BasicsStation      IntegrationType = iota
)

type Integration interface {
	Get() IntegrationType

	HandleUdpPacketForwarder(cf *pf.UdpPacketForwarder) error

	HandleBasicStation() error
}
