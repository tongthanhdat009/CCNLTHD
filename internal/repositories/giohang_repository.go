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
	CheckBienThe(maBienThe int, gia int, soLuong int) error
	CreateDonHang(tx *gorm.DB, donhang *models.DonHang) error
	BeginTransaction() *gorm.DB
	GetSanPham(tx *gorm.DB, maBienThe int, soLuong int) ([]models.SanPham, error)
	CreateChiTietDonHang(tx *gorm.DB, chiTiet []models.ChiTietDonHang) error
	XoaGioHangCuaNguoiDung(tx *gorm.DB, manguoidung int) error
	UpdateTonKhoBienThe(tx *gorm.DB, maBienThe int, soLuong int) error
}
type GioHangRepo struct {
	db *gorm.DB
}
func (r *GioHangRepo) BeginTransaction() *gorm.DB {
	return r.db.Begin()
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

func (s *GioHangRepo) CheckBienThe(maBienThe int, gia int, soluong int)  error {
	var bienThe models.BienThe
	err := s.db.Where("MaBienThe = ?", maBienThe).First(&bienThe).Error
	if err != nil {
		return fmt.Errorf("không tìm thấy sản phẩm mã biến thể %d", maBienThe)
	}
	if bienThe.Gia != float64(gia) {
		return fmt.Errorf("giá sản phẩm mã biến thể %d đã thay đổi (giỏ: %.2f, hiện tại: %.2f)",
			maBienThe, float64(gia), bienThe.Gia)
	}
	if bienThe.SoLuongTon < soluong {
		return fmt.Errorf("số lượng sản phẩm mã biến thể %d không đủ (giỏ: %d, hiện tại: %d)",
			maBienThe, soluong, bienThe.SoLuongTon)
	}
	return nil
}

func (r *GioHangRepo) CreateDonHang(tx *gorm.DB, donhang *models.DonHang) error {
	return tx.Create(donhang).Error
}

func (r *GioHangRepo) GetSanPham(tx *gorm.DB, maBienThe int, soLuong int) ([]models.SanPham, error) {
	var sanPhams []models.SanPham

	// ✅ 1. Lấy sản phẩm chưa bán thuộc biến thể cụ thể
	if err := tx.
		Table("SanPham").
		Joins("JOIN ChiTietPhieuNhap ON ChiTietPhieuNhap.MaChiTiet = SanPham.MaChiTietPhieuNhap").
		Where("ChiTietPhieuNhap.MaBienThe = ? AND SanPham.TrangThai = ?", maBienThe, "Chưa bán").
		Limit(soLuong).
		Find(&sanPhams).Error; err != nil {
		return nil, err
	}

	if len(sanPhams) < soLuong {
		return nil, fmt.Errorf("không đủ hàng trong kho cho biến thể %d", maBienThe)
	}

	// 2️⃣ Cập nhật trạng thái thành "Chờ duyệt"
	ids := make([]int, len(sanPhams))
	for i, sp := range sanPhams {
		ids[i] = sp.MaSanPham
	}

	if err := tx.Model(&models.SanPham{}).
		Where("MaSanPham IN ?", ids).
		Update("TrangThai", "Chờ duyệt").Error; err != nil {
		return nil, err
	}

	// 3️⃣ Trả về danh sách sản phẩm đã cập nhật
	return sanPhams, nil
}

func (r *GioHangRepo) CreateChiTietDonHang(tx *gorm.DB, chiTiet []models.ChiTietDonHang) error {
	return tx.Create(&chiTiet).Error
}

func (s *GioHangRepo) XoaGioHangCuaNguoiDung(tx *gorm.DB, manguoidung int) error {
	return tx.Where("MaNguoiDung = ?", manguoidung).Delete(&models.GioHang{}).Error
}
func (r *GioHangRepo) UpdateTonKhoBienThe(tx *gorm.DB, maBienThe int, soLuong int) error {
	return tx.Model(&models.BienThe{}).
		Where("MaBienThe = ?", maBienThe).
		Update("SoLuongTon", gorm.Expr("SoLuongTon - ?", soLuong)).Error
}
