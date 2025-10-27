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
		return models.SanPham{}, errors.New("không được để trống seri") 
	}

	var sanPham models.SanPham
	if result, err := s.repo.GetSanPhamBySeries(seri); err != nil {
		return sanPham, errors.New("sản phẩm không tồn tại")
	} else {
		sanPham = result
	}
	return sanPham, nil
}

func (s *traCuuAdminService) GetSanPhamByTrangThai(trangThai string) ([]models.SanPham, error) {
	if strings.TrimSpace(trangThai) == "" {
		return nil, errors.New("trạng thái không tồn tại")  
	}
	if trangThai != "Đã bán" && trangThai != "Chưa bán" && trangThai != "Chờ duyệt" {
		return nil, errors.New("trạng thái không hợp lệ")  
	}
	var sanPhams, err = s.repo.GetSanPhamByTrangThai(trangThai)
	if err != nil {
		return nil, errors.New("không có sản phẩm nào với trạng thái này")
	}
	return sanPhams, nil
}