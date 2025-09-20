package handlers

import (
    "net/http"

    "github.com/tongthanhdat009/CCNLTHD/internal/services"

    "github.com/gin-gonic/gin"
)

type HangHoaHandler struct {
    service services.HangHoaService
}

func NewHangHoaHandler(service services.HangHoaService) *HangHoaHandler {
    return &HangHoaHandler{service: service}
}

func (h *HangHoaHandler) GetAllHangHoa(c *gin.Context) {
    hangHoas, err := h.service.GetAllHangHoa()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch products"})
        return
    }
    c.JSON(http.StatusOK, hangHoas)
}