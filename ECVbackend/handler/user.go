package handler

import (
	"crypto/md5"
	"ecvbackend/lib"
	"ecvbackend/model"
	"log"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"regexp"
	"fmt"
	"time"
	"io"
	"strconv"
)

type LoginJSON struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterJSON struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
	Email string `json:"email" binding:"required"`
	InvitationCode string `json:"invitationCode" binding:"required"`
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
		if c.BindJSON(&data) != nil {
			c.JSON(400, gin.H{"msg": "Invalid Params!"})
			return
		}
		result := model.User{Username: data.Username}.Get()
		log.Println(result)
		if result != nil && lib.ComparePassword(data.Password, result.Password) {
			if _, err := lib.GetToken(data.Username); err == nil {
				session := sessions.Default(c)
				session.Set("loginuser", data.Username)
				session.Save()
				//c.Request.URL.Path = "/api/helloworld"
				c.JSON(200, gin.H{})
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
		if match,_:=regexp.MatchString("[\x21-\x7b]{8,16}",data.Username); match == false{
			c.JSON(400, gin.H{"msg": "Username Invalid!"})
			return
		}
		if match,_:=regexp.MatchString("[\x21-\x7b]{8,16}",data.Password); match == false{
			c.JSON(400, gin.H{"msg": "Password Invalid!"})
			return
		}
		if match,_:=regexp.MatchString("[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[\\w](?:[\\w-]*[\\w])?",data.Email); match == false{
			c.JSON(400, gin.H{"msg": "Email Invalid!"})
			return
		}
		if match,_:=regexp.MatchString("\\w{20}",data.InvitationCode); match == false{
			c.JSON(400, gin.H{"msg": "InvatationCode Invalid!"})
			return
		}
		if data.ConfirmPassword!=data.Password{
			c.JSON(400, gin.H{"msg": "Passworld not match!"})
			return
		}
		result := model.User{Username: data.Username}.Get()
		if result != nil {
			c.JSON(400, gin.H{"msg": "Username Exists!"})
			return
		}
		result2 := model.User{Email: data.Email}.Get()
		if result2 != nil {
			c.JSON(400, gin.H{"msg": "Email Exists!"})
			return
		}
		encryptedPassword, err := lib.EncryptPassword(data.Password)
		if err != nil || encryptedPassword == "" {
			c.JSON(400, gin.H{"msg": "Internal Error!"})
			return
		}
		result = model.User{
				Username: data.Username, 
				Password: encryptedPassword, 
				Email:data.Email,
				Md5:fmt.Sprintf("%x", md5.Sum([]byte(data.Email))),
				Nickname:"",
				Description:"",
				Organization:"",
				Birth:time.Now().Format("2006-01-02 15:04:05"),
		}.Insert()
		if result!=nil{
			c.JSON(200, gin.H{"msg": "Register Success!"})
		}else{
			c.JSON(200, gin.H{"msg": "internal error!"})
		}
	}
}

func GetProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		loginuser := session.Get("loginuser")
		username, ok := loginuser.(string)
		if !ok {
			c.JSON(400, gin.H{"msg": "Internal Error"})
			return
		}
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
		session := sessions.Default(c)
		loginuser := session.Get("loginuser")
		username, ok := loginuser.(string)
		if !ok {
			c.JSON(400, gin.H{"msg": "Internal Error"})
			return
		}
		var data UserProfileJson
		if c.BindJSON(&data) != nil {
			c.JSON(400, gin.H{"msg": "Invalid Params!"})
			return
		}

		Userdata := model.User{Username: username}.Get()
		
		var newPassword string
		if data.NewPassword != ""{
			newPassword,_=lib.EncryptPassword(data.NewPassword)
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

func GenToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data LoginJSON
		if c.BindJSON(&data) != nil {
			c.JSON(400, gin.H{"msg": "Invalid Params!"})
			return
		}
		result := model.User{Username: data.Username}.Get()
		log.Println(result)
		if result != nil && lib.ComparePassword(data.Password, result.Password) {
			if _, err := lib.GetToken(data.Username); err == nil {
				session := sessions.Default(c)
				session.Set("loginuser", data.Username)
				session.Save()
				h := md5.New()
				crutime := time.Now().Unix()
				io.WriteString(h, strconv.FormatInt(crutime, 10))
				token := fmt.Sprintf("%x", h.Sum(nil))
				c.JSON(200,gin.H{
					"token":string(token),
				},)
				return
			}
		}
		c.JSON(400, gin.H{"msg": "Get token Error!"})
	}
}