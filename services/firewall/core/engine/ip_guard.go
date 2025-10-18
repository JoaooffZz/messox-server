package engine

import (
	fc "firewall/cache"
	"time"
)

type ServiceIpGuard struct {
	checkInterval int
	reqLimit      int
	mapClients    map[string]*fc.ClientIP
}

func (service *ServiceIpGuard) analyzeIPTraffic() {
	for {
		duration := time.Duration(service.checkInterval) * time.Second
		timer := time.NewTimer(duration)
		<-timer.C
		for k, value := range service.mapClients {
			if *value.ReqCount > service.reqLimit {
				service.mapClients[k].IsSecure = false
			}
			var restart int = 0
			value.ReqCount = &restart
		}
	}
}

// VerifySecureIP(ip string) (bool, error)

// SetUpReqCount()
