package handlers

import (
	"net/http"
	"strconv"

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
	var f models.AdminReviewFilter
	if v := c.Query("maHangHoa"); v != "" {
		if n, _ := strconv.Atoi(v); n > 0 {
			f.MaHangHoa = &n
		}
	}
	if v := c.Query("maNguoiDung"); v != "" {
		if n, _ := strconv.Atoi(v); n > 0 {
			f.MaNguoiDung = &n
		}
	}
	if v := c.Query("trangThai"); v != "" {
		s := v
		f.TrangThai = &s
	}
	if v := c.Query("q"); v != "" {
		s := v
		f.Q = &s
	}
	if v := c.Query("page"); v != "" {
		if n, _ := strconv.Atoi(v); n > 0 {
			f.Page = n
		}
	}
	if v := c.Query("pageSize"); v != "" {
		if n, _ := strconv.Atoi(v); n > 0 {
			f.PageSize = n
		}
	}

	res, err := h.svc.List(c.Request.Context(), f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": res.Items, "pagination": gin.H{"page": f.Page, "pageSize": f.PageSize, "total": res.Total}})
}

// PUT /api/admin/reviews/:id/approve
func (h *AdminReviewHandler) Approve(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
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
	id, _ := strconv.Atoi(c.Param("id"))
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
	id, _ := strconv.Atoi(c.Param("id"))
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
