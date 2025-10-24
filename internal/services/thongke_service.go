package services

import (
	"context"

	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"github.com/tongthanhdat009/CCNLTHD/internal/repositories"
)

type ReportService interface {
	TopCustomers(ctx context.Context, r models.DateRange) ([]models.TopCustomer, error)
	PurchaseValue(ctx context.Context, r models.DateRange) (models.PurchaseValue, error)
	ImportedProducts(ctx context.Context, r models.DateRange) ([]models.ImportedProduct, error)
	ImportedBrands(ctx context.Context, r models.DateRange) ([]models.ImportedBrand, error)
	InvoiceStats(ctx context.Context, r models.DateRange) (models.InvoiceStats, error)
	BestSellers(ctx context.Context, r models.DateRange) ([]models.BestSeller, error)
	RevenueByBrand(ctx context.Context, r models.DateRange) ([]models.RevenueByBrand, error)
}

type reportService struct{ repo repositories.ReportRepository }

func NewReportService(r repositories.ReportRepository) ReportService { return &reportService{repo: r} }

func (s *reportService) TopCustomers(ctx context.Context, r models.DateRange) ([]models.TopCustomer, error) {
	return s.repo.TopCustomers(ctx, r)
}
func (s *reportService) PurchaseValue(ctx context.Context, r models.DateRange) (models.PurchaseValue, error) {
	return s.repo.PurchaseValue(ctx, r)
}
func (s *reportService) ImportedProducts(ctx context.Context, r models.DateRange) ([]models.ImportedProduct, error) {
	return s.repo.ImportedProducts(ctx, r)
}
func (s *reportService) ImportedBrands(ctx context.Context, r models.DateRange) ([]models.ImportedBrand, error) {
	return s.repo.ImportedBrands(ctx, r)
}
func (s *reportService) InvoiceStats(ctx context.Context, r models.DateRange) (models.InvoiceStats, error) {
	return s.repo.InvoiceStats(ctx, r)
}
func (s *reportService) BestSellers(ctx context.Context, r models.DateRange) ([]models.BestSeller, error) {
	return s.repo.BestSellers(ctx, r)
}
func (s *reportService) RevenueByBrand(ctx context.Context, r models.DateRange) ([]models.RevenueByBrand, error) {
	return s.repo.RevenueByBrand(ctx, r)
}
