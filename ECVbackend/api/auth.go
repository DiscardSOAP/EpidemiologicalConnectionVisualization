package api

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/astaxie/beego/validation"
    "ecvbackend/pkg/util"
    "ecvbackend/model"
)

type auth struct {
    Username string `valid:"Required; MaxSize(16)"`
    Password string `valid:"Required; MaxSize(16)"`
}

func GetAuth(c *gin.Context) {
    var userData LoginJSON
    if c.BindJSON(&userData) != nil {
        c.JSON(400, gin.H{"msg": "Invalid Params!"})
        return
    }

    valid := validation.Validation{}
    a := auth{Username: userData.Username, Password: userData.Password}
    ok, _ := valid.Valid(&a)

    //data := make(map[string]interface{})
	
    if ok {
        user := model.User{Username:userData.Username}.Get()
        if user!=nil && util.ComparePassword(userData.Password, user.Password){
            log.Println(userData)
            token, err := util.GenerateToken(userData.Username, userData.Password)
            if err != nil {
                //code = e.ERROR_AUTH_TOKEN
            } else {
                
                c.JSON(http.StatusOK, gin.H{
                    "token" : token,
                })
                return
                //data["token"] = token
                //code = e.SUCCESS
            }

        } else {
            //code = e.ERROR_AUTH
        }
    } else {
        for _, err := range valid.Errors {
            log.Println(err.Key, err.Message)
        }
    }
    
    c.JSON(400, gin.H{
        "msg" : "error code",
    })

}