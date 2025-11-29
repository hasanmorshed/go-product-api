package main

import (
	"fmt"
	"go-first-project/handlers"
	"go-first-project/product"
	"net/http"
)


func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /create-product", http.HandlerFunc(handlers.CreateProduct))
	mux.Handle("DELETE /delete-product", http.HandlerFunc(handlers.DeleteProduct))
	mux.Handle("PATCH /patch-product", http.HandlerFunc(handlers.PatchProduct))
	mux.Handle("PUT /update-product", http.HandlerFunc(handlers.UpdateProduct))

	fmt.Println("Server running on :3000")
	http.ListenAndServe(":3000", mux)
}

func init() {
	prd1 := product.Product{
		ID:          1,
		Title:       "Orange",
		Description: "The orange is a round, juicy citrus fruit known for its high vitamin C.",
		Price:       600,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTKsFReRVBHj3ajh2FvnGN3wEi9MWgQ3sizFg&s",
	}

	prd2 := product.Product{
		ID:          2,
		Title:       "Mango",
		Description: "he mango is a sweet, juicy, tropical stone fruit.",
		Price:       500,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRDh6RmUj-ZuZw_77mK_iQzGxg1R46_hVjSxg&s",
	}
	prd3 := product.Product{
		ID:          3,
		Title:       "Watermelon",
		Description: "The watermelon is a sweet, juicy, and refreshing summer fruit.",
		Price:       400,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQaZb1b0-iP2EeBtWfklHOdlVBh58g5mSgQ9A&s",
	}
	prd4 := product.Product{
		ID:          4,
		Title:       "Blueberry",
		Description: "Blueberries are small, round berries with a sweet and slightly.",
		Price:       300,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR-SxGRpJEWC_xvIsWKio837O1M1BBvilJfqQ&s",
	}
	prd5 := product.Product{
		ID:          5,
		Title:       "Cherry",
		Description: "The cherry is a small, round stone fruit that offers a burst of sweet.",
		Price:       200,
		ImgUrl:      "https://www.shutterstock.com/image-vector/two-cherries-isolated-on-white-600nw-2425952009.jpg",
	}

	product.ProductList = append(product.ProductList, prd1, prd2, prd3, prd4, prd5)
}

