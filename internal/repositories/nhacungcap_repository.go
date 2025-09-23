package repositories
import (
	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"gorm.io/gorm"
)
type NhaCungCapRepository interface {
	GetAll() ([]models.NhaCungCap, error)
	CreaateNhaCungCap(nhacungcap *models.NhaCungCap,) error
	UpdateNhaCungCap(nhacungcap *models.NhaCungCap,) error
	GetNhaCungCapByID(id int) (*models.NhaCungCap, error)
	GetNhaCungCapByName(name string) ([]models.NhaCungCap, error)
	DeleteNhaCungCap(id int) error
}
type NhaCungCapRepo struct {
	db *gorm.DB
}
func NewNhaCungCapRepository(db *gorm.DB) NhaCungCapRepository {
	return &NhaCungCapRepo{db: db}
}
func (r *NhaCungCapRepo) GetAll() ([]models.NhaCungCap, error) {
	var NhaCungCaps []models.NhaCungCap
	err := r.db.Find(&NhaCungCaps).Error
	return NhaCungCaps, err
}

func (r *NhaCungCapRepo) CreaateNhaCungCap(nhacungcap *models.NhaCungCap,) error {
	return r.db.Create(nhacungcap).Error
}

func (r *NhaCungCapRepo) UpdateNhaCungCap(nhacungcap *models.NhaCungCap,) error {
	return r.db.Save(nhacungcap).Error
}

func (r *NhaCungCapRepo) GetNhaCungCapByID(id int) (*models.NhaCungCap, error) {
	var nhacungcap models.NhaCungCap
	err := r.db.First(&nhacungcap, id).Error
	if err != nil {
		return nil, err
	}
	return &nhacungcap, nil
}

func (r *NhaCungCapRepo) GetNhaCungCapByName(name string) ([]models.NhaCungCap, error) {
	var nhacungcaps []models.NhaCungCap
	err := r.db.Where("tenncc LIKE ?", "%"+name+"%").Find(&nhacungcaps).Error
	if err != nil {
		return nil, err
	}
	return nhacungcaps, nil
}

func (r *NhaCungCapRepo) DeleteNhaCungCap(id int) error {
	return r.db.Delete(&models.NhaCungCap{}, id).Error
}