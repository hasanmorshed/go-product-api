package main

import (
	"encoding/json"
	"net/http"
)

func getProducts(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	handlePreflightReq(w, r)
	sendData(w, productList, 200)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	handlePreflightReq(w, r)

	var newProduct Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}

	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)
	sendData(w, newProduct, 201)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	handlePreflightReq(w, r)

	var data map[string]int
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}

	id := data["id"]

	for i, p := range productList {
		if p.ID == id {
			productList = append(productList[:i], productList[i+1:]...)
			sendData(w, map[string]string{"message": "Product deleted"}, 200)
			return
		}
	}

	http.Error(w, "Product not found", 404)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	handlePreflightReq(w, r)

	var updatedProduct Product
	err := json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}

	if updatedProduct.ID == 0 {
		http.Error(w, "Product ID is required", 400)
		return
	}

	for i, p := range productList {
		if p.ID == updatedProduct.ID {
			productList[i] = updatedProduct
			sendData(w, updatedProduct, 200)
			return
		}
	}

	http.Error(w, "Product not found", 404)
}

func patchProduct(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	handlePreflightReq(w, r)

	var patchData map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&patchData)
	if err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}

	idFloat := patchData["id"].(float64)
	id := int(idFloat)

	for i, p := range productList {
		if p.ID == id {
			if title, ok := patchData["title"].(string); ok {
				productList[i].Title = title
			}
			if desc, ok := patchData["description"].(string); ok {
				productList[i].Description = desc
			}
			if price, ok := patchData["price"].(float64); ok {
				productList[i].Price = price
			}
			if img, ok := patchData["imgurl"].(string); ok {
				productList[i].ImgUrl = img
			}
			sendData(w, productList[i], 200)
			return
		}
	}

	http.Error(w, "Product not found", 404)
}
