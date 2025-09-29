package models

import "database/sql"

type HangHoa struct {
    MaHangHoa   int           `gorm:"primaryKey;column:MaHangHoa" json:"ma_hang_hoa"`
    TenHangHoa  string        `gorm:"column:TenHangHoa" json:"ten_hang_hoa"`
    MaHang      int           `gorm:"column:MaHang" json:"-"`
    MaDanhMuc   int           `gorm:"column:MaDanhMuc" json:"-"`
    Mau         string        `gorm:"column:Mau" json:"mau"`
    MoTa        string        `gorm:"column:MoTa" json:"mo_ta"`
    TrangThai   string        `gorm:"column:TrangThai" json:"trang_thai"`
    MaKhuyenMai sql.NullInt64 `gorm:"column:MaKhuyenMai" json:"-"`
    AnhDaiDien  string        `gorm:"column:AnhDaiDien" json:"anh_dai_dien"`

    // --- Định nghĩa các mối quan hệ ---
    Hang      Hang      `gorm:"foreignKey:MaHang" json:"hang,omitempty"`
    DanhMuc   DanhMuc   `gorm:"foreignKey:MaDanhMuc" json:"danh_muc,omitempty"`
    KhuyenMai KhuyenMai `gorm:"foreignKey:MaKhuyenMai" json:"khuyen_mai,omitempty"`
    BienThes  []BienThe `gorm:"foreignKey:MaHangHoa" json:"bien_thes,omitempty"`
}

type BienThe struct {
    MaBienThe  int     `gorm:"primaryKey;column:MaBienThe" json:"ma_bien_the"`
    MaHangHoa  int     `gorm:"column:MaHangHoa" json:"ma_hang_hoa"`
    Size       string  `gorm:"column:Size" json:"size"`
    Gia        float64 `gorm:"column:Gia" json:"gia"`
    SoLuongTon int     `gorm:"column:SoLuongTon" json:"so_luong_ton"`
    TrangThai  string  `gorm:"column:TrangThai" json:"trang_thai"`

    HangHoa HangHoa `gorm:"foreignKey:MaHangHoa" json:"hang_hoa,omitempty"`
}

type DanhMuc struct {
    MaDanhMuc  int    `gorm:"primaryKey;column:MaDanhMuc" json:"ma_danh_muc"`
    TenDanhMuc string `gorm:"column:TenDanhMuc" json:"ten_danh_muc"`
}

type Hang struct {
    MaHang  int    `gorm:"primaryKey;column:MaHang" json:"ma_hang"`
    TenHang string `gorm:"column:TenHang" json:"ten_hang"`
}

type KhuyenMai struct {
    MaKhuyenMai  int     `gorm:"primaryKey;column:MaKhuyenMai" json:"ma_khuyen_mai"`
    TenKhuyenMai string  `gorm:"column:TenKhuyenMai" json:"ten_khuyen_mai"`
    MoTa         string  `gorm:"column:MoTa" json:"mo_ta"`
    GiaTri       float64 `gorm:"column:GiaTri" json:"gia_tri"`
}

// --- Cung cấp tên bảng cho GORM ---
func (HangHoa) TableName() string   { return "HangHoa" }
func (BienThe) TableName() string   { return "BienThe" }
func (DanhMuc) TableName() string   { return "DanhMuc" }
func (Hang) TableName() string      { return "Hang" }
func (KhuyenMai) TableName() string { return "KhuyenMai" }