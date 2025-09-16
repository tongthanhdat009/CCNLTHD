// cmd/api/main.go
package main

import (
	"log"
	"github.com/tongthanhdat009/CCNLTHD/internal/db"
	"github.com/tongthanhdat009/CCNLTHD/internal/handlers"
	"github.com/tongthanhdat009/CCNLTHD/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// ... (code tải .env và kết nối DB giữ nguyên)
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := db.ConnectDatabase()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer database.Close()

	// Khởi tạo service và handler cho Brand (Hãng)
	brandService := &services.HangService{DB: database}
	brandHandler := &handlers.HangHandler{Service: brandService}
	r := gin.Default()

	// Định nghĩa route mới cho Brand (Hãng)
	r.GET("/brands", brandHandler.GetAllBrands)
	// --- KẾT THÚC PHẦN CẬP NHẬT ---
	
	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Could not run server: %v", err)
	}
}