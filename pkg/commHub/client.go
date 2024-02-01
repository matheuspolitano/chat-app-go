package commHub

import (
	"log"
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

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

type message chan []byte

type Client struct {
	send     chan []byte
	conn     *websocket.Conn
	username string
	hub      *Hub
}

func NewClient(hub *Hub, conn *websocket.Conn) *Client {
	client := &Client{
		send:     make(chan []byte, 256),
		conn:     conn,
		hub:      hub,
		username: "MatheusPolitano",
	}
	hub.incoming <- client
	return client
}
func (c *Client) ReadPump() {
	defer func() {
		c.hub.leaving <- c
		c.conn.Close()
	}()
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("ReadPump error:", err)
			break
		}
		log.Println("ReadPump received:", string(message))
		c.hub.message <- message
	}
}

func (c *Client) WritePump() {
	defer func() {
		c.conn.Close()
	}()
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				log.Println("WritePump: send channel closed")
				return
			}
			err := c.conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Printf("Error writing to client %v", err)
				return
			}
			log.Println("WritePump sent:", string(message))

		case <-ticker.C:
			// This is just an example to demonstrate sending a message periodically.
			// In a real application, you would have logic to decide what messages to send.
			testMsg := "Ping from server"

			log.Println("WritePump ticker sent:", testMsg)
		}
	}
}
