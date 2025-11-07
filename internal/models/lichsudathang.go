package models

type OrderHistoryFilter struct {
	Status   string `form:"status"`   // optional: Đang xử lý, Đang giao, Đã giao, Hoàn tất, Đã hủy, ...
	Page     int    `form:"page"`     // default 1
	PageSize int    `form:"pageSize"` // default 10 (<=100)
}

type OrderSummary struct {
	MaDonHang int     `json:"maDonHang"`
	NgayTao   string  `json:"ngayTao"` // ISO8601
	TongTien  float64 `json:"tongTien"`
	TrangThai string  `json:"trangThai"`
}

type OrderItem struct {
	MaSanPham  int     `json:"maSanPham"`
	Seri       string  `json:"seri"`
	TenHangHoa string  `json:"tenHangHoa"`
	GiaBan     float64 `json:"giaBan"`
}

type OrderDetail struct {
	MaDonHang int     `json:"maDonHang"`
	NgayTao   string  `json:"ngayTao"`
	TongTien  float64 `json:"tongTien"`
	TrangThai string  `json:"trangThai"`

	// Địa chỉ giao hàng (từ donhang)
	TinhThanh  string `json:"tinhThanh"`
	QuanHuyen  string `json:"quanHuyen"`
	PhuongXa   string `json:"phuongXa"`
	DuongSoNha string `json:"duongSoNha"`
	SDT        string `json:"soDienThoai"`
	ThanhToan  string `json:"phuongThucThanhToan"`

	Items []OrderItem `json:"items"`
}

type OrderListResult struct {
	Items []OrderSummary `json:"items"`
	Total int64          `json:"total"`
}
