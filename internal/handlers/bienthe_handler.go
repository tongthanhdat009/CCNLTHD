package handlers

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"github.com/tongthanhdat009/CCNLTHD/internal/services"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "maBienThe không hợp lệ"})
		return
	}
	err = h.service.DeleteBienThe(bienTheIDint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể xóa biến thể"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Xóa biến thể thành công"})
}

func (h *QuanLyBienTheHandler) GetBienTheTheoHangHoa(c *gin.Context) {
    hangHoaIDStr := c.Param("maHangHoa")
    hangHoaIDint, err := strconv.Atoi(hangHoaIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "hangHoaID không hợp lệ"})
        return
    }
    bienThes, err := h.service.GetBienTheTheoMaHangHoa(hangHoaIDint)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch variants"})
        return
    }
	if len(bienThes) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Không có biến thể nào cho hàng hóa này"})
		return
	}
    c.JSON(http.StatusOK, bienThes)
}

func (h *QuanLyBienTheHandler) GetBienTheTheoMa(c *gin.Context) {
    bienTheIDStr := c.Param("maBienThe")  
    bienTheIDint, err := strconv.Atoi(bienTheIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "maBienThe không hợp lệ"})
        return
    }
    bienThe, err := h.service.GetBienTheTheoMa(bienTheIDint)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể lấy biến thể"})
        return
    }
    c.JSON(http.StatusOK, bienThe)
}

func (h *QuanLyBienTheHandler) CreateBienThe(c *gin.Context) {
	var bienThe models.BienThe
	if err := c.ShouldBindJSON(&bienThe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu đầu vào không hợp lệ"})
		return
	}
	if err := h.service.CreateBienTheTheoMaHangHoa(&bienThe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, bienThe)
}

func (h *QuanLyBienTheHandler) UpdateBienTheInfo(c *gin.Context) {
	maBienTheStr := c.Param("maBienThe")
	maBienThe, err := strconv.Atoi(maBienTheStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Mã biến thể không hợp lệ"})
		return
	}

	var bienThe models.BienThe
	if err := c.ShouldBindJSON(&bienThe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bienThe.MaBienThe = maBienThe

	if err := h.service.UpdateBienTheInfo(&bienThe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bienThe)
}

type UpdateStatusRequest struct {
	TrangThai string `json:"trangThai"`
}

func (h *QuanLyBienTheHandler) UpdateBienTheStatus(c *gin.Context) {
	maBienTheStr := c.Param("maBienThe")
	maBienThe, err := strconv.Atoi(maBienTheStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid maBienThe"})
		return
	}

	var req UpdateStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: 'trangThai' is required"})
		return
	}

	if err := h.service.UpdateBienTheStatus(maBienThe, req.TrangThai); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cập nhật trạng thái thành công"})
}

