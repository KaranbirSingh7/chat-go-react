package main

import (
	"fmt"
	"net/http"

	"github.com/karanbirsingh7/chat-go-react/backend/pkg/websocket"
)

// define out websocket route/path/endpoint
func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	//upgrade this connection to a websocket
	conn, err := websocket.Upgrade(w, r)

	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}

	//Create new client
	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	// Register client in pool
	pool.Register <- client

	// Listen for any messages from client
	client.Read()
}

func setupRoutes() {
	pool := websocket.NewPool()

	// This triggers a while loop in background where channels become active and listen if anything comes up
	go pool.Start()

	// Root handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server is running")
	})

	// WebSocket handler
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})

}

func main() {
	fmt.Println("Chat App v0.1.0")
	setupRoutes()
	http.ListenAndServe(":8000", nil)
}
