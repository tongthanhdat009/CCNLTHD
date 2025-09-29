package repositories

import (
    "errors"
    "github.com/tongthanhdat009/CCNLTHD/internal/models"
    "gorm.io/gorm"
)

type TimKiemSanPhamRepository interface {
    TimSanPham(
        tenHangHoa string,
        tenHang string,
        tenDanhMuc string,
        mau string,
        size string,
        giaToiThieu string,
        giaToiDa string,
    ) ([]models.SanPham, error)
}

type timKiemSanPhamRepository struct {
    db *gorm.DB
}

func NewTimKiemSanPhamRepository(db *gorm.DB) TimKiemSanPhamRepository {
    return &timKiemSanPhamRepository{db: db}
}

func (r *timKiemSanPhamRepository) TimSanPham(
    tenHangHoa string,
    tenHang string,
    tenDanhMuc string,
    mau string,
    size string,
    giaToiThieu string,
    giaToiDa string,
) ([]models.SanPham, error) {
    var sanphams []models.SanPham

    query := r.db.Model(&models.SanPham{}).
        Preload("ChiTietPhieuNhap").
        Preload("ChiTietPhieuNhap.BienThe").
        Preload("ChiTietPhieuNhap.BienThe.HangHoa").
        Preload("ChiTietPhieuNhap.BienThe.HangHoa.Hang").
        Preload("ChiTietPhieuNhap.BienThe.HangHoa.DanhMuc").
        Preload("ChiTietPhieuNhap.BienThe.HangHoa.KhuyenMai").
        Preload("ChiTietPhieuNhap.PhieuNhap").
        Preload("ChiTietPhieuNhap.PhieuNhap.NguoiDung").
        Preload("ChiTietPhieuNhap.PhieuNhap.NhaCungCap")

    needJoin := tenHangHoa != "" || tenHang != "" || tenDanhMuc != "" || mau != "" || size != "" || (giaToiThieu != "" && giaToiDa != "")

    if needJoin {
        query = query.Joins("JOIN chitietphieunhap ON sanpham.MaChiTietPhieuNhap = chitietphieunhap.MaChiTiet").
            Joins("JOIN bienthe ON chitietphieunhap.MaBienthe = bienthe.MaBienThe").
            Joins("JOIN hanghoa ON bienthe.MaHangHoa = hanghoa.MaHangHoa").
            Joins("JOIN hang ON hanghoa.MaHang = hang.MaHang").
            Joins("JOIN danhmuc ON hanghoa.MaDanhMuc = danhmuc.MaDanhMuc")
    }

    if tenHangHoa != "" {
        query = query.Where("hanghoa.TenHangHoa LIKE ?", "%"+tenHangHoa+"%")
    }
    if tenHang != "" {
        query = query.Where("hang.TenHang LIKE ?", "%"+tenHang+"%")
    }
    if tenDanhMuc != "" {
        query = query.Where("danhmuc.TenDanhMuc LIKE ?", "%"+tenDanhMuc+"%")
    }
    if mau != "" {
        query = query.Where("hanghoa.Mau LIKE ?", "%"+mau+"%")
    }
    if size != "" {
        query = query.Where("bienthe.Size LIKE ?", "%"+size+"%")
    }
    if giaToiThieu != "" && giaToiDa != "" {
        query = query.Where("bienthe.Gia BETWEEN ? AND ?", giaToiThieu, giaToiDa)
    } else if giaToiThieu != "" {
        query = query.Where("bienthe.Gia >= ?", giaToiThieu)
    } else if giaToiDa != "" {
        query = query.Where("bienthe.Gia <= ?", giaToiDa)
    }

    if err := query.Find(&sanphams).Error; err != nil {
        return nil, errors.New("Lỗi khi tìm kiếm sản phẩm: " + err.Error())
    }
    return sanphams, nil
}