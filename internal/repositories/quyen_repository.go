package repositories

import (
	"errors"

	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"gorm.io/gorm"
)

type QuyenRepository interface {
	GetAll() ([]models.Quyen, error)
	GetByID(id int) (*models.Quyen, error)
	CreateQuyen(TenQuyen string) (models.Quyen, error)
	GetMaChiTietChucNang() ([]int, error)
	PhanQuyen(maQuyen int, chiTietChucNangs []models.ChiTietChucNang) error
	CreateQuyenWithPermissions(tenQuyen string, maChiTietChucNangs []int) (models.Quyen, error)
	GetChucNangVaChiTiet(idQuyen int) ([]models.ChucNang, error)
}

type QuyenRepo struct {
	db *gorm.DB
}

func NewQuyenRepository(db *gorm.DB) QuyenRepository {
	return &QuyenRepo{db: db}
}

func (r *QuyenRepo) GetAll() ([]models.Quyen, error) {
	var Quyens []models.Quyen
	if err := r.db.Find(&Quyens).Error; err != nil {
		return nil, err
	}
	return Quyens, nil
}

func (r *QuyenRepo) GetByID(id int) (*models.Quyen, error) {
	var quyen models.Quyen
	if err := r.db.First(&quyen, "MaQuyen = ?", id).Error; err != nil {
		return nil, err
	}
	return &quyen, nil
}

func (r *QuyenRepo) GetChucNangVaChiTiet(idQuyen int) ([]models.ChucNang, error) {
	var chucNangs []models.ChucNang

	err := r.db.
		Model(&models.ChucNang{}).
		Joins("JOIN ChiTietChucNang ON ChiTietChucNang.MaChucNang = ChucNang.MaChucNang").
		Joins("JOIN PhanQuyen ON PhanQuyen.MaChiTietChucNang = ChiTietChucNang.MaChiTietChucNang").
		Where("PhanQuyen.MaQuyen = ?", idQuyen).
		Preload("ChiTietChucNangs", func(db *gorm.DB) *gorm.DB {
			// Chỉ load chi tiết nào quyền này có
			return db.Joins("JOIN PhanQuyen ON PhanQuyen.MaChiTietChucNang = ChiTietChucNang.MaChiTietChucNang").
				Where("PhanQuyen.MaQuyen = ?", idQuyen)
		}).
		Group("ChucNang.MaChucNang").
		Find(&chucNangs).Error

	if err != nil {
		return nil, err
	}

	return chucNangs, nil
}

func (r *QuyenRepo) CreateQuyen(TenQuyen string) (models.Quyen, error) {
	quyen := models.Quyen{
		TenQuyen: TenQuyen,
	}
	if err := r.db.Create(&quyen).Error; err != nil {
		return models.Quyen{}, err
	}
	return quyen, nil
}

func (r *QuyenRepo) GetMaChiTietChucNang() ([]int, error) {
	var maChiTietChucNangs []int
	err := r.db.Model(&models.ChiTietChucNang{}).
		Select("MaChiTietChucNang").
		Find(&maChiTietChucNangs).Error
	if err != nil {
		return nil, err
	}
	return maChiTietChucNangs, nil
}

func (r *QuyenRepo) PhanQuyen(maQuyen int, chiTietChucNangs []models.ChiTietChucNang) error {
	for _, chiTiet := range chiTietChucNangs {
		if err := r.db.Model(&models.PhanQuyen{}).Create(&models.PhanQuyen{
			MaQuyen:           maQuyen,
			MaChiTietChucNang: chiTiet.MaChiTietChucNang,
			TrangThai:         "Mở",
		}).Error; err != nil {
			return err
		}
	}
	return nil
}

func (r *QuyenRepo) CheckQuyenExists(tenQuyen string) (bool, error) {
	var count int64
	err := r.db.Model(&models.Quyen{}).
		Where("TenQuyen = ?", tenQuyen).
		Count(&count).Error

	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// CreateQuyenWithPermissions - Tạo quyền mới kèm phân quyền chi tiết
func (r *QuyenRepo) CreateQuyenWithPermissions(tenQuyen string, maChiTietChucNangs []int) (models.Quyen, error) {
	// 1. Kiểm tra quyền đã tồn tại chưa
	exists, err := r.CheckQuyenExists(tenQuyen)
	if err != nil {
		return models.Quyen{}, err
	}
	if exists {
		return models.Quyen{}, errors.New("tên quyền đã tồn tại")
	}

	// 2. Bắt đầu transaction
	tx := r.db.Begin()
	if tx.Error != nil {
		return models.Quyen{}, tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 3. Tạo quyền mới
	quyen := models.Quyen{
		TenQuyen: tenQuyen,
	}
	if err := tx.Create(&quyen).Error; err != nil {
		tx.Rollback()
		return models.Quyen{}, errors.New("không thể tạo quyền: " + err.Error())
	}

	// 4. Tạo phân quyền chi tiết
	if len(maChiTietChucNangs) > 0 {
		// Validate các mã chi tiết chức năng có tồn tại không
		var count int64
		err := tx.Model(&models.ChiTietChucNang{}).
			Where("MaChiTietChucNang IN ?", maChiTietChucNangs).
			Count(&count).Error

		if err != nil {
			tx.Rollback()
			return models.Quyen{}, errors.New("lỗi khi kiểm tra chi tiết chức năng: " + err.Error())
		}

		if int(count) != len(maChiTietChucNangs) {
			tx.Rollback()
			return models.Quyen{}, errors.New("một số mã chi tiết chức năng không hợp lệ")
		}

		// Tạo các bản ghi phân quyền
		for _, maChiTiet := range maChiTietChucNangs {
			phanQuyen := models.PhanQuyen{
				MaQuyen:           quyen.MaQuyen,
				MaChiTietChucNang: maChiTiet,
				TrangThai:         "Đóng",
			}
			if err := tx.Create(&phanQuyen).Error; err != nil {
				tx.Rollback()
				return models.Quyen{}, errors.New("không thể phân quyền: " + err.Error())
			}
		}
	}

	// 5. Commit transaction
	if err := tx.Commit().Error; err != nil {
		return models.Quyen{}, err
	}

	return quyen, nil
}
