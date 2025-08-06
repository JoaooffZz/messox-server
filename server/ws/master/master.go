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
}

type ServerWS struct {
	id string
	// client *connWS.Client
	hub *connWs.Hub
	w  http.ResponseWriter
	r *http.Request
	// msg model.Message
}

func (s *ServerWS)Run(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(s.w, s.r, nil)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, 
			gin.H{"error ws":"error up-conn-ws : %v", err},
		)
		return
	}

	client := connWs.NewClient(conn, s.hub, s.id)
	// go client.
	// client := &connWS.Client{
	// 	Id: ws.msg.FromID,
	// 	Hub: ws.hub, 
	// 	Conn: conn, 
	// 	Send: make(chan []byte, 256)}
	// s.client.Hub.Register <- s.client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	// go client.writePump()
	// go client.readPump()
}