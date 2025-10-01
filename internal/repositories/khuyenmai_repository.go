package repositories

import (
	"errors"
	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"gorm.io/gorm"
)

type KhuyenMaiRepository interface {
	TaoKhuyenMai(khuyenMai *models.KhuyenMai) error
	SuaKhuyenMai(makhuyenmai int, khuyenMai models.KhuyenMai) error
	XoaKhuyenMai(makhuyenmai int) error
	KiemTraTenTonTai(tenkhuyenmai string) (bool, error)
	GetAll() ([]models.KhuyenMai, error)
	GetByID(makhuyenmai int) (models.KhuyenMai, error)
}

type KhuyenMaiRepo struct {
	db *gorm.DB
}

func NewKhuyenMaiRepository(db *gorm.DB) KhuyenMaiRepository {
	return &KhuyenMaiRepo{db: db}
}

func (r *KhuyenMaiRepo) TaoKhuyenMai(khuyenMai *models.KhuyenMai) error {
	return r.db.Create(khuyenMai).Error
}

func (r *KhuyenMaiRepo) KiemTraTenTonTai(tenkhuyenmai string) (bool, error) {
	var khuyenMai models.KhuyenMai
	err := r.db.Where("TenKhuyenMai = ?", tenkhuyenmai).First(&khuyenMai).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil // Không tồn tại
	}
	if err != nil {
		return false, err // Lỗi DB
	}
	return true, nil // Đã tồn tại
}

func (r *KhuyenMaiRepo) SuaKhuyenMai(makhuyenmai int, khuyenMai models.KhuyenMai) error {
	return r.db.Model(&models.KhuyenMai{}).Where("MaKhuyenMai = ?", makhuyenmai).Updates(khuyenMai).Error
}

func (r *KhuyenMaiRepo) XoaKhuyenMai(makhuyenmai int) error {
	result := r.db.Where("MaKhuyenMai = ?", makhuyenmai).
		Delete(&models.KhuyenMai{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("không tìm thấy record để xóa")
	}
	return nil
}

func (r *KhuyenMaiRepo) GetAll() ([]models.KhuyenMai, error) {
	var khuyenMais []models.KhuyenMai
	err := r.db.Find(&khuyenMais).Error
	return khuyenMais, err
}

func (r *KhuyenMaiRepo) GetByID(makhuyenmai int) (models.KhuyenMai, error) {
	var khuyenMai models.KhuyenMai
	err := r.db.Where("MaKhuyenMai = ?", makhuyenmai).First(&khuyenMai).Error
	return khuyenMai, err
}

// func (r *GioHangRepo) GetAll(manguoidung int) ([]models.GioHang, error) {
// 	var gioHangs []models.GioHang

// 	err := r.db.Where("MaNguoiDung = ?", manguoidung).
// 		Preload("BienThe").
// 		Preload("BienThe.HangHoa").
// 		Preload("BienThe.HangHoa.KhuyenMai").
// 		Find(&gioHangs).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	// Tính toán giá có áp dụng khuyến mãi
// 	for i := range gioHangs {
// 		if gioHangs[i].BienThe.MaBienThe != 0 {
// 			// Lấy giá gốc từ BienThe
// 			giaGoc := gioHangs[i].BienThe.Gia

// 			// Kiểm tra có khuyến mãi không
// 			if gioHangs[i].BienThe.HangHoa.MaKhuyenMai.Valid {
// 				// Có khuyến mãi
// 				giaTriKhuyenMai := gioHangs[i].BienThe.HangHoa.KhuyenMai.GiaTri
// 				// Tính giá sau khuyến mãi: giá gốc - (giá gốc * % khuyến mãi / 100)
// 				gioHangs[i].Gia = giaGoc * (100 - giaTriKhuyenMai) / 100
// 			} else {
// 				// Không có khuyến mãi, giữ nguyên giá gốc
// 				gioHangs[i].Gia = giaGoc
// 			}
// 		}
// 	}

// 	return gioHangs, nil
// }

// func (r *GioHangRepo) GetAll(manguoidung int) ([]models.GioHang, error) {
// 	var results []struct {
// 		MaNguoiDung int     `json:"ma_nguoi_dung"`
// 		MaBienThe   int     `json:"ma_bien_the"`
// 		SoLuong     int     `json:"so_luong"`
// 		Gia         float64 `json:"gia"`
// 	}

// 	query := `
//         SELECT
//             gh.MaNguoiDung,
//             gh.MaBienThe,
//             gh.SoLuong,
//             CASE
//                 WHEN km.GiaTri IS NOT NULL THEN bt.Gia * (100 - km.GiaTri) / 100
//                 ELSE bt.Gia
//             END AS Gia
//         FROM GioHang gh
//         JOIN bienthe bt ON gh.MaBienThe = bt.MaBienThe
//         JOIN hanghoa hh ON bt.MaHangHoa = hh.MaHangHoa
//         LEFT JOIN khuyenmai km ON hh.MaKhuyenMai = km.MaKhuyenMai
//         WHERE gh.MaNguoiDung = ?
//     `

// 	err := r.db.Raw(query, manguoidung).Scan(&results).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Convert sang []models.GioHang
// 	var gioHangs []models.GioHang
// 	for _, result := range results {
// 		gioHangs = append(gioHangs, models.GioHang{
// 			MaNguoiDung: result.MaNguoiDung,
// 			MaBienThe:   result.MaBienThe,
// 			SoLuong:     result.SoLuong,
// 			Gia:         result.Gia,
// 		})
// 	}

// 	return gioHangs, nil
// }
