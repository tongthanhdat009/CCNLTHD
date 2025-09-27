package services

import (
	"errors"

	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"github.com/tongthanhdat009/CCNLTHD/internal/repositories"
)

type GioHangService interface {
	TaoGioHang(giohang models.GioHang) error
	SuaGioHang(giohang models.GioHang) error
	XoaGioHang(giohang models.GioHang) error
	GetAll(manguoidung int) ([]models.GioHang, error)
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
