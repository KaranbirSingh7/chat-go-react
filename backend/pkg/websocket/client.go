package websocket

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read() {
	// anonymous func runs ar end to empty channel/pool and close socket connection
	defer func() {
		// send value to channel
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		message := Message{
			Type: messageType,
			Body: string(p),
		}

		// send message to channel
		c.Pool.Broadcast <- message
		fmt.Printf("Message receieved: %+v\n", message)
	}
}
