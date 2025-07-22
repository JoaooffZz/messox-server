package connection

type Hub struct {
	// Registered clients.
	clients map[string]*Client

	// Inbound messages from the clients.
	Broadcast chan []byte

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		clients:    make(map[string]*Client),
	}
}

func (h *Hub) run() {
	for {
		select {
		    case client := <- h.Register:
				h.clients[client.Id] = client
			case client := <- h.Unregister:
				if _, ok := h.clients[client.Id]; ok {
					delete(h.clients, client.Id)
					close(client.Send)
				}
		}
	}
}