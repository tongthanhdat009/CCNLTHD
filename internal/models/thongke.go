package models

// Chuẩn time filter: dùng query `from` & `to` dạng YYYY-MM-DD.
// Ở repo mình sẽ chuyển thành [from 00:00:00, to 23:59:59].
type DateRange struct {
	From  string `form:"from" json:"from"`   // YYYY-MM-DD (optional)
	To    string `form:"to"   json:"to"`     // YYYY-MM-DD (optional)
	Limit int    `form:"limit" json:"limit"` // optional, default 10
}

// ====== Các DTO trả về ======

type TopCustomer struct {
	MaNguoiDung  int     `json:"maNguoiDung"`
	TenNguoiDung string  `json:"tenNguoiDung"`
	TongChi      float64 `json:"tongChi"`
	SoDon        int64   `json:"soDon"`
}

type PurchaseValue struct {
	// Tổng chi nhập trong khoảng
	Total float64 `json:"total"`
}

type ImportedProduct struct {
	MaHangHoa   int    `json:"maHangHoa"`
	TenHangHoa  string `json:"tenHangHoa"`
	SoLuongNhap int64  `json:"soLuongNhap"`
}

type ImportedBrand struct {
	MaHang  int    `json:"maHang"`
	TenHang string `json:"tenHang"`
	SoLuong int64  `json:"soLuong"`
}

type InvoiceStats struct {
	SoDonBan     int64   `json:"soDonBan"`
	TongDoanhThu float64 `json:"tongDoanhThu"`
}

type BestSeller struct {
	MaHangHoa  int     `json:"maHangHoa"`
	TenHangHoa string  `json:"tenHangHoa"`
	SoLuongBan int64   `json:"soLuongBan"`
	DoanhThu   float64 `json:"doanhThu"`
}

type RevenueByBrand struct {
	MaHang   int     `json:"maHang"`
	TenHang  string  `json:"tenHang"`
	DoanhThu float64 `json:"doanhThu"`
}
