package repositories

import (
    "github.com/tongthanhdat009/CCNLTHD/internal/models"
    "gorm.io/gorm"
)

type NguoiDungRepository interface {
    GetAll() ([]models.NguoiDung, error)
    Create(nguoiDung models.NguoiDung) error
    Update(maNguoiDung int, nguoiDung models.NguoiDung) error
    CheckNameExists(name string) (bool, error)
    FindQuyenKhachHang() (int, error)
    KiemTraDangNhap(tenDangNhap string) (*models.NguoiDung, error)
    LayChucNangTheoMaQuyen(maQuyen int) ([]models.ChucNang, error)
    CreateRefreshToken(refreshToken models.RefreshToken) error
    GetNguoiDungByID(maNguoiDung int) (*models.NguoiDung, error)
}

type NguoiDungRepo struct {
    db *gorm.DB
}

func NewNguoiDungRepository(db *gorm.DB) NguoiDungRepository {
    return &NguoiDungRepo{db: db}
}

func (r *NguoiDungRepo) GetAll() ([]models.NguoiDung, error) {
    var NguoiDungs []models.NguoiDung
    // Sử dụng Preload để tải các dữ liệu liên quan
    err := r.db.Preload("Quyen").Find(&NguoiDungs).Error
    return NguoiDungs, err
}
func (r *NguoiDungRepo) Create(nguoiDung models.NguoiDung) error {
    return r.db.Create(&nguoiDung).Error
}
func (r *NguoiDungRepo) Update(maNguoiDung int, nguoiDung models.NguoiDung) error {
    return r.db.Model(&models.NguoiDung{}).
        Where("MaNguoiDung = ?", maNguoiDung).
        Updates(nguoiDung).Error
}
func (r *NguoiDungRepo) CheckNameExists(name string) (bool, error) {
    var count int64
    err := r.db.Model(&models.NguoiDung{}).Where("TenDangNhap = ?", name).Count(&count).Error
    if err != nil {
        return false, err
    }
    return count > 0, nil
}
func (r *NguoiDungRepo) FindQuyenKhachHang() (int, error) {
    var quyen models.Quyen
    err := r.db.Where("TenQuyen = ?", "Khách hàng").First(&quyen).Error
    if err != nil {
        return 0, err
    }
    return quyen.MaQuyen, nil
}
// đăng nhập
func (r *NguoiDungRepo) KiemTraDangNhap(tenDangNhap string) (*models.NguoiDung, error) {
    var nguoiDung models.NguoiDung
    err := r.db.
        Preload("Quyen").
        Where("TenDangNhap = ?", tenDangNhap).
        First(&nguoiDung).Error
    if err != nil {
        return nil, err
    }
    return &nguoiDung, nil
}
// Lấy chức năng theo mã quyền
func (r *NguoiDungRepo) LayChucNangTheoMaQuyen(maQuyen int) ([]models.ChucNang, error) {
    var chucNangs []models.ChucNang

    err := r.db.
        Joins("JOIN chitietchucnang ON chitietchucnang.MaChucNang = chucnang.MaChucNang").
        Joins("JOIN phanquyen ON phanquyen.MaChiTietChucNang = chitietchucnang.MaChiTietChucNang").
        Where("phanquyen.MaQuyen = ? AND phanquyen.TrangThai = ?", maQuyen, "Mở").
        Preload("ChiTietChucNangs").
        Group("chucnang.MaChucNang").
        Find(&chucNangs).Error

    if err != nil {
        return nil, err
    }
    return chucNangs, nil
}
func (r *NguoiDungRepo) CreateRefreshToken(refreshToken models.RefreshToken) error {
    return r.db.Create(&refreshToken).Error
}
func (r *NguoiDungRepo) GetNguoiDungByID(maNguoiDung int) (*models.NguoiDung, error) {
    var nguoiDung models.NguoiDung
    err := r.db.Preload("Quyen").First(&nguoiDung, maNguoiDung).Error
    if err != nil {
        return nil, err
    }
    return &nguoiDung, nil
}
