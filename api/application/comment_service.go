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

func (cs *CommentService) Listen() ([]string, error) {
	commentModelList, err := cs.repos.Retrieve()
	if err != nil {
		return nil, err
	}
	var comments []string
	for _, v := range commentModelList {
		comments = append(comments, v.Content)
	}
	return comments, nil
}
