package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"github.com/tongthanhdat009/CCNLTHD/internal/services"
)

type ReviewHandler struct{ svc services.ReviewService }

func NewReviewHandler(svc services.ReviewService) *ReviewHandler { return &ReviewHandler{svc: svc} }

func (h *ReviewHandler) Create(c *gin.Context) {
	var dto models.CreateReviewDTO
	if err := c.ShouldBind(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	log.Printf("[Review/Create] uid=%v maHangHoa=%v", userIDFromCtx(c), dto.MaHangHoa)
	uid := userIDFromCtx(c)
	id, err := h.svc.Create(c.Request.Context(), uid, dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Đã gửi đánh giá, chờ duyệt", "id": id})
}

// GET /api/reviews/product/:id  (id ở đây là MaHangHoa)
func (h *ReviewHandler) GetByProduct(c *gin.Context) {
	maHH, err := strconv.Atoi(c.Param("id"))
	if err != nil || maHH <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id hàng hóa không hợp lệ"})
		return
	}
	items, err := h.svc.ListApprovedByHangHoa(c.Request.Context(), maHH)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *ReviewHandler) GetMine(c *gin.Context) {
	uid := userIDFromCtx(c)
	items, err := h.svc.ListMine(c.Request.Context(), uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

// helper
func userIDFromCtx(c *gin.Context) int {
	if v, ok := c.Get("userID"); ok {
		if id, ok2 := v.(int); ok2 && id > 0 {
			return id
		}
	}
	if h := c.GetHeader("X-User-ID"); h != "" {
		if id, err := strconv.Atoi(h); err == nil && id > 0 {
			return id
		}
	}
	return 0
}
