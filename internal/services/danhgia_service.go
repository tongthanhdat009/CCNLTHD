package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"github.com/tongthanhdat009/CCNLTHD/internal/repositories"
)

type ReviewService interface {
	Create(ctx context.Context, userID int, dto models.CreateReviewDTO) (int64, error)
	ListApprovedByHangHoa(ctx context.Context, maHH int) ([]models.Review, error)
	ListMine(ctx context.Context, userID int) ([]models.Review, error)
}

type reviewService struct{ repo repositories.ReviewRepository }

func NewReviewService(repo repositories.ReviewRepository) ReviewService {
	return &reviewService{repo: repo}
}

func (s *reviewService) Create(ctx context.Context, userID int, dto models.CreateReviewDTO) (int64, error) {
	if dto.Diem < 1 || dto.Diem > 5 {
		return 0, errors.New("điểm phải từ 1 đến 5")
	}
	ok, err := s.repo.UserPurchasedHangHoa(ctx, userID, dto.MaHangHoa)
	if err != nil {
		return 0, err
	}
	if !ok {
		return 0, fmt.Errorf("Bạn chưa mua hàng hóa này")
	}
	return s.repo.InsertHangHoa(ctx, userID, dto.MaHangHoa, dto)
}
func (s *reviewService) ListApprovedByHangHoa(ctx context.Context, maHH int) ([]models.Review, error) {
	return s.repo.ListApprovedByHangHoa(ctx, maHH)
}
func (s *reviewService) ListMine(ctx context.Context, userID int) ([]models.Review, error) {
	return s.repo.ListByUser(ctx, userID)
}
