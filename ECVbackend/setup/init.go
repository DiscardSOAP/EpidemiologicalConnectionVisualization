package setup

import (
	"ecvbackend/api"
	"ecvbackend/middleware"
	"ecvbackend/model"
	"ecvbackend/pkg/util"

	"github.com/gin-gonic/gin"
)

func Init() {
	vis := new(model.Visit)
	count := vis.Count()
	middleware.GlobalCount = count
	pass, _ := util.EncryptPassword("12345678")
	rootUser := model.User{
		Username: "rootroot",
		Root:     true,
	}
	if rootUser.Get() == nil {
		rootUser.Password = pass
		rootUser.Insert()
	}
}

func SetupRouter() *gin.Engine {
	gin.DefaultErrorWriter = util.ErrorLog.Writer()
	router := gin.Default()
	router.Use(middleware.AccessLoggerToFileAndDB())
	router.Use(middleware.CORSMiddleware())
	router.Use(util.EnableCookieSession())
	router.GET("/api/error/", api.TestError())
	router.GET("/api/visitnumber/", api.VisitNumber())
	router.POST("/api/register/", api.Register())
	router.POST("/api/login/", api.Login())
	router.POST("/api/auth/", api.GetAuth)
	router.POST("/api/track/", api.GetTrack())
	router.POST("/api/newslib/", api.GetNews())
	router.POST("/api/topology/", api.GetTopo())
	router.GET("/api/topology/", api.GetTopoInfo())
	router.GET("/api/news/:id", api.GetOne())
	User := router.Group("/api", middleware.JWT())
	{
		User.POST("/news", api.PostOne())
		User.POST("/news/:id", api.ModifyOne())
		User.DELETE("/news/:id", api.DeleteOne())
		User.DELETE("/logout/", api.Logout())
		User.GET("/profile/", api.GetProfile())
		User.GET("/profile/:username", api.GetUserProfile())
		User.POST("/profile/", api.ChangeProfile())
	}
	rootUser := router.Group("/api/root", middleware.JWT(), middleware.RootJWT())
	{
		rootUser.GET("/manage/", api.GetAllProfile())
		rootUser.DELETE("/manage/", api.DeleteProfile())
	}
	return router
}
