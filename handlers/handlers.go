package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleReq() {
	log.Println("Start development server localhost:8080")

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/user", InsertUser).Methods("OPTIONS", "POST")
	myRouter.HandleFunc("/user/{id}", UpdateUser).Methods("OPTIONS", "PUT")
	myRouter.HandleFunc("/user/{id}", DeleteUser).Methods("OPTIONS", "DELETE")
	myRouter.HandleFunc("/user/{id}", GetUser).Methods("OPTIONS", "GET")
	myRouter.HandleFunc("/user/{limit}/{offset}", GetUsers).Methods("OPTIONS", "GET")

	myRouter.HandleFunc("/product", InsertProduct).Methods("OPTIONS", "POST")
	myRouter.HandleFunc("/product/{id}", UpdateProduct).Methods("OPTIONS", "PUT")
	myRouter.HandleFunc("/product/{id}", DeleteProduk).Methods("OPTIONS", "DELETE")
	myRouter.HandleFunc("/product/{id}", GetProduct).Methods("OPTIONS", "GET")
	myRouter.HandleFunc("/product/{limit}/{offset}", GetProducts).Methods("OPTIONS", "GET")

	myRouter.HandleFunc("/login", loginUser).Methods("OPTIONS", "POST")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
