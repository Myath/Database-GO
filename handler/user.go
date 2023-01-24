package handler

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type users struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}

func (c *connection) createUser(w http.ResponseWriter, r *http.Request) {
	userForm(w, r, nil)
}

func (c *connection) storeUser(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Panic(err)
	}

	name := r.FormValue("name")

	fmt.Println(name)
	if name == "" {
		type errMsg struct {
			Err string
		}

		userForm(w, r, errMsg{
			Err: "Name is Required",
		})
		return
	}

	createUserQuery := `INSERT INTO users(name) VALUES($1)`
	res := c.db.MustExec(createUserQuery, name)

	if ok, err := res.RowsAffected(); err != nil || ok == 0 {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/user/create", http.StatusPermanentRedirect)
}

func userForm(w http.ResponseWriter, r *http.Request, data interface{}) {
	temp, err := template.ParseFiles("./template/create.html")
	if err != nil {
		log.Fatal(err)
	}

	temp.Execute(w, data)
}

func (c *connection) listUser(w http.ResponseWriter, r *http.Request) {
	var user users
	c.db.Select(&user, "SELECT * FROM users")
	fmt.Println(user)

	temp, err := template.ParseFiles("./template/list.html")
	if err != nil {
		log.Fatal(err)
	}

	temp.Execute(w, user)
}
