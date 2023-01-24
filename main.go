package main

import (
	"database/handler"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	tableMigration := `
    CREATE TABLE IF NOT EXISTS users (
        id serial,
        name varchar,
        primary key(id)
    );`

	db, err := sqlx.Connect("postgres", "user=postgres password=secret dbname=new sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(tableMigration)
	handler.New(db)
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Panic(err)
	}
}
