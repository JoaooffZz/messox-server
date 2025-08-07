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
   Id string
   // A conexão web socket
   Conn *websocket.Conn
   // Endereço de hub
   Hub *Hub
   // canal de mensagem
   Send chan *m.Message
}

func NewClient(conn *websocket.Conn, hub *Hub, id string) *Client {
   return &Client{
      Id: id,
      Conn: conn,
      Hub: hub,
      Send: make(chan *m.Message),
   }
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

      var msg m.Message
      err = json.Unmarshal(message, &msg)
      if err != nil {
         return
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

func (c *Client)WritePump() {
   ticker := time.NewTicker(pingPeriod)
   defer func() {
      ticker.Stop()
      c.Hub.Unregister <- c
      c.Conn.Close()
   }()
   for {
      select{        
         case msg, ok := <- c.Send:
            if !ok {
               c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
               // caso msg nao esteja vazia
               // chama uma go routine para salvar message no db BoxMessage
               return
            }
            // chama uma go routine para salvar a mensagem no banco de dados
            msgBytes, _ := json.Marshal(msg)
            c.Conn.WriteMessage(websocket.TextMessage, msgBytes)
         
         case <- ticker.C:
            c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
            if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
               return
            }
      }
   }
}      