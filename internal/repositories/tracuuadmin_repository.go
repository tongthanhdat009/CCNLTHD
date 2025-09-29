package repositories
import (
	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"gorm.io/gorm"
)
type TraCuuAdminRepository interface {
	GetSanPhamBySeries(seri string) (models.SanPham, error)
    GetSanPhamByTrangThai(trangThai string) ([]models.SanPham, error)
}

type traCuuAdminRepository struct {
	db *gorm.DB
}

func NewTraCuuAdminRepository(db *gorm.DB) TraCuuAdminRepository {
	return &traCuuAdminRepository{db: db}
}

func (r *traCuuAdminRepository) GetSanPhamBySeries(seri string) (models.SanPham, error) {
	var sanpham models.SanPham
	if err := r.db.Preload("ChiTietPhieuNhap").Preload("ChiTietPhieuNhap.PhieuNhap").Where("Seri = ?", seri).First(&sanpham).Error; err != nil {
		return models.SanPham{}, err
	}
	return sanpham, nil
}

func (r *traCuuAdminRepository) GetSanPhamByTrangThai(trangThai string) ([]models.SanPham, error) {
    var sanphams []models.SanPham
    if err := r.db.Where("TrangThai = ?", trangThai).Find(&sanphams).Error; err != nil {
        return nil, err
    }
    return sanphams, nil
}