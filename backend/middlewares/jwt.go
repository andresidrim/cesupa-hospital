package middlewares

import (
	"net/http"
	"strings"

	"github.com/andresidrim/cesupa-hospital/enums"
	us "github.com/andresidrim/cesupa-hospital/services/users"
	"github.com/andresidrim/cesupa-hospital/utils"
	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware(userService us.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Missing or invalid Authorization header"})
			return
		}
		tokenString := strings.TrimPrefix(auth, "Bearer ")

		userID, err := utils.ParseJWT(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token: " + err.Error()})
			return
		}

		user, err := userService.Get(uint64(userID))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "User not found"})
			return
		}

		c.Set("userID", userID)
		c.Set("role", enums.Role(user.Role))

		c.Next()
	}
}
