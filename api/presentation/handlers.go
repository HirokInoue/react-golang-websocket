package presentation

import (
	"time"
)

type Handler func(*Client, interface{})

func AddComment(c *Client, data interface{}) {
	go func() {
		c.send <- Body{Name: "add comment", Data: data}
		// TODO: insert to DB
	}()
}

func ListenComments(c *Client, data interface{}) {
	go func() {
		for {
			time.Sleep(time.Second * 5)
			// TODO: read database
			c.send <- Body{Name: "listen comment", Data: ""}
		}
	}()
}
