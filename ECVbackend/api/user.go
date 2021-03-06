package api

import (
	"crypto/md5"
	"ecvbackend/pkg/util"
	"ecvbackend/model"
	"github.com/gin-gonic/gin"
	"fmt"
	"time"
	"github.com/astaxie/beego/validation"
)

type LoginJSON struct {
	Username string `json:"username" binding:"required" validate:"max=16,min=8"`
	Password string `json:"password" binding:"required" validate:"max=16,min=8"`
}

type RegisterJSON struct {
	Username string `json:"username" binding:"required" validate:"max=16,min=8"`
	Password string `json:"password" binding:"required" validate:"max=16,min=8"`
	ConfirmPassword string `json:"confirmPassword" binding:"required" validate:"eqfield=Password"`
	Email string `json:"email" binding:"required" validate:"email"`
	InvitationCode string `json:"invitationCode" binding:"required" validate:"max=20,min=20"`
}


type UserProfileJson struct {
	Name string `json:"name"`
	Organization string `json:"organization" `
	Description string `json:"description"`
	Password string `json:"password" binding:"required"`
	NewPassword string `json:"newPassword" `
	ConfirmNewPassword string `json:"confirmNewPassword"`
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data LoginJSON
		util.RuntimeLog.Println("user login from "+c.ClientIP())
		if c.BindJSON(&data) != nil {
			c.JSON(400, gin.H{"msg": "Invalid Params!"})
			return
		}
		valid := validation.Validation{}
		ok, _ := valid.Valid(&data)
		if !ok{
			c.JSON(400, gin.H{"msg": "Invalid Params!"})
			util.RuntimeLog.Println("Login failed")
			return
		}
		result := model.User{Username: data.Username}.Get()
		if result != nil && util.ComparePassword(data.Password, result.Password) {
			util.SaveAuthSession(c,data.Username)
            token, _ := util.GenerateToken(data.Username, data.Password)
			c.JSON(200, gin.H{"token":token})
			util.RuntimeLog.Println("user %s login success, token is %d\n",data.Username,token)
			return
		}
		c.JSON(400, gin.H{"msg": "Login Error!"})
	}
}

func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		util.RuntimeLog.Println("user logout from "+c.ClientIP())
		util.ClearAuthSession(c)
		c.JSON(200, gin.H{})
	}
}

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data RegisterJSON
		if c.BindJSON(&data) != nil {
			c.JSON(400, gin.H{"msg": "Invalid Params!"})
			return
		}
		
		valid := validation.Validation{}
		ok, err := valid.Valid(&data)
		if !ok{
			c.JSON(400, gin.H{"msg": "Param Invalid"})
			return
		}
		checkUsername := model.User{Username: data.Username}.Get()
		if checkUsername != nil {
			c.JSON(400, gin.H{"msg": "Username Exists!"})
			return
		}
		checkEmail := model.User{Email: data.Email}.Get()
		if checkEmail != nil {
			c.JSON(400, gin.H{"msg": "Email Exists!"})
			return
		}
		encryptedPassword, err := util.EncryptPassword(data.Password)
		if err != nil || encryptedPassword == "" {
			c.JSON(400, gin.H{"msg": "Internal Error!"})
			return
		}
		insResult := model.User{
				Username: data.Username, 
				Password: encryptedPassword, 
				Email:data.Email,
				Md5:fmt.Sprintf("%x", md5.Sum([]byte(data.Email))),
				Nickname:data.Username,
				Description:"User",
				Organization:"Anonymous",
				Birth:time.Now().Format("2006-01-02 15:04:05"),
				Root:false,
		}.Insert()
		util.RuntimeLog.Println(insResult)
		if insResult!=nil{
			util.RuntimeLog.Println("IP ",c.ClientIP()," register success")
			c.JSON(200, gin.H{"msg": "Register Success!"})
		}else{
			util.RuntimeLog.Println("IP ",c.ClientIP()," register failed")
			c.JSON(200, gin.H{"msg": "internal error!"})
		}
	}
}
