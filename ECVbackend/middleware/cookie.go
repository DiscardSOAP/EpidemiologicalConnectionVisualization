package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"ecvbackend/pkg/util"
)


func GetSession(c *gin.Context) string {
    session := sessions.Default(c)
    loginuser := session.Get("loginuser").(string)
    return loginuser
}

func CookieAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if util.HasSession(c) {
			username := util.GetSessionUserName(c)
			c.Set("username", username)
			c.Next()
		}else {
			c.JSON(400, gin.H{"msg": "NOT LOGIN!"})
			c.Abort()
			return
		}
	}
}
