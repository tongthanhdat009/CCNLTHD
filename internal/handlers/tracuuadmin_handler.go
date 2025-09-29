package handlers

import (
    "github.com/gin-gonic/gin"
    "github.com/tongthanhdat009/CCNLTHD/internal/services"
    "strings"
    "log"
)

type TraCuuAdminHandler struct {
    service services.TraCuuAdminService
}

func NewTraCuuAdminHandler(service services.TraCuuAdminService) *TraCuuAdminHandler {
    return &TraCuuAdminHandler{service: service}  // Sửa: traCuuAdminHandler -> TraCuuAdminHandler
}

func (h *TraCuuAdminHandler) GetSanPhamBySeries(c *gin.Context) {
	seri := c.Param("seri")
	if strings.TrimSpace(seri) == "" {
		c.JSON(400, gin.H{"error": "Số seri không được để trống"})
		return
	}

	sanpham, err := h.service.GetSanPhamBySeries(seri)
	if err != nil {
		if strings.Contains(err.Error(), "không tồn tại") {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		log.Printf("Lỗi khi truy xuất sản phẩm: %v", err)
		c.JSON(500, gin.H{"error": "Đã xảy ra lỗi khi truy xuất sản phẩm"})
		return
	}

	c.JSON(200, sanpham)
}

func (h *TraCuuAdminHandler) GetSanPhamByTrangThai(c *gin.Context) {
    trangThai := c.Query("trangthai")  // Lấy từ query param
    if trangThai == "" {
        c.JSON(400, gin.H{"error": "thiếu tham số trangthai"})
        return
    }
    // Logic gọi service...
    sanphams, err := h.service.GetSanPhamByTrangThai(trangThai)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, sanphams)
}
