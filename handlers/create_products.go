package handlers

import (
	"encoding/json"
	"go-first-project/database"
	"go-first-project/global_router"
	"go-first-project/util"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	global_router.HandleCors(w)
	global_router.HandlePreflightReq(w, r)

	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}

	newProduct.ID = len(database.ProductList) + 1
	database.ProductList = append(database.ProductList, newProduct)
	util.SendData(w, newProduct, 201)
}
