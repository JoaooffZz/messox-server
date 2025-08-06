package connection

import (
	"encoding/json"
	m "ws/models"

	"github.com/gorilla/websocket"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)


type Client struct {
   // identificador
   Id string
   // A conexão web socket
   Conn *websocket.Conn
   // Endereço de hub
   Hub *Hub
   // canal de mensagem
   Send chan []byte
}

func NewClient(conn *websocket.Conn, hub *Hub, id string) *Client {
   return &Client{
      Id: id,
      Conn: conn,
      Hub: hub,
      Send: make(chan []byte, 256),
   }
}

func (c *Client)ReadPump() {
   defer func(){
      c.Hub.Unregister <- c
      c.Conn.Close()
   }()
   for {
      _, message, err := c.Conn.ReadMessage()
      if err != nil {
         break
      }

      var msg m.Message
      err = json.Unmarshal(message, &msg)
      if err != nil {
         break
      }
      c.Hub.Broadcast <- &msg
      
      // como fica o banco de dados?
      // seria melhor só repassar a mensagem para o destino,
      // o hub ja faz o cancelamento do registro dos cache-user-on
      // e quando a mensagem for enviada, e só chamar uma goroutine
      // que ira receber a mensagem e salvar no banco de dados
      // tratando (conversa nova, conversa antiga)

      // o usuario que receber a mensagem, o app devera tratala fazendo de uso
      // do (db-local)


   }
}