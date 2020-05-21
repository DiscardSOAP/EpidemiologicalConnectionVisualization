package middleware

import (
	"ecvbackend/pkg/e"
	"ecvbackend/pkg/util"
	"ecvbackend/model"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = e.SUCCESS
		tokens := strings.Split(c.Request.Header.Get("Authorization")," ")
		var header,token string;
		  if len(tokens)==2{
		      header = tokens[0]
		      token =  tokens[1]
		      if(header!="Bearer"){
		          util.RuntimeLog.Println("token invalid")
		          code = e.INVALID_PARAMS
		      }
		  }else{
		      token = tokens[0]
		  }
		if token == "" {
			util.RuntimeLog.Println("token invalid")
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
			c.Set("username", claims.Username)
		}

		if code != e.SUCCESS {
			c.JSON(401, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}
		c.Next()
	}
}


func RootJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		username:=c.MustGet("username").(string)
		user:=model.User{Username:username}.Get()
		if user.Root==false {
				c.JSON(401, gin.H{
				"msg":  "not root",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}