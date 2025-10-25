package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"github.com/tongthanhdat009/CCNLTHD/internal/services"

	"github.com/gin-gonic/gin"
)

type DonHangHandler struct {
	service services.DonHangService
}

func NewDonHangHandler(service services.DonHangService) *DonHangHandler {
	return &DonHangHandler{service: service}
}

// GetAllDonHang - Lấy tất cả đơn hàng
// GET /api/donhang
func (h *DonHangHandler) GetAllDonHang(c *gin.Context) {
	donHangs, err := h.service.GetAllDonHang()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Không thể lấy danh sách đơn hàng",
			"details": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":  donHangs,
		"total": len(donHangs),
	})
}

// GetDonHangByID - Lấy đơn hàng theo mã
// GET /api/donhang/:id
func (h *DonHangHandler) GetDonHangByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Mã đơn hàng không hợp lệ",
		})
		return
	}

	donHang, err := h.service.GetDonHangByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": donHang,
	})
}

// GetDetailByID - Xem chi tiết đầy đủ đơn hàng
// GET /api/donhang/:id/detail
func (h *DonHangHandler) GetDetailByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Mã đơn hàng không hợp lệ",
		})
		return
	}

	donHang, err := h.service.GetDetailByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": donHang,
	})
}

// CreateDonHang - Tạo đơn hàng mới
// POST /api/donhang
func (h *DonHangHandler) CreateDonHang(c *gin.Context) {
	var donHang models.DonHang
	if err := c.ShouldBindJSON(&donHang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dữ liệu không hợp lệ",
			"details": err.Error(),
		})
		return
	}

	if err := h.service.CreateDonHang(&donHang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Tạo đơn hàng thành công",
		"data":    donHang,
	})
}

// UpdateDonHang - Cập nhật thông tin đơn hàng
// PUT /api/donhang/:id
func (h *DonHangHandler) UpdateDonHang(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Mã đơn hàng không hợp lệ",
		})
		return
	}

	var donHang models.DonHang
	if err := c.ShouldBindJSON(&donHang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dữ liệu không hợp lệ",
			"details": err.Error(),
		})
		return
	}

	donHang.MaDonHang = id

	if err := h.service.UpdateDonHang(&donHang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Cập nhật đơn hàng thành công",
		"data":    donHang,
	})
}

// DeleteDonHang - Xóa đơn hàng
// DELETE /api/donhang/:id
func (h *DonHangHandler) DeleteDonHang(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Mã đơn hàng không hợp lệ",
		})
		return
	}

	if err := h.service.DeleteDonHang(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Xóa đơn hàng thành công",
	})
}

// ApproveOrder - Duyệt đơn hàng
// POST /api/donhang/:id/approve
func (h *DonHangHandler) ApproveOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Mã đơn hàng không hợp lệ",
		})
		return
	}

	if err := h.service.ApproveOrder(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Duyệt đơn hàng thành công",
	})
}

// UpdateOrderStatus - Cập nhật trạng thái đơn hàng
// PATCH /api/donhang/:id/status
func (h *DonHangHandler) UpdateOrderStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Mã đơn hàng không hợp lệ",
		})
		return
	}

	var request struct {
		TrangThai string `json:"trang_thai" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Trạng thái không được để trống",
		})
		return
	}

	if err := h.service.UpdateOrderStatus(id, request.TrangThai); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Cập nhật trạng thái thành công",
		"trang_thai": request.TrangThai,
	})
}

// CancelOrder - Hủy đơn hàng
// POST /api/donhang/:id/cancel
func (h *DonHangHandler) CancelOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Mã đơn hàng không hợp lệ",
		})
		return
	}

	var request struct {
		Reason string `json:"reason"`
	}

	// Lý do hủy không bắt buộc
	c.ShouldBindJSON(&request)

	if err := h.service.CancelOrder(id, request.Reason); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Hủy đơn hàng thành công",
	})
}

// SearchDonHang - Tìm kiếm đơn hàng
// GET /api/donhang/search
func (h *DonHangHandler) SearchDonHang(c *gin.Context) {
	// Lấy các tham số từ query string
	keyword := c.Query("keyword")
	trangThai := c.Query("trangThai")
	fromDateStr := c.Query("fromDate")
	toDateStr := c.Query("toDate")

	// Parse thời gian
	var fromDate, toDate time.Time
	var err error

	if fromDateStr != "" {
		fromDate, err = time.Parse(time.RFC3339, fromDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Định dạng fromDate không hợp lệ. Sử dụng RFC3339: 2025-10-01T00:00:00Z"})
			return
		}
	}

	if toDateStr != "" {
		toDate, err = time.Parse(time.RFC3339, toDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Định dạng toDate không hợp lệ. Sử dụng RFC3339: 2025-10-31T23:59:59Z"})
			return
		}
	}

	// Gọi service để tìm kiếm
	donHangs, err := h.service.SearchDonHang(keyword, trangThai, fromDate, toDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(donHangs) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "Không tìm thấy đơn hàng nào",
			"data":    []models.DonHang{},
			"total":   0,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Tìm kiếm đơn hàng thành công",
		"data":    donHangs,
		"total":   len(donHangs),
	})
}

// GetDonHangByNguoiDung - Lấy đơn hàng theo người dùng
// GET /api/donhang/nguoidung/:id
func (h *DonHangHandler) GetDonHangByNguoiDung(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Mã người dùng không hợp lệ",
		})
		return
	}

	donHangs, err := h.service.GetDonHangByNguoiDung(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  donHangs,
		"total": len(donHangs),
	})
}

// GetDonHangByStatus - Lấy đơn hàng theo trạng thái
// GET /api/donhang/status/:status
func (h *DonHangHandler) GetDonHangByStatus(c *gin.Context) {
	status := c.Param("status")

	donHangs, err := h.service.GetDonHangByStatus(status)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       donHangs,
		"total":      len(donHangs),
		"trang_thai": status,
	})
}
