package models

import "time"

type DonHang struct {
	MaDonHang          int       `json:"ma_don_hang" db:"MaDonHang"`
	MaNguoiDung        int       `json:"ma_nguoi_dung" db:"MaNguoiDung"`
	NgayTao            time.Time `json:"ngay_tao" db:"NgayTao"`
	TrangThai          string    `json:"trang_thai" db:"TrangThai"`
	TongTien           float64   `json:"tong_tien" db:"TongTien"`
	TinhThanh          string    `json:"tinh_thanh" db:"TinhThanh"`
	QuanHuyen          string    `json:"quan_huyen" db:"QuanHuyen"`
	PhuongXa           string    `json:"phuong_xa" db:"PhuongXa"`
	DuongSoNha         string    `json:"duong_so_nha" db:"DuongSoNha"`
	PhuongThucThanhToan string    `json:"phuong_thuc_thanh_toan" db:"PhuongThucThanhToan"`
}

type ChiTietDonHang struct {
	MaChiTiet int     `json:"ma_chi_tiet" db:"MaChiTiet"`
	MaDonHang int     `json:"ma_don_hang" db:"MaDonHang"`
	MaSanPham int     `json:"ma_san_pham" db:"MaSanPham"`
	GiaBan    float64 `json:"gia_ban" db:"GiaBan"`
}

type GioHang struct {
	MaNguoiDung int `json:"ma_nguoi_dung" db:"MaNguoiDung"`
	MaBienThe   int `json:"ma_bien_the" db:"MaBienThe"`
	SoLuong     int `json:"so_luong" db:"SoLuong"`
}