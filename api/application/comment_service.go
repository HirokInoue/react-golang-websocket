package application

import (
	d "github.com/HirokInoue/realtimeweb/domain"
)

func NewCommentService(r d.CommentsRepository) *CommentService {
	return &CommentService{
		repos: r,
	}
}

type CommentService struct {
	repos d.CommentsRepository
}

func (cs *CommentService) Add(data string) error {
	err := cs.repos.Save(d.Comment{Content: data})
	if err != nil {
		return err
	}
	return nil
}

func (cs *CommentService) Listen(s chan<- string, e chan error) {
	c := make(chan d.Comment)
	cs.repos.Feed(c, e)
	for {
		select {
		case comment := <-c:
			s <- comment.Content
		case <-e:
		}
	}
}
