package infra

import (
	"os"

	r "github.com/dancannon/gorethink"
)

func NewSession(dbName string) (*r.Session, error) {
	session, err := r.Connect(r.ConnectOpts{
		Address:  os.Getenv("DB_HOST"),
		Database: dbName,
	})
	if err != nil {
		return nil, err
	}
	return session, nil
}
