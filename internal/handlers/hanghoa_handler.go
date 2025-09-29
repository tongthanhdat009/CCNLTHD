package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/tongthanhdat009/CCNLTHD/internal/models"
    "github.com/tongthanhdat009/CCNLTHD/internal/services"
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

func (h *HangHoaHandler) CreateHangHoa(c *gin.Context) {
    var hh models.HangHoa
    if err := c.ShouldBindJSON(&hh); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
        return
    }
    if err := h.service.CreateHangHoa(&hh); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Thêm hàng hóa thành công"})
}

func (h *HangHoaHandler) UpdateHangHoa(c *gin.Context) {
    var hh models.HangHoa

    // Lấy id từ URL
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID không hợp lệ"})
        return
    }

    // Bind JSON vào hh
    if err := c.ShouldBindJSON(&hh); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
        return
    }

    // Gán id vào model
    hh.MaHangHoa = id

    // Gọi service cập nhật
    if err := h.service.UpdateHangHoa(&hh); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Cập nhật hàng hóa thành công"})
}

func (h *HangHoaHandler) SearchHangHoa(c *gin.Context) {
    tenHangHoa := c.Query("ten_hang_hoa")
    tenDanhMuc := c.Query("ten_danh_muc")
    tenHang := c.Query("ten_hang")
    mau := c.Query("mau")
    trangThai := c.Query("trang_thai")
    maKhuyenMai := c.Query("ma_khuyen_mai")

    hangHoas, err := h.service.SearchHangHoa(tenHangHoa, tenDanhMuc, tenHang, mau, trangThai, maKhuyenMai)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    if len(hangHoas) == 0 {
        c.JSON(http.StatusOK, gin.H{"message": "Không có hàng hóa phù hợp"})
        return
    }
    c.JSON(http.StatusOK, hangHoas)
}