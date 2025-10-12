package handlers

import (
	"fmt"
	"net/http"

	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"github.com/tongthanhdat009/CCNLTHD/internal/services"

	"github.com/gin-gonic/gin"
)

type GioHangHandler struct {
	service services.GioHangService
}

func NewGioHangHandler(service services.GioHangService) *GioHangHandler {
	return &GioHangHandler{service: service}
}

func (h *GioHangHandler) TaoGioHang(c *gin.Context) {
	var giohang models.GioHang
	if err := c.ShouldBindJSON(&giohang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := h.service.TaoGioHang(giohang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Tạo giỏ hàng thành công"})
}

func (h *GioHangHandler) SuaGioHang(c *gin.Context) {
	var giohang models.GioHang
	if err := c.ShouldBindJSON(&giohang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := h.service.SuaGioHang(giohang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Sửa giỏ hàng thành công"})
}

func (h *GioHangHandler) XoaGioHang(c *gin.Context) {
	var giohang models.GioHang
	if err := c.ShouldBindJSON(&giohang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := h.service.XoaGioHang(giohang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Xóa giỏ hàng thành công"})
}

func (h *GioHangHandler) GetAll(c *gin.Context) {
	// Lấy manguoidung từ URL parameter hoặc query parameter
	manguoidungStr := c.Param("id")

	if manguoidungStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Thiếu mã người dùng"})
		return
	}

	// Chuyển string thành int
	var manguoidung int
	if _, err := fmt.Sscanf(manguoidungStr, "%d", &manguoidung); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Mã người dùng không hợp lệ"})
		return
	}

	giohangs, err := h.service.GetAll(manguoidung)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể lấy danh sách giỏ hàng"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Lấy danh sách giỏ hàng thành công",
		"data":    giohangs,
	})
}
func (h *GioHangHandler) ThanhToan(c *gin.Context) {
	var req struct {
		GioHang    []models.GioHang `json:"giohang"`
		MaNguoiDung int             `json:"ma_nguoi_dung"`
		TinhThanh           string    `json:"tinh_thanh"`
		QuanHuyen           string    `json:"quan_huyen"`
		PhuongXa            string    `json:"phuong_xa"`
		DuongSoNha          string    `json:"duong_so_nha"`
		SoDienThoai        string    `json:"so_dien_thoai"`
		PhuongThucThanhToan string    `json:"phuong_thuc_thanh_toan"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if donhang, err := h.service.ThanhToan(req.GioHang, req.MaNguoiDung, req.TinhThanh, req.QuanHuyen, req.PhuongXa, req.DuongSoNha, req.PhuongThucThanhToan, req.SoDienThoai); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Thanh toán thành công", "data": donhang})
	}
}
