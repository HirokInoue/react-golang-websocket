package infra

import (
	"fmt"
	"time"

	d "github.com/HirokInoue/realtimeweb/domain"
	r "github.com/dancannon/gorethink"
)

const TIME_FORMAT = "2006-01-02T15:04:05+09:00"

type rethinkComment struct {
	Id        string `gorethink:"id,omitempty"`
	Content   string `gorethink:"content"`
	CreatedAt string `gorethink:"created_at"`
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
	rc := rethinkComment{Id: c.Id, Content: c.Content, CreatedAt: cr.now()}
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
		Pluck("id", "content", "created_at").
		OrderBy("created_at").
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

func (cr *RethinkCommentsRepository) now() string {
	japan, _ := time.LoadLocation("Asia/Tokyo")
	return time.Now().In(japan).Format(TIME_FORMAT)
}

func (cr *RethinkCommentsRepository) Feed(c chan<- d.Comment, e chan<- error) {
	res, err := r.Table("comments").
		OrderBy(r.OrderByOpts{Index: r.Asc("created_at")}).
		Changes(r.ChangesOpts{IncludeInitial: true}).
		Run(cr.session)
	if err != nil {
		e <- err
	}

	go func() {
		defer res.Close()
		var change r.ChangeResponse
		for res.Next(&change) {
			n := change.NewValue.(map[string]interface{})
			c <- d.Comment{
				Id:      fmt.Sprintf("%s", n["id"]),
				Content: fmt.Sprintf("%s", n["content"]),
			}
		}
		if res.Err() != nil {
			e <- res.Err()
		}
	}()
}
