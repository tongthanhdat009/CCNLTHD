package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tongthanhdat009/CCNLTHD/internal/db"
	"github.com/tongthanhdat009/CCNLTHD/internal/handlers"
	"github.com/tongthanhdat009/CCNLTHD/internal/middleware"
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

	// Middleware
	authRepo := repositories.NewAuthRepository(database)
	permissionService := services.NewPermissionService(authRepo)
	permissionMiddleware := middleware.NewPermissionMiddleware(permissionService)

	//Giỏ hàng
	gioHangRepo := repositories.NewGioHangRepository(database)
	gioHangService := services.NewGioHangService(gioHangRepo)
	gioHangHandler := handlers.NewGioHangHandler(gioHangService)
	
	// Tra cứu admin
    tracuuAdminRepo := repositories.NewTraCuuAdminRepository(database)
    tracuuAdminService := services.NewTraCuuAdminRepository(tracuuAdminRepo)
    traCuuAdminHandler := handlers.NewTraCuuAdminHandler(tracuuAdminService)

	//Khuyến mãi
	khuyenMaiRepo := repositories.NewKhuyenMaiRepository(database)
	khuyenMaiService := services.NewKhuyenMaiService(khuyenMaiRepo)
	khuyenMaiHandler := handlers.NewKhuyenMaiHandler(khuyenMaiService)

	//Tìm kiếm sản phẩm
	timKiemHangHoaRepo := repositories.NewTimKiemHangHoaRepository(database)
	timKiemHangHoaService := services.NewTimKiemHangHoaService(timKiemHangHoaRepo)
	timKiemHangHoaHandler := handlers.NewTimKiemHangHoaHandler(timKiemHangHoaService)

	//Quản lý biến thể
	bienTheRepo := repositories.NewBienTheRepository(database)
	bienTheService := services.NewBienTheService(bienTheRepo)
	bienTheHandler := handlers.NewQuanLyBienTheHandler(bienTheService)

	// Quản lý phiếu nhập
	phieuNhapRepo := repositories.NewPhieuNhapRepository(database)
	phieuNhapService := services.NewQuanLyPhieuNhapService(phieuNhapRepo)
	phieuNhapHandler := handlers.NewQuanLyPhieuNhapHandler(phieuNhapService)



	// --- Thiết lập server ---
	r := gin.Default()
	r.Static("/AnhHangHoa", "./static/AnhHangHoa")
	// Gọi hàm để thiết lập tất cả các routes
	routes.SetupRoutes(r, 
        hangHoaHandler, 
        donHangHandler, 
        nguoiDungHandler, 
        hangHandler, 
        nhaCungCapHandler, 
        dangKyHandler, 
        dangNhapHandler, 
        permissionMiddleware, 
		gioHangHandler,
		khuyenMaiHandler,
        traCuuAdminHandler,
		timKiemHangHoaHandler,
		bienTheHandler,
		phieuNhapHandler)

	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Could not run server: %v", err)
	}
}
