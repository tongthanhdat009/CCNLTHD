package repositories

import (
	"errors"
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
	SearchKhuyenMai(keyword string) ([]models.KhuyenMai, error)
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
	return r.db.Model(&models.KhuyenMai{}).Where("MaKhuyenMai = ?", makhuyenmai).Updates(khuyenMai).Error
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

func (r *KhuyenMaiRepo) SearchKhuyenMai(keyword string) ([]models.KhuyenMai, error) {
    var khuyenMais []models.KhuyenMai
    searchPattern := "%" + keyword + "%"
    
    err := r.db.Where("TenKhuyenMai LIKE ?", searchPattern).
        Or("MaKhuyenMai = ?", keyword). // Nếu keyword là số, tìm theo mã
        Find(&khuyenMais).Error
    
    return khuyenMais, err
}