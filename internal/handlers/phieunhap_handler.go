package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"github.com/tongthanhdat009/CCNLTHD/internal/services"

	"time"
)

type QuanLyPhieuNhapHandler struct {
    service services.QuanLyPhieuNhapService
}

func NewQuanLyPhieuNhapHandler(service services.QuanLyPhieuNhapService) *QuanLyPhieuNhapHandler {
    return &QuanLyPhieuNhapHandler{service: service}
}

// GetAllPhieuNhaps: Lấy tất cả phiếu nhập
func (h *QuanLyPhieuNhapHandler) GetAllPhieuNhaps(c *gin.Context) {
    phieuNhaps, err := h.service.GetAllPhieuNhaps()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch phiếu nhập"})
        return
    }
    c.JSON(http.StatusOK, phieuNhaps)
}

// CreatePhieuNhap: Tạo mới phiếu nhập
func (h *QuanLyPhieuNhapHandler) CreatePhieuNhap(c *gin.Context) {
    var phieuNhap models.PhieuNhap
    if err := c.ShouldBindJSON(&phieuNhap); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
        return
    }
    if err := h.service.CreatePhieuNhap(&phieuNhap); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"message": "Tạo phiếu nhập thành công", "data": phieuNhap})
}

func (h *QuanLyPhieuNhapHandler) SearchPhieuNhaps(c *gin.Context) {
    tenNguoiDung := c.Query("tenNguoiDung")
    tenNhaCungCap := c.Query("tenNhaCungCap")
    trangThai := c.Query("trangThai")

    var tuNgay, denNgay *time.Time
    if tuNgayStr := c.Query("tuNgay"); tuNgayStr != "" {
        t, err := time.Parse("2006-01-02", tuNgayStr)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tuNgay format. Use YYYY-MM-DD"})
            return
        }
        tuNgay = &t
    }
    if denNgayStr := c.Query("denNgay"); denNgayStr != "" {
        t, err := time.Parse("2006-01-02", denNgayStr)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid denNgay format. Use YYYY-MM-DD"})
            return
        }
        denNgay = &t
    }

    phieuNhaps, err := h.service.SearchPhieuNhaps(tenNguoiDung, tenNhaCungCap, trangThai, tuNgay, denNgay)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "count":   len(phieuNhaps),
        "data":    phieuNhaps,
    })
}

func (h *QuanLyPhieuNhapHandler) GetPhieuNhapByID(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    phieuNhap, err := h.service.GetPhieuNhapByID(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    if phieuNhap == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Phiếu nhập không tồn tại"})
        return
    }   
    c.JSON(http.StatusOK, phieuNhap)
}

func (h *QuanLyPhieuNhapHandler) DeletePhieuNhap(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID không hợp lệ"})
        return
    }
    if err := h.service.DeletePhieuNhap(id); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Xóa phiếu nhập thành công"})
}

// UpdatePhieuNhap: Cập nhật trạng thái phiếu nhập
func (h *QuanLyPhieuNhapHandler) UpdatePhieuNhap(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID không hợp lệ"})
        return
    }

    phieuNhap := models.PhieuNhap{
        MaPhieuNhap: id,
        TrangThai:   "Đã duyệt", 
    }
    
    if err := h.service.UpdatePhieuNhap(&phieuNhap, true); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Cập nhật phiếu nhập thành công",
        "data": phieuNhap,
        "approved": true,
    })
}

// ApprovePhieuNhap: Duyệt phiếu nhập (convenience method)
func (h *QuanLyPhieuNhapHandler) ApprovePhieuNhap(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID không hợp lệ"})
        return
    }
    
    phieuNhap := models.PhieuNhap{
        MaPhieuNhap: id,
        TrangThai:   "Đã duyệt",
    }
    
    if err := h.service.UpdatePhieuNhap(&phieuNhap, true); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{
        "message": "Duyệt phiếu nhập thành công",
        "data": phieuNhap,
    })
}

// ExistsInPhieuNhap: Kiểm tra nhà cung cấp có trong phiếu nhập không (dùng query param)
func (h *QuanLyPhieuNhapHandler) ExistsInPhieuNhap(c *gin.Context) {
    nhaCungCapIDStr := c.Query("nhaCungCapID")
    nhaCungCapID, err := strconv.Atoi(nhaCungCapIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid nhaCungCapID"})
        return
    }
    exists, err := h.service.ExistsInPhieuNhap(nhaCungCapID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"exists": exists})
}

// GetChiTietPhieuNhap: Lấy tất cả chi tiết của một phiếu nhập
func (h *QuanLyPhieuNhapHandler) GetChiTietPhieuNhap(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID không hợp lệ"})
        return
    }
    
    chiTiet, err := h.service.GetChiTietByPhieuNhap(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "count": len(chiTiet),
        "data": chiTiet,
    })
}

// DeleteChiTietPhieuNhap: Xóa tất cả chi tiết của một phiếu nhập
func (h *QuanLyPhieuNhapHandler) DeleteChiTietPhieuNhap(c *gin.Context) {
    idPhieuNhapStr := c.Param("id")
    idStr := c.Param("maChiTietPhieuNhap")

    idPhieuNhap, err := strconv.Atoi(idPhieuNhapStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID phiếu nhập không hợp lệ"})
        return
    }

    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID không hợp lệ"})
        return
    }
    log.Println("Deleting ChiTietPhieuNhap with MaChiTiet:", id, "from MaPhieuNhap:", idPhieuNhap)

    if err := h.service.DeleteChiTietByPhieuNhap(idPhieuNhap,id); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "message": "Đã xóa chi tiết phiếu nhập",
    })
}

func (h *QuanLyPhieuNhapHandler) DeleteAllChiTietPhieuNhap(c *gin.Context) {
    idStr := c.Param("id")

    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID không hợp lệ"})
        return
    }
    
    if err := h.service.DeleteAllChiTietByPhieuNhap(id); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "message": "Đã xóa tất cả chi tiết phiếu nhập",
    })
}

// CreateChiTietPhieuNhap: Tạo mới nhiều chi tiết cho phiếu nhập
func (h *QuanLyPhieuNhapHandler) CreateChiTietPhieuNhap(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID không hợp lệ"})
        return
    }
    
    var chiTietItems []models.ChiTietPhieuNhap
    if err := c.ShouldBindJSON(&chiTietItems); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ", "details": err.Error()})
        return
    }
    
    if err := h.service.CreateChiTietPhieuNhap(id, chiTietItems); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusCreated, gin.H{
        "success": true,
        "message": "Đã thêm chi tiết phiếu nhập thành công",
        "count": len(chiTietItems),
    })
}

// UpdateChiTietPhieuNhapSoLuong: Cập nhật số lượng của một chi tiết phiếu nhập
func (h *QuanLyPhieuNhapHandler) UpdateChiTietPhieuNhapSoLuong(c *gin.Context) {
    // Lấy mã phiếu nhập từ URL parameter
    idPhieuNhapStr := c.Param("id")
    idPhieuNhap, err := strconv.Atoi(idPhieuNhapStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID phiếu nhập không hợp lệ"})
        return
    }

    // Lấy mã chi tiết từ URL parameter
    idChiTietStr := c.Param("maChiTietPhieuNhap")
    idChiTiet, err := strconv.Atoi(idChiTietStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID chi tiết phiếu nhập không hợp lệ"})
        return
    }

    // Lấy số lượng mới từ request body
    var requestBody struct {
        SoLuong int `json:"so_luong"`
    }
    if err := c.ShouldBindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ", "details": err.Error()})
        return
    }

    // Kiểm tra số lượng
    if requestBody.SoLuong <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Số lượng phải lớn hơn 0"})
        return
    }

    // Gọi service để cập nhật số lượng
    if err := h.service.UpdateChiTietPhieuNhapSoLuong(idPhieuNhap, idChiTiet, requestBody.SoLuong); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Trả về kết quả thành công
    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "message": "Đã cập nhật số lượng chi tiết phiếu nhập thành công",
        "data": gin.H{
            "ma_phieu_nhap": idPhieuNhap,
            "ma_chi_tiet": idChiTiet,
            "so_luong_moi": requestBody.SoLuong,
        },
    })
}