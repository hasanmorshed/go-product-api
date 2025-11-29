package handlers

import (
	"encoding/json"
	"go-first-project/database"
	"go-first-project/global_router"
	"go-first-project/util"
	"net/http"
)

func PatchProduct(w http.ResponseWriter, r *http.Request) {
	global_router.HandleCors(w)
	global_router.HandlePreflightReq(w, r)

	var patchData map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&patchData)
	if err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}

	idFloat := patchData["id"].(float64)
	id := int(idFloat)

	for i, p := range database.ProductList {
		if p.ID == id {
			if title, ok := patchData["title"].(string); ok {
				database.ProductList[i].Title = title
			}
			if desc, ok := patchData["description"].(string); ok {
				database.ProductList[i].Description = desc
			}
			if price, ok := patchData["price"].(float64); ok {
				database.ProductList[i].Price = price
			}
			if img, ok := patchData["imgurl"].(string); ok {
				database.ProductList[i].ImgUrl = img
			}
			util.SendData(w, database.ProductList[i], 200)
			return
		}
	}

	http.Error(w, "Product not found", 404)
}