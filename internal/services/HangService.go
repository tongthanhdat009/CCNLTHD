// internal/services/brand_service.go
package services

import (
	"database/sql"
	"log"
	"github.com/tongthanhdat009/CCNLTHD/internal/models" // Thay bằng module path của bạn
)

// BrandService chứa logic liên quan đến hãng
type HangService struct {
	DB *sql.DB
}

// GetAllBrands lấy tất cả các hãng từ database
func (s *HangService) GetAllBrands() ([]models.Hang, error) {
	query := "SELECT MaHang, TenHang FROM hang"

	rows, err := s.DB.Query(query)
	if err != nil {
		log.Printf("Error querying brands: %v", err)
		return nil, err
	}
	defer rows.Close()

	var brands []models.Hang
	for rows.Next() {
		var b models.Hang
		if err := rows.Scan(&b.ID, &b.TenHang); err != nil {
			log.Printf("Error scanning brand: %v", err)
			continue
		}
		brands = append(brands, b)
	}

	return brands, nil
}