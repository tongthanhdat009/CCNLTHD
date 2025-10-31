package repositories

import (
	"context"
	"fmt"
	"log"
	"strings"

	"gorm.io/gorm"

	"github.com/tongthanhdat009/CCNLTHD/internal/models"
)

func normalizeStatus(input string) (string, error) {
	s := strings.TrimSpace(strings.ToLower(input))
	switch s {
	case "", "all":
		return "", nil
	case "pending", "chua_duyet", "chuaduyet", "chưa duyệt":
		return "Chưa duyệt", nil
	case "approved", "da_duyet", "daduyet", "đã duyệt":
		return "Đã duyệt", nil
	case "rejected", "tu_choi", "tuchoi", "da tu choi", "đã từ chối":
		return "Đã từ chối", nil
	default:
		return "", fmt.Errorf("Trạng thái không hợp lệ")
	}
}

// Kết quả trả về khi list (Exported)
type AdminReviewListResult struct {
	Items []models.Review
	Total int64
}

// Interface repo admin (Exported)
type AdminReviewRepository interface {
	AdminList(ctx context.Context, f models.AdminReviewFilter) (AdminReviewListResult, error)
	AdminUpdateStatus(ctx context.Context, id int, newStatus string) (int64, error)
	AdminDelete(ctx context.Context, id int) (int64, error)
}

// Implement
type adminReviewRepository struct{ db *gorm.DB }

// Hàm khởi tạo (Exported)
func NewAdminReviewRepository(db *gorm.DB) AdminReviewRepository {
	return &adminReviewRepository{db: db}
}

func (r *adminReviewRepository) AdminList(ctx context.Context, f models.AdminReviewFilter) (AdminReviewListResult, error) {
	if f.Page <= 0 {
		f.Page = 1
	}
	if f.PageSize <= 0 || f.PageSize > 200 {
		f.PageSize = 20
	}

	tx := r.db.WithContext(ctx).Table((models.Review{}).TableName())
	if f.MaHangHoa != nil {
		tx = tx.Where("MaHangHoa = ?", *f.MaHangHoa)
	}
	if f.MaNguoiDung != nil {
		tx = tx.Where("MaNguoiDung = ?", *f.MaNguoiDung)
	}
	if f.TrangThai != nil && strings.TrimSpace(*f.TrangThai) != "" {
		mapped, err := normalizeStatus(*f.TrangThai)
		if err != nil {
			return AdminReviewListResult{}, err
		}
		// LOG xem map ra gì
		log.Printf("[AdminList][repo] inputStatus=%q mapped=%q", *f.TrangThai, mapped)

		if mapped != "" {
			// So sánh “an toàn”: cắt khoảng trắng + không phân biệt hoa/thường
			tx = tx.Where("LOWER(TRIM(TrangThai)) = LOWER(TRIM(?))", mapped)
		}
	}
	if f.Q != nil && *f.Q != "" {
		tx = tx.Where("NoiDung LIKE ?", "%"+*f.Q+"%")
	}

	var total int64
	if err := tx.Count(&total).Error; err != nil {
		return AdminReviewListResult{}, fmt.Errorf("AdminList count: %w", err)
	}

	var items []models.Review
	if err := tx.Order("NgayDanhGia DESC").Limit(f.PageSize).Offset((f.Page - 1) * f.PageSize).Find(&items).Error; err != nil {
		return AdminReviewListResult{}, fmt.Errorf("AdminList find: %w", err)
	}
	return AdminReviewListResult{Items: items, Total: total}, nil
}

func (r *adminReviewRepository) AdminUpdateStatus(ctx context.Context, id int, newStatus string) (int64, error) {
	res := r.db.WithContext(ctx).Table((models.Review{}).TableName()).
		Where("MaDanhGia = ?", id).
		Update("TrangThai", newStatus)
	return res.RowsAffected, res.Error
}

func (r *adminReviewRepository) AdminDelete(ctx context.Context, id int) (int64, error) {
	res := r.db.WithContext(ctx).Table((models.Review{}).TableName()).
		Where("MaDanhGia = ?", id).
		Delete(&models.Review{})
	return res.RowsAffected, res.Error
}
