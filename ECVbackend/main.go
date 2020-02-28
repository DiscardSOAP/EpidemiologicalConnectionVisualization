package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/api/helloworld/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "Hello World from Gin backend!",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
