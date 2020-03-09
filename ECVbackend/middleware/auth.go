package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"

	"ecvbackend/lib"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := string(c.GetHeader("Authorization"))
		if authHeader == "" {
			c.JSON(400, gin.H{"msg": "Authorization NOT FOUND!"})
			c.Abort()
			return
		}
		tokenString := strings.Split(authHeader, " ")[1]
		if !lib.TokenAuthenticate(tokenString) {
			c.JSON(400, gin.H{"msg": "Invalid Token!"})
			c.Abort()
			return
		}
		c.Next()
	}
}
