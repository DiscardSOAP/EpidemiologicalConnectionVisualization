package test
import (
	"encoding/json"
	"ecvbackend/setup"
	"net/http/httptest"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
	"io/ioutil"
	"bytes"
	"fmt"
)

type User struct {
	UserName string `json:"user_name" form:"user_name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Age      int    `form:"age" json:"age" binding:"required"`
}

func Get(uri string, router *gin.Engine,t *testing.T) []byte {
	req := httptest.NewRequest("GET", uri, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	result := w.Result()
	defer result.Body.Close()
    body, _ := ioutil.ReadAll(result.Body)
    return body
}

func PostJson(uri string, param map[string]interface{}, router *gin.Engine) []byte {
    jsonByte,_ := json.Marshal(param)
    req := httptest.NewRequest("POST", uri, bytes.NewReader(jsonByte))
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    result := w.Result()
    defer result.Body.Close()
    body, _ := ioutil.ReadAll(result.Body)
    return body
}

func TestRegisterLogin(t *testing.T) {
	
	router := setup.SetupRouter()
	/*uri := "/api/register/"
	param := make(map[string]interface{})
	param["username"] = "jkxing1234"
	param["password"] = "123456"
    body := PostJson(uri, param, router)
	var response map[string]string
	json.Unmarshal(body, &response)
	value, _ := response["msg"]
	assert.Equal(t, "Register Success!", value)
	return*/
	uri := "/api/login/"
	param := make(map[string]interface{})
	param["username"] = "jkxing123"
	param["password"] = "123456"

	jsonByte,_ := json.Marshal(param)
	
    req := httptest.NewRequest("POST", uri, bytes.NewReader(jsonByte))
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    result := w.Result()
    defer result.Body.Close()
	//login_body, _ := ioutil.ReadAll(result.Body)
	cookies := result.Cookies()
	uri = "/api/user/profile/"
	w = httptest.NewRecorder()
	//request := &http.Request{Header: http.Header{"Cookie": recorder.HeaderMap["Set-Cookie"]}}
    req = httptest.NewRequest("GET", uri, nil)
	for k := range cookies{
		http.SetCookie(w, cookies[k])
		req.AddCookie(cookies[k])
		fmt.Println(cookies[k].Name)
	}
	router.ServeHTTP(w, req)
	result = w.Result()
	cookies = result.Cookies()
	for k := range cookies{
		fmt.Println(cookies[k])
	}
	var response map[string]string
    body, _ := ioutil.ReadAll(result.Body)
	json.Unmarshal(body, &response)
	value, _ := response["Username"]
	assert.Equal(t,"jkxing123",value)
}

func TestHelloWorld(t *testing.T){
	router := setup.SetupRouter()
	body := Get("/api/helloworld/",router,t)
	var response map[string]string
	json.Unmarshal(body, &response)
	value, _ := response["msg"]
	assert.Equal(t, "Hello Guest!", value)
}