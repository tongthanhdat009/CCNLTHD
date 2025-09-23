package handlers

import (
    "net/http"

    "github.com/tongthanhdat009/CCNLTHD/internal/services"

    "github.com/gin-gonic/gin"
)

type DonHangHandler struct {
    service services.DonHangService
}

func NewDonHangHandler(service services.DonHangService) *DonHangHandler {
    return &DonHangHandler{service: service}
}

func (h *DonHangHandler) GetAllDonHang(c *gin.Context) {
    DonHangs, err := h.service.GetAllDonHang()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch products"})
        return
    }
    c.JSON(http.StatusOK, DonHangs)
}
