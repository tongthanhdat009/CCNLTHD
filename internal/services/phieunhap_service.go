package services

import (
	"errors"
	"log"

	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"github.com/tongthanhdat009/CCNLTHD/internal/repositories"

	"time"
)

type QuanLyPhieuNhapService interface {
    GetAllPhieuNhaps() ([]models.PhieuNhap, error)
    CreatePhieuNhap(phieuNhap *models.PhieuNhap) error
    ExistsInPhieuNhap(nhaCungCapID int) (bool, error)
	GetPhieuNhapByID(id int) (*models.PhieuNhap, error)
	DeletePhieuNhap(id int) error
	UpdatePhieuNhap(phieuNhap *models.PhieuNhap, approve bool) error
    SearchPhieuNhaps(tenNguoiDung, tenNhaCungCap, trangThai string, tuNgay, denNgay *time.Time) ([]models.PhieuNhap, error)
	GetChiTietByPhieuNhap(maPhieuNhap int) ([]models.ChiTietPhieuNhap, error)
	DeleteChiTietByPhieuNhap(maPhieuNhap int, maChiTietPhieuNhap int) error
	CreateChiTietPhieuNhap(maPhieuNhap int, items []models.ChiTietPhieuNhap) error
    DeleteAllChiTietByPhieuNhap(maPhieuNhap int) error
    UpdateChiTietPhieuNhapSoLuong(maPhieuNhap int, maChiTiet int, soLuong int) error

}

type quanLyPhieuNhapService struct {
    phieuNhapRepo repositories.PhieuNhapRepository
}

func NewQuanLyPhieuNhapService(phieuNhapRepo repositories.PhieuNhapRepository) QuanLyPhieuNhapService {
    return &quanLyPhieuNhapService{phieuNhapRepo: phieuNhapRepo}
}

func (s *quanLyPhieuNhapService) GetAllPhieuNhaps() ([]models.PhieuNhap, error) {
    return s.phieuNhapRepo.GetAll()
}

func (s *quanLyPhieuNhapService) GetPhieuNhapByID(id int) (*models.PhieuNhap, error) {
	return s.phieuNhapRepo.GetPhieuNhapByID(id)
}

func (s *quanLyPhieuNhapService) CreatePhieuNhap(phieuNhap *models.PhieuNhap) error {
    // Kiểm tra nhà cung cấp tồn tại
	log.Println("Checking if NhaCungCap exists with MaNCC:", phieuNhap.MaNCC)
    exists, err := s.phieuNhapRepo.ExistsNhaCungCap(phieuNhap.MaNCC)
    if err != nil {
        return err
    }
    if !exists {
        return errors.New("nhà cung cấp không tồn tại")
    }
		
	// Kiểm tra người dùng tồn tại
	exists, err = s.phieuNhapRepo.ExistsNguoiDung(phieuNhap.MaNguoiDung)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("người dùng không tồn tại")
	}

	// Set NgayNhap mặc định nếu chưa có
	if phieuNhap.NgayNhap.IsZero() {
		phieuNhap.NgayNhap = time.Now()
	}
	phieuNhap.NgayNhap = phieuNhap.NgayNhap.Local()

    // Set trạng thái mặc định là "ChuaNhap"
    phieuNhap.TrangThai = "Chưa duyệt"

    return s.phieuNhapRepo.CreatePhieuNhap(phieuNhap)
}

func (s *quanLyPhieuNhapService) DeletePhieuNhap(id int) error {
	exists, err := s.phieuNhapRepo.GetPhieuNhapByID(id)
	if err != nil {
		return err
	}
	if exists == nil {
		return errors.New("phiếu nhập không tồn tại")
	}
	if exists.TrangThai == "Đã duyệt" {
		return errors.New("không thể xóa phiếu nhập đã duyệt")
	}
	return s.phieuNhapRepo.DeletePhieuNhap(id)
}

func (s *quanLyPhieuNhapService) UpdatePhieuNhap(phieuNhap *models.PhieuNhap, approve bool) error {
	// Kiểm tra phiếu nhập tồn tại
	exists, err := s.phieuNhapRepo.GetPhieuNhapByID(phieuNhap.MaPhieuNhap)
	log.Println("Updating PhieuNhap with MaPhieuNhap:", phieuNhap.MaPhieuNhap)
	if err != nil {
		return err
	}
	if exists == nil {
		return errors.New("phiếu nhập không tồn tại")
	}
	
	// Nếu phiếu nhập đã được duyệt trước đó, không cho phép cập nhật
	if exists.TrangThai == "Đã duyệt" {
		return errors.New("không thể cập nhật phiếu nhập đã duyệt")
	}
	
	// Nếu đang duyệt phiếu nhập, kiểm tra thêm điều kiện
	if approve {
		if phieuNhap.TrangThai != "Đã duyệt" {
			phieuNhap.TrangThai = "Đã duyệt"
		}
	}
	
	return s.phieuNhapRepo.UpdatePhieuNhap(phieuNhap, approve)
}

func (s *quanLyPhieuNhapService) SearchPhieuNhaps(tenNguoiDung, tenNhaCungCap, trangThai string, tuNgay, denNgay *time.Time) ([]models.PhieuNhap, error) {
    return s.phieuNhapRepo.SearchPhieuNhap(tenNguoiDung, tenNhaCungCap, trangThai, tuNgay, denNgay)
}

func (s *quanLyPhieuNhapService) ExistsInPhieuNhap(nhaCungCapID int) (bool, error) {
    return s.phieuNhapRepo.ExistsInPhieuNhap(nhaCungCapID)
}

func (s *quanLyPhieuNhapService) GetChiTietByPhieuNhap(maPhieuNhap int) ([]models.ChiTietPhieuNhap, error) {
    // Check if the receipt exists
    exists, err := s.phieuNhapRepo.GetPhieuNhapByID(maPhieuNhap)
    if err != nil {
        return nil, err
    }
    if exists == nil {
        return nil, errors.New("phiếu nhập không tồn tại")
    }
    
    // Get the details from repository
    return s.phieuNhapRepo.GetChiTietByPhieuNhap(maPhieuNhap)
}

func (s *quanLyPhieuNhapService) DeleteAllChiTietByPhieuNhap(maPhieuNhap int) error {
    // Check if the receipt exists
    exists, err := s.phieuNhapRepo.GetPhieuNhapByID(maPhieuNhap)
    if err != nil {
        return err
    }
    if exists == nil {
        return errors.New("phiếu nhập không tồn tại")
    }
    
    // Check if receipt is already approved
    if exists.TrangThai == "Đã duyệt" {
        return errors.New("không thể xóa chi tiết phiếu nhập đã duyệt")
    }
    
    // Delete the details through repository
    return s.phieuNhapRepo.DeleteAllChiTietByPhieuNhap(maPhieuNhap)
}

func (s *quanLyPhieuNhapService) DeleteChiTietByPhieuNhap(maPhieuNhap int, maChiTietPhieuNhap int) error {
    exists, err := s.phieuNhapRepo.GetPhieuNhapByID(maPhieuNhap)
    if err != nil {
        return err
    }
    if exists == nil {
        return errors.New("phiếu nhập không tồn tại")
    }
    if exists.TrangThai == "Đã duyệt" {
        return errors.New("không thể xóa chi tiết phiếu nhập trong phiếu nhập đã duyệt")
    }
    existsChiTiet, err := s.phieuNhapRepo.ExistsChiTietPhieuNhap(maChiTietPhieuNhap, maPhieuNhap)
    if err != nil {
        return err
    }
    if !existsChiTiet {
        return errors.New("chi tiết phiếu nhập không tồn tại")
    }   
    return s.phieuNhapRepo.DeleteChiTietByPhieuNhap(maPhieuNhap, maChiTietPhieuNhap)
}

func (s *quanLyPhieuNhapService) CreateChiTietPhieuNhap(maPhieuNhap int, items []models.ChiTietPhieuNhap) error {
    // Check if the receipt exists
    exists, err := s.phieuNhapRepo.GetPhieuNhapByID(maPhieuNhap)
    if err != nil {
        return err
    }
    if exists == nil {
        return errors.New("phiếu nhập không tồn tại")
    }
    
    // Check if receipt is already approved
    if exists.TrangThai == "Đã duyệt" {
        return errors.New("không thể thêm chi tiết vào phiếu nhập đã duyệt")
    }
    
    // Validate items
    if len(items) == 0 {
        return errors.New("danh sách chi tiết phiếu nhập không được rỗng")
    }
    
    // Set maPhieuNhap for all items to ensure consistency
    for i := range items {
        items[i].MaPhieuNhap = maPhieuNhap
        // Validate each item
        if items[i].MaBienthe <= 0 {
            return errors.New("mã biến thể không hợp lệ")
        }
        if items[i].SoLuong <= 0 {
            return errors.New("số lượng phải lớn hơn 0")
        }
        if items[i].GiaNhap <= 0 {
            return errors.New("giá nhập phải lớn hơn 0")
        }
		if items[i].ThoiGianBaoHanh.Valid && items[i].ThoiGianBaoHanh.Int64 < 0 {
			return errors.New("thời gian bảo hành không hợp lệ")
		}
		if items[i].NgaySanXuat.Valid && items[i].NgaySanXuat.Time.After(time.Now()) {
			return errors.New("ngày sản xuất không hợp lệ")
		}
    }
    
    // Create the details through repository
    return s.phieuNhapRepo.CreateChiTietPhieuNhap(maPhieuNhap, items)
}

func (s *quanLyPhieuNhapService) UpdateChiTietPhieuNhapSoLuong(maPhieuNhap int, maChiTiet int, soLuong int) error {
    // Check if the receipt exists
    exists, err := s.phieuNhapRepo.GetPhieuNhapByID(maPhieuNhap)
    if err != nil {
        return err
    }
    if exists == nil {
        return errors.New("phiếu nhập không tồn tại")
    }
    if exists.TrangThai == "Đã duyệt" {
        return errors.New("không thể cập nhật chi tiết phiếu nhập đã duyệt")
    }
    existsChiTiet, err := s.phieuNhapRepo.ExistsChiTietPhieuNhap(maChiTiet, maPhieuNhap)
    if err != nil {
        return err
    }
    if !existsChiTiet {
        return errors.New("chi tiết phiếu nhập không tồn tại")
    }
    if soLuong <= 0 {
        return errors.New("số lượng phải lớn hơn 0")
    }
    return s.phieuNhapRepo.UpdateChiTietPhieuNhapSoLuong(maPhieuNhap, maChiTiet, soLuong)
}