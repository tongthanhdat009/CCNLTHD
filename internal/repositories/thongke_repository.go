package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/tongthanhdat009/CCNLTHD/internal/models"
)

type ReportRepository interface {
	TopCustomers(ctx context.Context, r models.DateRange) ([]models.TopCustomer, error)
	PurchaseValue(ctx context.Context, r models.DateRange) (models.PurchaseValue, error)
	ImportedProducts(ctx context.Context, r models.DateRange) ([]models.ImportedProduct, error)
	ImportedBrands(ctx context.Context, r models.DateRange) ([]models.ImportedBrand, error)
	InvoiceStats(ctx context.Context, r models.DateRange) (models.InvoiceStats, error)
	BestSellers(ctx context.Context, r models.DateRange) ([]models.BestSeller, error)
	RevenueByBrand(ctx context.Context, r models.DateRange) ([]models.RevenueByBrand, error)
}

type reportRepository struct{ db *gorm.DB }

func NewReportRepository(db *gorm.DB) ReportRepository { return &reportRepository{db: db} }

// ---- helpers ----
func normRange(r models.DateRange) (from, to time.Time, has bool, err error) {
	if r.From == "" && r.To == "" {
		return time.Time{}, time.Time{}, false, nil
	}
	layout := "2006-01-02"
	var f, t time.Time
	if r.From != "" {
		if f, err = time.Parse(layout, r.From); err != nil {
			return
		}
	}
	if r.To != "" {
		if t, err = time.Parse(layout, r.To); err != nil {
			return
		}
		// inclusive đến cuối ngày
		t = t.Add(24 * time.Hour).Add(-time.Nanosecond)
	}
	// nếu chỉ có 1 đầu, cứ dùng đầu kia = đầu đó (1 ngày)
	if r.From != "" && r.To == "" {
		t = f.Add(24 * time.Hour).Add(-time.Nanosecond)
	}
	if r.From == "" && r.To != "" {
		f = t.Add(-24 * time.Hour).Add(time.Nanosecond)
	}
	return f, t, true, nil
}

func pickLimit(n int, def int) int {
	if n <= 0 {
		return def
	}
	if n > 100 {
		return 100
	}
	return n
}

// ---- queries ----

func (r *reportRepository) TopCustomers(ctx context.Context, dr models.DateRange) ([]models.TopCustomer, error) {
	from, to, has, err := normRange(dr)
	if err != nil {
		return nil, err
	}

	// NOTE: TrangThai lọc đơn bán; tuỳ DB bạn, chỉnh lại list trạng thái hoàn tất.
	q := `
		SELECT nd.MaNguoiDung,
		       COALESCE(nd.HoTen, nd.HoTen, CONCAT('User ', nd.MaNguoiDung)) AS HoTen,
		       SUM(dh.TongTien) AS TongChi,
		       COUNT(*) AS SoDon
		FROM donhang dh
		JOIN nguoidung nd ON nd.MaNguoiDung = dh.MaNguoiDung
		WHERE dh.TrangThai IN ('Hoàn tất','Đã giao','Đã thanh toán')
		%s
		GROUP BY nd.MaNguoiDung, HoTen
		ORDER BY TongChi DESC
		LIMIT ?
	`
	where := ""
	args := []any{}
	if has {
		where = "AND dh.NgayTao BETWEEN ? AND ?"
		args = append(args, from, to)
	}
	limit := pickLimit(dr.Limit, 10)
	args = append(args, limit)

	var out []models.TopCustomer
	if err := r.db.WithContext(ctx).Raw(fmt.Sprintf(q, where), args...).Scan(&out).Error; err != nil {
		return nil, fmt.Errorf("TopCustomers: %w", err)
	}
	return out, nil
}

func (r *reportRepository) PurchaseValue(ctx context.Context, dr models.DateRange) (models.PurchaseValue, error) {
	from, to, has, err := normRange(dr)
	if err != nil {
		return models.PurchaseValue{}, err
	}

	// Tổng chi nhập = SUM(ctpn.SoLuong * ctpn.DonGia) theo khoảng phieunhap.NgaySanXuat
	// NOTE: Nếu cột tên khác: thay ctpn.DonGia thành ctpn.GiaNhap (tuỳ schema).
	q := `
		SELECT COALESCE(SUM(ctpn.SoLuong * ctpn.GiaNhap), 0) AS Total
		FROM chitietphieunhap ctpn
		JOIN phieunhap pn ON pn.MaPhieuNhap = ctpn.MaPhieuNhap
		%s
	`
	where := ""
	args := []any{}
	if has {
		where = "WHERE pn.NgayNhap BETWEEN ? AND ?"
		args = append(args, from, to)
	}
	var res models.PurchaseValue
	if err := r.db.WithContext(ctx).Raw(fmt.Sprintf(q, where), args...).Scan(&res).Error; err != nil {
		return models.PurchaseValue{}, fmt.Errorf("PurchaseValue: %w", err)
	}
	return res, nil
}

func (r *reportRepository) ImportedProducts(ctx context.Context, dr models.DateRange) ([]models.ImportedProduct, error) {
	from, to, has, err := normRange(dr)
	if err != nil {
		return nil, err
	}

	q := `
		SELECT hh.MaHangHoa, hh.TenHangHoa, COALESCE(SUM(ctpn.SoLuong),0) AS SoLuongNhap
		FROM chitietphieunhap ctpn
		JOIN phieunhap pn ON pn.MaPhieuNhap = ctpn.MaPhieuNhap
		JOIN bienthe bt   ON bt.MaBienThe   = ctpn.MaBienThe
		JOIN hanghoa hh   ON hh.MaHangHoa   = bt.MaHangHoa
		%s
		GROUP BY hh.MaHangHoa, hh.TenHangHoa
		ORDER BY SoLuongNhap DESC
		LIMIT ?
	`
	where := ""
	args := []any{}
	if has {
		where = "WHERE pn.NgayNhap BETWEEN ? AND ?"
		args = append(args, from, to)
	}
	limit := pickLimit(dr.Limit, 20)
	args = append(args, limit)

	var out []models.ImportedProduct
	if err := r.db.WithContext(ctx).Raw(fmt.Sprintf(q, where), args...).Scan(&out).Error; err != nil {
		return nil, fmt.Errorf("ImportedProducts: %w", err)
	}
	return out, nil
}

func (r *reportRepository) ImportedBrands(ctx context.Context, dr models.DateRange) ([]models.ImportedBrand, error) {
	from, to, has, err := normRange(dr)
	if err != nil {
		return nil, err
	}

	q := `
		SELECT h.MaHang, h.TenHang, COALESCE(SUM(ctpn.SoLuong),0) AS SoLuong
		FROM chitietphieunhap ctpn
		JOIN phieunhap pn ON pn.MaPhieuNhap = ctpn.MaPhieuNhap
		JOIN bienthe bt   ON bt.MaBienThe   = ctpn.MaBienThe
		JOIN hanghoa hh   ON hh.MaHangHoa   = bt.MaHangHoa
		JOIN hang h      ON h.MaHang       = hh.MaHang
		%s
		GROUP BY h.MaHang, h.TenHang
		ORDER BY SoLuong DESC
		LIMIT ?
	`
	where := ""
	args := []any{}
	if has {
		where = "WHERE pn.NgayNhap BETWEEN ? AND ?"
		args = append(args, from, to)
	}
	limit := pickLimit(dr.Limit, 20)
	args = append(args, limit)

	var out []models.ImportedBrand
	if err := r.db.WithContext(ctx).Raw(fmt.Sprintf(q, where), args...).Scan(&out).Error; err != nil {
		return nil, fmt.Errorf("ImportedBrands: %w", err)
	}
	return out, nil
}

func (r *reportRepository) InvoiceStats(ctx context.Context, dr models.DateRange) (models.InvoiceStats, error) {
	from, to, has, err := normRange(dr)
	if err != nil {
		return models.InvoiceStats{}, err
	}

	q := `
		SELECT COUNT(*) AS SoDonBan, COALESCE(SUM(TongTien),0) AS TongDoanhThu
		FROM donhang dh
		WHERE dh.TrangThai IN ('Hoàn tất','Đã giao','Đã thanh toán')
		%s
	`
	where := ""
	args := []any{}
	if has {
		where = "AND dh.NgayTao BETWEEN ? AND ?"
		args = append(args, from, to)
	}
	var res models.InvoiceStats
	if err := r.db.WithContext(ctx).Raw(fmt.Sprintf(q, where), args...).Scan(&res).Error; err != nil {
		return models.InvoiceStats{}, fmt.Errorf("InvoiceStats: %w", err)
	}
	return res, nil
}

// TOP SP bán chạy
func (r *reportRepository) BestSellers(ctx context.Context, dr models.DateRange) ([]models.BestSeller, error) {
	from, to, has, err := normRange(dr)
	if err != nil {
		return nil, err
	}

	q := `
		SELECT hh.MaHangHoa, hh.TenHangHoa,
		       COUNT(*)                             AS SoLuongBan,
		       COALESCE(SUM(ct.GiaBan), 0)          AS DoanhThu
		FROM donhang dh
		JOIN chitietdonhang ct ON ct.MaDonHang = dh.MaDonHang
		JOIN sanpham sp        ON sp.MaSanPham = ct.MaSanPham
		JOIN chitietphieunhap ctpn ON ctpn.MaChiTiet = sp.MaChiTietPhieuNhap
		JOIN bienthe bt        ON bt.MaBienthe = ctpn.MaBienthe   -- chú ý tên cột trong DB của bạn
		JOIN hanghoa hh        ON hh.MaHangHoa = bt.MaHangHoa
		WHERE dh.TrangThai IN ('Hoàn tất','Đã giao','Đã thanh toán')
		%s
		GROUP BY hh.MaHangHoa, hh.TenHangHoa
		ORDER BY SoLuongBan DESC, DoanhThu DESC
		LIMIT ?
	`
	where := ""
	args := []any{}
	if has {
		where = "AND dh.NgayTao BETWEEN ? AND ?"
		args = append(args, from, to)
	}
	limit := pickLimit(dr.Limit, 20)
	args = append(args, limit)

	var out []models.BestSeller
	if err := r.db.WithContext(ctx).Raw(fmt.Sprintf(q, where), args...).Scan(&out).Error; err != nil {
		return nil, fmt.Errorf("BestSellers: %w", err)
	}
	return out, nil
}

// Doanh thu theo hãng
func (r *reportRepository) RevenueByBrand(ctx context.Context, dr models.DateRange) ([]models.RevenueByBrand, error) {
	from, to, has, err := normRange(dr)
	if err != nil {
		return nil, err
	}

	q := `
		SELECT h.MaHang, h.TenHang,
		       COALESCE(SUM(ct.GiaBan), 0) AS DoanhThu,
		       COUNT(*)                    AS SoLuongBan
		FROM donhang dh
		JOIN chitietdonhang ct ON ct.MaDonHang = dh.MaDonHang
		JOIN sanpham sp        ON sp.MaSanPham = ct.MaSanPham
		JOIN chitietphieunhap ctpn ON ctpn.MaChiTiet = sp.MaChiTietPhieuNhap
		JOIN bienthe bt        ON bt.MaBienthe = ctpn.MaBienthe
		JOIN hanghoa hh        ON hh.MaHangHoa = bt.MaHangHoa
		JOIN hang h            ON h.MaHang     = hh.MaHang
		WHERE dh.TrangThai IN ('Hoàn tất','Đã giao','Đã thanh toán')
		%s
		GROUP BY h.MaHang, h.TenHang
		ORDER BY DoanhThu DESC
		LIMIT ?
	`
	where := ""
	args := []any{}
	if has {
		where = "AND dh.NgayTao BETWEEN ? AND ?"
		args = append(args, from, to)
	}
	limit := pickLimit(dr.Limit, 20)
	args = append(args, limit)

	var out []models.RevenueByBrand
	if err := r.db.WithContext(ctx).Raw(fmt.Sprintf(q, where), args...).Scan(&out).Error; err != nil {
		return nil, fmt.Errorf("RevenueByBrand: %w", err)
	}
	return out, nil
}

// (Optional) helper scan for sql.NullString... nếu cần
func scanString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}
