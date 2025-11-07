package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"github.com/tongthanhdat009/CCNLTHD/internal/services"
)

type OrderHistoryHandler struct{ svc services.OrderHistoryService }

func NewOrderHistoryHandler(s services.OrderHistoryService) *OrderHistoryHandler {
	return &OrderHistoryHandler{svc: s}
}

// --- helper: LẤY userID KHÔNG CẦN SỬA MIDDLEWARE
func userIDFromRequest(c *gin.Context) int {
	if uid := c.GetInt("userID"); uid > 0 {
		return uid
	}
	// Fallback 1: header dev
	if x := strings.TrimSpace(c.GetHeader("X-User-ID")); x != "" {
		if n, err := strconv.Atoi(x); err == nil && n > 0 {
			return n
		}
	}
	// Fallback 2: parse Bearer và đọc claim "ma_nguoi_dung"
	if auth := c.GetHeader("Authorization"); strings.HasPrefix(auth, "Bearer ") {
		tokenStr := strings.TrimSpace(strings.TrimPrefix(auth, "Bearer "))
		if tokenStr != "" {
			secret := os.Getenv("JWT_SECRET")
			if secret == "" {
				secret = "dev-secret"
			} // dev only
			tok, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
				}
				return []byte(secret), nil
			})
			if err == nil && tok != nil && tok.Valid {
				if mc, ok := tok.Claims.(jwt.MapClaims); ok {
					if v, ok := mc["ma_nguoi_dung"]; ok {
						switch vv := v.(type) {
						case float64:
							return int(vv)
						case string:
							if n, err := strconv.Atoi(vv); err == nil {
								return n
							}
						}
					}
					// các fallback khác nếu hệ khác dùng uid/sub
					if v, ok := mc["uid"]; ok {
						switch vv := v.(type) {
						case float64:
							return int(vv)
						case string:
							if n, err := strconv.Atoi(vv); err == nil {
								return n
							}
						}
					}
					if v, ok := mc["sub"]; ok {
						switch vv := v.(type) {
						case float64:
							return int(vv)
						case string:
							if n, err := strconv.Atoi(vv); err == nil {
								return n
							}
						}
					}
				}
			}
		}
	}
	return 0
}

// GET /api/orders/me?status=&page=&pageSize=
func (h *OrderHistoryHandler) ListMine(c *gin.Context) {
	uid := userIDFromRequest(c)
	if uid == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var f models.OrderHistoryFilter
	_ = c.ShouldBindQuery(&f)

	res, err := h.svc.ListMine(c.Request.Context(), uid, f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GET /api/orders/:id
func (h *OrderHistoryHandler) GetDetail(c *gin.Context) {
	uid := userIDFromRequest(c)
	if uid == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id không hợp lệ"})
		return
	}

	dt, err := h.svc.GetDetail(c.Request.Context(), uid, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Không tìm thấy đơn hàng"})
		return
	}
	c.JSON(http.StatusOK, dt)
}
