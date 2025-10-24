package models

import "time"

type Review struct {
	MaDanhGia   int       `gorm:"column:MaDanhGia;primaryKey;autoIncrement" json:"maDanhGia"`
	MaHangHoa   int       `gorm:"column:MaHangHoa" json:"maHangHoa"`
	MaNguoiDung int       `gorm:"column:MaNguoiDung" json:"maNguoiDung"`
	Diem        int       `gorm:"column:Diem" json:"diem"`
	NoiDung     string    `gorm:"column:NoiDung" json:"noiDung"`
	TrangThai   string    `gorm:"column:TrangThai" json:"trangThai"`
	NgayDanhGia time.Time `gorm:"column:NgayDanhGia" json:"ngayDanhGia"`
}

func (Review) TableName() string { return "danhgia" }

// DTO KH
type CreateReviewDTO struct {
	MaHangHoa int    `json:"maHangHoa" form:"maHangHoa" binding:"required,min=1"`
	Diem      int    `json:"diem"      form:"diem"      binding:"required,min=1,max=5"`
	NoiDung   string `json:"noiDung"   form:"noiDung"   binding:"omitempty,max=255"`
}

func (CreateReviewDTO) Now() time.Time { return time.Now() } // dùng để set NgayDanhGia

// hằng trạng thái
const (
	ReviewStatusPending  = "Chưa duyệt"
	ReviewStatusApproved = "Đã duyệt"
	ReviewStatusRejected = "Đã từ chối"
	ReviewStatusHidden   = "Ẩn"
)
