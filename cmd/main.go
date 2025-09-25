package main

import (
    "log"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "github.com/tongthanhdat009/CCNLTHD/internal/db"
    "github.com/tongthanhdat009/CCNLTHD/internal/handlers"
    "github.com/tongthanhdat009/CCNLTHD/internal/repositories"
    "github.com/tongthanhdat009/CCNLTHD/internal/routes" // <-- Import package routes
    "github.com/tongthanhdat009/CCNLTHD/internal/services"
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

    // --- Khởi tạo các dependencies ---
    // Hàng hóa
    hangHoaRepo := repositories.NewHangHoaRepository(database)
    hangHoaService := services.NewHangHoaService(hangHoaRepo)
    hangHoaHandler := handlers.NewHangHoaHandler(hangHoaService)

    // Đơn hàng
    donHangRepo := repositories.NewDonHangRepository(database)
    donHangService := services.NewDonHangService(donHangRepo)
    donHangHandler := handlers.NewDonHangHandler(donHangService)

    // Người dùng
    nguoiDungRepo := repositories.NewNguoiDungRepository(database)
    nguoiDungService := services.NewNguoiDungService(nguoiDungRepo)
    nguoiDungHandler := handlers.NewNguoiDungHandler(nguoiDungService)

    //Hãng
    hangRepo := repositories.NewHangRepository(database)
    hangService := services.NewHangService(hangRepo, hangHoaRepo)
    hangHandler := handlers.NewHangHandler(hangService)

    //Nhà cung cấp
    nhaCungCapRepo := repositories.NewNhaCungCapRepository(database)
    nhaCungCapService := services.NewNhaCungCapService(nhaCungCapRepo, database)
    nhaCungCapHandler := handlers.NewNhaCungCapHandler(nhaCungCapService)

    // Đăng ký
    dangKyService := services.NewDangKyService(nguoiDungRepo)
    dangKyHandler := handlers.NewDangKyHandler(dangKyService)

    // Đăng nhập
    dangNhapService := services.NewDangNhapService(nguoiDungRepo)
    dangNhapHandler := handlers.NewDangNhapHandler(dangNhapService)

    // --- Thiết lập server ---
    r := gin.Default()

    // Gọi hàm để thiết lập tất cả các routes
    routes.SetupRoutes(r, hangHoaHandler, donHangHandler, nguoiDungHandler, hangHandler, nhaCungCapHandler, dangKyHandler, dangNhapHandler)

    log.Println("Starting server on :8080")
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Could not run server: %v", err)
    }
}