package presentation

import (
	"context"

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
		send:         make(chan Body),
		socket:       s,
		findHandler:  f,
		stopContexts: make(map[int]context.CancelFunc),
	}
}

type Client struct {
	send         chan Body
	socket       *websocket.Conn
	findHandler  FindHandler
	stopContexts map[int]context.CancelFunc
}

func (c *Client) NewStopContext(key int) context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	c.stopContexts[key] = cancel
	return ctx
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

func (c *Client) Close() {
	for _, cancel := range c.stopContexts {
		cancel()
	}
	close(c.send)
}
