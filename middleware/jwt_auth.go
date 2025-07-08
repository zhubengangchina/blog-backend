package middleware

import (
	"blog-backend/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "缺少Token"})
			c.Abort()
			return
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ParseJWT(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token无效"})
			c.Abort()
			return
		}

		// 存入上下文
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Next()
	}
}
