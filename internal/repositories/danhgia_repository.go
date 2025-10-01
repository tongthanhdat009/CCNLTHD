package repositories

import (
    "gorm.io/gorm"
    "github.com/tongthanhdat009/CCNLTHD/internal/models"
)

type ReviewRepository struct {
    DB *gorm.DB
}

func NewReviewRepository(db *gorm.DB) *ReviewRepository {
    return &ReviewRepository{DB: db}
}

func (r *ReviewRepository) Create(review *models.DanhGia) error {
    return r.DB.Create(review).Error
}

func (r *ReviewRepository) GetByProduct(maSanPham int) ([]models.DanhGia, error) {
    var list []models.DanhGia
    err := r.DB.Where("MaSanPham = ? AND TrangThai = 'Đã duyệt'", maSanPham).Find(&list).Error
    return list, err
}

func (r *ReviewRepository) GetByUser(maNguoiDung int) ([]models.DanhGia, error) {
    var list []models.DanhGia
    err := r.DB.Where("MaNguoiDung = ?", maNguoiDung).Order("NgayDanhGia DESC").Find(&list).Error
    return list, err
}

func (r *ReviewRepository) Approve(maDanhGia int) error {
    return r.DB.Model(&models.DanhGia{}).Where("MaDanhGia = ?", maDanhGia).
        Update("TrangThai", "Đã duyệt").Error
}

func (r *ReviewRepository) Reject(maDanhGia int) error {
    return r.DB.Model(&models.DanhGia{}).Where("MaDanhGia = ?", maDanhGia).
        Update("TrangThai", "Từ chối").Error
}

func (r *ReviewRepository) Delete(maDanhGia int) error {
    return r.DB.Delete(&models.DanhGia{}, maDanhGia).Error
}
