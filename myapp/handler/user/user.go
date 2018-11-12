package user

import (
	"html/template"
	"log"
	"net/http"

	"myapp/database/userData"

	validator "gopkg.in/go-playground/validator.v9"
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

func AddHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("add user start")

	// validate
	errorType := ""
	data := Data{
		Name: r.FormValue("name"),
	}
	validate := validator.New()
	err := validate.Struct(data)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		if validationErrors != nil {
			errorType = "length"
		}
	} else {
		registered := userData.GetUserByName(data.Name)
		if registered.Name != "" {
			errorType = "unique"
		}
	}
	if errorType != "" {
		log.Println("validation error")
		data.Error = errorType
		data.Users = userData.GetAllUsers()
		renderTemplate(w, "user", &data)
		return
	}

	userData.AddUser(r.FormValue("name"))

	http.Redirect(w, r, "/user/", http.StatusFound)
}

func InitializeHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("truncate start")

	userData.Truncate()

	log.Println("truncate end")

	http.Redirect(w, r, "/user/", http.StatusFound)

}

func renderTemplate(w http.ResponseWriter, tmpl string, d *Data) {
	t, _ := template.ParseFiles("template/" + tmpl + ".html")
	err := t.Execute(w, d)
	if err != nil {
		panic(err)
	}
}
