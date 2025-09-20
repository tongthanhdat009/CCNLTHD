package repositories

import (
    "github.com/tongthanhdat009/CCNLTHD/internal/models"
    "gorm.io/gorm"
)

type HangHoaRepository interface {
    GetAll() ([]models.HangHoa, error)
}

type hangHoaRepo struct {
    db *gorm.DB
}

func NewHangHoaRepository(db *gorm.DB) HangHoaRepository {
    return &hangHoaRepo{db: db}
}

func (r *hangHoaRepo) GetAll() ([]models.HangHoa, error) {
    var hangHoas []models.HangHoa
    // Sử dụng Preload để tải các dữ liệu liên quan
    err := r.db.Preload("Hang").Preload("DanhMuc").Preload("BienThes").Find(&hangHoas).Error
    return hangHoas, err
}