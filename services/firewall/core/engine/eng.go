package engine

import (
	fc "firewall/cache"
)

type EngineFirewall struct {
	config     *ConfigFirewal
	backup     *[]byte
	mapClients map[string]*fc.ClientIP
	isRun      bool
	IPGuard    ServiceIpGuard
}

// Default()

func New(config *ConfigFirewal, backup *[]byte) *EngineFirewall {
	return &EngineFirewall{
		config: config,
		backup: backup,
		isRun:  false,
	}
}

func (eng *EngineFirewall) Run() error {

	eng.IPGuard = ServiceIpGuard{
		checkInterval: *eng.config.CheckInterval,
		mapClients:    eng.mapClients,
	}
	eng.isRun = true
	eng.IPGuard.analyzeIPTraffic()
	return nil
}
