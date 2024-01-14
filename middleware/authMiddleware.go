package middleware

import (
	helper "golang-restaurant-management/helpers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			c.Abort()
			return
		}

		authHeader := c.GetHeader("Authorization")
		if token == "" {
			if authHeader == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "No Authorization header provided"})
				c.Abort()
				return
			}

			tokenParts := strings.Split(authHeader, " ")
			if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
				c.Abort()
				return
			}

			token = tokenParts[1] // Update this line to assign the value to the outer token variable

			claims, err := helper.ValidateToken(token)
			if err != "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
				c.Abort()
				return
			}

			c.Set("email", claims.Email)
			c.Set("first_name", claims.First_name)
			c.Set("last_name", claims.Last_name)
			c.Set("uid", claims.Uid)

			c.Next()
		}
	}
}
