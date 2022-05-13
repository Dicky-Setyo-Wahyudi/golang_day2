package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"training_golang_ketiga/connection"
	"training_golang_ketiga/structs"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func InsertUser(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var db_training_ketiga structs.Users

	json.Unmarshal(payloads, &db_training_ketiga)

	hashedPassword, err := HashPassword(db_training_ketiga.Password)
	db_training_ketiga.Password = hashedPassword
	connection.DB.Create(&db_training_ketiga)

	res := structs.Result{Code: 200, Data: db_training_ketiga, Message: "User Berhasil Ditambahkan"}

	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id_user := vars["id"]
	payloads, _ := ioutil.ReadAll(r.Body)

	var db_training_ketiga structs.Users
	connection.DB.First(&db_training_ketiga, id_user)

	json.Unmarshal(payloads, &db_training_ketiga)
	connection.DB.Model(&db_training_ketiga).Update(db_training_ketiga)

	if !db_training_ketiga.Status {
		connection.DB.Model(&db_training_ketiga).Updates(map[string]interface{}{"status": false})
	}

	res := structs.Result{Code: 200, Data: db_training_ketiga, Message: "User Berhasil Diupdate"}

	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id_user := vars["id"]

	var db_training_ketiga structs.Users

	connection.DB.First(&db_training_ketiga, id_user)
	connection.DB.Delete(&db_training_ketiga)

	res := structs.Result{Code: 200, Data: db_training_ketiga, Message: "User Berhasil Dihapus"}

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

	var db_training_ketiga structs.Users

	connection.DB.First(&db_training_ketiga, id_user)

	res := structs.Result{Code: 200, Data: db_training_ketiga, Message: "User Ada"}
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

	db_training_ketigas := []structs.Users{}

	connection.DB.Limit(limit).Offset(offset).Find(&db_training_ketigas)

	res := structs.Result{Code: 200, Data: db_training_ketigas, Message: "User Ada"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
