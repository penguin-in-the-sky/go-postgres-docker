package main

import (
	"log"
	"net/http"

	"myapp/handler/addPraise"
	"myapp/handler/addUser"
	"myapp/handler/approvePraise"
	"myapp/handler/home"
	"myapp/handler/login"
	"myapp/handler/user"

	"github.com/gorilla/mux"
)

func main() {

	log.Println("Go is running")

	router := mux.NewRouter()

	// login
	router.HandleFunc("/login", login.ViewHandler).Methods("GET")
	router.HandleFunc("/login", login.LoginHandler).Methods("POST")

	// addUser
	router.HandleFunc("/add/user", addUser.ViewHandler).Methods("GET")
	router.HandleFunc("/add/user", addUser.AddUserHandler).Methods("POST")

	// home
	router.HandleFunc("/home", home.ViewHandler).Methods("GET")
	router.HandleFunc("/display", home.DisplayPraiseHandler).Methods("GET")

	// addPraise
	router.HandleFunc("/praise/add", addPraise.ViewHandler).Methods("GET")
	router.HandleFunc("/praise/add", addPraise.AddHomeHandler).Methods("POST")

	// approvePraise
	router.HandleFunc("/praise/approve", approvePraise.ViewHandler).Methods("GET")
	router.HandleFunc("/praise/approve", approvePraise.ApproveHandler).Methods("POST")

	// user
	router.HandleFunc("/user/", user.ViewHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))

}
