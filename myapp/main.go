package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Data struct {
	Users []User
	Name  string
}

type User struct {
	ID         int    `db:"id"`
	Name       string `db:"user_name"`
	InvalidFlg bool   `db:"invalid_flg"`
}

func main() {

	log.Println("Go is running")

	router := mux.NewRouter()
	router.HandleFunc("/", viewHandler).Methods("GET")
	router.HandleFunc("/initialize", initializeHandler).Methods("GET")
	router.HandleFunc("/add", addHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))

}

func viewHandler(w http.ResponseWriter, r *http.Request) {

	users := getAllUsers()
	var data Data
	data.Users = users

	renderTemplate(w, "view", &data)

}

func addHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("add user start")

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

	http.Redirect(w, r, "/", http.StatusFound)
}

func initializeHandler(w http.ResponseWriter, r *http.Request) {

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

	http.Redirect(w, r, "/", http.StatusFound)

}

func getAllUsers() []User {
	db, err := sqlx.Connect("postgres", "host=postgres port=5432 user=postgres dbname=godb sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	users := []User{}
	db.Select(&users, "SELECT * FROM gogo.users")
	return users
}

func getDB() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "host=postgres port=5432 user=postgres dbname=godb sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return db
}

func renderTemplate(w http.ResponseWriter, tmpl string, d *Data) {
	t, _ := template.ParseFiles("template/" + tmpl + ".html")

	err := t.Execute(w, d)
	if err != nil {
		panic(err)
	}
}
