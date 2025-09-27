package services

import (
	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"github.com/tongthanhdat009/CCNLTHD/internal/repositories"
)

type KhuyenMaiService interface {
	TaoKhuyenMai(khuyenMai models.KhuyenMai) error
}

type khuyenMaiService struct {
	repo repositories.KhuyenMaiRepository
}

func NewKhuyenMaiService(repo repositories.KhuyenMaiRepository) KhuyenMaiService {
	return &khuyenMaiService{repo: repo}
}

func (s *khuyenMaiService) TaoKhuyenMai(khuyenMai models.KhuyenMai) error {
	return s.repo.TaoKhuyenMai(khuyenMai)
}
