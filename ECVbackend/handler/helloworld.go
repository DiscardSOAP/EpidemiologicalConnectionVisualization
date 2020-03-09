package handler

import "github.com/gin-gonic/gin"

func HelloToGuest() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "Hello Guest!"})
	}
}
func HelloToUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "Hello User!"})
	}
}
