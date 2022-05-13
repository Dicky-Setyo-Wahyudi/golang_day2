package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"training_golang_ketiga/connection"
	"training_golang_ketiga/structs"

	"github.com/gorilla/mux"
)

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var db_training_ketiga structs.Product
	fmt.Println(db_training_ketiga)

	json.Unmarshal(payloads, &db_training_ketiga)

	connection.DB.Create(&db_training_ketiga)

	res := structs.Result{Code: 200, Data: db_training_ketiga, Message: "Produk Berhasil Ditambahkan"}

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

	var db_training_ketiga structs.Product

	connection.DB.First(&db_training_ketiga, id_produk)

	json.Unmarshal(payloads, &db_training_ketiga)

	connection.DB.Model(&db_training_ketiga).Update(db_training_ketiga)
	if !db_training_ketiga.Status {
		connection.DB.Model(&db_training_ketiga).Updates(map[string]interface{}{"status": false})
	}

	res := structs.Result{Code: 200, Data: db_training_ketiga, Message: "Produk Berhasil Diupdate"}

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

	var db_training_ketiga structs.Product

	connection.DB.First(&db_training_ketiga, id_produk)
	connection.DB.Delete(&db_training_ketiga)

	res := structs.Result{Code: 200, Data: db_training_ketiga, Message: "Produk Berhasil Dihapus"}

	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id_produk := vars["id"]

	var db_training_ketiga structs.Product

	connection.DB.First(&db_training_ketiga, id_produk)

	res := structs.Result{Code: 200, Data: db_training_ketiga, Message: "Produk Ada"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	limit := vars["limit"]
	offset := vars["offset"]

	db_training_ketigas := []structs.Product{}

	connection.DB.Limit(limit).Offset(offset).Find(&db_training_ketigas)

	res := structs.Result{Code: 200, Data: db_training_ketigas, Message: "Produk Ada"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
