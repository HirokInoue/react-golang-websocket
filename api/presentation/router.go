package presentation

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

type FindHandler func(string) (Handler, bool)

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]Handler),
	}
}

type Router struct {
	rules map[string]Handler
}

func (ro *Router) Handle(handlerName string, handler Handler) {
	ro.rules[handlerName] = handler
}

func (ro *Router) FindHandler(handlerName string) (Handler, bool) {
	handler, ok := ro.rules[handlerName]
	return handler, ok
}

func (ro *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	defer socket.Close()

	client := NewClient(socket, ro.FindHandler)
	go client.Write()
	client.Read()
}
