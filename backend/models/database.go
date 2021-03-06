package models

import (
	r "gopkg.in/dancannon/gorethink.v2"
)

var session *r.Session

// Connects the application to the RethinkDB host and database.
func Connect(host, database string) (err error) {
	session, err = r.Connect(r.ConnectOpts{
		Address:  host,
		Database: database,
	})

	return err
}
