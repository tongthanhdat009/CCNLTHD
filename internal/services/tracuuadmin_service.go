package services

import (
    "github.com/tongthanhdat009/CCNLTHD/internal/models"
    "github.com/tongthanhdat009/CCNLTHD/internal/repositories"
	"strings"
	"errors"
)

type TraCuuAdminService interface {
	GetSanPhamBySeries(seri string) (models.SanPham, error)
	GetSanPhamByTrangThai(trangThai string) ([]models.SanPham, error)
}

type traCuuAdminService struct {
    repo repositories.TraCuuAdminRepository
}

func NewTraCuuAdminRepository(repo repositories.TraCuuAdminRepository) TraCuuAdminService {
    return &traCuuAdminService{repo: repo}
}

func (s *traCuuAdminService) GetSanPhamBySeries(seri string) (models.SanPham, error) {
	if strings.TrimSpace(seri) == "" {
		return models.SanPham{}, nil // Hoặc trả về lỗi tùy theo yêu cầu
	}
	return s.repo.GetSanPhamBySeries(seri)
}

func (s *traCuuAdminService) GetSanPhamByTrangThai(trangThai string) ([]models.SanPham, error) {
	if strings.TrimSpace(trangThai) == "" {
		return nil, errors.New("sản phẩm không tồn tại")  // Hoặc trả về lỗi tùy theo yêu cầu
	}
	return s.repo.GetSanPhamByTrangThai(trangThai)
}