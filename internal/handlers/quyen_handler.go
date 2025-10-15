package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tongthanhdat009/CCNLTHD/internal/services"
)

type QuyenHandler struct {
	service services.QuyenService
}

func NewQuyenHandler(service services.QuyenService) *QuyenHandler {
	return &QuyenHandler{service: service}
}

// CreateQuyenRequest - Request body cho API tạo quyền
type CreateQuyenRequest struct {
	TenQuyen           string `json:"ten_quyen" binding:"required"`
	MaChiTietChucNangs []int  `json:"ma_chi_tiet_chuc_nangs"`
}

// PhanQuyenRequest - Request body cho API phân quyền
type PhanQuyenRequest struct {
	MaChiTietChucNangs []int `json:"ma_chi_tiet_chuc_nangs" binding:"required"`
}

// Thêm các phương thức xử lý khác tại đây, ví dụ:
func (h *QuyenHandler) GetAll(c *gin.Context) {
	Quyens, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Quyens)
}
func (h *QuyenHandler) GetByID(c *gin.Context) {
	// Lấy ID từ tham số URL
	idParam := c.Param("id")
	// Chuyển đổi ID từ chuỗi sang số nguyên
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	Quyen, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Quyen)
}

// CreateQuyen - POST /api/quyen
func (h *QuyenHandler) CreateQuyen(c *gin.Context) {
	var req CreateQuyenRequest

	// Bind JSON request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dữ liệu không hợp lệ",
			"details": err.Error(),
		})
		return
	}

	// Tạo quyền với phân quyền chi tiết
	quyen, err := h.service.CreateQuyenWithPermissions(req.TenQuyen, req.MaChiTietChucNangs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Trả về kết quả
	c.JSON(http.StatusCreated, gin.H{
		"message": "Tạo quyền thành công",
		"data":    quyen,
	})
}

// GetAllChiTietChucNang - GET /api/quyen/chi-tiet-chuc-nang
func (h *QuyenHandler) GetAllChiTietChucNang(c *gin.Context) {
	maChiTiets, err := h.service.GetAllChiTietChucNang()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Lấy danh sách chi tiết chức năng thành công",
		"data":    maChiTiets,
	})
}

// PhanQuyen - POST /api/quyen/:id/phan-quyen
func (h *QuyenHandler) PhanQuyen(c *gin.Context) {
	// Lấy mã quyền từ URL
	idParam := c.Param("id")
	maQuyen, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Mã quyền không hợp lệ"})
		return
	}

	// Bind JSON request
	var req PhanQuyenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dữ liệu không hợp lệ",
			"details": err.Error(),
		})
		return
	}

	// Phân quyền
	if err := h.service.PhanQuyen(maQuyen, req.MaChiTietChucNangs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Phân quyền thành công",
	})
}
