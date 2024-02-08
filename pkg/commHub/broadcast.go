package commHub

type Hub struct {
	clients  map[*Client]bool
	message  chan MessageClient
	incoming chan *Client
	leaving  chan *Client
}

func NewHub() *Hub {
	return &Hub{
		clients:  make(map[*Client]bool),
		message:  make(chan MessageClient),
		incoming: make(chan *Client),
		leaving:  make(chan *Client),
	}
}

func (hub *Hub) Run() {
	for {
		select {
		case in := <-hub.incoming:
			hub.clients[in] = true
		case m := <-hub.message:
			for client := range hub.clients {
				if m.client != client {
					select {
					case client.send <- m:
					}
				}
			}

		case l := <-hub.leaving:
			if _, ok := hub.clients[l]; ok {
				delete(hub.clients, l)
				close(l.send)
			}
		}
	}
}
