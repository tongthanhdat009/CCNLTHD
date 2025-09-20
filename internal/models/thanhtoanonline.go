package models

import "time"

// ThanhToanOnline model
type ThanhToanOnline struct {
    MaThanhToan   int       `gorm:"primaryKey;column:MaThanhToan" json:"ma_thanh_toan"`
    MaDonHang     int       `gorm:"column:MaDonHang" json:"-"`
    MaGiaoDich    string    `gorm:"column:MaGiaoDich" json:"ma_giao_dich"`
    TrangThai     string    `gorm:"column:TrangThai" json:"trang_thai"`
    NgayThanhToan time.Time `gorm:"column:NgayThanhToan" json:"ngay_thanh_toan"`
    TongTien      float64   `gorm:"column:TongTien" json:"tong_tien"`
    MoTaGiaoDich  string    `gorm:"column:MoTaGiaoDich" json:"mo_ta_giao_dich"`

    // Mối quan hệ
    DonHang DonHang `gorm:"foreignKey:MaDonHang" json:"don_hang,omitempty"`
}

// --- Cung cấp tên bảng cho GORM ---
func (ThanhToanOnline) TableName() string { return "thanhtoanonline" }