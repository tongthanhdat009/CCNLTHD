package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tongthanhdat009/CCNLTHD/internal/services"
)

type ReviewHandler struct {
	Service *services.ReviewService
}

func (h *ReviewHandler) AdminList(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "AdminList: not implemented yet"})
}

func (h *ReviewHandler) Approve(c *gin.Context) {
	_, _ = strconv.Atoi(c.Param("id"))
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Approve: not implemented yet"})
}

func (h *ReviewHandler) Reject(c *gin.Context) {
	_, _ = strconv.Atoi(c.Param("id"))
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Reject: not implemented yet"})
}

func (h *ReviewHandler) Delete(c *gin.Context) {
	_, _ = strconv.Atoi(c.Param("id"))
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Delete: not implemented yet"})
}
func getUserID(c *gin.Context) (int, bool) {
	claimsRaw, ok := c.Get("user")
	if !ok {
		return 0, false
	}
	claims := claimsRaw.(jwt.MapClaims)
	v, ok := claims["ma_nguoi_dung"].(float64)
	if !ok {
		return 0, false
	}
	return int(v), true
}

// POST /api/reviews
func (h *ReviewHandler) Create(c *gin.Context) {
	uid, ok := getUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "chưa đăng nhập"})
		return
	}

	var req struct {
		MaSanPham int    `json:"ma_san_pham"`
		Diem      int    `json:"diem"`
		NoiDung   string `json:"noi_dung"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.Create(uid, req.MaSanPham, req.Diem, req.NoiDung); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Đánh giá đã gửi, chờ duyệt"})
}

// GET /api/reviews/product/:id
func (h *ReviewHandler) GetByProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	list, err := h.Service.GetByProduct(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
}

// GET /api/reviews/me
func (h *ReviewHandler) GetMine(c *gin.Context) {
	uid, ok := getUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "chưa đăng nhập"})
		return
	}
	list, err := h.Service.GetByUser(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)

}
func NewReviewHandler(reviewService *services.ReviewService) *ReviewHandler {
	return &ReviewHandler{
		Service: reviewService,
	}
}
