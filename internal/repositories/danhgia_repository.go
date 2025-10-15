package repositories

import (
	"context"
	"fmt"

	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"gorm.io/gorm"
)

type ReviewRepository interface {
	UserPurchasedHangHoa(ctx context.Context, userID, maHH int) (bool, error)
	InsertHangHoa(ctx context.Context, userID, maHH int, dto models.CreateReviewDTO) (int64, error)
	ListApprovedByHangHoa(ctx context.Context, maHH int) ([]models.Review, error)
	ListByUser(ctx context.Context, userID int) ([]models.Review, error)
}

type reviewRepository struct{ db *gorm.DB }

func NewReviewRepository(db *gorm.DB) ReviewRepository { return &reviewRepository{db: db} }

// JOIN: donhang → chitietdonhang → sanpham → chitietphieunhap → bienthe (MaHangHoa)
func (r *reviewRepository) UserPurchasedHangHoa(ctx context.Context, userID, maHH int) (bool, error) {
	var purchased int
	err := r.db.WithContext(ctx).Raw(`
		SELECT EXISTS(
		  SELECT 1
		  FROM donhang dh
		  JOIN chitietdonhang ct  ON ct.MaDonHang = dh.MaDonHang
		  JOIN sanpham sp         ON sp.MaSanPham = ct.MaSanPham
		  JOIN chitietphieunhap c ON c.MaChiTiet  = sp.MaChiTietPhieuNhap
		  JOIN bienthe bt         ON bt.MaBienThe = c.MaBienThe
		  WHERE dh.MaNguoiDung = ? AND bt.MaHangHoa = ?
		  LIMIT 1
		) AS purchased
	`, userID, maHH).Row().Scan(&purchased)
	if err != nil {
		return false, fmt.Errorf("UserPurchasedHangHoa: %w", err)
	}
	return purchased == 1, nil
}

func (r *reviewRepository) InsertHangHoa(ctx context.Context, userID, maHH int, dto models.CreateReviewDTO) (int64, error) {
	rec := models.Review{
		MaHangHoa:   maHH,
		MaNguoiDung: userID,
		Diem:        dto.Diem,
		NoiDung:     dto.NoiDung,
		TrangThai:   models.ReviewStatusPending,
		NgayDanhGia: dto.Now(), // xem helper ở models phía dưới
	}
	if err := r.db.WithContext(ctx).Table(rec.TableName()).Create(&rec).Error; err != nil {
		return 0, fmt.Errorf("InsertHangHoa: %w", err)
	}
	return int64(rec.MaDanhGia), nil
}

func (r *reviewRepository) ListApprovedByHangHoa(ctx context.Context, maHH int) ([]models.Review, error) {
	var out []models.Review
	err := r.db.WithContext(ctx).
		Table((models.Review{}).TableName()).
		Where("MaHangHoa = ? AND TrangThai = ?", maHH, models.ReviewStatusApproved).
		Order("NgayDanhGia DESC").
		Find(&out).Error
	return out, err
}
func (r *reviewRepository) ListByUser(ctx context.Context, userID int) ([]models.Review, error) {
	var out []models.Review
	err := r.db.WithContext(ctx).
		Table((models.Review{}).TableName()).
		Where("MaNguoiDung = ?", userID).
		Order("NgayDanhGia DESC").
		Find(&out).Error
	return out, err
}
