package services

import (
	"errors"

	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"github.com/tongthanhdat009/CCNLTHD/internal/repositories"
	"gorm.io/gorm"
)

type KhuyenMaiService interface {
	TaoKhuyenMai(khuyenMai *models.KhuyenMai) error
	SuaKhuyenMai(makhuyenmai int, khuyenMai models.KhuyenMai) error
	XoaKhuyenMai(makhuyenmai int) error
	GetAll() ([]models.KhuyenMai, error)
	GetByID(makhuyenmai int) (models.KhuyenMai, error)
}

type khuyenMaiService struct {
	repo repositories.KhuyenMaiRepository
}

func NewKhuyenMaiService(repo repositories.KhuyenMaiRepository) KhuyenMaiService {
	return &khuyenMaiService{repo: repo}
}

func (s *khuyenMaiService) TaoKhuyenMai(khuyenMai *models.KhuyenMai) error {
	exists, err := s.repo.KiemTraTenTonTai(khuyenMai.TenKhuyenMai)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("tên khuyến mãi đã tồn tại")
	}
	return s.repo.TaoKhuyenMai(khuyenMai)
}

func (s *khuyenMaiService) SuaKhuyenMai(makhuyenmai int, khuyenMai models.KhuyenMai) error {
	currentKhuyenMai, err := s.repo.GetByID(makhuyenmai)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("không tìm thấy khuyến mãi")
	}
	if err != nil {
		return err
	}

	// Nếu tên thay đổi → kiểm tra trùng
	if khuyenMai.TenKhuyenMai != "" && currentKhuyenMai.TenKhuyenMai != khuyenMai.TenKhuyenMai {
		exists, err := s.repo.KiemTraTenTonTai(khuyenMai.TenKhuyenMai)
		if err != nil {
			return err
		}
		if exists {
			return errors.New("tên khuyến mãi đã tồn tại")
		}
	}

	return s.repo.SuaKhuyenMai(makhuyenmai, khuyenMai)
}

// Xóa khuyến mãi
func (s *khuyenMaiService) XoaKhuyenMai(makhuyenmai int) error {
	return s.repo.XoaKhuyenMai(makhuyenmai)
}

// Lấy tất cả khuyến mãi
func (s *khuyenMaiService) GetAll() ([]models.KhuyenMai, error) {
	return s.repo.GetAll()
}

// Lấy khuyến mãi theo ID
func (s *khuyenMaiService) GetByID(makhuyenmai int) (models.KhuyenMai, error) {
	return s.repo.GetByID(makhuyenmai)
}
