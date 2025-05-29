package middlewares

import (
	"net/http"

	"github.com/andresidrim/cesupa-hospital/enums"
	"github.com/gin-gonic/gin"
	"slices"
)

func RoleMiddleware(allowedRoles ...enums.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		v, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "role not provided"})
			return
		}
		userRole, ok := v.(enums.Role)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "invalid role type"})
			return
		}
		if slices.Contains(allowedRoles, userRole) {
			c.Next()
			return
		}
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "access forbidden"})
	}
}
