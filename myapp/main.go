package main

import (
	"log"
	"net/http"

	"myapp/home"
	"myapp/user"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	log.Println("Go is running")

	router := mux.NewRouter()

	// home
	router.HandleFunc("/", home.ViewHandler).Methods("GET")

	// user
	router.HandleFunc("/user/", user.ViewHandler).Methods("GET")
	router.HandleFunc("/user/initialize", user.InitializeHandler).Methods("GET")
	router.HandleFunc("/user/add", user.AddHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))

}
