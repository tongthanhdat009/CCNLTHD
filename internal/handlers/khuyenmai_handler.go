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
	// Lấy id từ URL
	idParam := c.Param("id")
	makhuyenmai, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id không hợp lệ"})
		return
	}

	// Bind JSON request body vào struct
	khuyenMai := models.KhuyenMai{}
	if err := c.ShouldBindJSON(&khuyenMai); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Gọi service
	if err := h.service.UpdateKhuyenMai(makhuyenmai, khuyenMai); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả response
	c.JSON(http.StatusOK, gin.H{
		"message":   "Cập nhật thành công",
		"khuyenMai": khuyenMai,
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
	keyword := c.Query("keyword")
	khuyenMais, err := h.service.SearchKhuyenMai(keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, khuyenMais)
}
