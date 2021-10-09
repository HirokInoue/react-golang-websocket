package main

import (
	"fmt"
	"time"
)

type Body struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

// Client is responsible for read request and write response
type Client struct {
	send chan Body
}

func (c *Client) Read() {
	// TODO: read request
}

func (c *Client) Write() {
	for msg := range c.send {
		// TODO: write response
		fmt.Printf("%#v\n", msg)
	}
}

func (c *Client) Polling() {
	// TODO: read database
	for {
		time.Sleep(time.Second * 5)
		c.send <- Body{Name: "polling", Data: ""}
	}
}

func NewClient() *Client {
	return &Client{
		send: make(chan Body),
	}
}

func main() {
	client := NewClient()
	go client.Polling()
	client.Write()
}
