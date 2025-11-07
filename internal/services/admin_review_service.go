package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"github.com/tongthanhdat009/CCNLTHD/internal/repositories"
)

type AdminReviewService interface {
	List(ctx context.Context, f models.AdminReviewFilter) (repositories.AdminReviewListResult, error)
	SetStatus(ctx context.Context, id int, status string) error
	Delete(ctx context.Context, id int) error
}

type adminReviewService struct {
	repo repositories.AdminReviewRepository
}

func NewAdminReviewService(repo repositories.AdminReviewRepository) AdminReviewService {
	return &adminReviewService{repo: repo}
}

func (s *adminReviewService) List(ctx context.Context, f models.AdminReviewFilter) (repositories.AdminReviewListResult, error) {
	if f.Page <= 0 {
		f.Page = 1
	}
	if f.PageSize <= 0 || f.PageSize > 200 {
		f.PageSize = 20
	}
	return s.repo.AdminList(ctx, f)
}

func (s *adminReviewService) SetStatus(ctx context.Context, id int, status string) error {
	switch status {
	case models.ReviewStatusApproved, models.ReviewStatusRejected, models.ReviewStatusHidden, models.ReviewStatusPending:
	default:
		return errors.New("trạng thái không hợp lệ")
	}
	affected, err := s.repo.AdminUpdateStatus(ctx, id, status)
	if err != nil {
		return err
	}
	if affected == 0 {
		return fmt.Errorf("không tìm thấy đánh giá #%d", id)
	}
	return nil
}

func (s *adminReviewService) Delete(ctx context.Context, id int) error {
	affected, err := s.repo.AdminDelete(ctx, id)
	if err != nil {
		return err
	}
	if affected == 0 {
		return fmt.Errorf("không tìm thấy đánh giá #%d", id)
	}
	return nil
}
