package repositories

import (
	"errors"
	"strings"

	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"gorm.io/gorm"
)

type KhuyenMaiRepository interface {
	CreateKhuyenMai(khuyenMai *models.KhuyenMai) error
	UpdateKhuyenMai(makhuyenmai int, khuyenMai models.KhuyenMai) error
	DeleteKhuyenMai(makhuyenmai int) error
	ExistsTenKhuyenMai(tenkhuyenmai string) (bool, error)
	GetAll() ([]models.KhuyenMai, error)
	GetByID(makhuyenmai int) (models.KhuyenMai, error)
	SearchKhuyenMai(keyword string, maKhuyenMai *int, tenKhuyenMai string, minGiaTri, maxGiaTri *float64, sortBy, sortOrder string, page, pageSize int) ([]models.KhuyenMai, int64, error)
}

type KhuyenMaiRepo struct {
	db *gorm.DB
}

func NewKhuyenMaiRepository(db *gorm.DB) KhuyenMaiRepository {
	return &KhuyenMaiRepo{db: db}
}

func (r *KhuyenMaiRepo) CreateKhuyenMai(khuyenMai *models.KhuyenMai) error {
	return r.db.Create(khuyenMai).Error
}

func (r *KhuyenMaiRepo) ExistsTenKhuyenMai(tenkhuyenmai string) (bool, error) {
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

func (r *KhuyenMaiRepo) UpdateKhuyenMai(makhuyenmai int, khuyenMai models.KhuyenMai) error {
	updates := make(map[string]interface{})

	// Sử dụng reflection hoặc kiểm tra từng field
	if khuyenMai.TenKhuyenMai != "" {
		updates["TenKhuyenMai"] = khuyenMai.TenKhuyenMai
	}

	if khuyenMai.MoTa != "" {
		updates["MoTa"] = khuyenMai.MoTa
	}

	// GiaTri luôn được update nếu khác 0 hoặc được chỉ định rõ ràng
	updates["GiaTri"] = khuyenMai.GiaTri

	if len(updates) == 0 {
		return errors.New("không có thông tin nào để cập nhật")
	}

	result := r.db.Model(&models.KhuyenMai{}).
		Where("MaKhuyenMai = ?", makhuyenmai).
		Updates(updates)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("không tìm thấy record để cập nhật")
	}

	return nil
}

func (r *KhuyenMaiRepo) DeleteKhuyenMai(makhuyenmai int) error {
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

func (r *KhuyenMaiRepo) SearchKhuyenMai(keyword string, maKhuyenMai *int, tenKhuyenMai string, minGiaTri, maxGiaTri *float64, sortBy, sortOrder string, page, pageSize int) ([]models.KhuyenMai, int64, error) {
    var khuyenMais []models.KhuyenMai
    var total int64

    // ✅ Build query cơ bản
    query := r.db.Model(&models.KhuyenMai{})

    // ✅ 1. Tìm kiếm theo keyword (tổng quát - tất cả các trường)
    if keyword != "" {
        searchPattern := "%" + keyword + "%"
        query = query.Where(
            r.db.Where("TenKhuyenMai LIKE ?", searchPattern).
                Or("MoTa LIKE ?", searchPattern).
                Or("MaKhuyenMai = ?", keyword),
        )
    }

    // ✅ 2. Lọc theo MaKhuyenMai cụ thể (exact match)
    if maKhuyenMai != nil {
        query = query.Where("MaKhuyenMai = ?", *maKhuyenMai)
    }

    // ✅ 3. Lọc theo TenKhuyenMai (LIKE - partial match)
    if tenKhuyenMai != "" {
        query = query.Where("TenKhuyenMai LIKE ?", "%"+tenKhuyenMai+"%")
    }

    // ✅ 4. Filter theo giá trị min
    if minGiaTri != nil {
        query = query.Where("GiaTri >= ?", *minGiaTri)
    }

    // ✅ 5. Filter theo giá trị max
    if maxGiaTri != nil {
        query = query.Where("GiaTri <= ?", *maxGiaTri)
    }

    // ✅ 6. Đếm tổng số kết quả (trước khi phân trang)
    if err := query.Count(&total).Error; err != nil {
        return nil, 0, err
    }

    // ✅ 7. Sắp xếp
    if sortBy == "" {
        sortBy = "MaKhuyenMai"
    }
    if sortOrder == "" {
        sortOrder = "DESC"
    }

    // Validate sortBy (chống SQL injection)
    validSortFields := map[string]bool{
        "MaKhuyenMai":  true,
        "TenKhuyenMai": true,
        "GiaTri":       true,
    }

    if !validSortFields[sortBy] {
        sortBy = "MaKhuyenMai"
    }

    sortOrder = strings.ToUpper(sortOrder)
    if sortOrder != "ASC" && sortOrder != "DESC" {
        sortOrder = "DESC"
    }

    query = query.Order(sortBy + " " + sortOrder)

    // ✅ 8. Phân trang
    if page < 1 {
        page = 1
    }
    if pageSize < 1 || pageSize > 100 {
        pageSize = 10
    }

    offset := (page - 1) * pageSize
    query = query.Limit(pageSize).Offset(offset)

    // ✅ 9. Thực thi query
    if err := query.Find(&khuyenMais).Error; err != nil {
        return nil, 0, err
    }

    return khuyenMais, total, nil
}