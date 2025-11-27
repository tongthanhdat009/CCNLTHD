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

// UpdateQuyenRequest - Request body cho API cập nhật quyền
type UpdateQuyenRequest struct {
	TenQuyen string `json:"ten_quyen" binding:"required"`
}

// UpdateQuyenWithPermissionsRequest - Request body cho API cập nhật quyền kèm phân quyền
type UpdateQuyenWithPermissionsRequest struct {
	TenQuyen           string `json:"ten_quyen" binding:"required"`
	MaChiTietChucNangs []int  `json:"ma_chi_tiet_chuc_nangs"`
}

// PhanQuyenRequest - Request body cho API phân quyền
type PhanQuyenRequest struct {
	MaChiTietChucNangs []int `json:"ma_chi_tiet_chuc_nangs" binding:"required"`
}

// ==================== READ ====================

// GetAll - GET /api/quyen
func (h *QuyenHandler) GetAll(c *gin.Context) {
	quyens, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Lấy danh sách quyền thành công",
		"data":    quyens,
	})
}

// GetByID - GET /api/quyen/:id
func (h *QuyenHandler) GetByID(c *gin.Context) {
	// Lấy ID từ tham số URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Mã quyền không hợp lệ"})
		return
	}

	quyen, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy quyền"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Lấy thông tin quyền thành công",
		"data":    quyen,
	})
}

// GetAllChiTietChucNang - GET /api/quyen/chi-tiet-chuc-nang
func (h *QuyenHandler) GetAllChiTietChucNang(c *gin.Context) {
	maChiTiets, err := h.service.GetAllChiTietChucNang()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Lấy danh sách chi tiết chức năng thành công",
		"data":    maChiTiets,
	})
}

// ==================== CREATE ====================

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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Trả về kết quả
	c.JSON(http.StatusCreated, gin.H{
		"message": "Tạo quyền thành công",
		"data":    quyen,
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Phân quyền thành công",
	})
}

// ==================== UPDATE ====================

// UpdateQuyen - PUT /api/quyen/:id
// Cập nhật chỉ tên quyền, không sửa phân quyền
func (h *QuyenHandler) UpdateQuyen(c *gin.Context) {
	// Lấy mã quyền từ URL
	idParam := c.Param("id")
	maQuyen, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Mã quyền không hợp lệ"})
		return
	}

	// Bind JSON request
	var req UpdateQuyenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dữ liệu không hợp lệ",
			"details": err.Error(),
		})
		return
	}

	// Cập nhật quyền
	quyen, err := h.service.UpdateQuyen(maQuyen, req.TenQuyen)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Cập nhật quyền thành công",
		"data":    quyen,
	})
}

// UpdateQuyenWithPermissions - PUT /api/quyen/:id/full
// Cập nhật tên quyền + phân quyền
func (h *QuyenHandler) UpdateQuyenWithPermissions(c *gin.Context) {
	// Lấy mã quyền từ URL
	idParam := c.Param("id")
	maQuyen, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Mã quyền không hợp lệ"})
		return
	}

	// Bind JSON request
	var req UpdateQuyenWithPermissionsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dữ liệu không hợp lệ",
			"details": err.Error(),
		})
		return
	}

	// Cập nhật quyền kèm phân quyền
	quyen, err := h.service.UpdateQuyenWithPermissions(maQuyen, req.TenQuyen, req.MaChiTietChucNangs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Cập nhật quyền và phân quyền thành công",
		"data":    quyen,
	})
}

// ==================== DELETE ====================

// DeleteQuyen - DELETE /api/quyen/:id
func (h *QuyenHandler) DeleteQuyen(c *gin.Context) {
	// Lấy mã quyền từ URL
	idParam := c.Param("id")
	maQuyen, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Mã quyền không hợp lệ"})
		return
	}

	// Xóa quyền
	if err := h.service.DeleteQuyen(maQuyen); err != nil {
		// Phân biệt lỗi nghiệp vụ và lỗi hệ thống
		if err.Error() == "không thể xóa quyền vì đang có người dùng sử dụng" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Xóa quyền thành công",
	})
}
