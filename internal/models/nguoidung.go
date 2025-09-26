package models

import (
    "database/sql"
    "time"
)

type NguoiDung struct {
    MaNguoiDung int       `gorm:"primaryKey;column:MaNguoiDung" json:"ma_nguoi_dung"`
    TenDangNhap string    `gorm:"column:TenDangNhap" json:"ten_dang_nhap"`
    MatKhau     string    `gorm:"column:MatKhau" json:"-"`
    HoTen       string    `gorm:"column:HoTen" json:"ho_ten"`
    Email       string    `gorm:"column:Email" json:"email"`
    SoDienThoai string    `gorm:"column:SoDienThoai" json:"so_dien_thoai"`
    TinhThanh   sql.NullString `gorm:"column:TinhThanh" json:"tinh_thanh"`
    QuanHuyen   sql.NullString `gorm:"column:QuanHuyen" json:"quan_huyen"`
    PhuongXa    sql.NullString `gorm:"column:PhuongXa" json:"phuong_xa"`
    DuongSoNha  sql.NullString `gorm:"column:DuongSoNha" json:"duong_so_nha"`
    MaQuyen     int       `gorm:"column:MaQuyen" json:"ma_quyen"` // Ẩn đi để dùng struct Quyen bên dưới
    NgayTao     time.Time `gorm:"column:NgayTao;autoCreateTime" json:"ngay_tao"`
    NgayCapNhat time.Time `gorm:"column:NgayCapNhat;autoUpdateTime" json:"ngay_cap_nhat"`

    // --- Mối quan hệ Many-to-One ---
    // Người dùng này thuộc về quyền nào
    Quyen Quyen `gorm:"foreignKey:MaQuyen;references:MaQuyen" json:"quyen,omitempty"`
}

type Quyen struct {
    MaQuyen    int           `gorm:"primaryKey;column:MaQuyen" json:"ma_quyen"`
    TenQuyen   string        `gorm:"column:TenQuyen" json:"ten_quyen"`
    PhanQuyens []PhanQuyen   `gorm:"foreignKey:MaQuyen;references:MaQuyen" json:"phan_quyens,omitempty"`
    ChucNangs  []ChucNang    `gorm:"-" json:"chuc_nangs,omitempty"` // Không ánh xạ trực tiếp, sẽ lấy qua bảng PhanQuyen
}


// --- Cung cấp tên bảng cho GORM ---
func (NguoiDung) TableName() string {
    return "NguoiDung"
}

func (Quyen) TableName() string {
    return "Quyen"
}