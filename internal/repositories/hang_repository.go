package repositories

import (
    "github.com/tongthanhdat009/CCNLTHD/internal/models"
    "gorm.io/gorm"
)

type HangRepository interface {
    GetAll() ([]models.Hang, error)
    DeleteHang(id int) error
    CreateHang(hang *models.Hang) error
    UpdateHang(hang *models.Hang) error
    GetHangByID(id int) (*models.Hang, error)
    GetHangByName(name string) ([]models.Hang, error)  // Thay đổi thành slice
}

type HangRepo struct {
    db *gorm.DB
}

func NewHangRepository(db *gorm.DB) HangRepository {
    return &HangRepo{db: db}
}

func (r *HangRepo) GetAll() ([]models.Hang, error) {
    var Hangs []models.Hang
    err := r.db.Find(&Hangs).Error
    return Hangs, err
}

func (r *HangRepo) DeleteHang(id int) error {
    return r.db.Delete(&models.Hang{}, id).Error
}

func (r *HangRepo) CreateHang(hang *models.Hang) error {
    return r.db.Create(hang).Error
}

func (r *HangRepo) UpdateHang(hang *models.Hang) error {
    return r.db.Save(hang).Error
}

func (r *HangRepo) GetHangByID(id int) (*models.Hang, error) {
    var hang models.Hang
    err := r.db.First(&hang, id).Error
    if err != nil {
        return nil, err
    }
    return &hang, nil
}

func (r *HangRepo) GetHangByName(name string) ([]models.Hang, error) {
    var hangs []models.Hang
    err := r.db.Where("tenhang LIKE ?", "%"+name+"%").Find(&hangs).Error  // Dùng Find để lấy nhiều
    if err != nil {
        return nil, err
    }
    return hangs, nil
}