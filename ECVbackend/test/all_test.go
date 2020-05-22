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
	"log"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	f, err := os.OpenFile("test.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
    if err != nil {
        log.Fatal(err)
    }
    log.SetOutput(f)
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime | log.Lshortfile )
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

func Req(method string,uri string, param map[string]interface{}) *http.Response{
	jsonByte,_ := json.Marshal(param)
	req := httptest.NewRequest(method, uri, bytes.NewReader(jsonByte))
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
	result := w.Result()
	return result
}

func ReqWithCookies(method string, uri string, param map[string]interface{}, cookies []*http.Cookie,token string) *http.Response{
	jsonByte,_ := json.Marshal(param)
	req := httptest.NewRequest(method, uri, bytes.NewReader(jsonByte))
    req.Header.Add("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	for k := range cookies{
		http.SetCookie(w, cookies[k])
		req.AddCookie(cookies[k])
	}
    router.ServeHTTP(w, req)
	result := w.Result()
	return result
}


func TestRegister(t *testing.T){
	/*uri := "/api/register/"
	param := make(map[string]interface{})
	param["username"] = "jkxing12345"
	param["password"] = "123456789"
	param["confirmPassword"] = "123456789"
	param["email"] = "jkxing1211@gmail.com"
	param["invitationCode"] = "11111111111112111111"
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

func TestGetTopo(t *testing.T){
	return
	res1:=ReqWithCookies("POST","/api/newslib/",nil,nil,"")
	var response44 map[string]interface{}
	body,_ := ioutil.ReadAll(res1.Body)
	json.Unmarshal(body, &response44)
	fmt.Println("=============")
	fmt.Println(response44)
}

func TestGetTrack(t *testing.T){
	return
	param := make(map[string]interface{})
	param["category"] = "tianjin"
	res1:=ReqWithCookies("POST","/api/track/",param,nil,"")
	var response44 map[string]interface{}
	body,_ := ioutil.ReadAll(res1.Body)
	json.Unmarshal(body, &response44)
	fmt.Println("=============")
	fmt.Println(response44)
}

func TestGetTrend(t *testing.T){
	param := make(map[string]interface{})
	param["category"] = "tianjin"
	res1:=ReqWithCookies("POST","/api/trend/",param,nil,"")
	var response44 map[string]interface{}
	body,_ := ioutil.ReadAll(res1.Body)
	json.Unmarshal(body, &response44)
	fmt.Println("=============")
	fmt.Println(response44)
}

func TestLogin(t *testing.T) {
	return
	param := make(map[string]interface{})
	param["username"] = "jkxing1234"
	param["password"] = "12345678"
	var response map[string]interface{}
	res:=Req("POST","/api/login/",param)
	body,_ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	cookies := res.Cookies()
	json.Unmarshal(body, &response)
	token:=response["token"].(string)
	res = ReqWithCookies("GET","/api/profile/jkxing12345",nil,cookies,token)
    body, _ = ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	//value, _ := response["username"]
	assert.Equal(t,"jkxing12345",response["user"].(map[string]interface{})["username"])
	param = make(map[string]interface{})
	param["name"] = "skylines"
	param["password"] = "12345678"
	
	res = ReqWithCookies("POST","/api/profile/",param,cookies,token)
	body, _ = ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	//fmt.Println(response)
	value, _ := response["msg"]
	assert.Equal(t,"edit profile success",value)


	res = ReqWithCookies("GET","/api/root/manage/",param,cookies,token)
	body, _ = ioutil.ReadAll(res.Body)
	
	var response1 map[string]interface{}
	json.Unmarshal(body, &response1)
	//fmt.Println(response1)
	assert.Equal(t,"skylines",response1["users"].([]interface{})[0].(map[string]interface{})["name"].(string))
	
	
	/*param = make(map[string]interface{})
	param["users"] = []map[string]string{map[string]string{"username":"jkxing12345"}}
	res = ReqWithCookies("DELETE","/api/root/manage/",param,cookies,token)
	body, _ = ioutil.ReadAll(res.Body)
	
	var response2 map[string]interface{}
	json.Unmarshal(body, &response2)
	fmt.Println(response2)*/
	param = make(map[string]interface{})
	param["title"]="1"
	param["text"]="1"
	param["category"]="tianjin"
	tmp_rel1:=map[string]interface{}{
		"id":3,
		"relationship":"From",
	}
	tmp_rel2:=map[string]interface{}{
		"id":4,
		"relationship":"From",
	}
	param["relationships"]=[]map[string]interface{}{tmp_rel1,tmp_rel2}
	event := map[string]string{
		"startDate":"2020-01-22",
		"startTime":"1:23",
		"endDate":"2020-01-23",
		"endTime":"2:34",
		"addr":"123",
		"description":"321",
	}
	param["events"]=[]map[string]string{event}
	//fmt.Println(param)
	res = ReqWithCookies("POST","/api/news/",param,cookies,token)
	body, _ = ioutil.ReadAll(res.Body)
	var response3 map[string]interface{}
	json.Unmarshal(body, &response3)
	assert.Equal(t,"add news success",response3["msg"])
	//fmt.Println(response3)
	
	res = ReqWithCookies("GET","/api/news/5",nil,cookies,token)
	body, _ = ioutil.ReadAll(res.Body)
	var response4 map[string]interface{}
	json.Unmarshal(body, &response4)
	//fmt.Println(response4)
	
	param["title"]="2"
	param["publishTime"]="2020-03-03T12:48:48+08:00"
	res = ReqWithCookies("POST","/api/news/1",param,cookies,token)
	body, _ = ioutil.ReadAll(res.Body)
	var response5 map[string]string
	json.Unmarshal(body, &response5)
	assert.Equal(t,"change success",response5["msg"])
	
	res = ReqWithCookies("DELETE","/api/news/2",param,cookies,token)
	body, _ = ioutil.ReadAll(res.Body)
	var response6 map[string]string
	json.Unmarshal(body, &response6)
	assert.Equal(t,"delete success",response6["msg"])
	
	param=make(map[string]interface{})
	param["category"]="tianjin"
	res = ReqWithCookies("POST","/api/topology/",param,cookies,token)
	body, _ = ioutil.ReadAll(res.Body)
	var response7 []map[string]interface{}
	json.Unmarshal(body, &response7)
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

