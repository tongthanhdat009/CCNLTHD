package handlers

import (
	"net/http"

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

func (h *KhuyenMaiHandler) TaoKhuyenMai(c *gin.Context) {
	var khuyenMai models.KhuyenMai
	if err := c.ShouldBindJSON(&khuyenMai); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.TaoKhuyenMai(khuyenMai); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, khuyenMai)
}
