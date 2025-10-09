package services

import (
	"errors"
	"strconv"

	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"github.com/tongthanhdat009/CCNLTHD/internal/repositories"
)

type QuanLyBienTheService interface {
    GetBienTheTheoMaHangHoa(maHangHoa int) ([]models.BienThe, error)
    GetBienTheTheoMa(maBienThe int) (*models.BienThe, error)
    CreateBienTheTheoMaHangHoa(bienThe *models.BienThe) error
    UpdateBienTheInfo(bienThe *models.BienThe) error
    UpdateBienTheStatus(maBienThe int, trangThai string) error
    DeleteBienThe(maBienThe int) error
}

type BienTheSvc struct {
    repo repositories.BienTheRepository
}

func NewBienTheService(repo repositories.BienTheRepository) QuanLyBienTheService {
    return &BienTheSvc{repo: repo}
}

func (s *BienTheSvc) GetBienTheTheoMaHangHoa(maHangHoa int) ([]models.BienThe, error) {
    return s.repo.GetBienTheTheoMaHangHoa(maHangHoa)
}

func (s *BienTheSvc) GetBienTheTheoMa(maBienThe int) (*models.BienThe, error) {
    return s.repo.GetBienTheTheoMa(maBienThe)
}

func (s *BienTheSvc) DeleteBienThe(maBienThe int) error {
    existsBienThe, err := s.repo.GetBienTheTheoMa(maBienThe)
    if err != nil {
        return err
    }
    if existsBienThe == nil {
        return errors.New("biến thể không tồn tại")
    }
    used, err := s.repo.HasChiTietPhieuNhap(maBienThe)
    if err != nil {
        return err
    }
    if used {
        return errors.New("biến thể đang được sử dụng trong phiếu nhập")
    }
    return s.repo.DeleteBienThe(maBienThe)
}

func (s *BienTheSvc) CreateBienTheTheoMaHangHoa(bienThe *models.BienThe) error {
    // Kiểm tra Size phải là số lớn hơn 35
    // Kiểm tra MaHangHoa tồn tại
    exists, err := s.repo.ExistsHangHoa(bienThe.MaHangHoa)
    if err != nil {
        return err
    }
    if !exists {
        return errors.New("mã hàng hóa không tồn tại")
    }

    bienThe.TrangThai = "DangBan"

    sizeFloat, err := strconv.ParseFloat(bienThe.Size, 64)
    if err != nil {
        return errors.New("size phải là số")
    }

    if sizeFloat <= 0 {
        return errors.New("size phải lớn hơn 0")
    }

    bienThe.Size = strconv.FormatFloat(sizeFloat, 'f', -1, 64)

    existsBienThe, err := s.repo.ExistsBienTheByHangHoaAndSize(bienThe.MaHangHoa, bienThe.Size)
    if err != nil {
        return err
    }
    if existsBienThe {
        return errors.New("biến thể đã tồn tại")
    }

    // Kiểm tra Gia không nhỏ hơn 0
    if bienThe.Gia < 0 {
        return errors.New("giá không được nhỏ hơn 0")
    }


    return s.repo.CreateBienTheTheoMaHangHoa(bienThe)
}

func (s *BienTheSvc) UpdateBienTheInfo(bienThe *models.BienThe) error {
    // Kiểm tra MaHangHoa tồn tại
    exists, err := s.repo.ExistsHangHoa(bienThe.MaHangHoa)
    if err != nil {
        return err
    }
    if !exists {
        return errors.New("mã hàng hóa không tồn tại")
    }

    sizeFloat, err := strconv.ParseFloat(bienThe.Size, 64)
    if err != nil {
        return errors.New("size phải là số")
    }
    if sizeFloat <= 0 {
        return errors.New("size phải lớn hơn 0")
    }

    // Kiểm tra Gia không nhỏ hơn 0
    if bienThe.Gia < 0 {
        return errors.New("giá không được nhỏ hơn 0")
    }

    return s.repo.UpdateBienTheInfo(bienThe)
}

func (s *BienTheSvc) UpdateBienTheStatus(maBienThe int, trangThai string) error {
    if trangThai != "DangBan" && trangThai != "NgungBan" {
        return errors.New("trạng thái không hợp lệ")
    }

    _, err := s.repo.GetBienTheTheoMa(maBienThe)
    if err != nil {
        return errors.New("biến thể không tồn tại")
    }

    return s.repo.UpdateBienTheStatus(maBienThe, trangThai)
}