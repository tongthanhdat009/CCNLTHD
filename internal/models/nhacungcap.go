package models

type NhaCungCap struct {
    MaNCC       int    `gorm:"primaryKey;column:MaNCC" json:"ma_ncc"`
    TenNCC      string `gorm:"column:TenNCC" json:"ten_ncc"`
    DiaChi      string `gorm:"column:DiaChi" json:"dia_chi"`
    SoDienThoai string `gorm:"column:SoDienThoai" json:"so_dien_thoai"`
    Email       string `gorm:"column:Email" json:"email"`
}

// --- Cung cấp tên bảng cho GORM ---
func (NhaCungCap) TableName() string {
    return "NhaCungCap"
}