package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Router struct{}

func main() {
	http.Handle("/", &Router{})
	http.ListenAndServe(":8765", nil)
}

func (e *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer socket.Close()

	for {
		msgType, msg, err := socket.ReadMessage()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if err = socket.WriteMessage(msgType, msg); err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
