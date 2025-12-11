package cmd

import (
	"fmt"
	"go-first-project/config"
	"go-first-project/handlers"
	"net/http"
	"strconv"
)

func Serve() {
	cnf:=config.GetConfig()
	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("GET /products/{productID}", http.HandlerFunc(handlers.GetSingleProduct))
	mux.Handle("POST /create-product", http.HandlerFunc(handlers.CreateProduct))
	mux.Handle("DELETE /delete-product", http.HandlerFunc(handlers.DeleteProduct))
	mux.Handle("PATCH /patch-product", http.HandlerFunc(handlers.PatchProduct))
	mux.Handle("PUT /update-product", http.HandlerFunc(handlers.UpdateProduct))
	addr:=":"+ strconv.Itoa(cnf.HttpPort)
	fmt.Println("Server running on :",addr)
	http.ListenAndServe(addr, mux)
}