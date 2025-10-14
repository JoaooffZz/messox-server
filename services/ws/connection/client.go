package connection

import (
	"encoding/json"
	"time"
	m "ws/models"

	"github.com/gorilla/websocket"
)

const (
   // Time allowed to write a message to the peer.
   writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
)

type Client struct {
   // identificador
   ID int

   // A conexão web socket
   Conn *websocket.Conn

   // Endereço de hub
   Hub *Hub

   // Canal de mensagem
   Send chan []byte
}

func (c *Client)ReadPump() {
   defer func(){
      c.Hub.Unregister <- c
      c.Conn.Close()
   }()
   for {
      c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	   c.Conn.SetPongHandler(
         func(string) error {
            c.Conn.SetReadDeadline(time.Now().Add(pongWait));
            return nil
         },
      )

      _, message, err := c.Conn.ReadMessage()
      if err != nil {
         return
      }

      var event m.WsEvent
      err = json.Unmarshal(message, &event)
      if err != nil {
         return
      }
      event.Sender = &c.ID
      c.Hub.Broadcast <- &event
   }
}

func (c *Client)WritePump() {
   ticker := time.NewTicker(pingPeriod)
   defer func() {
      ticker.Stop()
      c.Hub.Unregister <- c
      c.Conn.Close()
   }()
   for {
      select{        
         case send, ok := <- c.Send:
            if !ok {
               c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
               // caso msg nao esteja vazia
               // chama uma go routine para salvar message no db BoxMessage
               return
            }
            c.Conn.WriteMessage(websocket.TextMessage, send)
         
         case <- ticker.C:
            c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
            if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
               return
            }
      }
   }
}   