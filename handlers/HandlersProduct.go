package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"training_golang_keempat/connection"
	"training_golang_keempat/structs"

	"github.com/gorilla/mux"
)

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var db_product structs.Product

	json.Unmarshal(payloads, &db_product)

	connection.DB.Create(&db_product)

	res := structs.Result{Code: 200, Data: db_product, Message: "Produk Berhasil Ditambahkan"}

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

	var db_product structs.Product

	connection.DB.First(&db_product, id_produk)

	json.Unmarshal(payloads, &db_product)

	connection.DB.Model(&db_product).Update(db_product)
	if !db_product.Status {
		connection.DB.Model(&db_product).Updates(map[string]interface{}{"status": false})
	}

	res := structs.Result{Code: 200, Data: db_product, Message: "Produk Berhasil Diupdate"}

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

	var db_product structs.Product

	connection.DB.First(&db_product, id_produk)
	connection.DB.Delete(&db_product)

	res := structs.Result{Code: 200, Data: db_product, Message: "Produk Berhasil Dihapus"}

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

	var db_product structs.Product

	connection.DB.First(&db_product, id_produk)

	res := structs.Result{Code: 200, Data: db_product, Message: "Produk Ada"}
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

	db_products := []structs.Product{}

	connection.DB.Limit(limit).Offset(offset).Find(&db_products)

	res := structs.Result{Code: 200, Data: db_products, Message: "Produk Ada"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
