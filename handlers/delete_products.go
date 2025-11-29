package handlers

import (
	"encoding/json"
	"go-first-project/database"
	"go-first-project/global_router"
	"go-first-project/util"
	"net/http"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	global_router.HandleCors(w)
	global_router.HandlePreflightReq(w, r)

	var data map[string]int
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}

	id := data["id"]

	for i, p := range database.ProductList {
		if p.ID == id {
			database.ProductList = append(database.ProductList[:i], database.ProductList[i+1:]...)
			util.SendData(w, map[string]string{"message": "Product deleted"}, 200)
			return
		}
	}

	http.Error(w, "Product not found", 404)
}
