package services

import (
    "errors"
    "github.com/tongthanhdat009/CCNLTHD/internal/models"
    "github.com/tongthanhdat009/CCNLTHD/internal/repositories"
    "gorm.io/gorm"
)

type HangService interface {
    GetAllHang() ([]models.Hang, error)
    DeleteHang(id int) error
    CreateHang(hang *models.Hang) error
    UpdateHang(hang *models.Hang) error
    GetHangByID(id int) (*models.Hang, error)  
    GetHangByName(name string) ([]models.Hang, error) 
}

type hangService struct {
    repo        repositories.HangRepository
    hangHoaRepo repositories.HangHoaRepository
}

func NewHangService(repo repositories.HangRepository, hangHoaRepo repositories.HangHoaRepository) HangService {
    return &hangService{repo: repo, hangHoaRepo: hangHoaRepo}
}

func (s *hangService) GetAllHang() ([]models.Hang, error) {
    return s.repo.GetAll()
}

func (s *hangService) DeleteHang(id int) error {
    if id <= 0 {
        return errors.New("mã hãng không hợp lệ")
    }

    // Kiểm tra hãng có tồn tại không
    _, err := s.repo.GetHangByID(id)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return errors.New("hãng không tồn tại")
        }
        return err
    }

    // Kiểm tra xem hãng có hàng hóa liên quan không
    count, err := s.hangHoaRepo.CountByHangID(id)
    if err != nil {
        return err
    }
    if count > 0 {
        return errors.New("không thể xóa hãng vì có hàng hóa đang sử dụng")
    }

    // Xóa hãng
    if err := s.repo.DeleteHang(id); err != nil {
        return err
    }

    return nil
}

func (s *hangService) CreateHang(hang *models.Hang) error {
    if hang.TenHang == "" {
        return errors.New("tên hãng không được để trống")
    }
    
    // Kiểm tra trùng lặp
    existingHangs, err := s.repo.GetHangByName(hang.TenHang)
    if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
        return err
    }
    if len(existingHangs) > 0 {
        return errors.New("hãng đã tồn tại")
    }

    return s.repo.CreateHang(hang)
}

func (s *hangService) UpdateHang(hang *models.Hang) error {
    if hang.MaHang <= 0 {
        return errors.New("mã hãng không hợp lệ")
    }
    // Kiểm tra xem hãng có hàng hóa liên quan không
    count, err := s.hangHoaRepo.CountByHangID(hang.MaHang)
    if err != nil {
        return err
    }
    if count > 0 {
        return errors.New("không thể cập nhật hãng vì có hàng hóa đang sử dụng")
    }
    
    // Kiểm tra hãng có tồn tại không
    _, err = s.repo.GetHangByID(hang.MaHang)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return errors.New("hãng không tồn tại")
        }
        return err
    }
    
    if hang.TenHang == "" {
        return errors.New("tên hãng không được để trống")
    }
    // Kiểm tra trùng lặp
    existingHangs, err := s.repo.GetHangByName(hang.TenHang)
    if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
        return err
    }
    if len(existingHangs) > 1 && existingHangs[0].MaHang != hang.MaHang {
        return errors.New("tên hãng đã tồn tại")
    }
    
    return s.repo.UpdateHang(hang)
}

func (s *hangService) GetHangByID(id int) (*models.Hang, error) {
    if id <= 0 {
        return nil, errors.New("mã hãng không hợp lệ")
    }

    hang, err := s.repo.GetHangByID(id)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errors.New("hãng không tồn tại")
        }
        return nil, err
    }

    return hang, nil
}

func (s *hangService) GetHangByName(name string) ([]models.Hang, error) {
    if name == "" {
        return nil, errors.New("tên hãng không được để trống")
    }

    return s.repo.GetHangByName(name)
}