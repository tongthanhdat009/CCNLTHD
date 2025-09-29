package models

import "time"

type DonHang struct {
    MaDonHang          int       `gorm:"primaryKey;column:MaDonHang" json:"ma_don_hang"`
    MaNguoiDung        int       `gorm:"column:MaNguoiDung" json:"ma_nguoi_dung"`
    NgayTao            time.Time `gorm:"column:NgayTao;autoCreateTime" json:"ngay_tao"`
    TrangThai          string    `gorm:"column:TrangThai" json:"trang_thai"`
    TongTien           float64   `gorm:"column:TongTien" json:"tong_tien"`
    TinhThanh          string    `gorm:"column:TinhThanh" json:"tinh_thanh"`
    QuanHuyen          string    `gorm:"column:QuanHuyen" json:"quan_huyen"`
    PhuongXa           string    `gorm:"column:PhuongXa" json:"phuong_xa"`
    DuongSoNha         string    `gorm:"column:DuongSoNha" json:"duong_so_nha"`
    PhuongThucThanhToan string    `gorm:"column:PhuongThucThanhToan" json:"phuong_thuc_thanh_toan"`

    // --- Mối quan hệ One-to-Many ---
    // Một đơn hàng có nhiều chi tiết đơn hàng
    ChiTietDonHangs []ChiTietDonHang `gorm:"foreignKey:MaDonHang" json:"chi_tiet_don_hangs,omitempty"`
}

type ChiTietDonHang struct {
    MaChiTiet int     `gorm:"primaryKey;column:MaChiTiet" json:"ma_chi_tiet"`
    MaDonHang int     `gorm:"column:MaDonHang" json:"ma_don_hang"`
    MaSanPham int     `gorm:"column:MaSanPham" json:"ma_san_pham"`
    SoLuong   int     `gorm:"column:SoLuong" json:"so_luong"`
    GiaBan    float64 `gorm:"column:GiaBan" json:"gia_ban"`

    // --- Mối quan hệ Many-to-One ---
    // Chi tiết này thuộc về biến thể sản phẩm nào
    DonHang DonHang `gorm:"foreignKey:MaDonHang" json:"don_hang,omitempty"`
}

type GioHang struct {
    MaNguoiDung int `gorm:"primaryKey;column:MaNguoiDung" json:"ma_nguoi_dung"`
    MaBienThe   int `gorm:"primaryKey;column:MaBienThe" json:"ma_bien_the"`
    SoLuong     int `gorm:"column:SoLuong" json:"so_luong"`
}

// --- Cung cấp tên bảng cho GORM ---
func (DonHang) TableName() string        { return "DonHang" }
func (ChiTietDonHang) TableName() string { return "ChiTietDonHang" }
func (GioHang) TableName() string        { return "GioHang" }