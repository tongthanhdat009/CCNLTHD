package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tongthanhdat009/CCNLTHD/internal/handlers"
	"github.com/tongthanhdat009/CCNLTHD/internal/middleware"
)

// SetupRoutes định nghĩa tất cả các route cho ứng dụng.
func SetupRoutes(r *gin.Engine,
	hangHoaHandler *handlers.HangHoaHandler,
	donHangHandler *handlers.DonHangHandler,
	nguoiDungHandler *handlers.NguoiDungHandler,
	hangHandler *handlers.HangHandler,
	nhaCungCap *handlers.NhaCungCapHandler,
	dangKyHandler *handlers.DangKyHandler,
	dangNhapHandler *handlers.DangNhapHandler,
	permissionMiddleware *middleware.PermissionMiddleware,
	gioHangHandler *handlers.GioHangHandler,
	khuyenMaiHandler *handlers.KhuyenMaiHandler,
	traCuuAdminHandler *handlers.TraCuuAdminHandler,
	timKiemSanPhamHandler *handlers.TimKiemSanPhamHandler,
	reviewHandler *handlers.ReviewHandler,
	adminReviewHandler *handlers.AdminReviewHandler,
) {
	// Các route không cần xác thực

	r.POST("/api/dangky", dangKyHandler.CreateNguoiDung)
	r.POST("/api/dangnhap", dangNhapHandler.KiemTraDangNhap)
	// Nhóm các API dưới tiền tố /api
	api := r.Group("/api", middleware.AuthMiddleware())
	{
		// Routes cho Hàng Hóa
		hangHoaRoutes := api.Group("/hanghoa")
		{
			hangHoaRoutes.GET("", permissionMiddleware.Require("Quản lý hàng hóa", "Xem"), hangHoaHandler.GetAllHangHoa)
			hangHoaRoutes.POST("", permissionMiddleware.Require("Quản lý hàng hóa", "Xử lý"), hangHoaHandler.CreateHangHoa)
			hangHoaRoutes.POST("/update/:id", permissionMiddleware.Require("Quản lý hàng hóa", "Xử lý"), hangHoaHandler.UpdateHangHoa)
			hangHoaRoutes.GET("/search", permissionMiddleware.Require("Quản lý hàng hóa", "Xem"), hangHoaHandler.SearchHangHoa)

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
			nguoiDungRoutes.PUT("/:id", permissionMiddleware.RequireUserIDMatch(), nguoiDungHandler.UpdateNguoiDung)
			nguoiDungRoutes.GET("/admin/:id", permissionMiddleware.Require("Quản lý người dùng", "Xem"), nguoiDungHandler.GetNguoiDungByID)
			nguoiDungRoutes.GET("/:id", permissionMiddleware.RequireUserIDMatch(), nguoiDungHandler.GetNguoiDungByID)
			nguoiDungRoutes.PUT("/admin/:id", permissionMiddleware.Require("Quản lý người dùng", "Sửa"), nguoiDungHandler.UpdateNguoiDungAdmin)
		}

		// Routes cho Giỏ Hàng
		gioHangRoutes := api.Group("/giohang")
		{
			gioHangRoutes.POST("", gioHangHandler.TaoGioHang)
			gioHangRoutes.PUT("", gioHangHandler.SuaGioHang)
			gioHangRoutes.DELETE("", gioHangHandler.XoaGioHang)
			gioHangRoutes.GET("/:id", gioHangHandler.GetAll)
		}

		// Routes cho Khuyến Mãi
		khuyenMaiRoutes := api.Group("/khuyenmai")
		{
			// khuyenMaiRoutes.GET("", khuyenMaiHandler.GetAllKhuyenMai)
			khuyenMaiRoutes.POST("", khuyenMaiHandler.TaoKhuyenMai)
			// khuyenMaiRoutes.PUT("", khuyenMaiHandler.SuaKhuyenMai)
			// khuyenMaiRoutes.DELETE("/:id", khuyenMaiHandler.XoaKhuyenMai)
			// khuyenMaiRoutes.GET("/:id", khuyenMaiHandler.GetKhuyenMaiByID)
			// khuyenMaiRoutes.GET("/search/:tenkm", khuyenMaiHandler.GetKhuyenMaiByName)
		}

		// Routes cho Hãng
		hangRoutes := api.Group("/hang")
		{
			hangRoutes.GET("", permissionMiddleware.Require("Quản lý hãng", "Xem"), hangHandler.GetAllHang)
			hangRoutes.DELETE("/:id", permissionMiddleware.Require("Quản lý hãng", "Xóa"), hangHandler.DeleteHang)
			hangRoutes.GET("/:id", permissionMiddleware.Require("Quản lý hãng", "Xem"), hangHandler.GetHangByID)
			hangRoutes.GET("/search/:tenhang", permissionMiddleware.Require("Quản lý hãng", "Xem"), hangHandler.GetHangByName)
			hangRoutes.POST("", permissionMiddleware.Require("Quản lý hãng", "Thêm"), hangHandler.CreateHang)
			hangRoutes.PUT("", permissionMiddleware.Require("Quản lý hãng", "Sửa"), hangHandler.UpdateHang)
		}

		// Routes cho Nhà Cung Cấp
		nhaCungCapRoutes := api.Group("/nhacungcap")
		{
			nhaCungCapRoutes.GET("", permissionMiddleware.Require("Quản lý nhà cung cấp", "Xem"), nhaCungCap.GetAllNhaCungCap)
			nhaCungCapRoutes.POST("", permissionMiddleware.Require("Quản lý nhà cung cấp", "Thêm"), nhaCungCap.CreateNhaCungCap)
			nhaCungCapRoutes.PUT("", permissionMiddleware.Require("Quản lý nhà cung cấp", "Sửa"), nhaCungCap.UpdateNhaCungCap)
			nhaCungCapRoutes.DELETE("/:id", permissionMiddleware.Require("Quản lý nhà cung cấp", "Xóa"), nhaCungCap.DeleteNhaCungCap)
			nhaCungCapRoutes.GET("/:id", permissionMiddleware.Require("Quản lý nhà cung cấp", "Xem"), nhaCungCap.GetNhaCungCapByID)
			nhaCungCapRoutes.GET("/search/:tenncc", permissionMiddleware.Require("Quản lý nhà cung cấp", "Xem"), nhaCungCap.GetNhaCungCapByName)
		}

		// Routes cho Tra cứu admin
		traCuuAdminRoutes := api.Group("/tracuuadmin")
		{
			traCuuAdminRoutes.GET("/:seri", permissionMiddleware.Require("Tra cứu sản phẩm", "Xem"), traCuuAdminHandler.GetSanPhamBySeries)
			traCuuAdminRoutes.GET("/trangthai", permissionMiddleware.Require("Tra cứu sản phẩm", "Xem"), traCuuAdminHandler.GetSanPhamByTrangThai)
		}

		// Routes cho Tìm kiếm sản phẩm
		timKiemSanPhamRoutes := api.Group("/timkiemsanpham")
		{
			timKiemSanPhamRoutes.GET("", timKiemSanPhamHandler.TimSanPham)
		}
		// KH đánh giá
		reviews := api.Group("/reviews")
		{
			// Khách hàng gửi đánh giá
			reviews.POST("", reviewHandler.Create)

			// Xem các đánh giá của chính tôi
			reviews.GET("/me", reviewHandler.GetMine)

			// Xem đánh giá theo 1 sản phẩm nhập kho (MaSanPham)
			reviews.GET("/product/:id", reviewHandler.GetByProduct)
		}
		// Routes ADMIN cho Đánh giá (lọc/duyệt/từ chối/xóa)
		admin := api.Group("/admin/reviews")
		{
			admin.GET("",
				permissionMiddleware.Require("Quản lý đánh giá", "Xử lý"),
				adminReviewHandler.List,
			)
			admin.PUT("/:id/approve",
				permissionMiddleware.Require("Quản lý đánh giá", "Xử lý"),
				adminReviewHandler.Approve,
			)
			admin.PUT("/:id/reject",
				permissionMiddleware.Require("Quản lý đánh giá", "Xử lý"),
				adminReviewHandler.Reject,
			)
			admin.DELETE("/:id",
				permissionMiddleware.Require("Quản lý đánh giá", "Xử lý"),
				adminReviewHandler.Delete,
			)
		}

	}
}
