package models

import "time"

// ChucNang model
type ChucNang struct {
    MaChucNang       int                 `gorm:"primaryKey;column:MaChucNang" json:"ma_chuc_nang"`
    TenChucNang      string              `gorm:"column:TenChucNang" json:"ten_chuc_nang"`
    ChiTietChucNangs []ChiTietChucNang   `gorm:"foreignKey:MaChucNang" json:"chi_tiet_chuc_nangs,omitempty"`
}

// ChiTietChucNang model
type ChiTietChucNang struct {
    MaChiTietChucNang int          `gorm:"primaryKey;column:MaChiTietChucNang" json:"ma_chi_tiet_chuc_nang"`
    MaChucNang        int          `gorm:"column:MaChucNang" json:"-"`
    TenChiTietChucNang string      `gorm:"column:TenChiTietChucNang" json:"ten_chi_tiet_chuc_nang"`
}

// PhanQuyen model (bảng trung gian)
type PhanQuyen struct {
    ID                 int                 `gorm:"primaryKey;autoIncrement" json:"id"`
    MaQuyen            int                 `gorm:"column:MaQuyen" json:"ma_quyen"`
    MaChiTietChucNang int                 `gorm:"column:MaChiTietChucNang" json:"ma_chi_tiet_chuc_nang"`
    NgayTao            time.Time           `gorm:"column:NgayTao;autoCreateTime" json:"ngay_tao"`
    TrangThai         string              `gorm:"column:TrangThai" json:"trang_thai"`
}

// --- Cung cấp tên bảng cho GORM ---
func (ChucNang) TableName() string { return "chucnang" }
func (ChiTietChucNang) TableName() string { return "chitietchucnang" }
func (PhanQuyen) TableName() string { return "phanquyen" }
