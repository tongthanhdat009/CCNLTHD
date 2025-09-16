package models

type NhaCungCap struct {
	MaNCC       int    `json:"ma_ncc" db:"MaNCC"`
	TenNCC      string `json:"ten_ncc" db:"TenNCC"`
	DiaChi      string `json:"dia_chi" db:"DiaChi"`
	SoDienThoai string `json:"so_dien_thoai" db:"SoDienThoai"`
	Email       string `json:"email" db:"Email"`
}