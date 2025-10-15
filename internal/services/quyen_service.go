package services

import (
	"errors"
    "github.com/tongthanhdat009/CCNLTHD/internal/models"
    "github.com/tongthanhdat009/CCNLTHD/internal/repositories"
)

type QuyenService interface {
    GetAll() ([]models.Quyen, error)
	GetByID(id int) (*models.Quyen, error)
	CreateQuyen(TenQuyen string) (models.Quyen, error)
	GetAllChiTietChucNang() ([]int, error)
	CreateQuyenWithPermissions(tenQuyen string, maChiTietChucNangs []int) (models.Quyen, error)
	PhanQuyen(maQuyen int, maChiTietChucNangs []int) error
}

type quyenService struct {
    repo repositories.QuyenRepository
}
func NewQuyenService(repo repositories.QuyenRepository) QuyenService {
	return &quyenService{repo: repo}
}

func (s *quyenService) GetAll() ([]models.Quyen, error) {
	quyen, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	for i := range quyen {
		chucNangs, err := s.repo.GetChucNangVaChiTiet(quyen[i].MaQuyen)
		if err != nil {
			return nil, err
		}
		quyen[i].ChucNangs = chucNangs
	}
	return quyen, nil
}

func (s *quyenService) GetByID(id int) (*models.Quyen, error) {
	quyen, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	chucNangs, err := s.repo.GetChucNangVaChiTiet(quyen.MaQuyen)
	if err != nil {
		return nil, err
	}
	quyen.ChucNangs = chucNangs
	return quyen, nil
}

func (s *quyenService) CreateQuyen(TenQuyen string) (models.Quyen, error) {
	quyen, err := s.repo.CreateQuyen(TenQuyen)
    if err != nil {
        return models.Quyen{}, err
    }

    return quyen, nil
	
}

func (s *quyenService) CreateQuyenWithPermissions(tenQuyen string, maChiTietChucNangs []int) (models.Quyen, error) {
    // Validate tên quyền
    if tenQuyen == "" {
        return models.Quyen{}, errors.New("tên quyền không được để trống")
    }

    // Tạo quyền với phân quyền chi tiết
    quyen, err := s.repo.CreateQuyenWithPermissions(tenQuyen, maChiTietChucNangs)
    if err != nil {
        return models.Quyen{}, err
    }

    return quyen, nil
}

func (s *quyenService) GetAllChiTietChucNang() ([]int, error) {
    maChiTiets, err := s.repo.GetMaChiTietChucNang()
    if err != nil {
        return nil, err
    }
    return maChiTiets, nil
}

func (s *quyenService) PhanQuyen(maQuyen int, maChiTietChucNangs []int) error {
    // Validate mã quyền
    _, err := s.repo.GetByID(maQuyen)
    if err != nil {
        return errors.New("quyền không tồn tại")
    }

    // Validate danh sách chi tiết chức năng
    if len(maChiTietChucNangs) == 0 {
        return errors.New("danh sách chi tiết chức năng trống")
    }

    // Chuyển đổi sang slice ChiTietChucNang
    chiTietChucNangs := make([]models.ChiTietChucNang, len(maChiTietChucNangs))
    for i, ma := range maChiTietChucNangs {
        chiTietChucNangs[i] = models.ChiTietChucNang{
            MaChiTietChucNang: ma,
        }
    }

    // Phân quyền
    return s.repo.PhanQuyen(maQuyen, chiTietChucNangs)
}
