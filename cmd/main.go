// cmd/api/main.go
package main

import (
    "log"

    "github.com/tongthanhdat009/CCNLTHD/internal/db"
    "github.com/tongthanhdat009/CCNLTHD/internal/handlers"
    "github.com/tongthanhdat009/CCNLTHD/internal/repositories"
    "github.com/tongthanhdat009/CCNLTHD/internal/services"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    // Kết nối database bằng GORM
    database, err := db.ConnectDatabase()
    if err != nil {
        log.Fatalf("Could not connect to the database: %v", err)
    }

    // Khởi tạo repository, service và handler cho hàng hóa
    hangHoaRepo := repositories.NewHangHoaRepository(database)
    hangHoaService := services.NewHangHoaService(hangHoaRepo)
    hangHoaHandler := handlers.NewHangHoaHandler(hangHoaService)

    r := gin.Default()

    // API routes cho hàng hóa
    r.GET("/api/hanghoa", hangHoaHandler.GetAllHangHoa)

    log.Println("Starting server on :8080")
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Could not run server: %v", err)
    }
}