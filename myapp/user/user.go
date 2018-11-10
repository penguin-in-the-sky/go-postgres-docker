package user

import (
	"html/template"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	validator "gopkg.in/go-playground/validator.v9"
)

type Data struct {
	Users []User `validate:"-"`
	Name  string `validate:"max=20,min=1"`
	Error string
}

type User struct {
	ID         int    `db:"id"`
	Name       string `db:"user_name"`
	InvalidFlg bool   `db:"invalid_flg"`
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {

	users := GetAllUsers()
	var data Data
	data.Users = users
	data.Error = ""

	renderTemplate(w, "user", &data)

}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("add user start")

	// validate
	data := Data{
		Name: r.FormValue("name"),
	}
	validate := validator.New()
	err := validate.Struct(data)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		if validationErrors != nil {
			log.Println("validation error")
			data.Error = "name"
			data.Users = GetAllUsers()
			renderTemplate(w, "user", &data)
			return
		}
	}

	db, err := sqlx.Connect("postgres", "host=postgres port=5432 user=postgres dbname=godb sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	name := r.FormValue("name")

	tx := db.MustBegin()
	tx.MustExec("INSERT INTO gogo.users (user_name, invalid_flg) VALUES ($1, $2)", name, false)
	tx.Commit()

	log.Println("add user end")

	http.Redirect(w, r, "/user/", http.StatusFound)
}

func InitializeHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("truncate start")

	db, err := sqlx.Connect("postgres", "host=postgres port=5432 user=postgres dbname=godb sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	tx := db.MustBegin()
	tx.MustExec("TRUNCATE TABLE gogo.users")
	tx.MustExec("INSERT INTO gogo.users (user_name, invalid_flg) VALUES ($1, $2)", "someone", false)
	tx.Commit()

	log.Println("truncate end")

	http.Redirect(w, r, "/user/", http.StatusFound)

}

func GetAllUsers() []User {
	db, err := sqlx.Connect("postgres", "host=postgres port=5432 user=postgres dbname=godb sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	users := []User{}
	db.Select(&users, "SELECT * FROM gogo.users")

	log.Println(users)
	return users
}

func renderTemplate(w http.ResponseWriter, tmpl string, d *Data) {
	t, _ := template.ParseFiles("template/" + tmpl + ".html")
	err := t.Execute(w, d)
	if err != nil {
		panic(err)
	}
}
