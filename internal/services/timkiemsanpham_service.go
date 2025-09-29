package services

import (
    "github.com/tongthanhdat009/CCNLTHD/internal/models"
    "github.com/tongthanhdat009/CCNLTHD/internal/repositories"
)

type TimKiemSanPhamService interface {
    TimSanPham(tenHangHoa string,
                tenHang string,
                tenDanhMuc string,
                mau string,
                size string,
                giaToiThieu string,
                giaToiDa string) ([]models.SanPham, error)
}

type timKiemSanPhamService struct {
    repo repositories.TimKiemSanPhamRepository
}

func NewTimKiemSanPhamService(repo repositories.TimKiemSanPhamRepository) TimKiemSanPhamService { // Sửa return type
    return &timKiemSanPhamService{repo: repo} // Sửa struct
}

func (s *timKiemSanPhamService) TimSanPham(tenHangHoa string, // Viết hoa
                                            tenHang string,
                                            tenDanhMuc string,
                                            mau string,
                                            size string,
                                            giaToiThieu string, // Giữ int
                                            giaToiDa string) ([]models.SanPham, error) { // Sửa signature
    return s.repo.TimSanPham(tenHangHoa, tenHang, tenDanhMuc, mau, size, giaToiThieu, giaToiDa)
}