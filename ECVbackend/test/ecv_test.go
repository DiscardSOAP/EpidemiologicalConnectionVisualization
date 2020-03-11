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

func TestRegister(t *testing.T){
	router := setup.SetupRouter()
	uri := "/api/register/"
	param := make(map[string]interface{})
	param["username"] = "jkxing1234"
	param["password"] = "12345678"
	param["confirmPassword"] = "12345678"
	param["email"] = "jkxing1210@gmail.com"
	param["invitationCode"] = "11111111111111111111"
    body := PostJson(uri, param, router)
	var response map[string]string
	json.Unmarshal(body, &response)
	value, _ := response["msg"]
	assert.Equal(t, "Register Success!", value)
	return
}
func TestLogin(t *testing.T) {
	
	router := setup.SetupRouter()
	uri := "/api/login/"
	param := make(map[string]interface{})
	param["username"] = "jkxing1234"
	param["password"] = "12345678"
	jsonByte,_ := json.Marshal(param)
    req := httptest.NewRequest("POST", uri, bytes.NewReader(jsonByte))
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    result := w.Result()
    defer result.Body.Close()
	cookies := result.Cookies()
	uri = "/api/profile/"
	w = httptest.NewRecorder()
    req = httptest.NewRequest("GET", uri, nil)
	for k := range cookies{
		http.SetCookie(w, cookies[k])
		req.AddCookie(cookies[k])
		fmt.Println(cookies[k].Name)
	}
	router.ServeHTTP(w, req)
	result = w.Result()
	var response map[string]string
    body, _ := ioutil.ReadAll(result.Body)
	json.Unmarshal(body, &response)
	fmt.Println(response)
	value, _ := response["username"]
	assert.Equal(t,"jkxing1234",value)
	
	param = make(map[string]interface{})
	param["name"] = "skylines"
	param["password"] = "12345678"
	
	jsonByte,_ = json.Marshal(param)
    req = httptest.NewRequest("POST", uri, bytes.NewReader(jsonByte))
	w = httptest.NewRecorder()
	for k := range cookies{
		http.SetCookie(w, cookies[k])
		req.AddCookie(cookies[k])
		fmt.Println(cookies[k].Name)
	}
	router.ServeHTTP(w, req)
	result = w.Result()
    body, _ = ioutil.ReadAll(result.Body)
	json.Unmarshal(body, &response)
	fmt.Println(response)
	value, _ = response["msg"]
	assert.Equal(t,"edit profile success",value)
}

func TestHelloWorld(t *testing.T){
	/*router := setup.SetupRouter()
	body := Get("/api/helloworld/",router,t)
	var response map[string]string
	json.Unmarshal(body, &response)
	value, _ := response["msg"]
	assert.Equal(t, "Hello Guest!", value)*/
}