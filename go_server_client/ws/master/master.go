package master

import (
	"net/http"
	connWS "ws/connection"
	model "ws/models"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

type ServerWS struct {
	hub *connWS.Hub
	w http.ResponseWriter
	r *http.Request
	msg model.Message
}

func (ws *ServerWS)serveWs() {
	conn, err := upgrader.Upgrade(ws.w, ws.r, nil)
	if err != nil {
		// log.Println(err)
		return
	}
	client := &connWS.Client{
		Id: ws.msg.FromID,
		Hub: ws.hub, 
		Conn: conn, 
		Send: make(chan []byte, 256)}
	client.Hub.Register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	// go client.writePump()
	// go client.readPump()
}