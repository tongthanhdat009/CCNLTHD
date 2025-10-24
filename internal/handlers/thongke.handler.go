package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"github.com/tongthanhdat009/CCNLTHD/internal/services"
)

type ReportHandler struct{ svc services.ReportService }

func NewReportHandler(s services.ReportService) *ReportHandler { return &ReportHandler{svc: s} }

// GET /api/reports/top-customers?from=&to=&limit=
func (h *ReportHandler) TopCustomers(c *gin.Context) {
	var f models.DateRange
	_ = c.ShouldBindQuery(&f)
	data, err := h.svc.TopCustomers(c.Request.Context(), f)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

// GET /api/reports/purchase-value?from=&to=
func (h *ReportHandler) PurchaseValue(c *gin.Context) {
	var f models.DateRange
	_ = c.ShouldBindQuery(&f)
	data, err := h.svc.PurchaseValue(c.Request.Context(), f)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

// GET /api/reports/imported-products?from=&to=&limit=
func (h *ReportHandler) ImportedProducts(c *gin.Context) {
	var f models.DateRange
	_ = c.ShouldBindQuery(&f)
	data, err := h.svc.ImportedProducts(c.Request.Context(), f)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

// GET /api/reports/imported-brands?from=&to=&limit=
func (h *ReportHandler) ImportedBrands(c *gin.Context) {
	var f models.DateRange
	_ = c.ShouldBindQuery(&f)
	data, err := h.svc.ImportedBrands(c.Request.Context(), f)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

// GET /api/reports/invoices?from=&to=
func (h *ReportHandler) Invoices(c *gin.Context) {
	var f models.DateRange
	_ = c.ShouldBindQuery(&f)
	data, err := h.svc.InvoiceStats(c.Request.Context(), f)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

// GET /api/reports/best-sellers?from=&to=&limit=
func (h *ReportHandler) BestSellers(c *gin.Context) {
	var f models.DateRange
	_ = c.ShouldBindQuery(&f)
	data, err := h.svc.BestSellers(c.Request.Context(), f)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

// GET /api/reports/revenue-by-brand?from=&to=&limit=
func (h *ReportHandler) RevenueByBrand(c *gin.Context) {
	var f models.DateRange
	_ = c.ShouldBindQuery(&f)
	data, err := h.svc.RevenueByBrand(c.Request.Context(), f)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}
