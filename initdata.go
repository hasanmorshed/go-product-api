package main

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "The orange is a round, juicy citrus fruit known for its high vitamin C.",
		Price:       600,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTKsFReRVBHj3ajh2FvnGN3wEi9MWgQ3sizFg&s",
	}

	prd2 := Product{
		ID:          2,
		Title:       "Mango",
		Description: "he mango is a sweet, juicy, tropical stone fruit.",
		Price:       500,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRDh6RmUj-ZuZw_77mK_iQzGxg1R46_hVjSxg&s",
	}
	prd3 := Product{
		ID:          3,
		Title:       "Watermelon",
		Description: "The watermelon is a sweet, juicy, and refreshing summer fruit.",
		Price:       400,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQaZb1b0-iP2EeBtWfklHOdlVBh58g5mSgQ9A&s",
	}
	prd4 := Product{
		ID:          4,
		Title:       "Blueberry",
		Description: "Blueberries are small, round berries with a sweet and slightly.",
		Price:       300,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR-SxGRpJEWC_xvIsWKio837O1M1BBvilJfqQ&s",
	}
	prd5 := Product{
		ID:          5,
		Title:       "Cherry",
		Description: "The cherry is a small, round stone fruit that offers a burst of sweet.",
		Price:       200,
		ImgUrl:      "https://www.shutterstock.com/image-vector/two-cherries-isolated-on-white-600nw-2425952009.jpg",
	}

	productList = append(productList, prd1, prd2, prd3, prd4, prd5)
}
