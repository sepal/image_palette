package models

import (
	r "github.com/dancannon/gorethink"
)

var session *r.Session

func Connect(host, database string) (err error) {
	session, err = r.Connect(r.ConnectOpts{
		Address:  host,
		Database: database,
	})

	return err
}
