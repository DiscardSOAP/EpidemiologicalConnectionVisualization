package handler

import (
	"ecvbackend/lib"
	"ecvbackend/model"
	"log"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginJSON struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterJSON struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data LoginJSON
		c.JSON(400, gin.H{"msg": "Invalid Params!"})
		if c.BindJSON(&data) != nil {
			return
		}
		result := model.User{Username: data.Username}.Get()
		log.Println(result)
		if result != nil && lib.ComparePassword(data.Password, result.Password) {
			if tokenString, err := lib.GetToken(data.Username); err == nil {
				//c.SetCookie("sessionID","1",60,"/","127.0.0.1",false,true)
				session := sessions.Default(c)
				session.Set("loginuser", data.Username)
				session.Save()
				c.JSON(200, gin.H{"token": tokenString})
				return
			}
		}
		c.JSON(400, gin.H{"msg": "Login Error!"})
	}
}

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data RegisterJSON
		if c.BindJSON(&data) != nil {
			c.JSON(400, gin.H{"msg": "Invalid Params!"})
			return
		}
		result := model.User{Username: data.Username}.Get()
		if result != nil {
			c.JSON(400, gin.H{"msg": "Register Error!"})
			return
		}
		encryptedPassword, err := lib.EncryptPassword(data.Password)
		if err != nil || encryptedPassword == "" {
			c.JSON(400, gin.H{"msg": "Register Error!"})
			return
		}
		result = model.User{Username: data.Username, Password: encryptedPassword}.Insert()
		c.JSON(200, gin.H{"msg": "Register Success!"})
	}
}

func Profile() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		loginuser := session.Get("loginuser")
		username, ok := loginuser.(string)
		if !ok {
			c.JSON(400, gin.H{"msg": "Internal Error"})
			return
		}
		pass := model.User{Username: username}.Get().Password
		c.JSON(200,gin.H{
			"Username":string(username),
			"Password":pass,
		},)
	}
}
