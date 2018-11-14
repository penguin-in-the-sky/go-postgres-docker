package login

import (
	"html/template"
	"net/http"
)

type Data struct {
	Name     string `validate:"max=20,min=1"`
	Password string `validate:"max=20,min=1"`
	Error    string
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "login")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/home", http.StatusFound)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string) { // TODO interfaceで共通化したい
	t, _ := template.ParseFiles("template/" + tmpl + ".html")
	err := t.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
