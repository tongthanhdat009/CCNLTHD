package repositories
import (
	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"gorm.io/gorm"
)
type PhieuNhapRepository interface {
	GetAll() ([]models.PhieuNhap, error)
	ExistsInPhieuNhap(nhaCungCapID int) (bool, error)
}

type phieuNhapRepository struct {
	db *gorm.DB
}

func NewPhieuNhapRepository(db *gorm.DB) PhieuNhapRepository {
	return &phieuNhapRepository{db: db}
}

func (r *phieuNhapRepository) GetAll() ([]models.PhieuNhap, error) {
	var phieunhaps []models.PhieuNhap
	if err := r.db.Preload("NhaCungCap").Find(&phieunhaps).Error; err != nil {
		return nil, err
	}
	return phieunhaps, nil
}

func (r *phieuNhapRepository) ExistsInPhieuNhap(nhaCungCapID int) (bool, error) {
    var count int64
    if err := r.db.Model(&models.PhieuNhap{}).Where("mancc = ?", nhaCungCapID).Count(&count).Error; err != nil {
        return false, err
    }
    return count > 0, nil
}