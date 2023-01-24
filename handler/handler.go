package handler

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

type connection struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) connection {
	c := connection{
		db: db,
	}

	http.HandleFunc("/user/create", c.createUser)
	http.HandleFunc("/user/store", c.storeUser)
	http.HandleFunc("/users", c.listUser)

	return c
}
