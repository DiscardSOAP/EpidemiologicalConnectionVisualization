package test
import (
	"encoding/json"
	"ecvbackend/setup"
	"net/http/httptest"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

type User struct {
	UserName string `json:"user_name" form:"user_name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Age      int    `form:"age" json:"age" binding:"required"`
}

func Get(uri string, router *gin.Engine,t *testing.T) string {
	req := httptest.NewRequest("GET", uri, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	result := w.Result()
	defer result.Body.Close()

	var response map[string]string
	json.Unmarshal([]byte(w.Body.String()), &response)
	value, _ := response["msg"]
	return value
}
func TestAll(t *testing.T){
	router := setup.SetupRouter()
	value := Get("/api/helloworld/",router,t)
	assert.Equal(t, "Hello Guest!", value)
}