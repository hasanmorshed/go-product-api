package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(getProducts))
	mux.Handle("POST /create-product", http.HandlerFunc(createProduct))
	mux.Handle("DELETE /delete-product", http.HandlerFunc(deleteProduct))
	mux.Handle("PATCH /patch-product", http.HandlerFunc(patchProduct))
	mux.Handle("PUT /update-product", http.HandlerFunc(updateProduct))

	fmt.Println("Server running on :3000")
	http.ListenAndServe(":3000", mux)
}
