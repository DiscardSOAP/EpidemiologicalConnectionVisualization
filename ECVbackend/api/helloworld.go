package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
)

func HelloToGuest() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "Hello Guest!"})
	}
}
func HelloToUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		loginuser := session.Get("loginuser")
		username,ok := loginuser.(string)
		if !ok {
			c.JSON(400, gin.H{"msg": "Internal Error"})
		}
		c.JSON(200, gin.H{"msg": "Hello user "+username})
	}
}
