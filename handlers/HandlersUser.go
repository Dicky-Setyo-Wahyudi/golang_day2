package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"training_golang_kedua/connection"
	"training_golang_kedua/structs"

	"github.com/gorilla/mux"
)

func Insert(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var db_training_kedua structs.Users

	json.Unmarshal(payloads, &db_training_kedua)

	connection.DB.Create(&db_training_kedua)

	res := structs.Result{Code: 200, Data: db_training_kedua, Message: "User Berhasil Ditambahkan"}

	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id_user := vars["id"]
	payloads, _ := ioutil.ReadAll(r.Body)

	var db_training_kedua structs.Users
	connection.DB.First(&db_training_kedua, id_user)

	json.Unmarshal(payloads, &db_training_kedua)
	connection.DB.Model(&db_training_kedua).Update(db_training_kedua)

	res := structs.Result{Code: 200, Data: db_training_kedua, Message: "User Berhasil Diupdate"}

	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id_user := vars["id"]

	var db_training_kedua structs.Users

	connection.DB.First(&db_training_kedua, id_user)
	connection.DB.Delete(&db_training_kedua)

	res := structs.Result{Code: 200, Data: db_training_kedua, Message: "User Berhasil Dihapus"}

	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id_user := vars["id"]

	var db_training_kedua structs.Users

	connection.DB.First(&db_training_kedua, id_user)

	res := structs.Result{Code: 200, Data: db_training_kedua, Message: "User Ada"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	limit := vars["limit"]
	offset := vars["offset"]

	db_training_keduas := []structs.Users{}

	connection.DB.Limit(limit).Offset(offset).Find(&db_training_keduas)

	res := structs.Result{Code: 200, Data: db_training_keduas, Message: "User Ada"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
