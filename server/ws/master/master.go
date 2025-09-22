package master

import (
	"net/http"
	connWs "ws/connection"

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
	UserID int
	Hub *connWs.Hub
	W  http.ResponseWriter
	R *http.Request
}

func (s *ServerWS)Run() {
	conn, err := upgrader.Upgrade(s.W, s.R, nil)
	if err != nil {
		return
	}

	client := &connWs.Client{
		ID: s.UserID,
		Conn: conn,
		Hub: s.Hub,
		Send: make(chan []byte),
	}
	s.Hub.Register <- client

	go client.ReadPump()
	go client.WritePump()
}