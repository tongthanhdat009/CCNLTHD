package services

import (
	"database/sql"
	"errors"

	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"github.com/tongthanhdat009/CCNLTHD/internal/repositories"
)

type ReviewService struct {
	Repo *repositories.ReviewRepository
}

func NewReviewService(repo *repositories.ReviewRepository) *ReviewService {
	return &ReviewService{Repo: repo}
}

func (s *ReviewService) Create(maNguoiDung, maSanPham, diem int, noiDung string) error {
	if diem < 1 || diem > 5 {
		return errors.New("Điểm phải từ 1 đến 5")
	}

	rv := &models.DanhGia{
		MaNguoiDung: maNguoiDung,
		MaSanPham:   maSanPham,
		Diem:        diem,
		NoiDung:     sql.NullString{String: noiDung, Valid: noiDung != ""},
		TrangThai:   "Chưa duyệt",
	}
	return s.Repo.Create(rv)
}

func (s *ReviewService) GetByProduct(maSanPham int) ([]models.DanhGia, error) {
	return s.Repo.GetByProduct(maSanPham)
}

func (s *ReviewService) GetByUser(maNguoiDung int) ([]models.DanhGia, error) {
	return s.Repo.GetByUser(maNguoiDung)
}

func (s *ReviewService) Approve(maDanhGia int) error {
	return s.Repo.Approve(maDanhGia)
}

func (s *ReviewService) Reject(maDanhGia int) error {
	return s.Repo.Reject(maDanhGia)
}

func (s *ReviewService) Delete(maDanhGia int) error {
	return s.Repo.Delete(maDanhGia)
}
