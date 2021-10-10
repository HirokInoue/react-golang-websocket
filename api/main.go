package main

import (
	"net/http"

	p "github.com/HirokInoue/realtimeweb/presentation"
)

func main() {
	router := p.NewRouter()
	router.Handle("add comment", p.AddComment)
	router.Handle("listen comments", p.ListenComments)
	http.Handle("/", router)
	http.ListenAndServe(":8765", nil)
}
