package repositories

import (
	"errors"
	"fmt"
	"time"

	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"gorm.io/gorm"
)

// PhieuNhapRepository định nghĩa interface cho các thao tác với phiếu nhập
type PhieuNhapRepository interface {
    GetAll() ([]models.PhieuNhap, error)
    CreatePhieuNhap(phieuNhap *models.PhieuNhap) error
    ExistsInPhieuNhap(nhaCungCapID int) (bool, error)
    ExistsNhaCungCap(maNCC int) (bool, error)
    ExistsNguoiDung(maND int) (bool, error)
    GetPhieuNhapByID(id int) (*models.PhieuNhap, error)
    DeletePhieuNhap(id int) error
    UpdatePhieuNhap(phieuNhap *models.PhieuNhap, createProducts bool) error
    SearchPhieuNhap(tenNguoiDung, tenNhaCungCap, trangThai string, tuNgay, denNgay *time.Time) ([]models.PhieuNhap, error)
    CreateChiTietPhieuNhap(maPhieuNhap int, items []models.ChiTietPhieuNhap) error
    GetChiTietByPhieuNhap(maPhieuNhap int) ([]models.ChiTietPhieuNhap, error)
    DeleteChiTietByPhieuNhap(maPhieuNhap int, maChiTietPhieuNhap int) error
    DeleteAllChiTietByPhieuNhap(maPhieuNhap int) error
    UpdateChiTietPhieuNhapSoLuong(maPhieuNhap int, maChiTiet int, soLuong int) error
    ExistsChiTietPhieuNhap(maChiTiet int, maPhieuNhap int) (bool, error)
}

type phieuNhapRepository struct {
    db *gorm.DB
}

// NewPhieuNhapRepository tạo mới một repository cho phiếu nhập
func NewPhieuNhapRepository(db *gorm.DB) PhieuNhapRepository {
    return &phieuNhapRepository{db: db}
}

// GetAll lấy tất cả phiếu nhập kèm thông tin liên quan
func (r *phieuNhapRepository) GetAll() ([]models.PhieuNhap, error) {
    var phieuNhaps []models.PhieuNhap
    if err := r.db.Preload("NhaCungCap").Preload("NguoiDung").Find(&phieuNhaps).Error; err != nil {
        return nil, err
    }
    return phieuNhaps, nil
}

// GetPhieuNhapByID lấy phiếu nhập theo ID
func (r *phieuNhapRepository) GetPhieuNhapByID(id int) (*models.PhieuNhap, error) {
    var phieuNhap models.PhieuNhap
    result := r.db.Preload("NhaCungCap").Preload("NguoiDung").First(&phieuNhap, id)
    
    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
        return nil, nil // Trả về nil khi không tìm thấy
    }
    
    if result.Error != nil {
        return nil, result.Error
    }
    
    return &phieuNhap, nil
}

// CreatePhieuNhap tạo mới phiếu nhập
func (r *phieuNhapRepository) CreatePhieuNhap(phieuNhap *models.PhieuNhap) error {
    return r.db.Create(phieuNhap).Error
}

// DeletePhieuNhap xóa phiếu nhập theo ID
func (r *phieuNhapRepository) DeletePhieuNhap(id int) error {
    // Sử dụng transaction để đảm bảo tính nhất quán dữ liệu
    return r.db.Transaction(func(tx *gorm.DB) error {
        // Xóa chi tiết phiếu nhập trước
        if err := tx.Where("MaPhieuNhap = ?", id).Delete(&models.ChiTietPhieuNhap{}).Error; err != nil {
            return err
        }
        
        // Sau đó xóa phiếu nhập
        if err := tx.Delete(&models.PhieuNhap{}, id).Error; err != nil {
            return err
        }
        
        return nil
    })
}

// UpdatePhieuNhap cập nhật phiếu nhập và tạo sản phẩm nếu được chỉ định
func (r *phieuNhapRepository) UpdatePhieuNhap(phieuNhap *models.PhieuNhap, createProducts bool) error {
    return r.db.Transaction(func(tx *gorm.DB) error {
        // Cập nhật trạng thái phiếu nhập
        if err := tx.Model(&models.PhieuNhap{}).
            Where("MaPhieuNhap = ?", phieuNhap.MaPhieuNhap).
            Update("TrangThai", phieuNhap.TrangThai).Error; err != nil {
            return err
        }

        // Nếu không cần tạo sản phẩm, kết thúc ở đây
        if !createProducts {
            return nil
        }

        // Lấy chi tiết phiếu nhập
        var details []models.ChiTietPhieuNhap
        if err := tx.Where("MaPhieuNhap = ?", phieuNhap.MaPhieuNhap).Find(&details).Error; err != nil {
            return err
        }
        
        if len(details) == 0 {
            return errors.New("phiếu nhập chưa có chi tiết")
        }

        // Cập nhật số lượng tồn kho cho các biến thể
        for _, detail := range details {
            if detail.SoLuong <= 0 {
                continue
            }
            
            // Cập nhật số lượng tồn kho của biến thể
            if err := tx.Model(&models.BienThe{}).
                Where("MaBienThe = ?", detail.MaBienthe).
                UpdateColumn("SoLuongTon", gorm.Expr("SoLuongTon + ?", detail.SoLuong)).Error; err != nil {
                return err
            }
        }

        // Tạo mã sê-ri cho sản phẩm
        timeStamp := time.Now().UnixNano()
        var sanPhams []models.SanPham
        
        // Tạo danh sách sản phẩm từ chi tiết phiếu nhập
        for _, detail := range details {
            if detail.SoLuong <= 0 {
                continue
            }
            
            for i := 0; i < detail.SoLuong; i++ {
                sanPhams = append(sanPhams, models.SanPham{
                    MaChiTietPhieuNhap: detail.MaChiTiet,
                    Seri:               fmt.Sprintf("PN%d-%d-%d", detail.MaChiTiet, timeStamp, i+1),
                    TrangThai:          "Đang bán",
                })
            }
        }
        
        // Tạo sản phẩm nếu có
        if len(sanPhams) > 0 {
            // Tối ưu bằng cách sử dụng CreateInBatches nếu có nhiều sản phẩm
            const batchSize = 100
            if err := tx.CreateInBatches(&sanPhams, batchSize).Error; err != nil {
                return err
            }
        }
        
        return nil
    })
}

// SearchPhieuNhap tìm kiếm phiếu nhập theo nhiều tiêu chí
func (r *phieuNhapRepository) SearchPhieuNhap(tenNguoiDung, tenNhaCungCap, trangThai string, tuNgay, denNgay *time.Time) ([]models.PhieuNhap, error) {
    query := r.db.Model(&models.PhieuNhap{})
    
    // Thêm điều kiện tìm kiếm theo tên người dùng
    if tenNguoiDung != "" {
        query = query.Joins("JOIN nguoidung ON nguoidung.manguoidung = phieunhap.manguoidung").
                     Where("nguoidung.hoten LIKE ?", "%"+tenNguoiDung+"%")
    }
    
    // Thêm điều kiện tìm kiếm theo tên nhà cung cấp
    if tenNhaCungCap != "" {
        query = query.Joins("JOIN nhacungcap ON nhacungcap.mancc = phieunhap.mancc").
                     Where("nhacungcap.tenncc LIKE ?", "%"+tenNhaCungCap+"%")
    }
    
    // Thêm các điều kiện tìm kiếm khác
    if trangThai != "" {
        query = query.Where("phieunhap.trangthai = ?", trangThai)
    }
    
    if tuNgay != nil {
        query = query.Where("phieunhap.ngaynhap >= ?", tuNgay)
    }
    
    if denNgay != nil {
        query = query.Where("phieunhap.ngaynhap <= ?", denNgay)
    }

    // Thực hiện truy vấn và trả về kết quả
    var phieuNhaps []models.PhieuNhap
    if err := query.Preload("NhaCungCap").Preload("NguoiDung").
                   Order("phieunhap.ngaynhap DESC").
                   Find(&phieuNhaps).Error; err != nil {
        return nil, err
    }
    
    return phieuNhaps, nil
}

// ExistsInPhieuNhap kiểm tra nhà cung cấp có trong phiếu nhập không
func (r *phieuNhapRepository) ExistsInPhieuNhap(nhaCungCapID int) (bool, error) {
    var count int64
    if err := r.db.Model(&models.PhieuNhap{}).Where("mancc = ?", nhaCungCapID).Count(&count).Error; err != nil {
        return false, err
    }
    return count > 0, nil
}

// ExistsNhaCungCap kiểm tra nhà cung cấp có tồn tại không
func (r *phieuNhapRepository) ExistsNhaCungCap(maNCC int) (bool, error) {
    var count int64
    if err := r.db.Model(&models.NhaCungCap{}).Where("MaNCC = ?", maNCC).Count(&count).Error; err != nil {
        return false, err
    }
    return count > 0, nil
}

// ExistsNguoiDung kiểm tra người dùng có tồn tại không
func (r *phieuNhapRepository) ExistsNguoiDung(maND int) (bool, error) {
    var count int64
    if err := r.db.Model(&models.NguoiDung{}).Where("MaNguoiDung = ?", maND).Count(&count).Error; err != nil {
        return false, err
    }
    return count > 0, nil
}

// CreateChiTietPhieuNhap tạo mới chi tiết phiếu nhập
func (r *phieuNhapRepository) CreateChiTietPhieuNhap(maPhieuNhap int, items []models.ChiTietPhieuNhap) error {
    if len(items) == 0 {
        return nil
    }
    
    return r.db.Transaction(func(tx *gorm.DB) error {
        for i := range items {
            items[i].MaPhieuNhap = maPhieuNhap
        }
        return tx.Create(&items).Error
    })
}

// GetChiTietByPhieuNhap lấy chi tiết của một phiếu nhập
func (r *phieuNhapRepository) GetChiTietByPhieuNhap(maPhieuNhap int) ([]models.ChiTietPhieuNhap, error) {
    var details []models.ChiTietPhieuNhap
    if err := r.db.Preload("BienThe").
                    Preload("BienThe.HangHoa").
                    Preload("BienThe.HangHoa.Hang").
                    Preload("BienThe.HangHoa.DanhMuc").
                    Preload("BienThe.HangHoa.KhuyenMai").
                    Preload("PhieuNhap").
                    Preload("PhieuNhap.NguoiDung").
                    Preload("PhieuNhap.NhaCungCap").
                    Where("MaPhieuNhap = ?", maPhieuNhap).Find(&details).Error; err != nil {
        return nil, err
    }
    return details, nil
}

// DeleteChiTietByPhieuNhap xóa tất cả chi tiết của một phiếu nhập
func (r *phieuNhapRepository) DeleteChiTietByPhieuNhap(maPhieuNhap int, maChiTietPhieuNhap int) error {
    return r.db.Where("MaPhieuNhap = ?", maPhieuNhap).Where("MaChiTiet = ?", maChiTietPhieuNhap).Delete(&models.ChiTietPhieuNhap{}).Error
}

//DeleteAllChiTietByPhieuNhap xóa tất cả chi tiết của một phiếu nhập
func (r *phieuNhapRepository) DeleteAllChiTietByPhieuNhap(maPhieuNhap int) error {
    return r.db.Where("MaPhieuNhap = ?", maPhieuNhap).Delete(&models.ChiTietPhieuNhap{}).Error
}

// UpdateChiTietPhieuNhapSoLuong cập nhật số lượng trong chi tiết phiếu nhập
func (r *phieuNhapRepository) UpdateChiTietPhieuNhapSoLuong(maPhieuNhap int, maChiTiet int, soLuong int) error{
    return r.db.Model(&models.ChiTietPhieuNhap{}).
        Where("MaChiTiet = ? AND MaPhieuNhap = ?", maChiTiet, maPhieuNhap).
        Update("SoLuong", soLuong).
        Error
}

func (r *phieuNhapRepository) ExistsChiTietPhieuNhap(maChiTiet int, maPhieuNhap int) (bool, error) {
    var count int64
    if err := r.db.Model(&models.ChiTietPhieuNhap{}).
        Where("MaChiTiet = ? AND MaPhieuNhap = ?", maChiTiet, maPhieuNhap).
        Count(&count).Error; err != nil {
        return false, err
    }
    return count > 0, nil
}