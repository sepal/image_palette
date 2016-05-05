package models

import (
	r "gopkg.in/dancannon/gorethink.v2"
)

var Session *r.Session

func Connect(host, database string) (err error) {
	Session, err = r.Connect(r.ConnectOpts{
		Address:  host,
		Database: database,
	})

	return err
}
