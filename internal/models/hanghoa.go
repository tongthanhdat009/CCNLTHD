package models

import "database/sql"

type HangHoa struct {
	MaHangHoa   int    `json:"ma_hang_hoa" db:"MaHangHoa"`
	TenHangHoa  string `json:"ten_hang_hoa" db:"TenHangHoa"`
	MaHang      int    `json:"ma_hang" db:"MaHang"`
	MaDanhMuc   int    `json:"ma_danh_muc" db:"MaDanhMuc"`
	Mau         string `json:"mau" db:"Mau"`
	MoTa        string `json:"mo_ta" db:"MoTa"`
	TrangThai   string `json:"trang_thai" db:"TrangThai"`
	MaKhuyenMai sql.NullInt64    `json:"ma_khuyen_mai" db:"MaKhuyenMai"`
	AnhDaiDien  string `json:"anh_dai_dien" db:"AnhDaiDien"`
}

type BienThe struct {
	MaBienThe  int     `json:"ma_bien_the" db:"MaBienThe"`
	MaHangHoa  int     `json:"ma_hang_hoa" db:"MaHangHoa"`
	Size       string  `json:"size" db:"Size"`
	Gia        float64 `json:"gia" db:"Gia"`
	SoLuongTon int     `json:"so_luong_ton" db:"SoLuongTon"`
	TrangThai  string  `json:"trang_thai" db:"TrangThai"`
}

type DanhMuc struct {
	MaDanhMuc  int    `json:"ma_danh_muc" db:"MaDanhMuc"`
	TenDanhMuc string `json:"ten_danh_muc" db:"TenDanhMuc"`
}

type Hang struct {
	MaHang  int    `json:"ma_hang" db:"MaHang"`
	TenHang string `json:"ten_hang" db:"TenHang"`
}

type KhuyenMai struct {
	MaKhuyenMai  int     `json:"ma_khuyen_mai" db:"MaKhuyenMai"`
	TenKhuyenMai string  `json:"ten_khuyen_mai" db:"TenKhuyenMai"`
	MoTa         string  `json:"mo_ta" db:"MoTa"`
	GiaTri       float64 `json:"gia_tri" db:"GiaTri"`
}