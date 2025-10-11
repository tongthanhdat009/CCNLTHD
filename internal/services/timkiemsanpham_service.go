package services

import (
	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"github.com/tongthanhdat009/CCNLTHD/internal/repositories"
)

type TimKiemHangHoaService interface {
	TimHangHoa(
		tenHangHoa string,
		tenHang string,
		tenDanhMuc string,
		mau string,
		size string,
		giaToiThieu *float64, // Interface yêu cầu kiểu con trỏ *float64
		giaToiDa *float64,
	) ([]models.HangHoa, error)
}

type timKiemHangHoaService struct {
	repo repositories.TimKiemHangHoaRepository
}

func NewTimKiemHangHoaService(repo repositories.TimKiemHangHoaRepository) TimKiemHangHoaService {
	return &timKiemHangHoaService{repo: repo}
}

// SỬA LẠI SIGNATURE CỦA HÀM NÀY ĐỂ KHỚP VỚI INTERFACE
func (s *timKiemHangHoaService) TimHangHoa(
	tenHangHoa string,
	tenHang string,
	tenDanhMuc string,
	mau string,
	size string,
	giaToiThieu *float64, // Phải là *float64 để khớp với interface
	giaToiDa *float64,   // Phải là *float64 để khớp với interface
) ([]models.HangHoa, error) {
	// Logic của service rất đơn giản, chỉ cần gọi xuống repository và truyền thẳng tham số
	return s.repo.TimHangHoa(tenHangHoa, tenHang, tenDanhMuc, mau, size, giaToiThieu, giaToiDa)
}