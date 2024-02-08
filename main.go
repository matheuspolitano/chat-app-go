package main

import (
	"net/http"
	"os"
	"time"

	"github.com/matheuspolitano/chat-app-go/pkg/commHub"
	"github.com/matheuspolitano/chat-app-go/pkg/handlers"
)

func main() {
	args := os.Args
	port := "8080"
	if len(args) >= 2 {
		port = args[1]
	}
	staticDir := http.Dir("./public")
	fileServer := http.FileServer(staticDir)
	hub := commHub.NewHub()
	go hub.Run()
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	http.Handle("/", fileServer)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeWs(w, r, hub)
	})

	server := &http.Server{
		Addr:              ":" + port,
		ReadHeaderTimeout: 3 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
