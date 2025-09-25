package services

import (
    "github.com/tongthanhdat009/CCNLTHD/internal/models"
    "github.com/tongthanhdat009/CCNLTHD/internal/repositories"
	"errors"
	"regexp"
	"strings"
	"golang.org/x/crypto/bcrypt"
)

type DangKyService interface {
    CreateNguoiDung(nguoiDung models.NguoiDung) error
}
type dangKyService struct {
    repo repositories.NguoiDungRepository
}
func NewDangKyService(repo repositories.NguoiDungRepository) DangKyService {
	return &dangKyService{repo: repo}
}
func validateDangKy(nguoiDung models.NguoiDung) error {
	// Trim khoảng trắng
	nguoiDung.TenDangNhap = strings.TrimSpace(nguoiDung.TenDangNhap)
	nguoiDung.MatKhau = strings.TrimSpace(nguoiDung.MatKhau)
	nguoiDung.HoTen = strings.TrimSpace(nguoiDung.HoTen)
	nguoiDung.Email = strings.TrimSpace(nguoiDung.Email)
	nguoiDung.SoDienThoai = strings.TrimSpace(nguoiDung.SoDienThoai)

	// --- Kiểm tra bắt buộc ---
	if nguoiDung.TenDangNhap == "" {
		return errors.New("ten dang nhap la bat buoc")
	}
	if nguoiDung.MatKhau == "" {
		return errors.New("mat khau la bat buoc")
	}
	if nguoiDung.HoTen == "" {
		return errors.New("ho ten la bat buoc")
	}
	if nguoiDung.Email == "" {
		return errors.New("email la bat buoc")
	}

	// --- Validate chi tiết ---
	// Tên đăng nhập >= 4 ký tự, không có khoảng trắng
	if len(nguoiDung.TenDangNhap) < 4 {
		return errors.New("ten dang nhap phai co it nhat 4 ky tu")
	}
	if strings.Contains(nguoiDung.TenDangNhap, " ") {
		return errors.New("ten dang nhap khong duoc chua khoang trang")
	}

	// Mật khẩu >= 6 ký tự
	if len(nguoiDung.MatKhau) < 6 {
		return errors.New("mat khau phai co it nhat 6 ky tu")
	}

	// Email hợp lệ
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(nguoiDung.Email) {
		return errors.New("email khong hop le")
	}
	return nil
}


func (s *dangKyService) CreateNguoiDung(nguoiDung models.NguoiDung) error {
    // Validate dữ liệu
    if err := validateDangKy(nguoiDung); err != nil {
        return err
    }

    // Kiểm tra tên người dùng đã tồn tại chưa
    exists, err := s.repo.CheckNameExists(nguoiDung.TenDangNhap)
    if err != nil {
        return err
    }
    if exists {
        return errors.New("tên người dùng đã tồn tại")
    }

    // Hash mật khẩu bằng bcrypt
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(nguoiDung.MatKhau), bcrypt.DefaultCost)
    if err != nil {
        return errors.New("không thể mã hóa mật khẩu")
    }
    nguoiDung.MatKhau = string(hashedPassword)

    // Gán quyền mặc định = khách hàng
    nguoiDung.MaQuyen, _ = s.repo.FindQuyenKhachHang()

    // Tạo người dùng mới
    return s.repo.Create(nguoiDung)
}