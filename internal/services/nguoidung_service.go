package services

import (
    "github.com/tongthanhdat009/CCNLTHD/internal/models"
    "github.com/tongthanhdat009/CCNLTHD/internal/repositories"
)

type NguoiDungService interface {
    GetAllNguoiDung() ([]models.NguoiDung, error)
    GetNguoiDungByID(maNguoiDung int) (*models.NguoiDung, error)
    UpdateNguoiDung(maNguoiDung int, nguoiDung models.NguoiDung) error
}

type nguoiDungService struct {
    repo repositories.NguoiDungRepository
}

func NewNguoiDungService(repo repositories.NguoiDungRepository) NguoiDungService {
    return &nguoiDungService{repo: repo}
}

func (s *nguoiDungService) GetAllNguoiDung() ([]models.NguoiDung, error) {
    return s.repo.GetAll()
}

func (s *nguoiDungService) GetNguoiDungByID(maNguoiDung int) (*models.NguoiDung, error) {
    return s.repo.GetNguoiDungByID(maNguoiDung)
}

func (s *nguoiDungService) UpdateNguoiDung(maNguoiDung int, nguoiDung models.NguoiDung) error {
    return s.repo.Update(maNguoiDung, nguoiDung)
}