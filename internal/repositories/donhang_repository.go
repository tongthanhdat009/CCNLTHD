package repositories

import (
	"errors"
	"time"

	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"gorm.io/gorm"
)

type DonHangRepository interface {
	// CRUD cơ bản
	CreateDonHang(donHang *models.DonHang) error
	GetByID(maDonHang int) (models.DonHang, error)
	GetAll() ([]models.DonHang, error)
	UpdateStatus(maDonHang int, trangThai string) error
	UpdateDonHang(donHang *models.DonHang) error
	DeleteDonHang(maDonHang int) error

	// Tìm kiếm đơn hàng
	SearchDonHang(keyword string, trangThai string, fromDate, toDate time.Time) ([]models.DonHang, error)
	GetByNguoiDung(maNguoiDung int) ([]models.DonHang, error)
	GetByStatus(trangThai string) ([]models.DonHang, error)

	// Xem chi tiết đơn hàng
	GetDetailByID(maDonHang int) (models.DonHang, error)

	// Validation
	ExistsDonHang(maDonHang int) (bool, error)
	GetCurrentStatus(maDonHang int) (string, error)
	CanUpdateStatus(currentStatus, newStatus string) bool
}

type DonHangRepo struct {
	db *gorm.DB
}

func NewDonHangRepository(db *gorm.DB) DonHangRepository {
	return &DonHangRepo{db: db}
}

// CreateDonHang - Tạo đơn hàng mới
func (r *DonHangRepo) CreateDonHang(donHang *models.DonHang) error {
	// Thiết lập trạng thái mặc định nếu chưa có

	return r.db.Create(donHang).Error
}

func (r *DonHangRepo) CreateChiTietDonHang(donHang *models.DonHang) error {
	return r.db.Create(&donHang.ChiTietDonHangs).Error
}

// GetByID - Lấy đơn hàng theo mã (không load chi tiết)
func (r *DonHangRepo) GetByID(maDonHang int) (models.DonHang, error) {
	var donHang models.DonHang
	err := r.db.Where("MaDonHang = ?", maDonHang).First(&donHang).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return donHang, errors.New("đơn hàng không tồn tại")
		}
		return donHang, err
	}
	return donHang, nil
}

// GetDetailByID - Xem chi tiết đầy đủ đơn hàng (bao gồm chi tiết và thông tin sản phẩm)
func (r *DonHangRepo) GetDetailByID(maDonHang int) (models.DonHang, error) {
	var donHang models.DonHang
	err := r.db.
		Preload("ChiTietDonHangs").
		Where("MaDonHang = ?", maDonHang).
		First(&donHang).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return donHang, errors.New("đơn hàng không tồn tại")
		}
		return donHang, err
	}
	return donHang, nil
}

// GetAll - Lấy tất cả đơn hàng
func (r *DonHangRepo) GetAll() ([]models.DonHang, error) {
	var donHangs []models.DonHang
	err := r.db.
		Preload("ChiTietDonHangs").
		Order("NgayTao DESC").
		Find(&donHangs).Error
	return donHangs, err
}

// UpdateStatus - Duyệt và tha			y đổi trạng thái đơn hàng
func (r *DonHangRepo) UpdateStatus(maDonHang int, trangThai string) error {
	// Kiểm tra đơn hàng có tồn tại không
	exists, err := r.ExistsDonHang(maDonHang)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("đơn hàng không tồn tại")
	}

	// Lấy trạng thái hiện tại
	currentStatus, err := r.GetCurrentStatus(maDonHang)
	if err != nil {
		return err
	}

	// Kiểm tra logic chuyển trạng thái
	if !r.CanUpdateStatus(currentStatus, trangThai) {
		return errors.New("không thể chuyển từ trạng thái '" + currentStatus + "' sang '" + trangThai + "'")
	}

	// Cập nhật trạng thái
	return r.db.Model(&models.DonHang{}).
		Where("MaDonHang = ?", maDonHang).
		Update("TrangThai", trangThai).Error
}

// UpdateDonHang - Cập nhật thông tin đơn hàng
func (r *DonHangRepo) UpdateDonHang(donHang *models.DonHang) error {
	return r.db.Save(donHang).Error
}

// DeleteDonHang - Xóa đơn hàng (soft delete hoặc hard delete tùy cấu hình)
func (r *DonHangRepo) DeleteDonHang(maDonHang int) error {
	return r.db.Delete(&models.DonHang{}, maDonHang).Error
}

// SearchDonHang - Tìm kiếm đơn hàng theo nhiều tiêu chí
func (r *DonHangRepo) SearchDonHang(keyword string, trangThai string, fromDate, toDate time.Time) ([]models.DonHang, error) {
	var donHangs []models.DonHang
	query := r.db.Model(&models.DonHang{}).Preload("ChiTietDonHangs")

	// Tìm theo mã đơn hàng hoặc mã người dùng
	if keyword != "" {
		// Thử convert keyword sang int để tìm theo mã
		// Nếu không convert được thì bỏ qua điều kiện này
		query = query.Where("MaDonHang = ? OR MaNguoiDung = ?", keyword, keyword)
	}

	// Lọc theo trạng thái
	if trangThai != "" {
		query = query.Where("TrangThai = ?", trangThai)
	}

	// Lọc theo khoảng thời gian
	if !fromDate.IsZero() {
		query = query.Where("NgayTao >= ?", fromDate)
	}
	if !toDate.IsZero() {
		query = query.Where("NgayTao <= ?", toDate)
	}

	err := query.Order("NgayTao DESC").Find(&donHangs).Error
	return donHangs, err
}

// GetByNguoiDung - Lấy đơn hàng theo người dùng
func (r *DonHangRepo) GetByNguoiDung(maNguoiDung int) ([]models.DonHang, error) {
	var donHangs []models.DonHang
	err := r.db.
		Preload("ChiTietDonHangs").
		Where("MaNguoiDung = ?", maNguoiDung).
		Find(&donHangs).Error
	return donHangs, err
}

// GetByStatus - Lấy đơn hàng theo trạng thái
func (r *DonHangRepo) GetByStatus(trangThai string) ([]models.DonHang, error) {
	var donHangs []models.DonHang
	err := r.db.
		Preload("ChiTietDonHangs").
		Where("TrangThai = ?", trangThai).
		Order("NgayTao DESC").
		Find(&donHangs).Error
	return donHangs, err
}

// ExistsDonHang - Kiểm tra đơn hàng có tồn tại không
func (r *DonHangRepo) ExistsDonHang(maDonHang int) (bool, error) {
	var count int64
	err := r.db.Model(&models.DonHang{}).
		Where("MaDonHang = ?", maDonHang).
		Count(&count).Error
	return count > 0, err
}

// GetCurrentStatus - Lấy trạng thái hiện tại của đơn hàng
func (r *DonHangRepo) GetCurrentStatus(maDonHang int) (string, error) {
	var donHang models.DonHang
	err := r.db.Select("TrangThai").
		Where("MaDonHang = ?", maDonHang).
		First(&donHang).Error
	if err != nil {
		return "", err
	}
	return donHang.TrangThai, nil
}

// CanUpdateStatus - Kiểm tra logic chuyển trạng thái hợp lệ
func (r *DonHangRepo) CanUpdateStatus(currentStatus, newStatus string) bool {
	// Định nghĩa flow trạng thái hợp lệ
	validTransitions := map[string][]string{
		"Đang xử lý":       {"Đang giao hàng", "Đã hủy"},
		"Đang giao hàng":     {"Đã giao hàng", "Giao hàng thất bại"},
		"Giao hàng thất bại": {"Đang giao hàng", "Đã hủy"},
		"Đã giao hàng":       {"Hoàn thành"},
		"Đã hủy":             {}, // Không thể chuyển sang trạng thái khác
		"Hoàn thành":         {}, // Không thể chuyển sang trạng thái khác
	}

	allowedStatuses, exists := validTransitions[currentStatus]
	if !exists {
		return false
	}

	for _, status := range allowedStatuses {
		if status == newStatus {
			return true
		}
	}
	return false
}
