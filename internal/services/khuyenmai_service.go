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
	SearchKhuyenMai(keyword string) ([]models.KhuyenMai, error)
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
		if currentKhuyenMai.TenKhuyenMai != khuyenMai.TenKhuyenMai {
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
	if khuyenMai.GiaTri != 0 { // Giả sử 0 nghĩa là không cập nhật
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

func (s *khuyenMaiService) SearchKhuyenMai(keyword string) ([]models.KhuyenMai, error) {
	// ✅ Validate keyword
	if strings.TrimSpace(keyword) == "" {
		return []models.KhuyenMai{}, nil // Trả về rỗng nếu không có từ khóa
	}
	return s.repo.SearchKhuyenMai(keyword)
}
