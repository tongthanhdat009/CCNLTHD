package services

import (
	"context"

	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"github.com/tongthanhdat009/CCNLTHD/internal/repositories"
)

type OrderHistoryService interface {
	ListMine(ctx context.Context, userID int, f models.OrderHistoryFilter) (models.OrderListResult, error)
	GetDetail(ctx context.Context, userID, orderID int) (models.OrderDetail, error)
}

type orderHistoryService struct {
	repo repositories.OrderHistoryRepository
}

func NewOrderHistoryService(r repositories.OrderHistoryRepository) OrderHistoryService {
	return &orderHistoryService{repo: r}
}

func (s *orderHistoryService) ListMine(ctx context.Context, userID int, f models.OrderHistoryFilter) (models.OrderListResult, error) {
	return s.repo.ListMine(ctx, userID, f)
}
func (s *orderHistoryService) GetDetail(ctx context.Context, userID, orderID int) (models.OrderDetail, error) {
	return s.repo.GetDetail(ctx, userID, orderID)
}
