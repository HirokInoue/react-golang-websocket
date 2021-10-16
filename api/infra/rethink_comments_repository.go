package infra

import (
	"fmt"

	d "github.com/HirokInoue/realtimeweb/domain"
	r "github.com/dancannon/gorethink"
)

type rethinkComment struct {
	Id      string `gorethink:"id,omitempty"`
	Content string `gorethink:"content"`
}

func NewCommentsRepository(s *r.Session) *RethinkCommentsRepository {
	return &RethinkCommentsRepository{
		session: s,
		table:   "comments",
	}
}

type RethinkCommentsRepository struct {
	session *r.Session
	table   string
}

func (cr *RethinkCommentsRepository) Save(c d.Comment) error {
	rc := rethinkComment{Id: c.Id, Content: c.Content}
	_, err := r.Table("comments").
		Insert(rc).
		RunWrite(cr.session)
	if err != nil {
		return err
	}
	return nil
}

func (cr *RethinkCommentsRepository) Retrieve() ([]d.Comment, error) {
	res, err := r.Table("comments").
		Pluck("id", "content").
		Run(cr.session)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	var comments []d.Comment
	var row map[string]interface{}
	for res.Next(&row) {
		comments = append(comments, d.Comment{
			Id:      fmt.Sprintf("%s", row["id"]),
			Content: fmt.Sprintf("%s", row["content"]),
		})
	}
	if res.Err() != nil {
		return nil, res.Err()
	}
	return comments, nil
}
