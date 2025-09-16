package models

type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Password string `json:"password"`
    Email    string `json:"email"`
}

type Product struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}

type Order struct {
    ID       int     `json:"id"`
    UserID   int     `json:"user_id"`
    ProductID int    `json:"product_id"`
    Quantity int     `json:"quantity"`
    Total    float64 `json:"total"`
}