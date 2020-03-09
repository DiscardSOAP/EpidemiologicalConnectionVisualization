package setup

import (
	"github.com/gin-gonic/gin"

	"ecvbackend/handler"
	"ecvbackend/middleware"

	"github.com/gin-contrib/sessions"

	"github.com/gin-contrib/sessions/cookie"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	store := cookie.NewStore([]byte("loginuser"))
	router.Use(middleware.CORSMiddleware())
	router.Use(sessions.Sessions("Sessions", store))
	router.GET("/api/helloworld/", handler.HelloToGuest())
	router.POST("/api/register/", handler.Register())
	router.POST("/api/login/", handler.Login())
	authorized := router.Group("/api/user", middleware.TokenAuthMiddleware())
	{
		authorized.GET("helloworld/", handler.HelloToUser())
		authorized.GET("profile/", handler.Profile())
	}
	return router
}
