package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// Define upgrader which requires a READ and WRITE size
var upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	CheckOrigin: 
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})
}

func main() {
	setupRoutes()
	http.ListenAndServe(":8000", nil)
}
