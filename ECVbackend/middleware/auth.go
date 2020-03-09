package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"

	"ecvbackend/lib"

	"github.com/gin-contrib/sessions"

	"fmt"
)


func GetSession(c *gin.Context) bool {
    session := sessions.Default(c)
    loginuser := session.Get("loginuser")
    fmt.Println("loginuser:", loginuser)
    if loginuser != nil {
        return true
    } else {
        return false
    }
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isLogin := GetSession(c)
		fmt.Println("isLogin:",isLogin)
		if isLogin == false {
			c.JSON(400, gin.H{"msg": "NOT LOGIN!"})
			c.Abort()
			return
		}
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
