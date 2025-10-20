package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/tongthanhdat009/CCNLTHD/internal/services"
	"strconv"
	"github.com/tongthanhdat009/CCNLTHD/internal/models"
)

type PhanQuyenHandler struct {
    service services.PhanQuyenService
}

func NewPhanQuyenHandler(service services.PhanQuyenService) *PhanQuyenHandler {
    return &PhanQuyenHandler{service: service}
}

// Thêm các phương thức xử lý khác tại đây, ví dụ:
func (h *PhanQuyenHandler) GetAll(c *gin.Context) {
	Quyens, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Quyens)
}
func (h *PhanQuyenHandler) GetByID(c *gin.Context) {
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
func (h *PhanQuyenHandler) UpdatePhanQuyen(c *gin.Context) {
	var phanquyen models.PhanQuyen
	if err := c.ShouldBindJSON(&phanquyen); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.UpdatePhanQuyen(&phanquyen); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, phanquyen)
}
