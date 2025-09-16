package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/tongthanhdat009/CCNLTHD/internal/services" // Thêm import này
)

// BrandHandler xử lý các yêu cầu liên quan đến hãng
type HangHandler struct {
    Service *services.HangService // Cần import services package
}

// GetAllBrands trả về danh sách tất cả các hãng
func (h *HangHandler) GetAllBrands(c *gin.Context) {
    brands, err := h.Service.GetAllBrands()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch brands"})
        return
    }

    c.JSON(http.StatusOK, brands)
}