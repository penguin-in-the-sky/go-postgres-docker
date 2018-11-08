package home

import (
	"html/template"
	"net/http"
)

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/home.html")
	err := t.Execute(w, nil)
	if err != nil {
		panic(err)
	}

}
