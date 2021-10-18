package infra

import (
	r "github.com/dancannon/gorethink"
)

func NewSession(dbName string) (*r.Session, error) {
	session, err := r.Connect(r.ConnectOpts{
		Address:  "db:28015",
		Database: dbName,
	})
	if err != nil {
		return nil, err
	}
	return session, nil
}
