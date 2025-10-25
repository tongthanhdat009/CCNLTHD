package services

import (
	"errors"
	"time"

	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"github.com/tongthanhdat009/CCNLTHD/internal/repositories"
)

type DonHangService interface {
	// CRUD cơ bản
	CreateDonHang(donHang *models.DonHang) error
	GetDonHangByID(maDonHang int) (models.DonHang, error)
	GetAllDonHang() ([]models.DonHang, error)
	UpdateDonHang(donHang *models.DonHang) error
	DeleteDonHang(maDonHang int) error

	// Quản lý trạng thái
	ApproveOrder(maDonHang int) error                        // Duyệt đơn hàng (Chờ xác nhận -> Đã xác nhận)
	UpdateOrderStatus(maDonHang int, trangThai string) error // Cập nhật trạng thái đơn hàng
	CancelOrder(maDonHang int, reason string) error          // Hủy đơn hàng

	// Tìm kiếm và lọc
	SearchDonHang(keyword string, trangThai string, fromDate, toDate time.Time) ([]models.DonHang, error)
	GetDonHangByNguoiDung(maNguoiDung int) ([]models.DonHang, error)
	GetDonHangByStatus(trangThai string) ([]models.DonHang, error)

	// Xem chi tiết
	GetDetailByID(maDonHang int) (models.DonHang, error)

	// Validation
	ValidateOrderData(donHang *models.DonHang) error
	CanModifyOrder(maDonHang int) (bool, error)
}

type donHangService struct {
	repo repositories.DonHangRepository
}

func NewDonHangService(repo repositories.DonHangRepository) DonHangService {
	return &donHangService{repo: repo}
}

// CreateDonHang - Tạo đơn hàng mới
func (s *donHangService) CreateDonHang(donHang *models.DonHang) error {
	// Validate dữ liệu đơn hàng
	if err := s.ValidateOrderData(donHang); err != nil {
		return err
	}

	// Thiết lập trạng thái mặc định
	if donHang.TrangThai == "" {
		donHang.TrangThai = "Chờ xác nhận"
	}

	// Tạo đơn hàng
	return s.repo.CreateDonHang(donHang)
}

// GetDonHangByID - Lấy thông tin đơn hàng theo mã
func (s *donHangService) GetDonHangByID(maDonHang int) (models.DonHang, error) {
	if maDonHang <= 0 {
		return models.DonHang{}, errors.New("mã đơn hàng không hợp lệ")
	}
	return s.repo.GetByID(maDonHang)
}

// GetAllDonHang - Lấy tất cả đơn hàng
func (s *donHangService) GetAllDonHang() ([]models.DonHang, error) {
	return s.repo.GetAll()
}

// UpdateDonHang - Cập nhật thông tin đơn hàng
func (s *donHangService) UpdateDonHang(donHang *models.DonHang) error {
	// Kiểm tra đơn hàng có tồn tại không
	exists, err := s.repo.ExistsDonHang(donHang.MaDonHang)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("đơn hàng không tồn tại")
	}

	// Kiểm tra có thể sửa đổi không
	canModify, err := s.CanModifyOrder(donHang.MaDonHang)
	if err != nil {
		return err
	}
	if !canModify {
		return errors.New("không thể sửa đổi đơn hàng ở trạng thái hiện tại")
	}

	// Lấy thông tin đơn hàng cũ để giữ nguyên MaNguoiDung
	oldOrder, err := s.repo.GetByID(donHang.MaDonHang)
	if err != nil {
		return err
	}

	// Giữ nguyên MaNguoiDung và NgayTao
	donHang.MaNguoiDung = oldOrder.MaNguoiDung
	donHang.NgayTao = oldOrder.NgayTao

	// Validate các trường khác (không validate MaNguoiDung)
	if donHang.TongTien < 0 {
		return errors.New("tổng tiền không hợp lệ")
	}

	validPaymentMethods := []string{"Tiền mặt", "Chuyển khoản", "Ví điện tử", "Thẻ tín dụng", "COD"}
	isValidPayment := false
	for _, method := range validPaymentMethods {
		if donHang.PhuongThucThanhToan == method {
			isValidPayment = true
			break
		}
	}
	if !isValidPayment {
		return errors.New("phương thức thanh toán không hợp lệ")
	}

	// Cập nhật đơn hàng
	return s.repo.UpdateDonHang(donHang)
}

// DeleteDonHang - Xóa đơn hàng
func (s *donHangService) DeleteDonHang(maDonHang int) error {
	// Kiểm tra đơn hàng có tồn tại không
	exists, err := s.repo.ExistsDonHang(maDonHang)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("đơn hàng không tồn tại")
	}

	// Kiểm tra trạng thái có cho phép xóa không
	currentStatus, err := s.repo.GetCurrentStatus(maDonHang)
	if err != nil {
		return err
	}

	// Chỉ cho phép xóa đơn hàng ở trạng thái "Đã hủy" hoặc "Chờ xác nhận"
	if currentStatus != "Đã hủy" && currentStatus != "Chờ xác nhận" {
		return errors.New("chỉ có thể xóa đơn hàng ở trạng thái 'Chờ xác nhận' hoặc 'Đã hủy'")
	}

	return s.repo.DeleteDonHang(maDonHang)
}

func (s *donHangService) ApproveOrder(maDonHang int) error {
	// Kiểm tra đơn hàng có tồn tại không
	exists, err := s.repo.ExistsDonHang(maDonHang)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("đơn hàng không tồn tại")
	}

	// Lấy trạng thái hiện tại
	currentStatus, err := s.repo.GetCurrentStatus(maDonHang)
	if err != nil {
		return err
	}

	// Chỉ có thể duyệt đơn hàng ở trạng thái "Chờ xác nhận"
	if currentStatus != "Chờ xác nhận" {
		return errors.New("chỉ có thể duyệt đơn hàng ở trạng thái 'Chờ xác nhận'")
	}

	// Cập nhật trạng thái sang "Đang giao hàng"
	return s.repo.UpdateStatus(maDonHang, "Đang giao hàng")
}

// UpdateOrderStatus - Cập nhật trạng thái đơn hàng
func (s *donHangService) UpdateOrderStatus(maDonHang int, trangThai string) error {
	if maDonHang <= 0 {
		return errors.New("mã đơn hàng không hợp lệ")
	}

	if trangThai == "" {
		return errors.New("trạng thái không được để trống")
	}

	// Kiểm tra trạng thái hợp lệ
	validStatuses := []string{
		"Chờ xác nhận",
		"Đang giao hàng",
		"Đã giao hàng",
		"Hoàn thành",
		"Đã hủy",
		"Giao hàng thất bại",
	}

	isValid := false
	for _, status := range validStatuses {
		if status == trangThai {
			isValid = true
			break
		}
	}

	if !isValid {
		return errors.New("trạng thái không hợp lệ")
	}

	return s.repo.UpdateStatus(maDonHang, trangThai)
}

// CancelOrder - Hủy đơn hàng
func (s *donHangService) CancelOrder(maDonHang int, reason string) error {
	// Kiểm tra đơn hàng có tồn tại không
	exists, err := s.repo.ExistsDonHang(maDonHang)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("đơn hàng không tồn tại")
	}

	// Kiểm tra trạng thái có cho phép hủy không
	currentStatus, err := s.repo.GetCurrentStatus(maDonHang)
	if err != nil {
		return err
	}

	// Chỉ cho phép hủy ở một số trạng thái nhất định
	allowedStatuses := []string{"Chờ xác nhận", "Đã xác nhận", "Đang chuẩn bị", "Giao hàng thất bại"}
	canCancel := false
	for _, status := range allowedStatuses {
		if status == currentStatus {
			canCancel = true
			break
		}
	}

	if !canCancel {
		return errors.New("không thể hủy đơn hàng ở trạng thái '" + currentStatus + "'")
	}

	// TODO: Lưu lý do hủy vào database nếu cần
	_ = reason

	// Cập nhật trạng thái thành "Đã hủy"
	return s.repo.UpdateStatus(maDonHang, "Đã hủy")
}

// SearchDonHang - Tìm kiếm đơn hàng theo nhiều tiêu chí
func (s *donHangService) SearchDonHang(keyword string, trangThai string, fromDate, toDate time.Time) ([]models.DonHang, error) {
	return s.repo.SearchDonHang(keyword, trangThai, fromDate, toDate)
}

// GetDonHangByNguoiDung - Lấy đơn hàng theo người dùng
func (s *donHangService) GetDonHangByNguoiDung(maNguoiDung int) ([]models.DonHang, error) {
	if maNguoiDung <= 0 {
		return nil, errors.New("mã người dùng không hợp lệ")
	}
	return s.repo.GetByNguoiDung(maNguoiDung)
}

// GetDonHangByStatus - Lấy đơn hàng theo trạng thái
func (s *donHangService) GetDonHangByStatus(trangThai string) ([]models.DonHang, error) {
	if trangThai == "" {
		return nil, errors.New("trạng thái không được để trống")
	}
	return s.repo.GetByStatus(trangThai)
}

// GetDetailByID - Xem chi tiết đầy đủ đơn hàng
func (s *donHangService) GetDetailByID(maDonHang int) (models.DonHang, error) {
	if maDonHang <= 0 {
		return models.DonHang{}, errors.New("mã đơn hàng không hợp lệ")
	}
	return s.repo.GetDetailByID(maDonHang)
}

func (s *donHangService) ValidateOrderData(donHang *models.DonHang) error {
	// Validate mã người dùng
	if donHang.MaNguoiDung <= 0 {
		return errors.New("mã người dùng không hợp lệ")
	}

	// Validate tổng tiền
	if donHang.TongTien < 0 {
		return errors.New("tổng tiền không hợp lệ")
	}

	// Validate địa chỉ (có thể bỏ qua nếu không bắt buộc)
	// if donHang.TinhThanh == "" || donHang.QuanHuyen == "" || donHang.PhuongXa == "" || donHang.DuongSoNha == "" {
	// 	return errors.New("địa chỉ không đầy đủ")
	// }

	// Validate phương thức thanh toán
	validPaymentMethods := []string{"Tiền mặt", "Chuyển khoản", "Ví điện tử", "Thẻ tín dụng", "COD"}
	isValidPayment := false
	for _, method := range validPaymentMethods {
		if donHang.PhuongThucThanhToan == method {
			isValidPayment = true
			break
		}
	}
	if !isValidPayment {
		return errors.New("phương thức thanh toán không hợp lệ")
	}

	// Validate trạng thái
	validStatuses := []string{"Chờ xác nhận", "Đang giao hàng", "Đã giao hàng", "Hoàn thành", "Đã hủy", "Giao hàng thất bại"}
	isValidStatus := false
	for _, status := range validStatuses {
		if donHang.TrangThai == status {
			isValidStatus = true
			break
		}
	}
	if !isValidStatus && donHang.TrangThai != "" {
		return errors.New("trạng thái không hợp lệ")
	}

	return nil
}

// CanModifyOrder - Kiểm tra có thể sửa đổi đơn hàng không
func (s *donHangService) CanModifyOrder(maDonHang int) (bool, error) {
	currentStatus, err := s.repo.GetCurrentStatus(maDonHang)
	if err != nil {
		return false, err
	}

	// Chỉ cho phép sửa đổi ở trạng thái "Chờ xác nhận" và "Đã xác nhận"
	modifiableStatuses := []string{"Chờ xác nhận", "Đang giao hàng", "Giao hàng thất bại"}
	for _, status := range modifiableStatuses {
		if status == currentStatus {
			return true, nil
		}
	}

	return false, nil
}
