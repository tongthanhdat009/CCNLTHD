package handlers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"github.com/tongthanhdat009/CCNLTHD/internal/services"
)

type AdminReviewHandler struct{ svc services.AdminReviewService }

func NewAdminReviewHandler(s services.AdminReviewService) *AdminReviewHandler {
	return &AdminReviewHandler{svc: s}
}

// GET /api/admin/reviews
func (h *AdminReviewHandler) List(c *gin.Context) {
	log.Printf("[AdminList][raw] status=%q trangThai=%q productId=%q maHangHoa=%q userId=%q maNguoiDung=%q page=%q size=%q",
		c.Query("status"), c.Query("trangThai"),
		c.Query("productId"), c.Query("maHangHoa"),
		c.Query("userId"), c.Query("maNguoiDung"),
		c.Query("page"), c.Query("pageSize"),
	)
	var f models.AdminReviewFilter

	// productId | maHangHoa
	if v := strings.TrimSpace(c.Query("productId")); v != "" {
		if n, _ := strconv.Atoi(v); n > 0 {
			f.MaHangHoa = &n
		}
	}
	if f.MaHangHoa == nil {
		if v := strings.TrimSpace(c.Query("maHangHoa")); v != "" {
			if n, _ := strconv.Atoi(v); n > 0 {
				f.MaHangHoa = &n
			}
		}
	}

	// userId | maNguoiDung
	if v := strings.TrimSpace(c.Query("userId")); v != "" {
		if n, _ := strconv.Atoi(v); n > 0 {
			f.MaNguoiDung = &n
		}
	}
	if f.MaNguoiDung == nil {
		if v := strings.TrimSpace(c.Query("maNguoiDung")); v != "" {
			if n, _ := strconv.Atoi(v); n > 0 {
				f.MaNguoiDung = &n
			}
		}
	}

	// status | trangThai (repo sẽ normalize "pending/approved/rejected" -> VN)
	if v := strings.TrimSpace(c.Query("status")); v != "" {
		s := v
		f.TrangThai = &s
	}
	if f.TrangThai == nil {
		if v := strings.TrimSpace(c.Query("trangThai")); v != "" {
			s := v
			f.TrangThai = &s
		}
	}

	// q (tìm trong Nội dung)
	if v := strings.TrimSpace(c.Query("q")); v != "" {
		s := v
		f.Q = &s
	}

	// page / pageSize (clamp)
	if v := strings.TrimSpace(c.Query("page")); v != "" {
		if n, _ := strconv.Atoi(v); n > 0 {
			f.Page = n
		}
	}
	if v := strings.TrimSpace(c.Query("pageSize")); v != "" {
		if n, _ := strconv.Atoi(v); n > 0 {
			f.PageSize = n
		}
	}
	if f.Page <= 0 {
		f.Page = 1
	}
	if f.PageSize <= 0 || f.PageSize > 100 {
		f.PageSize = 10
	}
	pStr := func(s *string) string {
		if s == nil {
			return "<nil>"
		}
		return *s
	}
	pInt := func(i *int) any {
		if i == nil {
			return "<nil>"
		}
		return *i
	}
	log.Printf("[AdminList][parsed] TrangThai=%s MaHangHoa=%v MaNguoiDung=%v Page=%d Size=%d",
		pStr(f.TrangThai), pInt(f.MaHangHoa), pInt(f.MaNguoiDung), f.Page, f.PageSize,
	)
	res, err := h.svc.List(c.Request.Context(), f)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"items": res.Items,
		"pagination": gin.H{
			"page": f.Page, "pageSize": f.PageSize, "total": res.Total,
		},
	})
}

// PUT /api/admin/reviews/:id/approve
func (h *AdminReviewHandler) Approve(c *gin.Context) {
	idStr := strings.TrimSpace(c.Param("id")) // <-- TRIM
	id, _ := strconv.Atoi(idStr)
	if id <= 0 {
		c.JSON(400, gin.H{"message": "id không hợp lệ"})
		return
	}
	if err := h.svc.SetStatus(c.Request.Context(), id, "Đã duyệt"); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Đã duyệt"})
}

// PUT /api/admin/reviews/:id/reject
func (h *AdminReviewHandler) Reject(c *gin.Context) {
	idStr := strings.TrimSpace(c.Param("id")) // <-- TRIM
	id, _ := strconv.Atoi(idStr)
	if id <= 0 {
		c.JSON(400, gin.H{"message": "id không hợp lệ"})
		return
	}
	if err := h.svc.SetStatus(c.Request.Context(), id, "Đã từ chối"); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Đã từ chối"})
}

// DELETE /api/admin/reviews/:id
func (h *AdminReviewHandler) Delete(c *gin.Context) {
	idStr := strings.TrimSpace(c.Param("id")) // <-- TRIM
	id, _ := strconv.Atoi(idStr)
	if id <= 0 {
		c.JSON(400, gin.H{"message": "id không hợp lệ"})
		return
	}
	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Đã xóa"})
}
