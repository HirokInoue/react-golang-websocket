package domain

import (
	"context"
)

type CommentsRepository interface {
	Save(Comment) error
	Retrieve() ([]Comment, error)
	Feed(chan<- Comment, chan<- error, context.Context)
}
