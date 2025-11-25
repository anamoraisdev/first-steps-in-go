package main

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
