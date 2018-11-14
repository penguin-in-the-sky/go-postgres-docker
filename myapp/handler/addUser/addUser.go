package addUser

import (
	"html/template"
	"log"
	"net/http"

	"myapp/database/userData"

	validator "gopkg.in/go-playground/validator.v9"
)

type Data struct {
	Authorities []userData.Authority `validate:"-"`
	Name        string               `validate:"max=20,min=1"`
	Password    string               `validate:"max=20,min=1"`
	Authority   string               `validate:"min=1"`
	Error       string
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	data := Data{
		Authorities: userData.GetAllAuthorities(),
	}
	log.Println(data.Authorities)
	renderTemplate(w, "addUser", &data)
}

func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("add user start")

	// validate TODO 権限のバリデーション未実装
	errorType := ""
	data := Data{
		Name:      r.FormValue("name"),
		Password:  r.FormValue("password"),
		Authority: r.FormValue("authority"),
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
		data.Authorities = userData.GetAllAuthorities()
		renderTemplate(w, "addUser", &data)
		return
	}

	userData.AddUser(data.Name, data.Password, data.Authority)
	log.Println("add user end")
	http.Redirect(w, r, "/add/user", http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, d *Data) {
	t, _ := template.ParseFiles("template/" + tmpl + ".html")
	err := t.Execute(w, d)
	if err != nil {
		panic(err)
	}
}
