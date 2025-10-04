package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"github.com/tongthanhdat009/CCNLTHD/internal/services"
	"strconv"
)

type QuanLyBienTheHandler struct {
	service services.QuanLyBienTheService
}

func NewQuanLyBienTheHandler(service services.QuanLyBienTheService) *QuanLyBienTheHandler {
	return &QuanLyBienTheHandler{service: service}
}

func (h *QuanLyBienTheHandler) DeleteBienThe(c *gin.Context) {
	bienTheIDStr := c.Param("maBienThe")
	bienTheIDint, err := strconv.Atoi(bienTheIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid maBienThe"})
		return
	}
	err = h.service.DeleteBienThe(bienTheIDint)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (h *QuanLyBienTheHandler) GetBienTheTheoHangHoa(c *gin.Context) {
    hangHoaIDStr := c.Param("maHangHoa")
    hangHoaIDint, err := strconv.Atoi(hangHoaIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hangHoaID"})
        return
    }
    bienThes, err := h.service.GetBienTheTheoMaHangHoa(hangHoaIDint)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch variants"})
        return
    }
    c.JSON(http.StatusOK, bienThes)
}

func (h *QuanLyBienTheHandler) GetBienTheTheoMa(c *gin.Context) {
    bienTheIDStr := c.Param("maBienThe")  
    bienTheIDint, err := strconv.Atoi(bienTheIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid maBienThe"})
        return
    }
    bienThe, err := h.service.GetBienTheTheoMa(bienTheIDint)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch variant"})
        return
    }
    c.JSON(http.StatusOK, bienThe)
}

func (h *QuanLyBienTheHandler) CreateBienThe(c *gin.Context) {
	var bienThe models.BienThe
	if err := c.ShouldBindJSON(&bienThe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := h.service.CreateBienTheTheoMaHangHoa(&bienThe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, bienThe)
}

func (h *QuanLyBienTheHandler) UpdateBienThe(c *gin.Context) {
	var bienThe models.BienThe
	if err := c.ShouldBindJSON(&bienThe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := h.service.UpdateBienThe(&bienThe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bienThe)
}

