package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"training_golang_kedua/connection"
	"training_golang_kedua/structs"

	"github.com/gorilla/mux"
)

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var db_training_kedua structs.Product

	json.Unmarshal(payloads, &db_training_kedua)

	connection.DB.Create(&db_training_kedua)

	res := structs.Result{Code: 200, Data: db_training_kedua, Message: "Produk Berhasil Ditambahkan"}

	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id_produk := vars["id"]
	payloads, _ := ioutil.ReadAll(r.Body)

	var db_training_kedua structs.Product

	connection.DB.First(&db_training_kedua, id_produk)

	json.Unmarshal(payloads, &db_training_kedua)

	connection.DB.Model(&db_training_kedua).Update(db_training_kedua)

	res := structs.Result{Code: 200, Data: db_training_kedua, Message: "Produk Berhasil Diupdate"}

	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func DeleteProduk(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id_produk := vars["id"]

	var db_training_kedua structs.Product

	connection.DB.First(&db_training_kedua, id_produk)
	connection.DB.Delete(&db_training_kedua)

	res := structs.Result{Code: 200, Data: db_training_kedua, Message: "Produk Berhasil Dihapus"}

	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetProduk(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id_produk := vars["id"]

	var db_training_kedua structs.Product

	connection.DB.First(&db_training_kedua, id_produk)

	res := structs.Result{Code: 200, Data: db_training_kedua, Message: "Produk Ada"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetProduks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	limit := vars["limit"]
	offset := vars["offset"]

	db_training_keduas := []structs.Product{}

	connection.DB.Limit(limit).Offset(offset).Find(&db_training_keduas)

	res := structs.Result{Code: 200, Data: db_training_keduas, Message: "Produk Ada"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
