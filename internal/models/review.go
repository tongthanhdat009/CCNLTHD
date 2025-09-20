package models

import (
    "database/sql"
    "time"
)

type DanhGia struct {
    MaDanhGia   int          `gorm:"primaryKey;column:MaDanhGia" json:"ma_danh_gia"`
    MaSanPham   int          `gorm:"column:MaSanPham" json:"-"`
    MaNguoiDung int          `gorm:"column:MaNguoiDung" json:"-"`
    Diem        int          `gorm:"column:Diem" json:"diem"`
    NoiDung     sql.NullString `gorm:"column:NoiDung" json:"noi_dung"`
    TrangThai   string       `gorm:"column:TrangThai" json:"trang_thai"`
    NgayDanhGia time.Time    `gorm:"column:NgayDanhGia;autoCreateTime" json:"ngay_danh_gia"`

    // --- Mối quan hệ Many-to-One ---
    NguoiDung NguoiDung `gorm:"foreignKey:MaNguoiDung" json:"nguoi_dung,omitempty"`
    HangHoa   HangHoa   `gorm:"foreignKey:MaSanPham" json:"hang_hoa,omitempty"`
}
// --- Cung cấp tên bảng cho GORM ---
func (DanhGia) TableName() string {
    return "DanhGia"
}