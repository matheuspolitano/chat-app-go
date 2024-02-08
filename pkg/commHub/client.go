package commHub

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/matheuspolitano/chat-app-go/pkg/model"
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

type MessageClient struct {
	client   *Client
	messagem []byte
}

type Client struct {
	send     chan MessageClient
	conn     *websocket.Conn
	username string
	hub      *Hub
}

func NewClient(hub *Hub, conn *websocket.Conn, username string) *Client {
	client := &Client{
		send:     make(chan MessageClient),
		conn:     conn,
		hub:      hub,
		username: username,
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
		log.Printf("received from %s :%s", c.username, string(message))
		c.hub.message <- MessageClient{c, message}
	}
}

func (c *Client) WritePump() {
	defer func() {
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				log.Printf("%s send channel closed", c.username)
				return
			}
			userMessage := model.UserMessage{
				Username: message.client.username,
				Message:  string(message.messagem),
			}
			sendMessage, err := json.Marshal(userMessage)
			if err != nil {
				log.Fatal("Error marshalling user message to JSON")
			}
			err = c.conn.WriteMessage(websocket.TextMessage, sendMessage)
			if err != nil {
				log.Printf("Error writing to client %v", err)
				return
			}
			log.Println("WritePump sent:", string(message.messagem))

		}
	}
}
