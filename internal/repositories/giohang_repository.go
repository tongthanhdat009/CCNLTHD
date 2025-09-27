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
	var Giohangs []models.GioHang
	err := r.db.Find(&Giohangs).Error
	return Giohangs, err
}
