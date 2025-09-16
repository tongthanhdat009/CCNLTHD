package services

import (
	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"github.com/tongthanhdat009/CCNLTHD/internal/repositories"
)

type HangHoaService struct {
	repo *repositories.HangHoaRepository
}

func NewHangHoaService(repo *repositories.HangHoaRepository) *HangHoaService {
	return &HangHoaService{repo: repo}
}

func (s *HangHoaService) GetAllHangHoa() ([]models.HangHoa, error) {
	return s.repo.GetAllHangHoa()
}