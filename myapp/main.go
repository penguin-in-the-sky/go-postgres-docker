package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {

	log.Println("Go is running")

	router := mux.NewRouter()
	router.HandleFunc("/", HelloHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	type User struct {
		ID         int
		Name       string
		InvalidFlg bool
	}

	db, err := sqlx.Connect("postgres", "host=postgres port=5432 user=postgres dbname=godb sslmode=disable")
	defer db.Close()

	if err != nil {
		log.Fatalln(err)
	}

	// insert
	tx := db.MustBegin()
	tx.MustExec("INSERT INTO gogo.users (user_name, invalid_flg) VALUES ($1, $2)", "testUser", false)
	tx.Commit()

	// select
	users := []User{}
	db.Select(&users, "SELECT * FROM gogo.users")

	log.Println("users:" + strconv.Itoa(len(users)))

	// fmt.Fprintf(w, "Hello,"+users[0].Name)
}
