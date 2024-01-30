package main

import (
	"net/http"
	"time"
)

func main() {
	staticDir := http.Dir("./public")
	fileServer := http.FileServer(staticDir)

	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// Serve index.html for root URL
	http.Handle("/", fileServer)

	server := &http.Server{
		Addr:              ":8089",
		ReadHeaderTimeout: 3 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
