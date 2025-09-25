package models

import "time"

type RefreshToken struct {
    MaToken     int       `gorm:"primaryKey;column:MaToken" json:"ma_token"`
    MaNguoiDung int       `gorm:"column:MaNguoiDung" json:"ma_nguoi_dung"`
    Token       string    `gorm:"column:Token" json:"token"`
    NgayTao     time.Time `gorm:"column:NgayTao;autoCreateTime" json:"ngay_tao"`
    NgayHetHan  time.Time `gorm:"column:NgayHetHan" json:"ngay_het_han"`
    TrangThai   string    `gorm:"column:TrangThai" json:"trang_thai"`

    // --- Mối quan hệ Many-to-One ---
    NguoiDung NguoiDung `gorm:"foreignKey:MaNguoiDung" json:"nguoi_dung,omitempty"`
}

func (RefreshToken) TableName() string {
    return "refreshtoken" 
}