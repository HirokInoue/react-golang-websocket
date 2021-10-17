package presentation

import (
	"fmt"
	"log"
	"time"

	a "github.com/HirokInoue/realtimeweb/application"
	i "github.com/HirokInoue/realtimeweb/infra"
)

type Handler interface {
	exec(*Client, interface{})
}

func NewAddCommentHandler() (*AddCommentHandler, error) {
	session, err := i.NewSession("realtimeweb")
	if err != nil {
		return nil, err
	}
	repos := i.NewCommentsRepository(session)
	s := a.NewCommentService(repos)
	return &AddCommentHandler{
		service: s,
	}, nil
}

type AddCommentHandler struct {
	service *a.CommentService
}

func (ah *AddCommentHandler) exec(c *Client, data interface{}) {
	go func() {
		isOk := true
		err := ah.service.Add(fmt.Sprintf("%s", data))
		if err != nil {
			log.Print(err)
			isOk = false
		}
		c.send <- Body{Name: "add comment", Ok: isOk}
	}()
}

func NewListenCommentsHandler() (*ListenCommentsHandler, error) {
	// FIXME: Don't Repeat Yourself!
	session, err := i.NewSession("realtimeweb")
	if err != nil {
		return nil, err
	}
	repos := i.NewCommentsRepository(session)
	s := a.NewCommentService(repos)
	return &ListenCommentsHandler{
		service: s,
	}, nil
}

type ListenCommentsHandler struct {
	service *a.CommentService
}

func (lh *ListenCommentsHandler) exec(c *Client, data interface{}) {
	go func() {
		for {
			isOk := true
			comments, err := lh.service.Listen()
			if err != nil {
				log.Print(err)
				isOk = false
			}
			c.send <- Body{Name: "listen comments", Ok: isOk, Data: comments}
			time.Sleep(time.Second * 3)
		}
	}()
}
