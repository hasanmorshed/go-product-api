package handlers

import (
	"go-first-project/database"
	"go-first-project/global_router"
	"go-first-project/util"
	"net/http"
	"strconv"
)

func GetSingleProduct(w http.ResponseWriter, r *http.Request) {
	global_router.HandleCors(w)
	global_router.HandlePreflightReq(w, r)

	productID:=r.PathValue("productID")
	pId,err:=strconv.Atoi(productID)
	if err!=nil{
		http.Error(w,"Plz give me a valid product id",400)
		return
	}
	for _,product:=range database.ProductList{
		if product.ID ==pId{
			util.SendData(w,product,200)
			return
		}
	}
	util.SendData(w,"Data Not Found",404)
}
