package handlers

import (
	"encoding/json"
	"go-first-project/database"
	"go-first-project/global_router"
	"go-first-project/util"
	"net/http"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	global_router.HandleCors(w)
	global_router.HandlePreflightReq(w, r)

	var updatedProduct database.Product
	err := json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}

	if updatedProduct.ID == 0 {
		http.Error(w, "Product ID is required", 400)
		return
	}

	for i, p := range database.ProductList {
		if p.ID == updatedProduct.ID {
			database.ProductList[i] = updatedProduct
			util.SendData(w, updatedProduct, 200)
			return
		}
	}

	http.Error(w, "Product not found", 404)
}
