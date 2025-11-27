package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Order struct {
	ID       int     `json:"id"`
	UserID   int     `json:"user_id"`
	Products []int   `json:"products"`
	Total    float64 `json:"total"`
}

var users = []User{}
var products = []Product{}
var orders = []Order{}
var nextUserID = 1
var nextProductID = 1
var nextOrderID = 1

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var u User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Invalid user JSON", http.StatusBadRequest)
		return
	}

	u.ID = nextUserID
	nextUserID++

	users = append(users, u)

	json.NewEncoder(w).Encode(u)
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	idUserString := r.URL.Path[len("/users/"):]
	idUser, err := strconv.Atoi(idUserString)

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	for i, u := range users {
		if u.ID == idUser {
			deletedUser := users[i]
			users = append(users[:i], users[i+1:]...)
			json.NewEncoder(w).Encode(deletedUser)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func createProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid user JSON", http.StatusBadRequest)
		return
	}

	product.ID = nextProductID
	nextProductID++

	products = append(products, product)

	json.NewEncoder(w).Encode(product)
}

func deleteProductHandler(w http.ResponseWriter, r *http.Request) {
	idProductString := r.URL.Path[len("/products/"):]
	idProduct, err := strconv.Atoi(idProductString)

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	for i, product := range products {
		if product.ID == idProduct {
			deletedProduct := products[i]
			products = append(products[:i], products[i+1:]...)
			json.NewEncoder(w).Encode(deletedProduct)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

func getOrdersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func createOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order Order
	var user *User
	var total float64

	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid user JSON", http.StatusBadRequest)
		return
	}

	for i := range users {
		if users[i].ID == order.UserID {
			user = &users[i]
			break
		}
	}
	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	productMap := make(map[int]Product)
	for _, product := range products {
		productMap[product.ID] = product
	}

	for _, productID := range order.Products {
		product, exists := productMap[productID]
		if !exists {
			http.Error(w, fmt.Sprintf("Product ID %d not found", productID), http.StatusBadRequest)
			return
		}
		total += product.Price
	}

	if user.Balance < total {
		http.Error(w, "Insufficient balance", http.StatusPaymentRequired)
		return
	} else {
		user.Balance -= total
	}

	order.ID = nextOrderID
	order.Total = total
	nextOrderID++

	orders = append(orders, order)

	json.NewEncoder(w).Encode(order)

}
func main() {
	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getUsersHandler(w, r)
		case http.MethodPost:
			createUserHandler(w, r)
		case http.MethodDelete:
			deleteUserHandler(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/products/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getProductsHandler(w, r)
		case http.MethodPost:
			createProductHandler(w, r)
		case http.MethodDelete:
			deleteProductHandler(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/orders/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getOrdersHandler(w, r)
		case http.MethodPost:
			createOrderHandler(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server running on http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
