package handlers

import (
    "net/http"

    "github.com/tongthanhdat009/CCNLTHD/internal/services"
    "github.com/tongthanhdat009/CCNLTHD/internal/models"
    "github.com/gin-gonic/gin"
    "strconv"
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
func (h *NguoiDungHandler) UpdateNguoiDung(c *gin.Context) {
    maNguoiDungStr := c.Param("id")
    maNguoiDung, err := strconv.Atoi(maNguoiDungStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }
    // Bind dữ liệu update (có thể chỉ một vài trường)
    var input models.NguoiDung

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.service.UpdateNguoiDung(maNguoiDung, input); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
func (h *NguoiDungHandler) GetNguoiDungByID(c *gin.Context) {
    maNguoiDungStr := c.Param("id")
    maNguoiDung, err := strconv.Atoi(maNguoiDungStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Mã người dùng không hợp lệ"})
        return
    }
    nguoiDung, err := h.service.GetNguoiDungByID(maNguoiDung)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, nguoiDung)
}

func (h *NguoiDungHandler) UpdateNguoiDungAdmin(c *gin.Context) {
    maNguoiDungStr := c.Param("id")
    maNguoiDung, err := strconv.Atoi(maNguoiDungStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    var input models.NguoiDung

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.service.UpdateNguoiDungAdmin(maNguoiDung, input); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
}
func (h *NguoiDungHandler) CreateNguoiDung(c *gin.Context) {
    var nd models.NguoiDung
    if err := c.ShouldBindJSON(&nd); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.service.CreateNguoiDung(nd); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Người dùng đã được tạo thành công"})
}
