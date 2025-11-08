package services

import (
	"errors"
	"strings"

	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"github.com/tongthanhdat009/CCNLTHD/internal/repositories"
	"gorm.io/gorm"
)

type KhuyenMaiService interface {
	CreateKhuyenMai(khuyenMai *models.KhuyenMai) error
	UpdateKhuyenMai(makhuyenmai int, khuyenMai models.KhuyenMai) error
	DeleteKhuyenMai(makhuyenmai int) error
	GetAll() ([]models.KhuyenMai, error)
	GetByID(makhuyenmai int) (models.KhuyenMai, error)
	SearchKhuyenMai(keyword string, maKhuyenMai *int, tenKhuyenMai string, minGiaTri, maxGiaTri *float64, sortBy, sortOrder string, page, pageSize int) ([]models.KhuyenMai, int64, error)
}

type khuyenMaiService struct {
	repo repositories.KhuyenMaiRepository
}

func NewKhuyenMaiService(repo repositories.KhuyenMaiRepository) KhuyenMaiService {
	return &khuyenMaiService{repo: repo}
}

func (s *khuyenMaiService) CreateKhuyenMai(khuyenMai *models.KhuyenMai) error {
	// Validate tên khuyến mãi
	if strings.TrimSpace(khuyenMai.TenKhuyenMai) == "" {
		return errors.New("tên khuyến mãi không được để trống")
	}

	// Validate giá trị khuyến mãi
	if khuyenMai.GiaTri < 0 {
		return errors.New("giá trị khuyến mãi phải lớn hơn hoặc bằng 0")
	}

	// Kiểm tra trùng tên
	exists, err := s.repo.ExistsTenKhuyenMai(khuyenMai.TenKhuyenMai)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("tên khuyến mãi đã tồn tại")
	}

	return s.repo.CreateKhuyenMai(khuyenMai)
}

func (s *khuyenMaiService) UpdateKhuyenMai(makhuyenmai int, khuyenMai models.KhuyenMai) error {
	if makhuyenmai <= 0 {
        return errors.New("id không hợp lệ")
    }

    // Kiểm tra khuyến mãi có tồn tại không
    currentKhuyenMai, err := s.repo.GetByID(makhuyenmai)
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return errors.New("không tìm thấy khuyến mãi")
    }
    if err != nil {
        return err
    }

    // Validate tên khuyến mãi mới (nếu có)
    if khuyenMai.TenKhuyenMai != "" {
        if strings.TrimSpace(khuyenMai.TenKhuyenMai) == "" {
            return errors.New("tên khuyến mãi không được để trống")
        }

        // Nếu tên thay đổi → kiểm tra trùng
        if khuyenMai.TenKhuyenMai != currentKhuyenMai.TenKhuyenMai {
            exists, err := s.repo.ExistsTenKhuyenMai(khuyenMai.TenKhuyenMai)
            if err != nil {
                return err
            }
            if exists {
                return errors.New("tên khuyến mãi đã tồn tại")
            }
        }
    }

    // Validate giá trị khuyến mãi mới (nếu có)
    if khuyenMai.GiaTri != 0 {
        if khuyenMai.GiaTri < 0 {
            return errors.New("giá trị khuyến mãi phải lớn hơn hoặc bằng 0")
        }
    }

    return s.repo.UpdateKhuyenMai(makhuyenmai, khuyenMai)
}

func (s *khuyenMaiService) DeleteKhuyenMai(makhuyenmai int) error {
	// Kiểm tra khuyến mãi có tồn tại không
	_, err := s.repo.GetByID(makhuyenmai)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("không tìm thấy khuyến mãi")
	}
	if err != nil {
		return err
	}

	return s.repo.DeleteKhuyenMai(makhuyenmai)
}

func (s *khuyenMaiService) GetAll() ([]models.KhuyenMai, error) {
	return s.repo.GetAll()
}

func (s *khuyenMaiService) GetByID(makhuyenmai int) (models.KhuyenMai, error) {
	khuyenMai, err := s.repo.GetByID(makhuyenmai)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.KhuyenMai{}, errors.New("không tìm thấy khuyến mãi")
	}
	return khuyenMai, err
}

func (s *khuyenMaiService) SearchKhuyenMai(keyword string, maKhuyenMai *int, tenKhuyenMai string, minGiaTri, maxGiaTri *float64, sortBy, sortOrder string, page, pageSize int) ([]models.KhuyenMai, int64, error) {
    // ✅ Validation
    if maKhuyenMai != nil && *maKhuyenMai <= 0 {
        return nil, 0, errors.New("mã khuyến mãi không hợp lệ")
    }

    if minGiaTri != nil && *minGiaTri < 0 {
        return nil, 0, errors.New("giá trị min không hợp lệ")
    }

    if maxGiaTri != nil && *maxGiaTri < 0 {
        return nil, 0, errors.New("giá trị max không hợp lệ")
    }

    if minGiaTri != nil && maxGiaTri != nil && *minGiaTri > *maxGiaTri {
        return nil, 0, errors.New("giá trị min không được lớn hơn max")
    }

    return s.repo.SearchKhuyenMai(keyword, maKhuyenMai, tenKhuyenMai, minGiaTri, maxGiaTri, sortBy, sortOrder, page, pageSize)
}