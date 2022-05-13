package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"training_golang_ketiga/connection"
	"training_golang_ketiga/structs"

	"golang.org/x/crypto/bcrypt"
)

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
