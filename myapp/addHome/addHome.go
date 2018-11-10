package addHome

import (
	"html/template"
	"net/http"

	"myapp/user"
)

type Data struct {
	Users   []user.User `validate:"-"`
	Name    string      `validate:"max=20,min=1"`
	Content string      `validate:"max=600,min=1"`
	Error   string
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	data := Data{
		Users: user.GetAllUsers(),
		Error: "",
	}

	t, _ := template.ParseFiles("template/add.html")
	err := t.Execute(w, data)
	if err != nil {
		panic(err)
	}
}
