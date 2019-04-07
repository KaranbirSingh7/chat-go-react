package websocket

import (
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

// Upgrade func helps endppoint to upgrade to a websocket
func Upgrade(w http.ResponseWriter, r *http.Request) (conn *websocket.Conn, err error) {
	//upgrade connection to a websocket
	conn, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// listen indefinitely for new message coming through websocket connection
	return
}

// // Reader used for reading incoming messages
// func Reader(conn *websocket.Conn) {
// 	// GoLang while loop
// 	for {
// 		// read message
// 		messageType, p, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}

// 		// print what is being recived on endpoint
// 		fmt.Println(string(p))
// 		if err := conn.WriteMessage(messageType, p); err != nil {
// 			log.Println(err)
// 			return
// 		}
// 	}
// }

// // Writer used for writing outsgoing messages
// func Writer(conn *websocket.Conn) {
// 	for {
// 		fmt.Println("Sending")

// 		messageType, r, err := conn.NextReader()
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		w, err := conn.NextWriter(messageType)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		if _, err := io.Copy(w, r); err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		if err := w.Close(); err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 	}
// }
