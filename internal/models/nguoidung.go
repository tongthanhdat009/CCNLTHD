package models

import "time"

type NguoiDung struct {
	MaNguoiDung int       `json:"ma_nguoi_dung" db:"MaNguoiDung"`
	TenDangNhap string    `json:"ten_dang_nhap" db:"TenDangNhap"`
	MatKhau     string    `json:"mat_khau" db:"MatKhau"`
	HoTen       string    `json:"ho_ten" db:"HoTen"`
	Email       string    `json:"email" db:"Email"`
	SoDienThoai string    `json:"so_dien_thoai" db:"SoDienThoai"`
	TinhThanh   string    `json:"tinh_thanh" db:"TinhThanh"`
	QuanHuyen   string    `json:"quan_huyen" db:"QuanHuyen"`
	PhuongXa    string    `json:"phuong_xa" db:"PhuongXa"`
	DuongSoNha  string    `json:"duong_so_nha" db:"DuongSoNha"`
	MaQuyen     int       `json:"ma_quyen" db:"MaQuyen"`
	NgayTao     time.Time `json:"ngay_tao" db:"NgayTao"`
	NgayCapNhat time.Time `json:"ngay_cap_nhat" db:"NgayCapNhat"`
}

// Quyen represents the 'quyen' (role) table
type Quyen struct {
	MaQuyen  int    `json:"ma_quyen" db:"MaQuyen"`
	TenQuyen string `json:"ten_quyen" db:"TenQuyen"`
}