package connection

import (
	"encoding/json"
	m "ws/models"
)
type Hub struct {
	// Registered clients.
	clients map[string]*Client

	// Inbound messages from the clients.
	Broadcast chan *m.WsEvent

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan *m.WsEvent),
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

			case event := <- h.Broadcast:    
				adderess, ok := h.clients[event.Adderess]
				if !ok {
	                // chama a goroutine para salvar a mensagem no
				    // banco de dados de notificação.
					return
				}
				msg, err :=json.Marshal(event)
				if err != nil {
					// chama log
					return
				}
				select {
					case adderess.Send <- msg:
						// chama uma go routine para salvar os dados
						// do evento no db.
					default:
						h.Unregister <- adderess
				}
		}
	}
}