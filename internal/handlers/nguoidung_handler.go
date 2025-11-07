package handlers

import (
	"database/sql"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"github.com/tongthanhdat009/CCNLTHD/internal/services"
)

type NguoiDungHandler struct {
	service services.NguoiDungService
}

func NewNguoiDungHandler(service services.NguoiDungService) *NguoiDungHandler {
	return &NguoiDungHandler{service: service}
}

// Payload tạo người dùng để bind JSON (bao gồm cả mật khẩu)
type createNguoiDungRequest struct {
	TenDangNhap string `json:"ten_dang_nhap"`
	MatKhau     string `json:"mat_khau"`
	HoTen       string `json:"ho_ten"`
	Email       string `json:"email"`
	SoDienThoai string `json:"so_dien_thoai"`
	TinhThanh   string `json:"tinh_thanh"`
	QuanHuyen   string `json:"quan_huyen"`
	PhuongXa    string `json:"phuong_xa"`
	DuongSoNha  string `json:"duong_so_nha"`
	MaQuyen     int    `json:"ma_quyen"`
}

// Payload update người dùng (client gửi string cho các trường nullable)
type updateNguoiDungRequest struct {
	HoTen       *string `json:"ho_ten"`
	Email       *string `json:"email"`
	SoDienThoai *string `json:"so_dien_thoai"`
	TinhThanh   *string `json:"tinh_thanh"`
	QuanHuyen   *string `json:"quan_huyen"`
	PhuongXa    *string `json:"phuong_xa"`
	DuongSoNha  *string `json:"duong_so_nha"`
}

// Payload update người dùng (admin có thể cập nhật cả quyền)
type updateNguoiDungAdminRequest struct {
	HoTen       *string `json:"ho_ten"`
	Email       *string `json:"email"`
	SoDienThoai *string `json:"so_dien_thoai"`
	TinhThanh   *string `json:"tinh_thanh"`
	QuanHuyen   *string `json:"quan_huyen"`
	PhuongXa    *string `json:"phuong_xa"`
	DuongSoNha  *string `json:"duong_so_nha"`
	MaQuyen     *int    `json:"ma_quyen"`
}

func toNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{String: "", Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
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
	// Bind dữ liệu update (client gửi string cho các trường địa chỉ, có thể là chuỗi rỗng)
	var req updateNguoiDungRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Lấy dữ liệu hiện tại để chỉ cập nhật các trường có xuất hiện trong JSON
	current, err := h.service.GetNguoiDungByID(maNguoiDung)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	nd := *current
	if req.HoTen != nil {
		nd.HoTen = *req.HoTen
	}
	if req.Email != nil {
		nd.Email = *req.Email
	}
	if req.SoDienThoai != nil {
		nd.SoDienThoai = *req.SoDienThoai
	}
	if req.TinhThanh != nil {
		nd.TinhThanh = sql.NullString{String: *req.TinhThanh, Valid: true}
	}
	if req.QuanHuyen != nil {
		nd.QuanHuyen = sql.NullString{String: *req.QuanHuyen, Valid: true}
	}
	if req.PhuongXa != nil {
		nd.PhuongXa = sql.NullString{String: *req.PhuongXa, Valid: true}
	}
	if req.DuongSoNha != nil {
		nd.DuongSoNha = sql.NullString{String: *req.DuongSoNha, Valid: true}
	}

	if err := h.service.UpdateNguoiDung(maNguoiDung, nd); err != nil {
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

	var req updateNguoiDungAdminRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Lấy dữ liệu hiện tại và chỉ cập nhật các trường có xuất hiện
	current, err := h.service.GetNguoiDungByID(maNguoiDung)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	nd := *current
	if req.HoTen != nil {
		nd.HoTen = *req.HoTen
	}
	if req.Email != nil {
		nd.Email = *req.Email
	}
	if req.SoDienThoai != nil {
		nd.SoDienThoai = *req.SoDienThoai
	}
	if req.TinhThanh != nil {
		nd.TinhThanh = sql.NullString{String: *req.TinhThanh, Valid: true}
	}
	if req.QuanHuyen != nil {
		nd.QuanHuyen = sql.NullString{String: *req.QuanHuyen, Valid: true}
	}
	if req.PhuongXa != nil {
		nd.PhuongXa = sql.NullString{String: *req.PhuongXa, Valid: true}
	}
	if req.DuongSoNha != nil {
		nd.DuongSoNha = sql.NullString{String: *req.DuongSoNha, Valid: true}
	}
	if req.MaQuyen != nil {
		nd.MaQuyen = *req.MaQuyen
	}

	if err := h.service.UpdateNguoiDungAdmin(maNguoiDung, nd); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
func (h *NguoiDungHandler) CreateNguoiDung(c *gin.Context) {
	var req createNguoiDungRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	nd := models.NguoiDung{
		TenDangNhap: req.TenDangNhap,
		MatKhau:     req.MatKhau,
		HoTen:       req.HoTen,
		Email:       req.Email,
		SoDienThoai: req.SoDienThoai,
		TinhThanh:   toNullString(req.TinhThanh),
		QuanHuyen:   toNullString(req.QuanHuyen),
		PhuongXa:    toNullString(req.PhuongXa),
		DuongSoNha:  toNullString(req.DuongSoNha),
		MaQuyen:     req.MaQuyen,
	}

	if err := h.service.CreateNguoiDung(&nd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Người dùng đã được tạo thành công", "created": nd})
}
