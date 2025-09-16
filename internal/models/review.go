// models/review.go
package models

import "time"

// DanhGia represents the 'danhgia' (review) table
type DanhGia struct {
	MaDanhGia   int       `json:"ma_danh_gia" db:"MaDanhGia"`
	MaSanPham   int       `json:"ma_san_pham" db:"MaSanPham"`
	MaNguoiDung int       `json:"ma_nguoi_dung" db:"MaNguoiDung"`
	Diem        int       `json:"diem" db:"Diem"`
	NoiDung     string    `json:"noi_dung" db:"NoiDung"`
	TrangThai   string    `json:"trang_thai" db:"TrangThai"`
	NgayDanhGia time.Time `json:"ngay_danh_gia" db:"NgayDanhGia"`
}