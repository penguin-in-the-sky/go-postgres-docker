package home

import (
	"html/template"
	"net/http"

	"myapp/user"
)

type Data struct {
	Users  []user.User
	Praise Praise
	Error  string
}

type Praise struct {
	Id          int       `db:"id"`
	User        user.User `db:"user"`
	Content     string    `db:"content"`
	HasApproved bool      `db:"has_approved"`
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	data := Data{
		Users: user.GetAllUsers(),
	}
	renderTemplate(w, "home", &data)
}

func DisplayPraiseHandler(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func renderTemplate(w http.ResponseWriter, tmpl string, d *Data) { // TODO interfaceで共通化したい
	t, _ := template.ParseFiles("template/" + tmpl + ".html")
	err := t.Execute(w, d)
	if err != nil {
		panic(err)
	}
}
