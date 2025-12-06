
---

# 📦 Go First Project — Simple Product CRUD API

A lightweight modular **REST API** built using **Golang**, implementing full **CRUD operations** for product management.
The project demonstrates clean folder structure, handlers separation, custom routing, and JSON-based API responses.

---

## 🚀 Features

* ✔ **GET** all products
* ✔ **POST** create new product
* ✔ **PUT** update full product
* ✔ **PATCH** update partial product fields
* ✔ **DELETE** delete product
* ✔ Custom `ServeMux` router
* ✔ Modular file structure
* ✔ JSON responses with a reusable `sendData()` utility
* ✔ In-memory product database
* ✔ CORS support (if added)

---

## 📁 Project Structure

```
GO_FIRST_PROJECT/
│
├── cmd/
│   └── server.go               # Server entry point 
│
├── database/
│   └── product.go              # Product struct & product data storage
│
├── global_router/
│   └── global_router.go        # All routes registered here
│
├── handlers/
│   ├── create_products.go      # Handler for POST request
│   ├── delete_products.go      # Handler for DELETE request
│   ├── get_products.go         # Handler for fetching all products
│   ├── get_single_product.go   # Handler for fetching a single product by ID
│   ├── patch_products.go       # Handler for PATCH request
│   └── put_products.go         # Handler for PUT request
│
├── util/
│   └── send_data.go            # Utility for sending JSON responses
│
├── go.mod                      # Go module file
├── main.go                     # Starts the server
└── README.md                   # Project documentation

```

---

## ⚙️ Installation & Run

### 1. Clone the repository

```bash
git clone https://github.com/YourUserName/go-product-api.git
cd go-product-api
```

### 2. Run the project

```
go run main.go
```

Server will start at:

```
http://localhost:3000
```

---

## 📡 API Endpoints

### **Get all products**

```
GET /products
```

---


### **Get single product**

```
GET /products/{productID}
Example:
GET /products/2
```

---

### **Create a new product**

```
POST /create-product
```

**Body (JSON):**

```json
{
  "title": "Mango",
  "description": "Sweet fruit",
  "price": 300,
  "imgurl": "https://example.com/mango.jpg"
}
```

---

### **Update full product (PUT)**

```
PUT /update-product
```

**Body:**

```json
{
  "id": 1,
  "title": "Updated Mango",
  "description": "Updated description",
  "price": 350,
  "imgurl": "https://example.com/new.jpg"
}
```

---

### **Update partial product (PATCH)**

```
PATCH /patch-product
```

**Body (any fields allowed):**

```json
{
  "id": 1,
  "price": 400
}
```

---

### **Delete product**

```
DELETE /delete-product
```

**Body:**

```json
{
  "id": 1
}
```

---

## 🛠 Tech Used

* **Go 1.22+**
* **net/http**
* **encoding/json**

---

## 📘 Learning Goals

This project helps you understand:

* Routing with `ServeMux`
* Handlers in separate files
* JSON marshaling/unmarshaling
* Clean folder structure in Go projects
* CRUD API design

---

## 🤝 Contributing

Contributions and suggestions are welcome!

---

## 📜 License

This project is open-source and free to use.

---


