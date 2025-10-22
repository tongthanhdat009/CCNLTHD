package services

import (
    "github.com/tongthanhdat009/CCNLTHD/internal/models"
    "github.com/tongthanhdat009/CCNLTHD/internal/repositories"
	"errors"
)

type PhanQuyenService interface {
    GetAll() ([]models.Quyen, error)
	GetByID(id int) (*models.Quyen, error)
	UpdatePhanQuyen(phanquyen *models.PhanQuyen) error
}

type phanQuyenService struct {
    repo repositories.PhanQuyenRepository
}
func NewPhanQuyenService(repo repositories.PhanQuyenRepository) PhanQuyenService {
	return &phanQuyenService{repo: repo}
}

func (s *phanQuyenService) GetAll() ([]models.Quyen, error) {
	quyen, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	for i := range quyen {
		chucNangs, err := s.repo.GetChucNangVaChiTiet(quyen[i].MaQuyen)
		if err != nil {
			return nil, err
		}
		quyen[i].ChucNangs = chucNangs
	}
	return quyen, nil
}

func (s *phanQuyenService) GetByID(id int) (*models.Quyen, error) {
	quyen, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	chucNangs, err := s.repo.GetChucNangVaChiTiet(quyen.MaQuyen)
	if err != nil {
		return nil, err
	}
	quyen.ChucNangs = chucNangs
	return quyen, nil
}
func (s *phanQuyenService) UpdatePhanQuyen(phanquyen *models.PhanQuyen) error {
	if phanquyen.TrangThai == "Đóng" || phanquyen.TrangThai == "Mở" {
		return s.repo.UpdatePhanQuyen(phanquyen)
	}
	return errors.New("trạng thái không hợp lệ, chỉ được phép 'Mở' hoặc 'Đóng'")
}