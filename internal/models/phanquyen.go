package models

// MaChucNang model
type MaChucNang struct {
    MaChucNang int    `gorm:"primaryKey;column:MaChucNang" json:"ma_chuc_nang"`
    TenChucNang string `gorm:"column:TenChucNang" json:"ten_chuc_nang"`

    // Mối quan hệ
    ChiTietChucNangs []ChiTietChucNang `gorm:"foreignKey:MaChucNang" json:"chi_tiet_chuc_nangs,omitempty"`
}

// ChiTietChucNang model
type ChiTietChucNang struct {
    MaChiTietChucNang int    `gorm:"primaryKey;column:MaChiTietChucNang" json:"ma_chi_tiet_chuc_nang"`
    MaChucNang        int    `gorm:"column:MaChucNang" json:"-"`
    TenChiTietChucNang string `gorm:"column:TenChiTietChucNang" json:"ten_chi_tiet_chuc_nang"`

    // Mối quan hệ nhiều-nhiều với Quyen
    Quyens []*Quyen `gorm:"many2many:phanquyen;foreigsnKey:MaChiTietChucNang;joinForeignKey:MaChiTietChucNang;References:MaQuyen;joinReferences:MaQuyen" json:"quyens,omitempty"`
}

// --- Cung cấp tên bảng cho GORM ---
func (MaChucNang) TableName() string      { return "machucnang" }
func (ChiTietChucNang) TableName() string { return "chitietchucnang" }
// Bảng phanquyen được GORM quản lý tự động qua gorm:"many2many:phanquyen"