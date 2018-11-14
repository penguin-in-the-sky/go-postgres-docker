package user

import (
	"html/template"
	"net/http"

	"myapp/database/userData"
)

type Data struct {
	Users []userData.User `validate:"-"`
	Name  string          `validate:"max=20,min=1"`
	Error string
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {

	users := userData.GetAllUsers()
	var data Data
	data.Users = users
	data.Error = ""

	renderTemplate(w, "user", &data)

}

func renderTemplate(w http.ResponseWriter, tmpl string, d *Data) {
	t, _ := template.ParseFiles("template/" + tmpl + ".html")
	err := t.Execute(w, d)
	if err != nil {
		panic(err)
	}
}
