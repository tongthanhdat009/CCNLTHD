package handlers

import (
	"net/http"
	"strconv"

	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"github.com/tongthanhdat009/CCNLTHD/internal/services"

	"github.com/gin-gonic/gin"
)

type KhuyenMaiHandler struct {
	service services.KhuyenMaiService
}

func NewKhuyenMaiHandler(service services.KhuyenMaiService) *KhuyenMaiHandler {
	return &KhuyenMaiHandler{service: service}
}

func (h *KhuyenMaiHandler) CreateKhuyenMai(c *gin.Context) {
	khuyenMai := models.KhuyenMai{}
	if err := c.ShouldBindJSON(&khuyenMai); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateKhuyenMai(&khuyenMai); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, khuyenMai)
}

func (h *KhuyenMaiHandler) UpdateKhuyenMai(c *gin.Context) {
	var req models.KhuyenMai
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dữ liệu không hợp lệ"})
		return
	}

	// Validate MaKhuyenMai phải > 0
	if req.MaKhuyenMai <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id không hợp lệ"})
		return
	}

	// Gọi service để update
	if err := h.service.UpdateKhuyenMai(req.MaKhuyenMai, req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ✅ Lấy lại thông tin khuyến mãi sau khi cập nhật
	updatedKhuyenMai, err := h.service.GetByID(req.MaKhuyenMai)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cập nhật thành công nhưng không lấy được thông tin"})
		return
	}

	// ✅ Trả về thông tin đầy đủ
	c.JSON(http.StatusOK, gin.H{
		"message": "Cập nhật khuyến mãi thành công",
		"data":    updatedKhuyenMai,
	})
}

func (h *KhuyenMaiHandler) DeleteKhuyenMai(c *gin.Context) {
	idParam := c.Param("id")
	makhuyenmai, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id không hợp lệ"})
		return
	}

	if err := h.service.DeleteKhuyenMai(makhuyenmai); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Xóa khuyến mãi thành công",
	})
}

func (h *KhuyenMaiHandler) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	makhuyenmai, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id không hợp lệ"})
		return
	}

	khuyenMai, err := h.service.GetByID(makhuyenmai)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy khuyến mãi"})
		return
	}

	c.JSON(http.StatusOK, khuyenMai)
}

func (h *KhuyenMaiHandler) GetAll(c *gin.Context) {
	khuyenMais, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, khuyenMais)
}

func (h *KhuyenMaiHandler) SearchKhuyenMai(c *gin.Context) {
	// ✅ 1. Lấy keyword (tìm kiếm tổng quát)
	keyword := c.Query("keyword")

	// ✅ 2. Lấy ma_khuyen_mai (lọc chính xác theo mã)
	var maKhuyenMai *int
	if maStr := c.Query("ma_khuyen_mai"); maStr != "" {
		if val, err := strconv.Atoi(maStr); err == nil && val > 0 {
			maKhuyenMai = &val
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ma_khuyen_mai không hợp lệ"})
			return
		}
	}

	// ✅ 3. Lấy ten_khuyen_mai (lọc theo tên - LIKE)
	tenKhuyenMai := c.Query("ten_khuyen_mai")

	// ✅ 4. Lấy min_gia_tri, max_gia_tri (optional)
	var minGiaTri, maxGiaTri *float64

	if minStr := c.Query("min_gia_tri"); minStr != "" {
		if val, err := strconv.ParseFloat(minStr, 64); err == nil {
			minGiaTri = &val
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "min_gia_tri không hợp lệ"})
			return
		}
	}

	if maxStr := c.Query("max_gia_tri"); maxStr != "" {
		if val, err := strconv.ParseFloat(maxStr, 64); err == nil {
			maxGiaTri = &val
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "max_gia_tri không hợp lệ"})
			return
		}
	}

	// ✅ 5. Lấy sort_by và sort_order (optional)
	sortBy := c.DefaultQuery("sort_by", "MaKhuyenMai")
	sortOrder := c.DefaultQuery("sort_order", "DESC")

	// ✅ 6. Lấy page và page_size (optional)
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	// ✅ 7. Gọi service
	khuyenMais, total, err := h.service.SearchKhuyenMai(keyword, maKhuyenMai, tenKhuyenMai, minGiaTri, maxGiaTri, sortBy, sortOrder, page, pageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ✅ 8. Tính total_pages
	totalPages := int(total) / pageSize
	if int(total)%pageSize != 0 {
		totalPages++
	}

	// ✅ 9. Response với đầy đủ thông tin
	c.JSON(http.StatusOK, gin.H{
		"message": "Tìm kiếm khuyến mãi thành công",
		"filters": gin.H{
			"keyword":        keyword,
			"ma_khuyen_mai":  maKhuyenMai,
			"ten_khuyen_mai": tenKhuyenMai,
			"min_gia_tri":    minGiaTri,
			"max_gia_tri":    maxGiaTri,
			"sort_by":        sortBy,
			"sort_order":     sortOrder,
		},
		"pagination": gin.H{
			"page":        page,
			"page_size":   pageSize,
			"total":       total,
			"total_pages": totalPages,
		},
		"data": khuyenMais,
	})
}
