package models
import (
    "database/sql"
    "time"
)

// PhieuNhap model
type PhieuNhap struct {
    MaPhieuNhap int          `gorm:"primaryKey;column:MaPhieuNhap" json:"ma_phieu_nhap"`
    MaNguoiDung int          `gorm:"column:MaNguoiDung" json:"-"`
    MaNCC       int          `gorm:"column:MaNCC" json:"-"`
    NgayNhap    time.Time    `gorm:"column:NgayNhap;autoCreateTime" json:"ngay_nhap"`
    TrangThai   string       `gorm:"column:TrangThai" json:"trang_thai"`
    MoTa        sql.NullString `gorm:"column:MoTa" json:"mo_ta"`

    // Mối quan hệ
    NguoiDung        NguoiDung          `gorm:"foreignKey:MaNguoiDung" json:"nguoi_dung,omitempty"`
    NhaCungCap       NhaCungCap         `gorm:"foreignKey:MaNCC" json:"nha_cung_cap,omitempty"`
    ChiTietPhieuNhap []ChiTietPhieuNhap `gorm:"foreignKey:MaPhieuNhap" json:"chi_tiet_phieu_nhap,omitempty"`
}

// ChiTietPhieuNhap model
type ChiTietPhieuNhap struct {
    MaChiTiet       int           `gorm:"primaryKey;column:MaChiTiet" json:"ma_chi_tiet"`
    MaPhieuNhap     int           `gorm:"column:MaPhieuNhap" json:"-"`
    MaBienthe       int           `gorm:"column:MaBienthe" json:"-"`
    SoLuong         int           `gorm:"column:SoLuong" json:"so_luong"`
    GiaNhap         float64       `gorm:"column:GiaNhap" json:"gia_nhap"`
    NgaySanXuat     sql.NullTime  `gorm:"column:NgaySanXuat" json:"ngay_san_xuat"`
    ThoiGianBaoHanh sql.NullInt64 `gorm:"column:ThoiGianBaoHanh" json:"thoi_gian_bao_hanh"`

    PhieuNhap PhieuNhap `gorm:"foreignKey:MaPhieuNhap" json:"phieu_nhap,omitempty"`
    // Mối quan hệ
    BienThe  BienThe   `gorm:"foreignKey:MaBienthe" json:"bien_the,omitempty"`
    SanPhams []SanPham `gorm:"foreignKey:MaChiTietPhieuNhap" json:"san_phams,omitempty"`
}

// SanPham model (sản phẩm vật lý với số seri)
type SanPham struct {
    MaSanPham          int `gorm:"primaryKey;column:MaSanPham" json:"ma_san_pham"`
    MaChiTietPhieuNhap int `gorm:"column:MaChiTietPhieuNhap" json:"-"`
    Seri               string `gorm:"column:Seri" json:"seri"`
    TrangThai          string `gorm:"column:TrangThai" json:"trang_thai"`

    ChiTietPhieuNhap ChiTietPhieuNhap `gorm:"foreignKey:MaChiTietPhieuNhap" json:"chi_tiet_phieu_nhap,omitempty"`
    ChiTietDonHangs []ChiTietDonHang `gorm:"foreignKey:MaSanPham" json:"chi_tiet_don_hangs,omitempty"`
}

// --- Cung cấp tên bảng cho GORM ---
func (PhieuNhap) TableName() string        { return "phieunhap" }
func (ChiTietPhieuNhap) TableName() string { return "chitietphieunhap" }
func (SanPham) TableName() string          { return "sanpham" }