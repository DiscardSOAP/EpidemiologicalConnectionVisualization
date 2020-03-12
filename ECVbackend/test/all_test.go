package test
import (
	"encoding/json"
	"ecvbackend/setup"
	"net/http/httptest"
	"net/http"
	"github.com/stretchr/testify/assert"
	"testing"
	"io/ioutil"
	"bytes"
	"os"
	"path"
	"runtime"
	"fmt"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

var router = setup.SetupRouter()

type userProfileJson struct {
	Username string `json:"username"`
	Name string `json:"name"`
	Birth string `json:"birth"`
	Organization string `json:"organization" `
	Description string `json:"description"`
	Email string `json:"email" binding:"required"`
	Md5 string `json:"email_md5" `
}

func Post(uri string, param map[string]interface{}) *http.Response{
	jsonByte,_ := json.Marshal(param)
	req := httptest.NewRequest("POST", uri, bytes.NewReader(jsonByte))
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
	result := w.Result()
	return result
}

func ReqWithCookies(method string, uri string, param map[string]interface{}, cookies []*http.Cookie,token string) *http.Response{
	jsonByte,_ := json.Marshal(param)
	req := httptest.NewRequest(method, uri, bytes.NewReader(jsonByte))
    req.Header.Add("token", token)
	w := httptest.NewRecorder()
	for k := range cookies{
		http.SetCookie(w, cookies[k])
		req.AddCookie(cookies[k])
	}
    router.ServeHTTP(w, req)
	result := w.Result()
	return result
}

func TestLogin(t *testing.T) {
	
	param := make(map[string]interface{})
	param["username"] = "jkxing1234"
	param["password"] = "12345678"
	var response map[string]interface{}
	res:=Post("/api/login/",param)
	body,_ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	cookies := res.Cookies()
	json.Unmarshal(body, &response)
	token:=response["token"].(string)
	res = ReqWithCookies("GET","/api/profile/",nil,cookies,token)
    body, _ = ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	//value, _ := response["username"]
	assert.Equal(t,"jkxing1234",response["user"].(map[string]interface{})["username"])
	
	param = make(map[string]interface{})
	param["name"] = "skylines"
	param["password"] = "12345678"
	
	res = ReqWithCookies("POST","/api/profile/",param,cookies,token)
	body, _ = ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	fmt.Println(response)
	value, _ := response["msg"]
	assert.Equal(t,"edit profile success",value)
}

func TestHelloWorld(t *testing.T){
	/*req := httptest.NewRequest("GET", "/api/helloworld/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	result := w.Result()
    body, _ := ioutil.ReadAll(result.Body)
	var response map[string]string
	json.Unmarshal(body, &response)
	value, _ := response["msg"]
	assert.Equal(t, "Hello Guest!", value)*/
}


func TestRegister(t *testing.T){
	/*uri := "/api/register/"
	param := make(map[string]interface{})
	param["username"] = "jkxing1234"
	param["password"] = "12345678"
	param["confirmPassword"] = "12345678"
	param["email"] = "jkxing1210@gmail.com"
	param["invitationCode"] = "11111111111111111111"
	jsonByte,_ := json.Marshal(param)
	req := httptest.NewRequest("POST", uri, bytes.NewReader(jsonByte))
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    result := w.Result()
    defer result.Body.Close()
    body, _ := ioutil.ReadAll(result.Body)
	var response map[string]string
	json.Unmarshal(body, &response)
	value, _ := response["msg"]
	assert.Equal(t, "Register Success!", value)
	return*/
}