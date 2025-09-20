package services

import (
    "github.com/tongthanhdat009/CCNLTHD/internal/models"
    "github.com/tongthanhdat009/CCNLTHD/internal/repositories"
)

type HangHoaService interface {
    GetAllHangHoa() ([]models.HangHoa, error)
}

type hangHoaService struct {
    repo repositories.HangHoaRepository
}

func NewHangHoaService(repo repositories.HangHoaRepository) HangHoaService {
    return &hangHoaService{repo: repo}
}

func (s *hangHoaService) GetAllHangHoa() ([]models.HangHoa, error) {
    return s.repo.GetAll()
}