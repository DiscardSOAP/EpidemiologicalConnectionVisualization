package api

import (
	"github.com/gin-gonic/gin"
	"ecvbackend/model"
	"ecvbackend/pkg/util"
)

func GetProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.MustGet("username").(string)
		Userdata := model.User{Username: username}.Get()
		c.JSON(200,gin.H{"user":gin.H{
			"username":string(Userdata.Username),
			"name":string(Userdata.Nickname),
			"birth":string(Userdata.Birth),
			"organization":string(Userdata.Organization),
			"description":string(Userdata.Description),
			"email":string(Userdata.Email),
			"email_md5":string(Userdata.Md5),},
		},)
	}
}

func ChangeProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.MustGet("username").(string)
		var data UserProfileJson
		if c.BindJSON(&data) != nil {
			c.JSON(400, gin.H{"msg": "Invalid Params!"})
			return
		}

		Userdata := model.User{Username: username}.Get()
		
		var newPassword string
		if data.NewPassword != ""{
			newPassword,_=util.EncryptPassword(data.NewPassword)
		} else{
			newPassword = Userdata.Password
		}

		model.User{
			Username: username,
			Password: newPassword,
			Email: Userdata.Email,
			Md5: Userdata.Md5,
			Nickname: data.Name,
			Organization: data.Organization,
			Birth: Userdata.Birth,
			Description: data.Description,
			}.Update()
		c.JSON(200,gin.H{
			"msg":"edit profile success",
		},)
	}
}