package repositories

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"github.com/tongthanhdat009/CCNLTHD/internal/models"
)

type OrderHistoryRepository interface {
	ListMine(ctx context.Context, userID int, f models.OrderHistoryFilter) (models.OrderListResult, error)
	GetDetail(ctx context.Context, userID, orderID int) (models.OrderDetail, error)
}

type orderHistoryRepository struct{ db *gorm.DB }

func NewOrderHistoryRepository(db *gorm.DB) OrderHistoryRepository {
	return &orderHistoryRepository{db: db}
}

func (r *orderHistoryRepository) ListMine(ctx context.Context, userID int, f models.OrderHistoryFilter) (models.OrderListResult, error) {
	if f.Page <= 0 {
		f.Page = 1
	}
	if f.PageSize <= 0 || f.PageSize > 100 {
		f.PageSize = 10
	}

	base := r.db.WithContext(ctx).Table("donhang").
		Where("MaNguoiDung = ?", userID)

	if f.Status != "" {
		base = base.Where("TrangThai = ?", f.Status)
	}

	var total int64
	if err := base.Count(&total).Error; err != nil {
		return models.OrderListResult{}, fmt.Errorf("ListMine count: %w", err)
	}

	var rows []struct {
		MaDonHang int
		NgayTao   string
		TongTien  float64
		TrangThai string
	}
	if err := base.
		Select("MaDonHang, NgayTao, TongTien, TrangThai").
		Order("NgayTao DESC, MaDonHang DESC").
		Limit(f.PageSize).
		Offset((f.Page - 1) * f.PageSize).
		Scan(&rows).Error; err != nil {
		return models.OrderListResult{}, fmt.Errorf("ListMine query: %w", err)
	}

	items := make([]models.OrderSummary, 0, len(rows))
	for _, r := range rows {
		items = append(items, models.OrderSummary{
			MaDonHang: r.MaDonHang,
			NgayTao:   r.NgayTao,
			TongTien:  r.TongTien,
			TrangThai: r.TrangThai,
		})
	}
	return models.OrderListResult{Items: items, Total: total}, nil
}

func (r *orderHistoryRepository) GetDetail(ctx context.Context, userID, orderID int) (models.OrderDetail, error) {
	// 1) Lấy thông tin đơn, kiểm tra thuộc user
	var head struct {
		MaDonHang  int
		NgayTao    string
		TongTien   float64
		TrangThai  string
		TinhThanh  string
		QuanHuyen  string
		PhuongXa   string
		DuongSoNha string
		SDT        string `gorm:"column:SoDienThoai"`
		ThanhToan  string `gorm:"column:PhuongThucThanhToan"`
	}
	qh := r.db.WithContext(ctx).Table("donhang").
		Where("MaDonHang = ? AND MaNguoiDung = ?", orderID, userID).
		Select("MaDonHang, NgayTao, TongTien, TrangThai, TinhThanh, QuanHuyen, PhuongXa, DuongSoNha, SoDienThoai, PhuongThucThanhToan")
	if err := qh.Take(&head).Error; err != nil {
		return models.OrderDetail{}, fmt.Errorf("GetDetail head: %w", err)
	}

	// 2) Lấy items
	// ct (chitietdonhang) -> sp (sanpham) -> ctpn (chitietphieunhap) -> bt (bienthe) -> hh (hanghoa)
	var items []models.OrderItem
	if err := r.db.WithContext(ctx).
		Table("chitietdonhang AS ct").
		Joins("JOIN sanpham sp ON sp.MaSanPham = ct.MaSanPham").
		Joins("JOIN chitietphieunhap ctpn ON ctpn.MaChiTiet = sp.MaChiTietPhieuNhap").
		Joins("JOIN bienthe bt ON bt.MaBienthe = ctpn.MaBienthe").
		Joins("JOIN hanghoa hh ON hh.MaHangHoa = bt.MaHangHoa").
		Where("ct.MaDonHang = ?", orderID).
		Select("ct.MaSanPham AS MaSanPham, sp.Seri AS Seri, hh.TenHangHoa AS TenHangHoa, ct.GiaBan AS GiaBan").
		Scan(&items).Error; err != nil {
		return models.OrderDetail{}, fmt.Errorf("GetDetail items: %w", err)
	}

	return models.OrderDetail{
		MaDonHang: head.MaDonHang,
		NgayTao:   head.NgayTao,
		TongTien:  head.TongTien,
		TrangThai: head.TrangThai,

		TinhThanh:  head.TinhThanh,
		QuanHuyen:  head.QuanHuyen,
		PhuongXa:   head.PhuongXa,
		DuongSoNha: head.DuongSoNha,
		SDT:        head.SDT,
		ThanhToan:  head.ThanhToan,

		Items: items,
	}, nil
}
