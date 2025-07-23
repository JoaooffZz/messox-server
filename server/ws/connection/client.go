package connection

import (
	"github.com/gorilla/websocket"
)

type Client struct {
   // identificador
   Id string
   // A conexão web socket
   Conn *websocket.Conn
   // Endereço de hub
   Hub *Hub
   // canal de escrever
   Send chan []byte
}