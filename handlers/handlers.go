package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleReq() {
	log.Println("Start development server localhost:8000")

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/user", Insert).Methods("OPTIONS", "POST")
	myRouter.HandleFunc("/user/{id}", Update).Methods("OPTIONS", "PUT")
	myRouter.HandleFunc("/user/{id}", Delete).Methods("OPTIONS", "DELETE")
	myRouter.HandleFunc("/user/{id}", GetUser).Methods("OPTIONS", "GET")
	myRouter.HandleFunc("/user/{limit}/{offset}", GetUsers).Methods("OPTIONS", "GET")

	myRouter.HandleFunc("/product", InsertProduct).Methods("OPTIONS", "POST")
	myRouter.HandleFunc("/product/{id}", UpdateProduct).Methods("OPTIONS", "PUT")
	myRouter.HandleFunc("/product/{id}", DeleteProduk).Methods("OPTIONS", "DELETE")
	myRouter.HandleFunc("/product/{id}", GetProduk).Methods("OPTIONS", "GET")
	myRouter.HandleFunc("/product/{limit}/{offset}", GetProduks).Methods("OPTIONS", "GET")

	log.Fatal(http.ListenAndServe(":8000", myRouter))
}
