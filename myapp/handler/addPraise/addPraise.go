package addPraise

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"myapp/database/praiseData"
	"myapp/database/userData"

	validator "gopkg.in/go-playground/validator.v9"
)

type Data struct {
	Users  []userData.User `validate:"-"`
	User   string          `validate:"min=1"`
	Praise string          `validate:"max=600,min=1"`
	Error  string
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	data := Data{
		Users: userData.GetAllUsers(),
		Error: "",
	}

	renderTemplate(w, "add", &data)
}

func AddHomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("add praise start")

	// validate
	data := Data{
		User:   r.FormValue("user"),
		Praise: r.FormValue("praise"),
	}

	log.Println(data.User)

	validate := validator.New()
	err := validate.Struct(data)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		if validationErrors != nil {
			log.Println(validationErrors)
			data.Error = "name"
			data.Users = userData.GetAllUsers()
			renderTemplate(w, "add", &data)
			return
		}
	}

	var id int
	id, _ = strconv.Atoi(data.User)
	user := userData.GetUserById(id)
	// TODO registered_user
	praiseData.AddPraise(data.Praise, user.ID)
	http.Redirect(w, r, "/praise/add", http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, d *Data) {
	t, _ := template.ParseFiles("template/" + tmpl + ".html")
	err := t.Execute(w, d)
	if err != nil {
		panic(err)
	}
}
