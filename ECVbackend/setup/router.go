package setup

import (
	"github.com/gin-gonic/gin"

	"ecvbackend/handler"
	"ecvbackend/middleware"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	router.GET("/api/helloworld/", handler.HelloToGuest())
	router.POST("/api/register/", handler.Register())
	router.POST("/api/login/", handler.Login())
	authorized := router.Group("/api/user", middleware.TokenAuthMiddleware())
	{
		authorized.GET("helloworld/", handler.HelloToUser())
	}
	return router
}
