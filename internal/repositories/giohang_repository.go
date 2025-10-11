package repositories

import (
	"errors"
	"fmt"

	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"gorm.io/gorm"
)

type GioHangRepository interface {
	TaoGioHang(giohang models.GioHang) error
	LayGia(mabienthe int) (float64, error)
	CheckStatus(mabienthe int) error
	KiemTraTonTai(manguoidung int, mabienthe int) bool
	CheckSoLuong(soluong int, mabienthe int) bool
	SuaGioHang(giohang models.GioHang) error
	XoaGioHang(giohang models.GioHang) error
	GetAll(manguoidung int) ([]models.GioHang, error)
	GetAllGia(manguoidung int) (float64, error)
	GetByID(mabienthe int) ([]models.GioHang, error)
	CreateDH(donHang models.DonHang) error
	GetSanPham(mabienthe int, soluong int) ([]models.SanPham, error)
}
type GioHangRepo struct {
	db *gorm.DB
}

func NewGioHangRepository(db *gorm.DB) GioHangRepository {
	return &GioHangRepo{db: db}
}

func (r *GioHangRepo) TaoGioHang(giohang models.GioHang) error {
	return r.db.Create(&giohang).Error
}

func (r *GioHangRepo) LayGia(mabienthe int) (float64, error) {
	var bienThe models.BienThe
	err := r.db.Where("MaBienThe = ?", mabienthe).First(&bienThe).Error
	if err != nil {
		return 0, err
	}
	return bienThe.Gia, nil
}

func (r *GioHangRepo) CheckStatus(mabienthe int) error {
	var bienThe models.BienThe
	err := r.db.Where("MaBienThe = ? AND TrangThai = ?", mabienthe, "DangBan").First(&bienThe).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *GioHangRepo) KiemTraTonTai(manguoidung int, mabienthe int) bool {
	var gioHang models.GioHang
	err := r.db.Where("MaNguoiDung = ? AND MaBienThe = ?", manguoidung, mabienthe).First(&gioHang).Error
	if err != nil {
		return false // Không tồn tại
	}
	return true // Đã tồn tại
}

func (r *GioHangRepo) CheckSoLuong(soluong int, mabienthe int) bool {
	var bienThe models.BienThe
	err := r.db.Where("MaBienThe = ? ", mabienthe).First(&bienThe).Error
	if err != nil {
		return false
	}
	if soluong > bienThe.SoLuongTon {
		return false
	}
	return true
}

func (r *GioHangRepo) SuaGioHang(giohang models.GioHang) error {
	return r.db.Where("MaNguoiDung = ? AND MaBienThe = ?", giohang.MaNguoiDung, giohang.MaBienThe).Updates(&giohang).Error
}

func (r *GioHangRepo) XoaGioHang(giohang models.GioHang) error {
	// Cách 1: Xóa theo điều kiện (recommended)
	result := r.db.Where("MaNguoiDung = ? AND MaBienThe = ?", giohang.MaNguoiDung, giohang.MaBienThe).Delete(&models.GioHang{})
	if result.Error != nil {
		return result.Error
	}
	fmt.Print(result)
	if result.RowsAffected == 0 {
		return errors.New("không tìm thấy record để xóa")
	}
	return nil
}

func (r *GioHangRepo) GetAll(manguoidung int) ([]models.GioHang, error) {
	var results []struct {
		MaNguoiDung int     `json:"ma_nguoi_dung"`
		MaBienThe   int     `json:"ma_bien_the"`
		SoLuong     int     `json:"so_luong"`
		Gia         float64 `json:"gia"`
	}

	query := `
        SELECT 
            gh.MaNguoiDung,
            gh.MaBienThe,
            gh.SoLuong,
            CASE 
                WHEN km.GiaTri IS NOT NULL THEN bt.Gia * (100 - km.GiaTri) / 100
                ELSE bt.Gia 
            END AS Gia
        FROM GioHang gh
        JOIN bienthe bt ON gh.MaBienThe = bt.MaBienThe
        JOIN hanghoa hh ON bt.MaHangHoa = hh.MaHangHoa
        LEFT JOIN khuyenmai km ON hh.MaKhuyenMai = km.MaKhuyenMai
        WHERE gh.MaNguoiDung = ?
    `

	err := r.db.Raw(query, manguoidung).Scan(&results).Error
	if err != nil {
		return nil, err
	}

	// Convert sang []models.GioHang
	var gioHangs []models.GioHang
	for _, result := range results {
		gioHangs = append(gioHangs, models.GioHang{
			MaNguoiDung: result.MaNguoiDung,
			MaBienThe:   result.MaBienThe,
			SoLuong:     result.SoLuong,
			Gia:         result.Gia,
		})
	}

	return gioHangs, nil
}

func (r *GioHangRepo) GetSanPham(mabienthe int, soluong int) ([]models.SanPham, error) {
	var sanPham []models.SanPham
	err := r.db.Table("san_pham").
		Select("san_pham.*, chi_tiet_phieu_nhap.GiaNhap as gia_ban").
		Joins("JOIN chi_tiet_phieu_nhap ON san_pham.MaChiTietPhieuNhap = chi_tiet_phieu_nhap.MaChiTiet").
		Joins("JOIN bien_the ON chi_tiet_phieu_nhap.MaBienthe = bien_the.MaBienThe").
		Where("bien_the.MaBienThe = ? AND san_pham.TrangThai = ?", mabienthe, "Chưa bán").
		Limit(soluong).
		First(&sanPham).Error

	if err != nil {
		return sanPham, err
	}
	if len(sanPham) < soluong {
		return sanPham, errors.New("không đủ sản phẩm trong kho")
	}
	return sanPham, nil
}

func (r *GioHangRepo) GetAllGia(manguoidung int) (float64, error) {
	var totalGia float64
	err := r.db.Model(&models.GioHang{}).
		Where("MaNguoiDung = ?", manguoidung).
		Select("SUM(Gia * SoLuong)").
		Scan(&totalGia).Error
	if err != nil {
		return 0, err
	}
	return totalGia, nil
}

func (r *GioHangRepo) GetByID(manguoidung int) ([]models.GioHang, error) {
	var gioHang []models.GioHang
	err := r.db.Where("MaNguoiDung = ?", manguoidung).Find(&gioHang).Error
	if err != nil {
		return nil, err
	}
	return gioHang, nil
}

func (r *GioHangRepo) CreateDH(donHang models.DonHang) error {
	return r.db.Create(&donHang).Error
}

func (r *GioHangRepo) CreateChiTietDonHang(madonhang int, sanpham []models.SanPham) error {
	for _, sp := range sanpham {
		if err := r.db.Create(&models.ChiTietDonHang{MaDonHang: madonhang, MaSanPham: sp.MaSanPham, GiaBan: sp.GiaBan}).Error; err != nil {
			return err
		}
	}
	return nil
}