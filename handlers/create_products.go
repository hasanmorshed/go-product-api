package handlers

import (
	"encoding/json"
	"go-first-project/global_router"
	"go-first-project/product"
	"go-first-project/util"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	global_router.HandleCors(w)
	global_router.HandlePreflightReq(w, r)

	var newProduct product.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}

	newProduct.ID = len(product.ProductList) + 1
	product.ProductList = append(product.ProductList, newProduct)
	util.SendData(w, newProduct, 201)
}
