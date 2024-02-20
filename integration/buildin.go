package integration

import "github.com/NephogramX/hwconfig/configfile/pf"

type BuildinIntegration struct {
}

func NewBuildinIntegration() *BuildinIntegration {
	return &BuildinIntegration{}
}

func (b *BuildinIntegration) Get() IntegrationType {
	return Buildin
}

func (b *BuildinIntegration) HandleUdpPacketForwarder(cf *pf.UdpPacketForwarder) error {
	cf.GateWayConfig.ServerAddress = "localhost"
	cf.GateWayConfig.ServPortUp = 1700
	cf.GateWayConfig.ServPortDown = 1700
	return nil
}

func (b *BuildinIntegration) HandleBasicStation() error {
	return nil
}
