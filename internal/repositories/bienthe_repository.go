package repositories

import (
	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"gorm.io/gorm"
)

type BienTheRepository interface {
	GetBienTheTheoMaHangHoa(maHangHoa int) ([]models.BienThe, error)
	GetBienTheTheoMa(maBienThe int) (*models.BienThe, error)
	CreateBienTheTheoMaHangHoa(bienThe *models.BienThe) error
	UpdateBienTheInfo(bienThe *models.BienThe) error
	UpdateBienTheStatus(maBienThe int, trangThai string) error
	DeleteBienThe(maBienThe int) error
	ExistsHangHoa(maHangHoa int) (bool, error)
	ExistsBienTheByHangHoaAndSize(maHangHoa int, size int) (bool, error)
	HasChiTietPhieuNhap(maBienThe int) (bool, error)
	IncrementSoLuongTon(maBienThe int, amount int) error
	UpdateSanPhamStatus(maSanPham int, trangThai string) error
	GetMaBienTheFromSanPham(maSanPham int) (int, error)
}

type BienTheRepo struct {
	db *gorm.DB
}

func NewBienTheRepository(db *gorm.DB) BienTheRepository {
	return &BienTheRepo{db: db}
}

func (r *BienTheRepo) HasChiTietPhieuNhap(maBienThe int) (bool, error) {
	var count int64
	if err := r.db.Model(&models.ChiTietPhieuNhap{}).Where("MaBienthe = ?", maBienThe).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *BienTheRepo) GetBienTheTheoMaHangHoa(maHangHoa int) ([]models.BienThe, error) {
	var bienThe []models.BienThe
	err := r.db.Where("MaHangHoa = ?", maHangHoa).Find(&bienThe).Error
	if err != nil {
		return nil, err
	}
	return bienThe, nil
}

func (r *BienTheRepo) GetBienTheTheoMa(maBienThe int) (*models.BienThe, error) {
	var bienThe models.BienThe
	err := r.db.Preload("HangHoa").
		Preload("HangHoa.Hang").
		Preload("HangHoa.DanhMuc").
		Preload("HangHoa.KhuyenMai").Where("MaBienThe = ?", maBienThe).First(&bienThe).Error
	if err != nil {
		return nil, err
	}
	return &bienThe, nil
}

func (r *BienTheRepo) CreateBienTheTheoMaHangHoa(bienThe *models.BienThe) error {
	return r.db.Create(bienThe).Error
}

func (r *BienTheRepo) UpdateBienTheInfo(bienThe *models.BienThe) error {
	return r.db.Model(&models.BienThe{}).Where("mabienthe = ?", bienThe.MaBienThe).Updates(map[string]interface{}{
		"Size": bienThe.Size,
		"Gia":  bienThe.Gia,
	}).Error
}

func (r *BienTheRepo) UpdateBienTheStatus(maBienThe int, trangThai string) error {
	return r.db.Model(&models.BienThe{}).Where("mabienthe = ?", maBienThe).Update("TrangThai", trangThai).Error
}

func (r *BienTheRepo) DeleteBienThe(maBienThe int) error {
	return r.db.Delete(&models.BienThe{}, maBienThe).Error
}

func (r *BienTheRepo) ExistsHangHoa(maHangHoa int) (bool, error) {
	var count int64
	err := r.db.Model(&models.HangHoa{}).Where("MaHangHoa = ?", maHangHoa).Count(&count).Error
	return count > 0, err
}

func (r *BienTheRepo) ExistsBienTheByHangHoaAndSize(maHangHoa int, size int) (bool, error) {
	var count int64
	err := r.db.Model(&models.BienThe{}).Where("MaHangHoa = ? AND Size = ?", maHangHoa, size).Count(&count).Error
	return count > 0, err
}

func (r *BienTheRepo) IncrementSoLuongTon(maBienThe int, amount int) error {
	return r.db.Model(&models.BienThe{}).Where("MaBienThe = ?", maBienThe).Update("SoLuongTon", gorm.Expr("SoLuongTon + ?", amount)).Error
}

func (r *BienTheRepo) UpdateSanPhamStatus(maSanPham int, trangThai string) error {
	return r.db.Model(&models.SanPham{}).Where("MaSanPham = ?", maSanPham).Update("TrangThai", trangThai).Error
}

func (r *BienTheRepo) GetMaBienTheFromSanPham(maSanPham int) (int, error) {
	var maBienThe int
	// Join SanPham -> ChiTietPhieuNhap -> BienThe
	err := r.db.Table("sanpham").
		Select("chitietphieunhap.MaBienthe").
		Joins("JOIN chitietphieunhap ON sanpham.MaChiTietPhieuNhap = chitietphieunhap.MaChiTiet").
		Where("sanpham.MaSanPham = ?", maSanPham).
		Scan(&maBienThe).Error
	return maBienThe, err
}
