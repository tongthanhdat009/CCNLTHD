package services

import (
    "github.com/tongthanhdat009/CCNLTHD/internal/models"
    "github.com/tongthanhdat009/CCNLTHD/internal/repositories"
)

type DonHangService interface {
    GetAllDonHang() ([]models.DonHang, error)
}

type donHangService struct {
    repo repositories.DonHangRepository
}

func NewDonHangService(repo repositories.DonHangRepository) DonHangService {
    return &donHangService{repo: repo}
}

func (s *donHangService) GetAllDonHang() ([]models.DonHang, error) {
    return s.repo.GetAll()
}