package master

import (
	"net/http"
	connWs "ws/connection"

	// model "ws/models"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // aceita qualquer origem
	},
}

type ServerWS struct {
	Id string
	Hub *connWs.Hub
	W  http.ResponseWriter
	R *http.Request
}

func (s *ServerWS)Run(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(s.W, s.R, nil)
	if err != nil {
		return
	}

	// if err != nil {
	// 	ctx.Writer.WriteHeader(http.StatusInternalServerError)
	// 	// ctx.JSON(http.StatusInternalServerError, 
	// 	// 	gin.H{"error ws":"error up-conn-ws"},
	// 	// )
	// 	return
	// }
	// ctx.JSON(http.StatusSwitchingProtocols, nil)

	client := connWs.NewClient(conn, s.Hub, s.Id)
	s.Hub.Register <- client

	go client.ReadPump()
	go client.WritePump()
}