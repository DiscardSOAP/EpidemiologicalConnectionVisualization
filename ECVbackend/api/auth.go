package api

import (
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
    }else{
        util.RuntimeLog.Println("login data binding error")
    }
    valid := validation.Validation{}
    a := auth{Username: userData.Username, Password: userData.Password}
    ok, err := valid.Valid(&a)
    if err!=nil{
        util.RuntimeLog.Println("login data valid error")
    }
    //data := make(map[string]interface{})
    if ok {
        user := model.User{Username:userData.Username}.Get()
        if user!=nil && util.ComparePassword(userData.Password, user.Password){
            util.RuntimeLog.Println(userData)
            token, err := util.GenerateToken(userData.Username, userData.Password)
            if err != nil {
                util.RuntimeLog.Println(err)
            } else {
                util.RuntimeLog.Println("login success")
                c.JSON(http.StatusOK, gin.H{
                    "token" : token,
                })
                return
            }

        }
    } else {
        for _, err := range valid.Errors {
            util.RuntimeLog.Println(err.Key, err.Message)
        }
    }
    c.JSON(400, gin.H{
        "msg" : "error code",
    })

}