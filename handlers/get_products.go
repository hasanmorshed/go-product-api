package handlers

import (
	"go-first-project/global_router"
	"go-first-project/product"
	"go-first-project/util"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	global_router.HandleCors(w)
	global_router.HandlePreflightReq(w, r)
	util.SendData(w, product.ProductList, 200)
}