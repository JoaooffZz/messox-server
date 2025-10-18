package connection

import (
	"net/http"

	mb "mb/ports"

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
	ClientID  int
	HandlerMB mb.HandlerMB
	W         http.ResponseWriter
	R         *http.Request
}

func (s *ServerWS) Run() {
	conn, err := upgrader.Upgrade(s.W, s.R, nil)
	if err != nil {
		return
	}

	client := Client{
		ID:        s.ClientID,
		Conn:      conn,
		HandlerMB: s.HandlerMB,
		Send:      make(chan []byte),
	}

	go client.ReadPump()
	go client.WritePump()
}
