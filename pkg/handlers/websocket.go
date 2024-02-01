package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/matheuspolitano/chat-app-go/pkg/commHub"
)

var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow all origins. In a production environment, replace this with proper origin validation
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ServeWs(w http.ResponseWriter, r *http.Request, hub *commHub.Hub) {
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalln(err)
		return
	}
	client := commHub.NewClient(hub, conn)
	go client.WritePump()
	go client.ReadPump()

}
