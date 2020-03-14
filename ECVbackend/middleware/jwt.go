package middleware

import (
	"ecvbackend/pkg/e"
	"ecvbackend/pkg/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.Request.Header.Get("Authorization")
		/*var header string, token string
		  if len(tokens)==2{
		      header = tokens[0]
		      token =  tokens[1]
		      if(header!="Bearer"){
		          log.Println("token invalid")
		          code = e.INVALID_PARAMS
		      }
		  }else{
		      token = tokens[0]
		  }*/
		if token == "" {
			log.Println("token invalid")
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
			c.JSON(http.StatusUnauthorized, gin.H{
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
