package services

import (
    "github.com/tongthanhdat009/CCNLTHD/internal/models"
    "github.com/tongthanhdat009/CCNLTHD/internal/repositories"
)

type QuyenService interface {
    GetAll() ([]models.Quyen, error)
	GetByID(id int) (*models.Quyen, error)
}

type quyenService struct {
    repo repositories.QuyenRepository
}
func NewQuyenService(repo repositories.QuyenRepository) QuyenService {
	return &quyenService{repo: repo}
}

func (s *quyenService) GetAll() ([]models.Quyen, error) {
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

func (s *quyenService) GetByID(id int) (*models.Quyen, error) {
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