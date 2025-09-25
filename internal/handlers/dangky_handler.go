package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"github.com/tongthanhdat009/CCNLTHD/internal/services"
)

// DTO cho request đăng ký (4 trường)
type DangKyRequest struct {
    TenDangNhap string `json:"ten_dang_nhap" binding:"required"`
    MatKhau     string `json:"mat_khau" binding:"required"`
    HoTen       string `json:"ho_ten" binding:"required"`
    Email       string `json:"email" binding:"required,email"`
}

type DangKyHandler struct {
    service services.DangKyService
}

func NewDangKyHandler(service services.DangKyService) *DangKyHandler {
    return &DangKyHandler{service: service}
}

func (h *DangKyHandler) CreateNguoiDung(c *gin.Context) {
    var req DangKyRequest

    // Bind JSON vào DTO
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Chuyển DTO -> model
    nguoiDung := models.NguoiDung{
        TenDangNhap: req.TenDangNhap,
        MatKhau:     req.MatKhau, // service sẽ hash mật khẩu
        HoTen:       req.HoTen,
        Email:       req.Email,
    }

    // Gọi service để tạo người dùng
    if err := h.service.CreateNguoiDung(nguoiDung); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Tạo tài khoản thành công"})
}
