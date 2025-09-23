package services

import (
    "errors" // Thêm import này
    "github.com/tongthanhdat009/CCNLTHD/internal/models"
    "github.com/tongthanhdat009/CCNLTHD/internal/repositories"
	"gorm.io/gorm"
    "regexp"
    // "log"
)

type NhaCungCapService interface {
    GetAllNhaCungCap() ([]models.NhaCungCap, error)
    CreateNhaCungCap(nhacungcap *models.NhaCungCap) error
    UpdateNhaCungCap(nhacungcap *models.NhaCungCap) error
    GetNhaCungCapByID(id int) (*models.NhaCungCap, error)
    GetNhaCungCapByName(name string) ([]models.NhaCungCap, error)
    DeleteNhaCungCap(id int) error
}

type nhaCungCapService struct {
    repo                repositories.NhaCungCapRepository
    phieuNhapRepository repositories.PhieuNhapRepository
}

func NewNhaCungCapService(repo repositories.NhaCungCapRepository, db *gorm.DB) NhaCungCapService {
    return &nhaCungCapService{
        repo:                repo,
        phieuNhapRepository: repositories.NewPhieuNhapRepository(db),
    }
}

func (s *nhaCungCapService) GetAllNhaCungCap() ([]models.NhaCungCap, error) {
    return s.repo.GetAll()
}

func (s *nhaCungCapService) CreateNhaCungCap(nhacungcap *models.NhaCungCap) error {
    // Kiểm tra tên nhà cung cấp không được để trống
    if nhacungcap.TenNCC == "" {
        return errors.New("tên nhà cung cấp không được để trống")
    }

    // Kiểm tra địa chỉ không được để trống
    if nhacungcap.DiaChi == "" {
        return errors.New("địa chỉ không được để trống")
    }

    // Kiểm tra số điện thoại hợp lệ
    phoneRegex := regexp.MustCompile(`^(84|0[3|5|7|8|9])[0-9]{8}$`)
    if !phoneRegex.MatchString(nhacungcap.SoDienThoai) {
        return errors.New("số điện thoại không hợp lệ")
    }

    // Kiểm tra email hợp lệ (nếu có)
    if nhacungcap.Email != "" {
        emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
        if !emailRegex.MatchString(nhacungcap.Email) {
            return errors.New("email không hợp lệ")
        }
    }else{
        return errors.New("email không được để trống")
    }

    // Kiểm tra trùng lặp
    existingNhaCungCaps, err := s.repo.GetNhaCungCapByName(nhacungcap.TenNCC)
    if err != nil {
        return err // Lỗi hệ thống
    }
    if len(existingNhaCungCaps) > 0 {
        return errors.New("nhà cung cấp đã tồn tại")
    }

    // Thêm nhà cung cấp mới
    return s.repo.CreaateNhaCungCap(nhacungcap)
}

func (s *nhaCungCapService) UpdateNhaCungCap(nhacungcap *models.NhaCungCap) error {
    // Kiểm tra ID hợp lệ
    if nhacungcap.MaNCC <= 0 {
        return errors.New("ID nhà cung cấp không hợp lệ")
    }

    // Kiểm tra tên nhà cung cấp không được để trống
    if nhacungcap.TenNCC == "" {
        return errors.New("tên nhà cung cấp không được để trống")
    }

    // Kiểm tra địa chỉ không được để trống
    if nhacungcap.DiaChi == "" {
        return errors.New("địa chỉ không được để trống")
    }

    // Kiểm tra số điện thoại hợp lệ
    phoneRegex := regexp.MustCompile(`^(84|0[3|5|7|8|9])[0-9]{8}$`)
    if !phoneRegex.MatchString(nhacungcap.SoDienThoai) {
        return errors.New("số điện thoại không hợp lệ")
    }

    // Kiểm tra email hợp lệ (nếu có)
    if nhacungcap.Email != "" {
        emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
        if !emailRegex.MatchString(nhacungcap.Email) {
            return errors.New("email không hợp lệ")
        }
    }

    // Kiểm tra nhà cung cấp có tồn tại không
    existingNhaCungCap, err := s.repo.GetNhaCungCapByID(nhacungcap.MaNCC)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return errors.New("nhà cung cấp không tồn tại")
        }
        return err // Lỗi hệ thống
    }

    // Kiểm tra trùng lặp tên nhà cung cấp (nếu tên mới khác tên hiện tại)
    if existingNhaCungCap.TenNCC != nhacungcap.TenNCC {
        duplicateNhaCungCaps, err := s.repo.GetNhaCungCapByName(nhacungcap.TenNCC)
        if err != nil {
            return err // Lỗi hệ thống
        }
        if len(duplicateNhaCungCaps) > 0 {
            return errors.New("tên nhà cung cấp đã tồn tại")
        }
    }

    // Gọi repository để cập nhật nhà cung cấp
    return s.repo.UpdateNhaCungCap(nhacungcap)
}

func (s *nhaCungCapService) GetNhaCungCapByID(id int) (*models.NhaCungCap, error) {
    return s.repo.GetNhaCungCapByID(id)
}

func (s *nhaCungCapService) GetNhaCungCapByName(name string) ([]models.NhaCungCap, error) {
    return s.repo.GetNhaCungCapByName(name)
}

func (s *nhaCungCapService) DeleteNhaCungCap(id int) error {
    // log.Printf("Bắt đầu xóa nhà cung cấp với ID: %d", id)

    // Kiểm tra ID hợp lệ
    if id <= 0 {
        // log.Printf("ID không hợp lệ: %d", id)
        return errors.New("ID nhà cung cấp không hợp lệ")
    }

    // Kiểm tra nhà cung cấp có tồn tại không
    existingNhaCungCap, err := s.repo.GetNhaCungCapByID(id)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            // log.Printf("Nhà cung cấp không tồn tại với ID: %d", id)
            return errors.New("nhà cung cấp không tồn tại")
        }
        // log.Printf("Lỗi khi kiểm tra nhà cung cấp: %v", err)
        return err
    }
    // log.Printf("Nhà cung cấp tồn tại: %+v", existingNhaCungCap)

    // Kiểm tra nhà cung cấp có tồn tại trong phiếu nhập không
    exists, err := s.phieuNhapRepository.ExistsInPhieuNhap(existingNhaCungCap.MaNCC)
    if err != nil {
        // log.Printf("Lỗi khi kiểm tra phiếu nhập: %v", err)
        return err
    }
    // log.Printf("Nhà cung cấp tồn tại trong phiếu nhập: %v", exists)
    if exists {
        // log.Printf("Không thể xóa nhà cung cấp vì đã tồn tại trong phiếu nhập")
        return errors.New("không thể xóa nhà cung cấp vì đã tồn tại trong phiếu nhập")
    }

    // Gọi repository để xóa nhà cung cấp
    if err := s.repo.DeleteNhaCungCap(existingNhaCungCap.MaNCC); err != nil {
        // log.Printf("Lỗi khi xóa nhà cung cấp: %v", err)
        return err
    }

    // log.Printf("Xóa nhà cung cấp thành công với ID: %d", id)
    return nil
}