package repositories

import (
    "github.com/tongthanhdat009/CCNLTHD/internal/models"
    "gorm.io/gorm"
)

type PhanQuyenRepository interface {
	GetAll() ([]models.Quyen, error)
	GetByID(id int) (*models.Quyen, error)
	GetChucNangVaChiTiet(idQuyen int) ([]models.ChucNang, error)
	UpdatePhanQuyen(phanquyen *models.PhanQuyen) error
}

type PhanQuyenRepo struct {
    db *gorm.DB
}

func NewPhanQuyenRepository(db *gorm.DB) PhanQuyenRepository {
    return &PhanQuyenRepo{db: db}
}

func (r *PhanQuyenRepo) GetAll() ([]models.Quyen, error) {
	var Quyens []models.Quyen
	if err := r.db.Find(&Quyens).Error; err != nil {
		return nil, err
	}
	return Quyens, nil
}

func (r *PhanQuyenRepo) GetByID(id int) (*models.Quyen, error) {
	var quyen models.Quyen
	if err := r.db.First(&quyen, "MaQuyen = ?", id).Error; err != nil {
		return nil, err
	}
	return &quyen, nil
}

func (r *PhanQuyenRepo) GetChucNangVaChiTiet(idQuyen int) ([]models.ChucNang, error) {
    var chucNangs []models.ChucNang

    err := r.db.
        Model(&models.ChucNang{}).
        Joins("JOIN ChiTietChucNang ON ChiTietChucNang.MaChucNang = ChucNang.MaChucNang").
        Joins("JOIN PhanQuyen ON PhanQuyen.MaChiTietChucNang = ChiTietChucNang.MaChiTietChucNang").
        Where("PhanQuyen.MaQuyen = ?", idQuyen).
        Preload("ChiTietChucNangs", func(db *gorm.DB) *gorm.DB {
            // preload ChiTietChucNang chỉ thuộc quyền này
            return db.
                Joins("JOIN PhanQuyen ON PhanQuyen.MaChiTietChucNang = ChiTietChucNang.MaChiTietChucNang").
                Where("PhanQuyen.MaQuyen = ?", idQuyen).
                Preload("PhanQuyens", "MaQuyen = ?", idQuyen)
        }).
        Group("ChucNang.MaChucNang").
        Find(&chucNangs).Error

    if err != nil {
        return nil, err
    }

    return chucNangs, nil
}
func (r *PhanQuyenRepo) UpdatePhanQuyen(phanquyen *models.PhanQuyen) error {
	return r.db.Save(phanquyen).Error
}