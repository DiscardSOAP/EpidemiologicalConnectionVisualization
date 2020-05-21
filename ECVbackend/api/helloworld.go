package api

import (
	"ecvbackend/model"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func TestError() gin.HandlerFunc {
	return func(c *gin.Context) {
		ints := []int{1, 2, 5}
		fmt.Println(ints[5])
	}
}

func HelloToUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		loginuser := session.Get("loginuser")
		username, ok := loginuser.(string)
		if !ok {
			c.JSON(400, gin.H{"msg": "Internal Error"})
		}
		c.JSON(200, gin.H{"msg": "Hello user " + username})
	}
}

func VisitNumber() gin.HandlerFunc {
	return func(c *gin.Context) {
		vis := new(model.Visit)
		vis.StatusCode = 200
		count := vis.Count()
		c.JSON(200, gin.H{"count": count})
	}
}
