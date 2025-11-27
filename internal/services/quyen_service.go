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
	UpdateQuyen(maQuyen int, tenQuyen string) (*models.Quyen, error)
	DeleteQuyen(maQuyen int) error
	UpdateQuyenWithPermissions(maQuyen int, tenQuyen string, maChiTietChucNangs []int) (*models.Quyen, error)
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

// UpdateQuyen - Cập nhật tên quyền
func (s *quyenService) UpdateQuyen(maQuyen int, tenQuyen string) (*models.Quyen, error) {
	// 1. Validate tên quyền
	if tenQuyen == "" {
		return nil, errors.New("tên quyền không được để trống")
	}

	// 2. Kiểm tra quyền có tồn tại không
	exists, err := s.repo.CheckQuyenExists(maQuyen)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("quyền không tồn tại")
	}

	// 3. Thực hiện cập nhật
	quyen := &models.Quyen{
		MaQuyen:  maQuyen,
		TenQuyen: tenQuyen,
	}

	err = s.repo.UpdateQuyen(quyen)
	if err != nil {
		return nil, err
	}

	// 4. Trả về quyền đã cập nhật
	return s.GetByID(maQuyen)
}

// DeleteQuyen - Xóa quyền
func (s *quyenService) DeleteQuyen(maQuyen int) error {
	// 1. Kiểm tra quyền có tồn tại không
	exists, err := s.repo.CheckQuyenExists(maQuyen)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("quyền không tồn tại")
	}

	// 2. Kiểm tra có người dùng đang sử dụng quyền này không
	hasNguoiDung, err := s.repo.CheckQuyenHasNguoiDung(maQuyen)
	if err != nil {
		return err
	}
	if hasNguoiDung {
		return errors.New("không thể xóa quyền vì đang có người dùng sử dụng")
	}

	// 3. Thực hiện xóa (cascade sẽ tự động xóa phân quyền)
	return s.repo.DeleteQuyen(maQuyen)
}

// UpdateQuyenWithPermissions - Cập nhật quyền kèm phân quyền
func (s *quyenService) UpdateQuyenWithPermissions(maQuyen int, tenQuyen string, maChiTietChucNangs []int) (*models.Quyen, error) {
	// 1. Validate tên quyền
	if tenQuyen == "" {
		return nil, errors.New("tên quyền không được để trống")
	}

	// 2. Kiểm tra quyền tồn tại
	exists, err := s.repo.CheckQuyenExists(maQuyen)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("quyền không tồn tại")
	}

	// 3. Cập nhật tên quyền
	quyen := &models.Quyen{
		MaQuyen:  maQuyen,
		TenQuyen: tenQuyen,
	}

	err = s.repo.UpdateQuyen(quyen)
	if err != nil {
		return nil, err
	}

	// 4. Cập nhật phân quyền nếu có
	if len(maChiTietChucNangs) > 0 {
		// Xóa toàn bộ phân quyền cũ
		err = s.repo.DeletePhanQuyenByMaQuyen(maQuyen)
		if err != nil {
			return nil, err
		}

		// Thêm phân quyền mới
		chiTietChucNangs := make([]models.ChiTietChucNang, len(maChiTietChucNangs))
		for i, ma := range maChiTietChucNangs {
			chiTietChucNangs[i] = models.ChiTietChucNang{
				MaChiTietChucNang: ma,
			}
		}

		err = s.repo.PhanQuyen(maQuyen, chiTietChucNangs)
		if err != nil {
			return nil, err
		}
	}

	// 5. Trả về quyền đã cập nhật
	return s.GetByID(maQuyen)
}
