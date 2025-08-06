package connection

import (
	m "ws/models"
)

type Hub struct {
	// Registered clients.
	clients map[string]*Client

	// Inbound messages from the clients.
	Broadcast chan *m.Message

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan *m.Message),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		clients:    make(map[string]*Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		    case client := <- h.Register:
				h.clients[client.Id] = client
				
			case client := <- h.Unregister:
				if _, ok := h.clients[client.Id]; ok {
					delete(h.clients, client.Id)
					close(client.Send)
				}

			case message := <-h.Broadcast:
				to, ok := h.clients[message.To]
				if !ok {
				    // chama a goroutine para salvar a mensagem no
					// banco de dados de notificação.
				}

				// chama to.WritePump
		}
	}
}