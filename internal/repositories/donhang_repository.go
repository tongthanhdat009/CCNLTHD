package repositories

import (
    "github.com/tongthanhdat009/CCNLTHD/internal/models"
    "gorm.io/gorm"
)

type DonHangRepository interface {
    GetAll() ([]models.DonHang, error)
}

type DonHangRepo struct {
    db *gorm.DB
}

func NewDonHangRepository(db *gorm.DB) DonHangRepository {
    return &DonHangRepo{db: db}
}

func (r *DonHangRepo) GetAll() ([]models.DonHang, error) {
    var DonHangs []models.DonHang
    // Sử dụng Preload để tải các dữ liệu liên quan
    err := r.db.Preload("ChiTietDonHangs.BienThe").Find(&DonHangs).Error
    return DonHangs, err
}