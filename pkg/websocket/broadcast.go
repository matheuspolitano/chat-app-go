package websocket

import "github.com/gorilla/websocket"

type message chan<- string

type Client struct {
	send     message
	conn     *websocket.Conn
	username string
}

type Hub struct {
	clients  map[*Client]bool
	message  chan string
	incoming chan *Client
	leaving  chan *Client
}

func NewHub() *Hub {
	return &Hub{
		clients:  make(map[*Client]bool),
		message:  make(chan string),
		incoming: make(chan *Client),
		leaving:  make(chan *Client),
	}
}

func (hub *Hub) Run() {
	for {
		select {
		case m := <-hub.message:
			for client := range hub.clients {
				select {
				case client.send <- m:
				}
			}

		case in := <-hub.incoming:
			hub.clients[in] = true

		case l := <-hub.leaving:
			if _, ok := hub.clients[l]; ok {
				delete(hub.clients, l)
				close(l.send)
			}
		}
	}
}
