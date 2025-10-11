package repositories

import (
    "github.com/tongthanhdat009/CCNLTHD/internal/models"
    "gorm.io/gorm"
)

type QuyenRepository interface {
	GetAll() ([]models.Quyen, error)
	GetByID(id int) (*models.Quyen, error)
	GetChucNangVaChiTiet(idQuyen int) ([]models.ChucNang, error)
}

type QuyenRepo struct {
    db *gorm.DB
}

func NewQuyenRepository(db *gorm.DB) QuyenRepository {
    return &QuyenRepo{db: db}
}

func (r *QuyenRepo) GetAll() ([]models.Quyen, error) {
	var Quyens []models.Quyen
	if err := r.db.Find(&Quyens).Error; err != nil {
		return nil, err
	}
	return Quyens, nil
}

func (r *QuyenRepo) GetByID(id int) (*models.Quyen, error) {
	var quyen models.Quyen
	if err := r.db.First(&quyen, "MaQuyen = ?", id).Error; err != nil {
		return nil, err
	}
	return &quyen, nil
}

func (r *QuyenRepo) GetChucNangVaChiTiet(idQuyen int) ([]models.ChucNang, error) {
    var chucNangs []models.ChucNang

    err := r.db.
        Model(&models.ChucNang{}).
        Joins("JOIN ChiTietChucNang ON ChiTietChucNang.MaChucNang = ChucNang.MaChucNang").
        Joins("JOIN PhanQuyen ON PhanQuyen.MaChiTietChucNang = ChiTietChucNang.MaChiTietChucNang").
        Where("PhanQuyen.MaQuyen = ?", idQuyen).
        Preload("ChiTietChucNangs", func(db *gorm.DB) *gorm.DB {
            // Chỉ load chi tiết nào quyền này có
            return db.Joins("JOIN PhanQuyen ON PhanQuyen.MaChiTietChucNang = ChiTietChucNang.MaChiTietChucNang").
                Where("PhanQuyen.MaQuyen = ?", idQuyen)
        }).
        Group("ChucNang.MaChucNang").
        Find(&chucNangs).Error

    if err != nil {
        return nil, err
    }

    return chucNangs, nil
}



