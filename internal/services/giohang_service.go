package services

import (
	"errors"
	"fmt"

	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"github.com/tongthanhdat009/CCNLTHD/internal/repositories"
	"time"
)

type GioHangService interface {
	TaoGioHang(giohang models.GioHang) error
	SuaGioHang(giohang models.GioHang) error
	XoaGioHang(giohang models.GioHang) error
	GetAll(manguoidung int) ([]models.GioHang, error)
	ThanhToan(giohang []models.GioHang, manguoidung int,tinh string ,quan string, phuong string, sonha string, phuongthucthanhtoan string, sodienthoai string) (models.DonHang, error)
}

type gioHangService struct {
	repo repositories.GioHangRepository
}

func NewGioHangService(repo repositories.GioHangRepository) GioHangService {
	return &gioHangService{repo: repo}
}

func (s *gioHangService) TaoGioHang(giohang models.GioHang) error {
	if giohang.SoLuong <= 0 {
		return errors.New("giỏ hàng không hợp lệ")
	}

	var err error
	giohang.Gia, err = s.repo.LayGia(giohang.MaBienThe)
	if err != nil {
		return errors.New("giỏ hàng không hợp lệ")
	}
	err = s.repo.CheckStatus(giohang.MaBienThe)
	if err != nil {
		return errors.New("biến thể không bán")
	}

	// Kiểm tra sản phẩm đã tồn tại trong giỏ hàng chưa
	if s.repo.KiemTraTonTai(giohang.MaNguoiDung, giohang.MaBienThe) {
		return errors.New("biến thể đã có trong giỏ hàng")
	}

	if !s.repo.CheckSoLuong(giohang.SoLuong, giohang.MaBienThe) {
		return errors.New("không đủ hàng")
	}
	return s.repo.TaoGioHang(giohang)
}

func (s *gioHangService) SuaGioHang(giohang models.GioHang) error {
	if giohang.SoLuong <= 0 {
		return errors.New("giỏ hàng không hợp lệ")
	}

	// Kiểm tra sản phẩm đã tồn tại trong giỏ hàng chưa
	if !s.repo.KiemTraTonTai(giohang.MaNguoiDung, giohang.MaBienThe) {
		return errors.New("chưa có sản phẩm trong giỏ hàng")
	}

	if !s.repo.CheckSoLuong(giohang.SoLuong, giohang.MaBienThe) {
		return errors.New("không đủ hàng")
	}
	return s.repo.SuaGioHang(giohang)
}

func (s *gioHangService) XoaGioHang(giohang models.GioHang) error {
	if giohang.SoLuong <= 0 {
		return errors.New("giỏ hàng không hợp lệ")
	}

	if err := s.repo.XoaGioHang(giohang); err != nil {
		return err
	}
	return nil
}

func (s *gioHangService) GetAll(manguoidung int) ([]models.GioHang, error) {
	return s.repo.GetAll(manguoidung)
}

func (s *gioHangService) ThanhToan(
	giohang []models.GioHang,
	manguoidung int,
	tinh string,
	quan string,
	phuong string,
	sonha string,
	phuongthucthanhtoan string,
	sodienthoai string) (models.DonHang, error) {

	if (tinh == "") || (quan == "") || (phuong == "") || (sonha == "") || (phuongthucthanhtoan == "") || (sodienthoai == "") {
		return models.DonHang{}, errors.New("thiếu thông tin thanh toán")
	}

	tx := s.repo.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 1️⃣ Kiểm tra biến thể
	for _, item := range giohang {
		if err := s.repo.CheckBienThe(item.MaBienThe, int(item.Gia), item.SoLuong); err != nil {
			return models.DonHang{}, err
		}
	}

	// 3️⃣ Tạo đơn hàng
	donhang := models.DonHang{
		MaNguoiDung:         manguoidung,
		NgayTao:             time.Now(),
		TrangThai:           "Đang xử lý",
		TinhThanh:           tinh,
		QuanHuyen:           quan,
		PhuongXa:            phuong,
		DuongSoNha:          sonha,
		PhuongThucThanhToan: phuongthucthanhtoan,
		SoDienThoai:         sodienthoai,
		TongTien:            0, // Sẽ tính sau
	}

	var chiTietDonHangs []models.ChiTietDonHang
	for _, item := range giohang {
		donhang.TongTien += item.Gia * float64(item.SoLuong)
	}
	if err := s.repo.CreateDonHang(tx, &donhang); err != nil {
		tx.Rollback()
		return models.DonHang{}, err
	}

	for _, item := range giohang {
		
		sanPhams, err := s.repo.GetSanPham(tx, item.MaBienThe, item.SoLuong)
		if err != nil {
			tx.Rollback()
			return models.DonHang{}, err
		}

		for _, sp := range sanPhams {
			chiTietDonHangs = append(chiTietDonHangs, models.ChiTietDonHang{
				MaDonHang: donhang.MaDonHang,
				MaSanPham: sp.MaSanPham,
				GiaBan:    item.Gia, // Giá tại thời điểm thanh toán
			})
		}
	}
	fmt.Println("Chi tiết đơn hàng:", chiTietDonHangs)

	// 5️⃣ Tạo tất cả chi tiết đơn hàng một lượt
	if err := s.repo.CreateChiTietDonHang(tx, chiTietDonHangs); err != nil {
		tx.Rollback()
		return models.DonHang{}, err
	}

	if err := s.repo.XoaGioHangCuaNguoiDung(tx, manguoidung); err != nil {
		tx.Rollback()
		return models.DonHang{}, err
	}


	// 5️⃣ Commit transaction
	if err := tx.Commit().Error; err != nil {
		return models.DonHang{}, err
	}

	donhang.ChiTietDonHangs = chiTietDonHangs

	return donhang, nil
}

