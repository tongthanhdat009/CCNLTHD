package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tongthanhdat009/CCNLTHD/internal/services"
)

type HangHoaHandler struct {
	service *services.HangHoaService
}

func NewHangHoaHandler(service *services.HangHoaService) *HangHoaHandler {
	return &HangHoaHandler{service: service}
}

func (h *HangHoaHandler) GetAllHangHoa(c *gin.Context) {
	hangHoas, err := h.service.GetAllHangHoa()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Lỗi khi lấy danh sách hàng hóa",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    hangHoas,
		"message": "Lấy danh sách hàng hóa thành công",
	})
}