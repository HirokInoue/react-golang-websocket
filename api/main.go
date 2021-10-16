package main

import (
	"log"
	"net/http"

	p "github.com/HirokInoue/realtimeweb/presentation"
)

func main() {
	addCommentHandler, err := p.NewAddCommentHandler()
	if err != nil {
		log.Panic(err.Error())
	}
	listenCommentsHandler, err := p.NewListenCommentsHandler()
	if err != nil {
		log.Panic(err.Error())
	}

	router := p.NewRouter()
	router.Handle("add comment", addCommentHandler)
	router.Handle("listen comments", listenCommentsHandler)
	http.Handle("/", router)
	http.ListenAndServe(":8765", nil)
}
