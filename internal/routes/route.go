package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/tongthanhdat009/CCNLTHD/internal/handlers"
    "github.com/tongthanhdat009/CCNLTHD/internal/middleware"
)

// SetupRoutes định nghĩa tất cả các route cho ứng dụng.
func SetupRoutes(r *gin.Engine, hangHoaHandler *handlers.HangHoaHandler, donHangHandler *handlers.DonHangHandler, nguoiDungHandler  *handlers.NguoiDungHandler, hangHandler *handlers.HangHandler, nhaCungCap *handlers.NhaCungCapHandler, dangKyHandler *handlers.DangKyHandler, dangNhapHandler *handlers.DangNhapHandler, permissionMiddleware *middleware.PermissionMiddleware) {
    // Các route không cần xác thực

    r.POST("/api/dangky", dangKyHandler.CreateNguoiDung)
    r.POST("/api/dangnhap", dangNhapHandler.KiemTraDangNhap)
    // Nhóm các API dưới tiền tố /api
    api := r.Group("/api", middleware.AuthMiddleware())
    {
        // Routes cho Hàng Hóa
        hangHoaRoutes := api.Group("/hanghoa")
        {
            hangHoaRoutes.GET("", hangHoaHandler.GetAllHangHoa)
        }

        // Routes cho Đơn Hàng
        donHangRoutes := api.Group("/donhang")
        {
            donHangRoutes.GET("", donHangHandler.GetAllDonHang)
            // Thêm các route khác cho đơn hàng ở đây
        }

        // Routes cho Người Dùng
        nguoiDungRoutes := api.Group("/nguoidung")
        {
            nguoiDungRoutes.GET("", permissionMiddleware.Require("Quản lý người dùng", "Xem"), nguoiDungHandler.GetAllNguoiDung)
            nguoiDungRoutes.PATCH("/:id", nguoiDungHandler.UpdateNguoiDung)
        }

        // Routes cho Hãng
        hangRoutes := api.Group("/hang")
        {
            hangRoutes.GET("", hangHandler.GetAllHang)
            hangRoutes.DELETE("/:id", hangHandler.DeleteHang)
            hangRoutes.GET("/:id", hangHandler.GetHangByID)
            hangRoutes.GET("/search/:tenhang", hangHandler.GetHangByName)  
            hangRoutes.POST("", hangHandler.CreateHang)
            hangRoutes.PUT("", hangHandler.UpdateHang)
        }

        // Routes cho Nhà Cung Cấp
        nhaCungCapRoutes := api.Group("/nhacungcap")
        {
            nhaCungCapRoutes.GET("", nhaCungCap.GetAllNhaCungCap)
            nhaCungCapRoutes.POST("", nhaCungCap.CreateNhaCungCap)
            nhaCungCapRoutes.PUT("", nhaCungCap.UpdateNhaCungCap)
            nhaCungCapRoutes.DELETE("/:id", nhaCungCap.DeleteNhaCungCap)
            nhaCungCapRoutes.GET("/:id", nhaCungCap.GetNhaCungCapByID)
            nhaCungCapRoutes.GET("/search/:tenncc", nhaCungCap.GetNhaCungCapByName)
        }
    }
}