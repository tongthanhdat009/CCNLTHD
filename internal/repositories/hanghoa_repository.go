package repositories

import (
    "github.com/tongthanhdat009/CCNLTHD/internal/models"
    "gorm.io/gorm"
)

type HangHoaRepository interface {
    GetAll() ([]models.HangHoa, error)
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
    CountByHangID(hangID int) (int64, error)
    ExistsHang(maHang int) (bool, error)
    ExistsDanhMuc(maDanhMuc int) (bool, error)
    ExistsKhuyenMai(maKhuyenMai int64) (bool, error)
}

type hangHoaRepo struct {
    db *gorm.DB
}

func NewHangHoaRepository(db *gorm.DB) HangHoaRepository {
    return &hangHoaRepo{db: db}
}

func (r *hangHoaRepo) GetHangHoaByID(maHangHoa int) (*models.HangHoa, error) {
    var hangHoa models.HangHoa
    err := r.db.Preload("Hang").Preload("DanhMuc").Preload("BienThes").First(&hangHoa, maHangHoa).Error
    if err != nil {
        return nil, err
    }
    return &hangHoa, nil
}

func (r *hangHoaRepo) GetAll() ([]models.HangHoa, error) {
    var hangHoas []models.HangHoa
    err := r.db.Preload("Hang").Preload("DanhMuc").Preload("BienThes").Find(&hangHoas).Error
    return hangHoas, err
}

func (r *hangHoaRepo) CountByHangID(hangID int) (int64, error) {
    var count int64
    err := r.db.Model(&models.HangHoa{}).Where("mahang = ?", hangID).Count(&count).Error
    return count, err
}

func (r *hangHoaRepo) CreateHangHoa(hangHoa *models.HangHoa) error {
    return r.db.Create(hangHoa).Error
}

func (r *hangHoaRepo) UpdateHangHoa(hangHoa *models.HangHoa) error {
    return r.db.Save(hangHoa).Error
}

// Kiểm tra mã hãng có tồn tại không
func (r *hangHoaRepo) ExistsHang(maHang int) (bool, error) {
    var count int64
    err := r.db.Model(&models.Hang{}).Where("MaHang = ?", maHang).Count(&count).Error
    return count > 0, err
}

// Kiểm tra mã danh mục có tồn tại không
func (r *hangHoaRepo) ExistsDanhMuc(maDanhMuc int) (bool, error) {
    var count int64
    err := r.db.Model(&models.DanhMuc{}).Where("MaDanhMuc = ?", maDanhMuc).Count(&count).Error
    return count > 0, err
}

// Kiểm tra mã khuyến mãi có tồn tại không
func (r *hangHoaRepo) ExistsKhuyenMai(maKhuyenMai int64) (bool, error) {
    var count int64
    err := r.db.Model(&models.KhuyenMai{}).Where("MaKhuyenMai = ?", maKhuyenMai).Count(&count).Error
    return count > 0, err
}

func (r *hangHoaRepo) SearchHangHoa(
    tenHangHoa string,
    tenDanhMuc string,
    tenHang string,
    mau string,
    trangThai string,
    maKhuyenMai string,
) ([]models.HangHoa, error) {
    var hangHoas []models.HangHoa

    query := r.db.Model(&models.HangHoa{}).
        Preload("Hang").
        Preload("DanhMuc").
        Preload("KhuyenMai").
        Preload("BienThes")

    // Join với bảng liên quan nếu có điều kiện tìm kiếm
    if tenDanhMuc != "" {
        query = query.Joins("JOIN danhmuc ON hanghoa.MaDanhMuc = danhmuc.MaDanhMuc")
    }
    if tenHang != "" {
        query = query.Joins("JOIN hang ON hanghoa.MaHang = hang.MaHang")
    }
    if maKhuyenMai != "" {
        query = query.Joins("JOIN khuyenmai ON hanghoa.MaKhuyenMai = khuyenmai.MaKhuyenMai")
    }

    // Điều kiện tìm kiếm
    if tenHangHoa != "" {
        query = query.Where("hanghoa.TenHangHoa LIKE ?", "%"+tenHangHoa+"%")
    }
    if tenDanhMuc != "" {
        query = query.Where("danhmuc.TenDanhMuc LIKE ?", "%"+tenDanhMuc+"%")
    }
    if tenHang != "" {
        query = query.Where("hang.TenHang LIKE ?", "%"+tenHang+"%")
    }
    if mau != "" {
        query = query.Where("hanghoa.Mau LIKE ?", "%"+mau+"%")
    }
    if trangThai != "" {
        query = query.Where("hanghoa.TrangThai = ?", trangThai)
    }
    if maKhuyenMai != "" {
        query = query.Where("hanghoa.MaKhuyenMai = ?", maKhuyenMai)
    }

    if err := query.Find(&hangHoas).Error; err != nil {
        return nil, err
    }
    return hangHoas, nil
}