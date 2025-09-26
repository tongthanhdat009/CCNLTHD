package handlers

import (
    "net/http"

    "github.com/tongthanhdat009/CCNLTHD/internal/services"

    "github.com/gin-gonic/gin"
    "strconv"
    "database/sql"
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
    var input struct {
        HoTen       *string `json:"ho_ten"`
        Email       *string `json:"email"`
        SoDienThoai *string `json:"so_dien_thoai"`
        TinhThanh   *string `json:"tinh_thanh"`
        QuanHuyen   *string `json:"quan_huyen"`
        PhuongXa    *string `json:"phuong_xa"`
        DuongSoNha  *string `json:"duong_so_nha"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Lấy dữ liệu hiện tại của người dùng
    nguoiDungCu, err := h.service.GetNguoiDungByID(maNguoiDung)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    // Chỉ update những trường có dữ liệu gửi lên
    if input.HoTen != nil {
        nguoiDungCu.HoTen = *input.HoTen
    }
    if input.Email != nil {
        nguoiDungCu.Email = *input.Email
    }
    if input.SoDienThoai != nil {
        nguoiDungCu.SoDienThoai = *input.SoDienThoai
    }
    if input.TinhThanh != nil {
        nguoiDungCu.TinhThanh = sql.NullString{String: *input.TinhThanh, Valid: true}
    }
    if input.QuanHuyen != nil {
        nguoiDungCu.QuanHuyen = sql.NullString{String: *input.QuanHuyen, Valid: true}
    }
    if input.PhuongXa != nil {
        nguoiDungCu.PhuongXa = sql.NullString{String: *input.PhuongXa, Valid: true}
    }
    if input.DuongSoNha != nil {
        nguoiDungCu.DuongSoNha = sql.NullString{String: *input.DuongSoNha, Valid: true}
    }

    if err := h.service.UpdateNguoiDung(maNguoiDung, *nguoiDungCu); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
