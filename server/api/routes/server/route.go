package server

import (
	portsDB "ports/db"
	ping "server/ping"
	ws "server/ws"
	connWs "ws/connection"

	"github.com/gin-gonic/gin"
)

type RoutesServer struct {
	Ping ping.RoutePing
	Ws ws.RouteWs
}

func New(eng *gin.Engine, db portsDB.DB, keyPem *[]byte, 
	apiKey *string, hub *connWs.Hub) RoutesServer {
	return RoutesServer{
		Ping: ping.RoutePing{
			Eng: eng,
			ApiKey: apiKey,
		},
		Ws: ws.RouteWs{
			Eng: eng,
			DB: db,
			Hub: hub,
			KeyPem: keyPem,
		},
	}
}