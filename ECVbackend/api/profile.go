package api

import (
	"ecvbackend/model"
	"ecvbackend/pkg/util"
	"github.com/gin-gonic/gin"
)

func GetProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.MustGet("username").(string)
		Userdata := model.User{Username: username}.Get()
		c.JSON(200, gin.H{"user": gin.H{
			"username":     string(Userdata.Username),
			"name":         string(Userdata.Nickname),
			"birth":        string(Userdata.Birth),
			"organization": string(Userdata.Organization),
			"description":  string(Userdata.Description),
			"email":        string(Userdata.Email),
			"email_md5":    string(Userdata.Md5)},
		})
	}
}

func GetUserProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")
		Userdata := model.User{Username: username}.Get()
		util.RuntimeLog.Println(username,"request userdata")
		if Userdata!=nil{
			c.JSON(200, gin.H{"user": gin.H{
				"username":     string(Userdata.Username),
				"name":         string(Userdata.Nickname),
				"birth":        string(Userdata.Birth),
				"organization": string(Userdata.Organization),
				"description":  string(Userdata.Description),
				"email":        string(Userdata.Email),
				"email_md5":    string(Userdata.Md5)},
			})
		}else{
			c.JSON(200, gin.H{"msg": "No such user"})
		}
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
		if data.NewPassword != "" {
			newPassword, _ = util.EncryptPassword(data.NewPassword)
		} else {
			newPassword = Userdata.Password
		}

		model.User{
			Username:     username,
			Password:     newPassword,
			Email:        Userdata.Email,
			Md5:          Userdata.Md5,
			Nickname:     data.Name,
			Organization: data.Organization,
			Birth:        Userdata.Birth,
			Description:  data.Description,
		}.Update()
		c.JSON(200, gin.H{
			"msg": "edit profile success",
		})
	}
}

type UserProfile struct {
	Username     string `json:"username"`
	Name         string `json:"name"`
	Organization string `json:"organization" `
	Birth        string `json:"birth"`
	Description  string `json:"description"`
	Email        string `json:"email" binding:"required"`
	Md5          string `json:"email_md5" `
}

func GetAllProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		Users := model.User{}.FindAll()
		userProfiles := make([]UserProfile, 0)
		for i := 0; i < len(Users); i++ {
			userProfiles = append(userProfiles, UserProfile{
				Username:     Users[i].Username,
				Name:         Users[i].Nickname,
				Organization: Users[i].Organization,
				Birth:        Users[i].Birth,
				Description:  Users[i].Description,
				Email:        Users[i].Email,
				Md5:          Users[i].Md5,
			})
		}
		c.JSON(200, gin.H{"users": userProfiles})
	}
}

func DeleteProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data map[string][]map[string]string
		if c.BindJSON(&data) != nil {
			c.JSON(400, gin.H{"msg": "Invalid Params!"})
			return
		}
		users := data["users"]
		for i := 0; i < len(users); i++ {
			model.User{Username: users[i]["username"]}.Delete()
		}
		c.JSON(200, gin.H{"msg": "delete user success"})
	}
}
