package services

import (
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"github.com/tongthanhdat009/CCNLTHD/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type NguoiDungService interface {
    GetAllNguoiDung() ([]models.NguoiDung, error)
    GetNguoiDungByID(maNguoiDung int) (*models.NguoiDung, error)
    UpdateNguoiDung(maNguoiDung int, nguoiDung models.NguoiDung) error
    UpdateNguoiDungAdmin(maNguoiDung int, nguoiDung models.NguoiDung) error
    ValidateNguoiDung(nguoiDung models.NguoiDung) error
    CreateNguoiDung(nguoiDung *models.NguoiDung) error
}

type nguoiDungService struct {
    repo repositories.NguoiDungRepository
}

func NewNguoiDungService(repo repositories.NguoiDungRepository) NguoiDungService {
    return &nguoiDungService{repo: repo}
}

func (s *nguoiDungService) GetAllNguoiDung() ([]models.NguoiDung, error) {
    return s.repo.GetAll()
}

func (s *nguoiDungService) GetNguoiDungByID(maNguoiDung int) (*models.NguoiDung, error) {
    return s.repo.GetNguoiDungByID(maNguoiDung)
}

func (s *nguoiDungService) UpdateNguoiDung(maNguoiDung int, nd models.NguoiDung) error {
    // Lấy người dùng cũ
    nguoiDungCu, err := s.repo.GetNguoiDungByID(maNguoiDung)
    if err != nil {
        return err
    }

    // Cập nhật các trường
    nguoiDungCu.HoTen = nd.HoTen
    nguoiDungCu.Email = nd.Email
    nguoiDungCu.SoDienThoai = nd.SoDienThoai

    // Xử lý NULL cho sql.NullString
    if strings.TrimSpace(nd.TinhThanh.String) == "" {
        nguoiDungCu.TinhThanh = sql.NullString{String: "", Valid: false}
    } else {
        nguoiDungCu.TinhThanh = sql.NullString{String: nd.TinhThanh.String, Valid: true}
    }

    if strings.TrimSpace(nd.QuanHuyen.String) == "" {
        nguoiDungCu.QuanHuyen = sql.NullString{String: "", Valid: false}
    } else {
        nguoiDungCu.QuanHuyen = sql.NullString{String: nd.QuanHuyen.String, Valid: true}
    }

    if strings.TrimSpace(nd.PhuongXa.String) == "" {
        nguoiDungCu.PhuongXa = sql.NullString{String: "", Valid: false}
    } else {
        nguoiDungCu.PhuongXa = sql.NullString{String: nd.PhuongXa.String, Valid: true}
    }

    if strings.TrimSpace(nd.DuongSoNha.String) == "" {
        nguoiDungCu.DuongSoNha = sql.NullString{String: "", Valid: false}
    } else {
        nguoiDungCu.DuongSoNha = sql.NullString{String: nd.DuongSoNha.String, Valid: true}
    }

    if err := s.ValidateNguoiDung(*nguoiDungCu); err != nil {
        return err
    }

    return s.repo.Update(maNguoiDung, *nguoiDungCu)
}

func (s *nguoiDungService) UpdateNguoiDungAdmin(maNguoiDung int, nd models.NguoiDung) error {
    // Lấy người dùng cũ

    nguoiDungCu, err := s.repo.GetNguoiDungByID(maNguoiDung)
    if err != nil {
        return err
    }
    // Cập nhật các trường
    nguoiDungCu.HoTen = nd.HoTen
    nguoiDungCu.Email = nd.Email
    nguoiDungCu.SoDienThoai = nd.SoDienThoai

    // Xử lý NULL cho sql.NullString
    if strings.TrimSpace(nd.TinhThanh.String) == "" {
        nguoiDungCu.TinhThanh = sql.NullString{String: "", Valid: false}
    } else {
        nguoiDungCu.TinhThanh = sql.NullString{String: nd.TinhThanh.String, Valid: true}
    }

    if strings.TrimSpace(nd.QuanHuyen.String) == "" {
        nguoiDungCu.QuanHuyen = sql.NullString{String: "", Valid: false}
    } else {
        nguoiDungCu.QuanHuyen = sql.NullString{String: nd.QuanHuyen.String, Valid: true}
    }

    if strings.TrimSpace(nd.PhuongXa.String) == "" {
        nguoiDungCu.PhuongXa = sql.NullString{String: "", Valid: false}
    } else {
        nguoiDungCu.PhuongXa = sql.NullString{String: nd.PhuongXa.String, Valid: true}
    }

    if strings.TrimSpace(nd.DuongSoNha.String) == "" {
        nguoiDungCu.DuongSoNha = sql.NullString{String: "", Valid: false}
    } else {
        nguoiDungCu.DuongSoNha = sql.NullString{String: nd.DuongSoNha.String, Valid: true}
    }

    // Cập nhật MaQuyen nếu khác 0
    if nd.MaQuyen != 0 {
        nguoiDungCu.MaQuyen = nd.MaQuyen
    }

    if err := s.ValidateNguoiDung(*nguoiDungCu); err != nil {
        return err
    }
    return s.repo.Update(maNguoiDung, *nguoiDungCu)
}


func (s *nguoiDungService) CreateNguoiDung(nguoiDung *models.NguoiDung) error {
    // Kiểm tra trường bắt buộc
    if err := s.ValidateNguoiDung(*nguoiDung); err != nil {
        return err
    }
    if nguoiDung.MaQuyen == 0 {
        // Gán quyền "Khách hàng" nếu không có quyền nào được chỉ định
        maQuyen, err := s.repo.FindQuyenKhachHang()
        if err != nil {
            return errors.New("không tìm thấy quyền 'Khách hàng'")
        }
        nguoiDung.MaQuyen = maQuyen
    }
    if exists, err := s.repo.CheckNameExists(nguoiDung.TenDangNhap); err != nil {
        return err
    } else if exists {
        return errors.New("tên đăng nhập đã tồn tại")
    }
    // Mã hóa mật khẩu
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(nguoiDung.MatKhau), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    nguoiDung.MatKhau = string(hashedPassword)

    return s.repo.Create(nguoiDung)
}

func (s *nguoiDungService) ValidateNguoiDung(nd models.NguoiDung) error {
    fmt.Print("Validating user: ", nd)
    // Kiểm tra trường bắt buộc
    if nd.TenDangNhap == "" {
        return errors.New("tên đăng nhập là bắt buộc")
    }
    if nd.MatKhau == "" {
        return errors.New("mật khẩu là bắt buộc")
    }
    if nd.HoTen == "" {
        return errors.New("họ tên là bắt buộc")
    }
    if nd.Email == "" {
        return errors.New("email là bắt buộc")
    }

    // Validate chi tiết
    if len(nd.TenDangNhap) < 4 {
        return errors.New("tên đăng nhập phải có ít nhất 4 ký tự")
    }
    if strings.Contains(nd.TenDangNhap, " ") {
        return errors.New("tên đăng nhập không được chứa khoảng trắng")
    }
    if len(nd.MatKhau) < 6 {
        return errors.New("mật khẩu phải có ít nhất 6 ký tự")
    }
    emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
    if !emailRegex.MatchString(nd.Email) {
        return errors.New("email không hợp lệ")
    }
    if nd.SoDienThoai != "" {
        phoneRegex := regexp.MustCompile(`^(0|\+84)(3[2-9]|5[6|8|9]|7[0|6-9]|8[1-5]|9[0-4|6-9])[0-9]{7}$`)
        if !phoneRegex.MatchString(nd.SoDienThoai) {
            return errors.New("số điện thoại không hợp lệ")
        }
    }

    return nil
}

