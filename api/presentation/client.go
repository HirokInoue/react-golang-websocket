package presentation

import (
	"github.com/gorilla/websocket"
)

type Body struct {
	Name string      `json:"name"`
	Ok   bool        `json:"ok"`
	Data interface{} `json:"data"`
}

// Client is responsible for read request and write response
func NewClient(s *websocket.Conn, f FindHandler) *Client {
	return &Client{
		send:        make(chan Body),
		socket:      s,
		findHandler: f,
	}
}

type Client struct {
	send        chan Body
	socket      *websocket.Conn
	findHandler FindHandler
}

func (c *Client) Read() {
	var body Body
	for {
		if err := c.socket.ReadJSON(&body); err != nil {
			break
		}
		if handler, ok := c.findHandler(body.Name); ok {
			handler.exec(c, body.Data)
		}
	}
	c.socket.Close()
}

func (c *Client) Write() {
	for msg := range c.send {
		if err := c.socket.WriteJSON(msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
