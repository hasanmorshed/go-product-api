package cmd

import (
	"fmt"
	"go-first-project/handlers"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /create-product", http.HandlerFunc(handlers.CreateProduct))
	mux.Handle("DELETE /delete-product", http.HandlerFunc(handlers.DeleteProduct))
	mux.Handle("PATCH /patch-product", http.HandlerFunc(handlers.PatchProduct))
	mux.Handle("PUT /update-product", http.HandlerFunc(handlers.UpdateProduct))

	fmt.Println("Server running on :3000")
	http.ListenAndServe(":3000", mux)
}