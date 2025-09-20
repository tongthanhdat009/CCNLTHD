package handlers

import (
    "net/http"

    "github.com/tongthanhdat009/CCNLTHD/internal/services"

    "github.com/gin-gonic/gin"
)

type NguoiDungHandler struct {
    service services.NguoiDungService
}

func NewNguoiDungHandler(service services.NguoiDungService) *NguoiDungHandler {
    return &NguoiDungHandler{service: service}
}

func (h *NguoiDungHandler) GetAllNguoiDung(c *gin.Context) {
    NguoiDungs, err := h.service.GetAllNguoiDung()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch products"})
        return
    }
    c.JSON(http.StatusOK, NguoiDungs)
}