package addUser

import (
	"html/template"
	"net/http"
)

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "addUser")
}

func renderTemplate(w http.ResponseWriter, tmpl string) { // TODO interfaceで共通化したい
	t, _ := template.ParseFiles("template/" + tmpl + ".html")
	err := t.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
