package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/tongthanhdat009/CCNLTHD/internal/handlers"
)

// SetupRoutes định nghĩa tất cả các route cho ứng dụng.
func SetupRoutes(r *gin.Engine, hangHoaHandler *handlers.HangHoaHandler, donHangHandler *handlers.DonHangHandler, nguoiDungHandler  *handlers.NguoiDungHandler) {
    // Nhóm các API dưới tiền tố /api
    api := r.Group("/api")
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
            nguoiDungRoutes.GET("", nguoiDungHandler.GetAllNguoiDung)
        }
    }
}