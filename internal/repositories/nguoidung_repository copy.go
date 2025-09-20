package repositories

import (
    "github.com/tongthanhdat009/CCNLTHD/internal/models"
    "gorm.io/gorm"
)

type NguoiDungRepository interface {
    GetAll() ([]models.NguoiDung, error)
}

type NguoiDungRepo struct {
    db *gorm.DB
}

func NewNguoiDungRepository(db *gorm.DB) NguoiDungRepository {
    return &NguoiDungRepo{db: db}
}

func (r *NguoiDungRepo) GetAll() ([]models.NguoiDung, error) {
    var NguoiDungs []models.NguoiDung
    // Sử dụng Preload để tải các dữ liệu liên quan
    err := r.db.Preload("Quyen").Find(&NguoiDungs).Error
    return NguoiDungs, err
}