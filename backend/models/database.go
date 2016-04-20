package models

import (
	r "github.com/dancannon/gorethink"
)

var session *r.Session

func Connect(host, database string) (err error) {
	session, err = r.Connect(r.ConnectOpts{
		Address:  "192.168.99.100:28015",
		Database: "color_space",
	})

	return err
}
