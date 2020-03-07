package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type HelloWorldJSON struct {
	Msg string
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestPublicHelloWorld(t *testing.T) {
	router := SetupRouter()
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/helloworld/", nil)
	router.ServeHTTP(response, request)
	var data HelloWorldJSON
	err := json.Unmarshal([]byte(response.Body.String()), &data)
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "Hello Guest!", data.Msg)
}

func TestMain(m *testing.M) {
	m.Run()
}
