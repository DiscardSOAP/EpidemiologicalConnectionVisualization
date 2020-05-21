package api

import (
	"ecvbackend/model"
	"ecvbackend/pkg/util"
	"github.com/gin-gonic/gin"
	//"encoding/json"
	//"time"
)

func GetTopo() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		var data map[string]string
		if c.BindJSON(&data) != nil {
			c.JSON(400, gin.H{"msg": "Invalid Params!"})
			return
		}
		res := model.FindRel(data["category"])
		util.RuntimeLog.Println(c.ClientIP(),"request topo")
		c.JSON(200, res)
	}
}

func GetTopoInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := model.GetCatCount()
		util.RuntimeLog.Println(c.ClientIP(),"request topoinfo")
		c.JSON(200, res)
	}
}