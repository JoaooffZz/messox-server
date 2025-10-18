package connection

import (
	"context"
	"encoding/json"
	mb "mb/ports"
	"strconv"
	wsModels "ws/models"

	"time"

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
	ID        int
	Conn      *websocket.Conn
	HandlerMB mb.HandlerMB
	Send      chan []byte
	closed    chan struct{}
}

func (c *Client) ReadPump() {
	c.closed = make(chan struct{})
	defer func() {
		close(c.closed)
		c.Conn.Close()
	}()
	for {
		c.Conn.SetReadDeadline(time.Now().Add(pongWait))
		c.Conn.SetPongHandler(
			func(string) error {
				c.Conn.SetReadDeadline(time.Now().Add(pongWait))
				return nil
			},
		)

		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			return
		}

		var event wsModels.WsEvent
		err = json.Unmarshal(message, &event)
		if err != nil {
			continue
		}
		event.Sender = &c.ID

		eventByte, _ := json.Marshal(event)

		c.HandlerMB.CreatePubWsEvents(mb.WsPubMsg{
			ContentType: "application/json",
			Body:        eventByte,
			RoutingKey:  c.HandlerMB.Build().PubWsEvents(c.ID, *event.Adderess),
		})
	}
}

func (c *Client) WritePump() {
	ctx, cancel := context.WithCancel(context.Background())
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		cancel()
		ticker.Stop()
		c.Conn.Close()
	}()
	msg, err := c.HandlerMB.OpenConsumerWsEvents(ctx,
		mb.WsConsumer{
			QueueName:  strconv.Itoa(c.ID),
			RoutingKey: c.HandlerMB.Build().ConWsEvents(c.ID),
		})
	if err != nil {
		return
	}
	for {
		select {
		case m, ok := <-msg:
			if !ok {
				return
			}
			c.Conn.WriteMessage(websocket.TextMessage, m)

		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		case <-c.closed:
			return
		}
	}
}
