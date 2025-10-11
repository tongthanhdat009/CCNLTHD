package repositories

import (
	"errors"

	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"gorm.io/gorm"
)

type TimKiemHangHoaRepository interface {
	TimHangHoa(
		tenHangHoa string,
		tenHang string,
		tenDanhMuc string,
		mau string,
		size string,
		giaToiThieu *float64,
		giaToiDa *float64,
	) ([]models.HangHoa, error)
}

type timKiemHangHoaRepository struct {
	db *gorm.DB
}

func NewTimKiemHangHoaRepository(db *gorm.DB) TimKiemHangHoaRepository {
	return &timKiemHangHoaRepository{db: db}
}

func (r *timKiemHangHoaRepository) TimHangHoa(
	tenHangHoa string,
	tenHang string,
	tenDanhMuc string,
	mau string,
	size string,
	giaToiThieu *float64,
	giaToiDa *float64,
) ([]models.HangHoa, error) {

	var hangHoas []models.HangHoa

	// --- BƯỚC 1: TẠO SUBQUERY ĐỂ LẤY DANH SÁCH ID HÀNG HÓA HỢP LỆ ---
	// Query từ BienThe và join ngược lên để áp dụng tất cả bộ lọc.
	subQuery := r.db.Model(&models.BienThe{}).
		Select("BienThe.MaHangHoa"). // Chỉ lấy cột MaHangHoa
		Joins("JOIN HangHoa ON BienThe.MaHangHoa = HangHoa.MaHangHoa").
		Joins("JOIN Hang ON HangHoa.MaHang = Hang.MaHang").
		Joins("JOIN DanhMuc ON HangHoa.MaDanhMuc = DanhMuc.MaDanhMuc").
		Where("BienThe.TrangThai = ?", "DangBan") // Điều kiện cơ bản

	// Áp dụng tất cả các bộ lọc vào subquery để tìm ID
	if tenHangHoa != "" {
		subQuery = subQuery.Where("HangHoa.TenHangHoa LIKE CONCAT('%', ?, '%') COLLATE utf8mb4_unicode_ci", tenHangHoa)
	}
	if tenHang != "" {
		subQuery = subQuery.Where("Hang.TenHang LIKE CONCAT('%', ?, '%') COLLATE utf8mb4_unicode_ci", tenHang)
	}
	if tenDanhMuc != "" {
		subQuery = subQuery.Where("DanhMuc.TenDanhMuc LIKE CONCAT('%', ?, '%') COLLATE utf8mb4_unicode_ci", tenDanhMuc)
	}
	if mau != "" {
		subQuery = subQuery.Where("HangHoa.Mau LIKE CONCAT('%', ?, '%') COLLATE utf8mb4_unicode_ci", mau)
	}
	if size != "" {
		subQuery = subQuery.Where("BienThe.Size = ?", size)
	}
	if giaToiThieu != nil && giaToiDa != nil {
		subQuery = subQuery.Where("BienThe.Gia BETWEEN ? AND ?", *giaToiThieu, *giaToiDa)
	} else if giaToiThieu != nil {
		subQuery = subQuery.Where("BienThe.Gia >= ?", *giaToiThieu)
	} else if giaToiDa != nil {
		subQuery = subQuery.Where("BienThe.Gia <= ?", *giaToiDa)
	}

	// --- BƯỚC 2 & 3: QUERY CHÍNH - LẤY HANGHOA DỰA TRÊN SUBQUERY VÀ PRELOAD DỮ LIỆU ---
	query := r.db.Model(&models.HangHoa{}).
		Where("HangHoa.MaHangHoa IN (?)", subQuery). // Lọc theo danh sách ID từ subquery
		Preload("BienThes", func(db *gorm.DB) *gorm.DB {
			// Áp dụng các điều kiện lọc cho BienThe
			db = db.Where("TrangThai = ?", "DangBan")
			if size != "" {
				db = db.Where("Size = ?", size)
			}
			if giaToiThieu != nil && giaToiDa != nil {
				db = db.Where("Gia BETWEEN ? AND ?", *giaToiThieu, *giaToiDa)
			} else if giaToiThieu != nil {
				db = db.Where("Gia >= ?", *giaToiThieu)
			} else if giaToiDa != nil {
				db = db.Where("Gia <= ?", *giaToiDa)
			}
			return db
		}).
		Preload("Hang").
		Preload("DanhMuc")

	if err := query.Find(&hangHoas).Error; err != nil {
		return nil, errors.New("Lỗi khi tìm kiếm hàng hóa: " + err.Error())
	}

	return hangHoas, nil
}