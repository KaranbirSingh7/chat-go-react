package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Define upgrader which requires a READ and WRITE size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// CORS probably??
	CheckOrigin: func(r *http.Request) bool { return true },
}

// Reader: listen for new messages being sent to our WebSocket endpoint
func reader(conn *websocket.Conn) {
	// GoLang while loop
	for {
		// read message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// print what is being recived on endpoint
		fmt.Println(string(p))
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

// define out websocket route/path/endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Message recieved from: ", r.Host)

	//upgrade this connection to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	// listen indefinitely for new message coming through websocket connection
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})

	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Chat App v0.1.0")
	setupRoutes()
	http.ListenAndServe(":8000", nil)
}
