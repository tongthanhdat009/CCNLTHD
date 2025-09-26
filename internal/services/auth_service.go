package services

import (
	"github.com/tongthanhdat009/CCNLTHD/internal/repositories"
)

type PermissionService interface {
	 KiemTraQuyen(maQuyen int, tenChucNang string, tenHanhDong string) (bool, error)
}

type permissionService struct {
	repo repositories.AuthRepository
}

func NewPermissionService(repo repositories.AuthRepository) PermissionService {
	return &permissionService{repo: repo}
}

func (s *permissionService) KiemTraQuyen(maQuyen int, tenChucNang string, tenHanhDong string) (bool, error) {
	return s.repo.KiemTraQuyen(maQuyen, tenChucNang, tenHanhDong)
}
