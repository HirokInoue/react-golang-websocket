package infra

import (
	"testing"

	d "github.com/HirokInoue/realtimeweb/domain"
	r "github.com/dancannon/gorethink"
)

var session *r.Session

func init() {
	session, _ = NewSession("realtimeweb_test")
}

func TestRethinkCommentsRepository_Save(t *testing.T) {
	tests := []struct {
		name  string
		c     d.Comment
		setUp func()
	}{
		{
			name: "1件保存する",
			c:    d.Comment{Content: "RethinkDB is the open-source, scalable database that makes building realtime apps dramatically easier."},
			setUp: func() {
				r.Table("comments").Delete().Run(session)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setUp()

			cr := NewCommentsRepository(session)
			cr.Save(d.Comment{Content: tt.c.Content})

			s := assertForSave{session, t}
			s.assertCount(1)
			s.assert(tt.c.Content)
		})
	}
}

type assertForSave struct {
	session *r.Session
	t       *testing.T
}

func (a assertForSave) assertCount(want int) {
	cur, err := r.Table("comments").
		Count().
		Run(a.session)
	if err != nil {
		a.t.Fatal(err)
	}
	var cnt int
	_ = cur.One(&cnt)
	if cnt != want {
		a.t.Errorf("\ngot: %v \nwant: %v \n", cnt, 1)
	}
}

func (a assertForSave) assert(want string) {
	res, err := r.Table("comments").
		Pluck("id", "content", "created_at").
		Run(a.session)
	if err != nil {
		a.t.Fatal(err)
	}
	var row map[string]interface{}
	for res.Next(&row) {
		if row["id"] == "" ||
			row["content"] != want ||
			row["created_at"] == "" {
			a.t.Errorf("\ngot: %v\nwant: %v \n", row["content"], want)
		}
	}
}

func TestRethinkCommentsRepository_Retrieve(t *testing.T) {
	comment := d.Comment{Content: "RethinkDB is the open-source, scalable database that makes building realtime apps dramatically easier."}
	tests := []struct {
		name  string
		want  []d.Comment
		setUp func()
	}{
		{
			name: "DBに0件のデータがある",
			want: []d.Comment{},
			setUp: func() {
				r.Table("comments").Delete().Run(session)
			},
		},
		{
			name: "DBにn件のデータがある",
			want: []d.Comment{
				comment,
			},
			setUp: func() {
				r.Table("comments").Delete().Run(session)
				cr := NewCommentsRepository(session)
				_ = cr.Save(comment)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setUp()

			cr := NewCommentsRepository(session)
			got, err := cr.Retrieve()
			if err != nil {
				t.Fatal(err)
			}

			r := assertForRetrieve{t}
			r.assert(got, tt.want)
		})
	}
}

type assertForRetrieve struct {
	t *testing.T
}

func (a *assertForRetrieve) assert(got, want []d.Comment) {
	if len(got) != len(want) {
		a.t.Errorf("wrong length.\ngot: %v\nwant: %v\n", len(got), len(want))
	}

	for _, g := range got {
		if !a.inSlice(g.Content, want) {
			a.t.Errorf("wrong content.\ngot: %v\nwant: %v\n", g.Content, want)
		}
	}
}

func (a *assertForRetrieve) inSlice(n string, h []d.Comment) bool {
	for _, v := range h {
		if v.Content == n {
			return true
		}
	}
	return false
}
