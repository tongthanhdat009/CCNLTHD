package handlers

import (
    "github.com/gin-gonic/gin"
    "github.com/tongthanhdat009/CCNLTHD/internal/services"
    "net/http"
)

type TimKiemSanPhamHandler struct {
	service services.TimKiemSanPhamService
}

func NewTimKiemSanPhamHandler(service services.TimKiemSanPhamService) *TimKiemSanPhamHandler {
	return &TimKiemSanPhamHandler{service: service}
}


func (h *TimKiemSanPhamHandler) TimSanPham(c *gin.Context) {
    tenHangHoa := c.Query("tenHangHoa")
    tenHang := c.Query("tenHang")
    tenDanhMuc := c.Query("tenDanhMuc")
    mau := c.Query("mau")
    size := c.Query("size")
    giaToiThieuStr := c.Query("giaToiThieu")
    giaToiDaStr := c.Query("giaToiDa")

    sanPhams, err := h.service.TimSanPham(tenHangHoa, tenHang, tenDanhMuc, mau, size, giaToiThieuStr, giaToiDaStr)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if len(sanPhams) == 0 {
        c.JSON(http.StatusOK, gin.H{"message": "Không có sản phẩm phù hợp"})
        return
    }

    c.JSON(http.StatusOK, sanPhams)
}