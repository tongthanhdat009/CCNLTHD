package services

import (
    "github.com/tongthanhdat009/CCNLTHD/internal/models"
    "github.com/tongthanhdat009/CCNLTHD/internal/repositories"
    "errors"
)

type HangHoaService interface {
    GetAllHangHoa() ([]models.HangHoa, error)
    CreateHangHoa(hangHoa *models.HangHoa) error
    UpdateHangHoa(hangHoa *models.HangHoa) error
    SearchHangHoa(
        tenHangHoa string,
        tenDanhMuc string,
        tenHang string,
        mau string,
        trangThai string,
        maKhuyenMai string,
    ) ([]models.HangHoa, error)
    GetHangHoaByID(maHangHoa int) (*models.HangHoa, error)
}

type hangHoaService struct {
    repo repositories.HangHoaRepository
}

func NewHangHoaService(repo repositories.HangHoaRepository) HangHoaService {
    return &hangHoaService{repo: repo}
}

func (s *hangHoaService) GetHangHoaByID(maHangHoa int) (*models.HangHoa, error) {
    return s.repo.GetHangHoaByID(maHangHoa)
}

func (s *hangHoaService) GetAllHangHoa() ([]models.HangHoa, error) {
    return s.repo.GetAll()
}

func (s *hangHoaService) CreateHangHoa(hangHoa *models.HangHoa) error {
    // Kiểm tra các trường bắt buộc
    if hangHoa.TenHangHoa == "" ||
        hangHoa.MaHang == 0 ||
        hangHoa.MaDanhMuc == 0 ||
        hangHoa.Mau == ""{
        return errors.New("tên hàng hóa, mã hãng, mã danh mục, màu sắc không được để trống")
    }

    // Kiểm tra mã hãng có tồn tại không
    exists, err := s.repo.ExistsHang(hangHoa.MaHang)
    if err != nil {
        return err
    }
    if !exists {
        return errors.New("mã hãng không tồn tại")
    }

    // Kiểm tra mã danh mục có tồn tại không
    exists, err = s.repo.ExistsDanhMuc(hangHoa.MaDanhMuc)
    if err != nil {
        return err
    }
    if !exists {
        return errors.New("mã danh mục không tồn tại")
    }

    // Kiểm tra mã khuyến mãi nếu có
    if hangHoa.MaKhuyenMai.Valid {
        exists, err = s.repo.ExistsKhuyenMai(hangHoa.MaKhuyenMai.Int64)
        if err != nil {
            return err
        }
        if !exists {
            return errors.New("mã khuyến mãi không tồn tại")
        }
    }

    if hangHoa.AnhDaiDien == "" {
        hangHoa.AnhDaiDien = "no-product.png"
    }

    return s.repo.CreateHangHoa(hangHoa)
}

func (s *hangHoaService) UpdateHangHoa(hangHoa *models.HangHoa) error {
    // Kiểm tra các trường bắt buộc
    if hangHoa.TenHangHoa == "" ||
        hangHoa.MaHang == 0 ||
        hangHoa.MaDanhMuc == 0 ||
        hangHoa.Mau == "" {
        return errors.New("tên hàng hóa, mã hãng, mã danh mục, màu sắc và ảnh đại diện không được để trống")
    }

    // Kiểm tra mã hãng có tồn tại không
    exists, err := s.repo.ExistsHang(hangHoa.MaHang)
    if err != nil {
        return err
    }
    if !exists {
        return errors.New("mã hãng không tồn tại")
    }

    // Kiểm tra mã danh mục có tồn tại không
    exists, err = s.repo.ExistsDanhMuc(hangHoa.MaDanhMuc)
    if err != nil {
        return err
    }
    if !exists {
        return errors.New("mã danh mục không tồn tại")
    }

    // Kiểm tra mã khuyến mãi nếu có
    if hangHoa.MaKhuyenMai.Valid {
        exists, err = s.repo.ExistsKhuyenMai(hangHoa.MaKhuyenMai.Int64)
        if err != nil {
            return err
        }
        if !exists {
            return errors.New("mã khuyến mãi không tồn tại")
        }
    }

    return s.repo.UpdateHangHoa(hangHoa)
}

func (s *hangHoaService) SearchHangHoa(
    tenHangHoa string,
    tenDanhMuc string,
    tenHang string,
    mau string,
    trangThai string,
    maKhuyenMai string,
) ([]models.HangHoa, error) {
    return s.repo.SearchHangHoa(tenHangHoa, tenDanhMuc, tenHang, mau, trangThai, maKhuyenMai)
}