package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/tongthanhdat009/CCNLTHD/internal/services"
    "time"
)

const (
	RefreshTokenDuration = time.Hour * 24 * 7
)
type DangNhapHandler struct {
    service services.DangNhapService
}

func NewDangNhapHandler(service services.DangNhapService) *DangNhapHandler {
    return &DangNhapHandler{service: service}
}

func (h *DangNhapHandler) KiemTraDangNhap(c *gin.Context) {
    var req struct {
        TenDangNhap string `json:"ten_dang_nhap" binding:"required"`
        MatKhau     string `json:"mat_khau" binding:"required"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    nguoiDung, accessToken, refreshToken, err := h.service.KiemTraDangNhap(req.TenDangNhap, req.MatKhau)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    // Tạo cookie Refresh Token
    c.SetCookie("refresh_token", refreshToken, int(RefreshTokenDuration.Seconds()), "/", "", false, true)

    // Trả Access Token trong header
    c.Header("Authorization", "Bearer "+accessToken)

    // Trả dữ liệu người dùng
    c.JSON(http.StatusOK, gin.H{"user": nguoiDung})
}
