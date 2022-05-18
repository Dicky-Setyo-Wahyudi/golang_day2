package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"training_golang_keempat/connection"
	"training_golang_keempat/structs"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func InsertUser(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var db_user structs.Users

	json.Unmarshal(payloads, &db_user)

	hashedPassword, err := HashPassword(db_user.Password)
	db_user.Password = hashedPassword
	connection.DB.Create(&db_user)

	res := structs.Result{Code: 200, Data: db_user, Message: "User Berhasil Ditambahkan"}

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

	var db_user structs.Users
	connection.DB.First(&db_user, id_user)

	json.Unmarshal(payloads, &db_user)
	connection.DB.Model(&db_user).Update(db_user)

	if !db_user.Status {
		connection.DB.Model(&db_user).Updates(map[string]interface{}{"status": false})
	}

	res := structs.Result{Code: 200, Data: db_user, Message: "User Berhasil Diupdate"}

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

	var db_user structs.Users

	connection.DB.First(&db_user, id_user)
	connection.DB.Delete(&db_user)

	res := structs.Result{Code: 200, Data: db_user, Message: "User Berhasil Dihapus"}

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

	var db_user structs.Users

	connection.DB.First(&db_user, id_user)

	res := structs.Result{Code: 200, Data: db_user, Message: "User Ada"}
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

	db_users := []structs.Users{}

	connection.DB.Limit(limit).Offset(offset).Find(&db_users)

	res := structs.Result{Code: 200, Data: db_users, Message: "User Ada"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var db_login structs.Users
	var cekUser structs.UsersLogin

	res := structs.Result{Code: 200, Data: cekUser, Message: "Gagal Login"}

	json.Unmarshal(payloads, &cekUser)
	connection.DB.Where("username=?", &cekUser.Username).Find(&db_login)
	if CheckPasswordHash(cekUser.Password, db_login.Password) {
		res = structs.Result{Code: 200, Data: cekUser, Message: "Berhasil Login"}
	}

	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
