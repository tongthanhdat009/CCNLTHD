package repositories

import (
    "gorm.io/gorm"
)

type AuthRepository interface {
    KiemTraQuyen(maQuyen int, tenChucNang string, tenHanhDong string) (bool, error) 
}

type AuthRepo struct {
    db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
    return &AuthRepo{db: db}
}

func (r *AuthRepo) KiemTraQuyen(maQuyen int, tenChucNang string, tenHanhDong string) (bool, error) {
    var count int64

    err := r.db.
        Table("phanquyen").
        Joins("JOIN chitietchucnang ON chitietchucnang.MaChiTietChucNang = phanquyen.MaChiTietChucNang").
        Joins("JOIN chucnang ON chitietchucnang.MaChucNang = chucnang.MaChucNang").
        Where("phanquyen.MaQuyen = ? AND phanquyen.TrangThai = ? AND chucnang.TenChucNang = ? AND chitietchucnang.TenChiTietChucNang = ?", maQuyen, "Mở", tenChucNang, tenHanhDong).
        Count(&count).Error

    if err != nil {
        return false, err
    }

    return count > 0, nil
}
