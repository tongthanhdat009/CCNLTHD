package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tongthanhdat009/CCNLTHD/internal/services"
)

type PermissionMiddleware struct {
	permissionService services.PermissionService
}

func NewPermissionMiddleware(ps services.PermissionService) *PermissionMiddleware {
	return &PermissionMiddleware{permissionService: ps}
}

// Middleware kiểm tra quyền
func (pm *PermissionMiddleware) Require(maChucNang string, chiTiet string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claimsRaw, exists := c.Get("user")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		claims, ok := claimsRaw.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}

		maQuyen, ok := claims["ma_quyen"].(float64)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing role info"})
			return
		}

		// Kiểm tra quyền
		hasPermission, err := pm.permissionService.KiemTraQuyen(int(maQuyen), maChucNang, chiTiet)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if !hasPermission {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "No permission"})
			return
		}

		c.Next()
	}
}
