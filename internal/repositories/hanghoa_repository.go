package repositories

import (
	"database/sql"
	"fmt"

	"github.com/tongthanhdat009/CCNLTHD/internal/models"
)

type HangHoaRepository struct {
	db *sql.DB
}

func NewHangHoaRepository(db *sql.DB) *HangHoaRepository {
	return &HangHoaRepository{db: db}
}

func (r *HangHoaRepository) GetAllHangHoa() ([]models.HangHoa, error) {
	query := `
		SELECT *
		FROM hanghoa
	`
	
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hangHoas []models.HangHoa
	for rows.Next() {
		var hh models.HangHoa
		err := rows.Scan(
			&hh.MaHangHoa,
			&hh.TenHangHoa,
			&hh.MaHang,
			&hh.MaDanhMuc,
			&hh.Mau,
			&hh.MoTa,
			&hh.TrangThai,
			&hh.MaKhuyenMai,
			&hh.AnhDaiDien,
		)
		if err != nil {
			return nil, err
		}
		hangHoas = append(hangHoas, hh)
		fmt.Printf("Retrieved HangHoa: %+v\n", hh)
	}

	return hangHoas, nil
}