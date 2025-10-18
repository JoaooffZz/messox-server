package cache

type Storage struct {
	ClientsIPs []ClientIP `json:"clients_ips"`
}
type ClientIP struct {
	IP       string `json:"ip"`
	IsSecure bool   `json:"is_secure"`
	ReqCount *int   `json:"request_count"`
}

type FirewallCacheStorage struct{}

func (cache *FirewallCacheStorage) New() map[string]*ClientIP {
	return map[string]*ClientIP{}
}

func (cache *FirewallCacheStorage) RestoreBackup(storage Storage) (map[string]*ClientIP, error) {
	mapClients := make(map[string]*ClientIP)

	for _, c := range storage.ClientsIPs {
		mapClients[c.IP] = &c
	}

	return mapClients, nil
}
