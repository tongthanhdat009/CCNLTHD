package models

type AdminReviewFilter struct {
	MaHangHoa   *int    `form:"maHangHoa" json:"maHangHoa"`
	MaNguoiDung *int    `form:"maNguoiDung" json:"maNguoiDung"`
	TrangThai   *string `form:"trangThai"  json:"trangThai"`
	Q           *string `form:"q"          json:"q"`
	Page        int     `form:"page"       json:"page"`
	PageSize    int     `form:"pageSize"   json:"pageSize"`
}
