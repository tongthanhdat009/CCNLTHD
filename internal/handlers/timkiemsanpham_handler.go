package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tongthanhdat009/CCNLTHD/internal/services"
	"net/http"
	"strconv" // Thêm thư viện để chuyển đổi chuỗi
)

type TimKiemHangHoaHandler struct {
	service services.TimKiemHangHoaService
}

func NewTimKiemHangHoaHandler(service services.TimKiemHangHoaService) *TimKiemHangHoaHandler {
	return &TimKiemHangHoaHandler{service: service}
}

func (h *TimKiemHangHoaHandler) TimHangHoa(c *gin.Context) {
	// Lấy các tham số dạng chuỗi như bình thường
	tenHangHoa := c.Query("tenHangHoa")
	tenHang := c.Query("tenHang")
	tenDanhMuc := c.Query("tenDanhMuc")
	mau := c.Query("mau")
	size := c.Query("size")

	// --- PHẦN SỬA ĐỔI QUAN TRỌNG ---
	// 1. Khai báo các biến giá trị với kiểu con trỏ, giá trị mặc định là nil
	var giaToiThieu, giaToiDa *float64

	// 2. Đọc và chuyển đổi chuỗi sang float64 một cách an toàn
	giaToiThieuStr := c.Query("giaToiThieu")
	if giaToiThieuStr != "" {
		val, err := strconv.ParseFloat(giaToiThieuStr, 64)
		if err != nil {
			// Nếu người dùng nhập không phải là số, trả về lỗi 400
			c.JSON(http.StatusBadRequest, gin.H{"error": "Giá tối thiểu không hợp lệ"})
			return
		}
		giaToiThieu = &val // Gán địa chỉ của giá trị đã chuyển đổi
	}

	giaToiDaStr := c.Query("giaToiDa")
	if giaToiDaStr != "" {
		val, err := strconv.ParseFloat(giaToiDaStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Giá tối đa không hợp lệ"})
			return
		}
		giaToiDa = &val
	}
	// --- KẾT THÚC PHẦN SỬA ĐỔI ---

	// Gọi service với các tham số đã được xử lý đúng kiểu
	HangHoas, err := h.service.TimHangHoa(tenHangHoa, tenHang, tenDanhMuc, mau, size, giaToiThieu, giaToiDa)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(HangHoas) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Không có sản phẩm phù hợp"})
		return
	}

	c.JSON(http.StatusOK, HangHoas)
}