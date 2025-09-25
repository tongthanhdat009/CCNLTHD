package services

import (
    "github.com/tongthanhdat009/CCNLTHD/internal/models"
    "github.com/tongthanhdat009/CCNLTHD/internal/repositories"
	"errors"
	"strings"
	"golang.org/x/crypto/bcrypt"
    "github.com/golang-jwt/jwt/v5"
    "time"
)
const (
	JwtSecretKey         = "bismaatj"
	RefreshTokenDuration  = time.Hour * 24 * 7 // 7 ngày
	AccessTokenDuration   = time.Minute * 15    // 15 phút
)

type DangNhapService interface {
    KiemTraDangNhap(tenTaiKhoan string, matKhau string) (*models.NguoiDung, string, string, error)
}
type dangNhapService struct {
    repo repositories.NguoiDungRepository
}
func NewDangNhapService(repo repositories.NguoiDungRepository) DangNhapService {
	return &dangNhapService{repo: repo}
}
func (s *dangNhapService) KiemTraDangNhap(tenTaiKhoan string, matKhau string) (*models.NguoiDung, string, string, error) {
    // Kiểm tra tên đăng nhập
    tenTaiKhoan = strings.TrimSpace(tenTaiKhoan)
    if tenTaiKhoan == "" {
        return nil, "", "", errors.New("tên đăng nhập không được để trống")
    }

    // Kiểm tra mật khẩu
    matKhau = strings.TrimSpace(matKhau)
    if matKhau == "" {
        return nil, "", "", errors.New("mật khẩu không được để trống")
    }

    // Lấy người dùng
    nguoiDung, err := s.repo.KiemTraDangNhap(tenTaiKhoan)
    if err != nil {
        return nil, "", "", errors.New("tên đăng nhập hoặc mật khẩu không đúng")
    }

    // Kiểm tra mật khẩu
    if err := bcrypt.CompareHashAndPassword([]byte(nguoiDung.MatKhau), []byte(matKhau)); err != nil {
        return nil, "", "", errors.New("tên đăng nhập hoặc mật khẩu không đúng")
    }

	chucNangs, err := s.repo.LayChucNangTheoMaQuyen(nguoiDung.MaQuyen)
    if err != nil {
        return nil, "", "", err
    }
    
    nguoiDung.Quyen.ChucNangs = chucNangs

    rfToken, err := s.GenerateRefreshToken(nguoiDung.MaNguoiDung)
    if err != nil {
        return nil, "", "", err
    }

    err = s.repo.CreateRefreshToken(rfToken)
    if err != nil {
        return nil, "", "", err
    }

    accessToken, err := s.GenerateAccessToken(nguoiDung.MaNguoiDung, nguoiDung.Quyen.TenQuyen)
    if err != nil {
        return nil, "", "", err
    }

    return nguoiDung, accessToken, rfToken.Token, nil
}

func (s *dangNhapService) GenerateRefreshToken(maNguoiDung int) (models.RefreshToken, error) {
	// Tạo payload token
	claims := jwt.MapClaims{
		"ma_nguoi_dung": maNguoiDung,
		"exp":     time.Now().Add(RefreshTokenDuration).Unix(),
		"type":    "refresh",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JwtSecretKey))
	if err != nil {
		return models.RefreshToken{}, err
	}

	// Tạo bản ghi RefreshToken
	refreshToken := models.RefreshToken{
		MaNguoiDung: maNguoiDung,
		Token:       tokenString,
		NgayTao:     time.Now(),
		NgayHetHan:  time.Now().Add(RefreshTokenDuration),
		TrangThai:   "Hoạt Động",
	}

	return refreshToken, nil
}
func (s *dangNhapService) GenerateAccessToken(maNguoiDung int, maQuyen string) (string, error) {

	// Tạo payload token
	claims := jwt.MapClaims{
		"ma_nguoi_dung": maNguoiDung,
		"role":          maQuyen,
		"exp":          time.Now().Add(AccessTokenDuration).Unix(),
		"type":         "access",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(JwtSecretKey))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}