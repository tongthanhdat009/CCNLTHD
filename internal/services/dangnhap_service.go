package services

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"github.com/tongthanhdat009/CCNLTHD/internal/repositories"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const (
	RefreshTokenDuration = time.Hour * 24 * 7 // 7 ngày
	AccessTokenDuration  = time.Minute * 100000   // 100000 phút
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
	tenTaiKhoan = strings.TrimSpace(tenTaiKhoan)
	if tenTaiKhoan == "" {
		return nil, "", "", errors.New("tên đăng nhập không được để trống")
	}

	matKhau = strings.TrimSpace(matKhau)
	if matKhau == "" {
		return nil, "", "", errors.New("mật khẩu không được để trống")
	}

	nguoiDung, err := s.repo.KiemTraDangNhap(tenTaiKhoan)
	if err != nil {
		return nil, "", "", errors.New("tên đăng nhập hoặc mật khẩu không đúng")
	}

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

	accessToken, err := s.GenerateAccessToken(nguoiDung.MaNguoiDung, nguoiDung.MaQuyen)
	if err != nil {
		return nil, "", "", err
	}

	err = s.repo.CreateRefreshToken(rfToken)
	if err != nil {
		return nil, "", "", err
	}

	return nguoiDung, accessToken, rfToken.Token, nil
}

func (s *dangNhapService) GenerateRefreshToken(maNguoiDung int) (models.RefreshToken, error) {
	claims := jwt.MapClaims{
		"ma_nguoi_dung": maNguoiDung,
		"exp":           time.Now().Add(RefreshTokenDuration).Unix(),
		"type":          "refresh",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return models.RefreshToken{}, errors.New("JWT_SECRET not set in environment variables")
	}
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return models.RefreshToken{}, err
	}

	return models.RefreshToken{
		MaNguoiDung: maNguoiDung,
		Token:       tokenString,
		NgayTao:     time.Now(),
		NgayHetHan:  time.Now().Add(RefreshTokenDuration),
		TrangThai:   "Hoạt Động",
	}, nil
}

func (s *dangNhapService) GenerateAccessToken(maNguoiDung int, maQuyen int) (string, error) {
	claims := jwt.MapClaims{
		"ma_nguoi_dung": maNguoiDung,
		"ma_quyen":      maQuyen,
		"exp":           time.Now().Add(AccessTokenDuration).Unix(),
		"type":          "access",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("JWT_SECRET not set in environment variables")
	}
	return token.SignedString([]byte(secret))
}
