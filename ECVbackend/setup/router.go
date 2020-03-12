package setup

import (
	"github.com/gin-gonic/gin"

	"ecvbackend/api"
	"ecvbackend/middleware"
	"ecvbackend/pkg/util"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
    router.Use(util.EnableCookieSession())
	router.Use(middleware.CORSMiddleware())
	router.GET("/api/helloworld/", api.HelloToGuest())
	router.POST("/api/register/", api.Register())
	router.POST("/api/login/", api.Login())
	router.POST("/api/auth/",api.GetAuth)
	router.GET("/api/profile/", middleware.CookieAuth(),middleware.JWT(),api.GetProfile())
	router.POST("/api/profile/", middleware.CookieAuth(),middleware.JWT(), api.ChangeProfile())
	return router
}
