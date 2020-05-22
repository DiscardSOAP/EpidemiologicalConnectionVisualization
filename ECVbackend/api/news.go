package api

import (
	"ecvbackend/model"
	"ecvbackend/pkg/util"
	"strconv"
	"github.com/gin-gonic/gin"
	"time"
	"strings"
	"sort"
)
func GetTrend() gin.HandlerFunc{
	return func(c *gin.Context) {
		var data struct{
			Category string `json:"category"`;
		}
		if c.BindJSON(&data) != nil {
			c.JSON(400, gin.H{"msg": "Invalid Params!"})
			return
		}
		ans := model.FindTrend(data.Category)
		type Trend struct {
			Open int
			Close int
		}
		ds := make(map[string]Trend,0)
		for i:=0;i<len(ans);i++{
			ed := ""
			st := ""
			for j:=0;j<len(ans[i].Events);j++{
				if len(ed)==0 || (len(ans[i].Events[j].EndDate)>=10&&ans[i].Events[j].EndDate>ed){
					ed = ans[i].Events[j].EndDate
				}
				if len(st)==0 || (len(ans[i].Events[j].StartDate)>=10&&ans[i].Events[j].StartDate<st){
					st = ans[i].Events[j].StartDate
				}
			}
			if st==""{
				st = ans[i].PublishTime
			}
			if ed==""{
				ed = ans[i].PublishTime
			}
			if len(st)<10 || len(ed)<10{
				continue
			}
			st = strings.Replace(st[0:10],"-","/",-1)
			ed = strings.Replace(ed[0:10],"-","/",-1)
			pre := ds[st]
			pre.Open = pre.Open+1 
			ds[st] = pre
			pre = ds[ed]
			pre.Close = pre.Close+1 
			ds[ed] = pre
		}
		tmp1:=0
		tmp2:=0
		dates := make([]string,0)
		opens := make([]int,0)
		closes := make([]int,0)
		keys := make([]string, 0)
		for v := range ds {
			keys = append(keys, v)
		}
		sort.Strings(keys)
		for _,key := range keys {
			v:=ds[key]
			tmp1+=v.Open
			tmp2+=v.Close
			tmp1-=v.Close
			dates = append(dates,key)
			opens = append(opens,tmp1)
			closes = append(closes,tmp2)
		}
		c.JSON(200, gin.H{"date": dates,
							"active": opens,
							"locked": closes})
	}
}
func GetTrack() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data struct{
			Category string `json:"category"`;
		}
		if c.BindJSON(&data) != nil {
			c.JSON(400, gin.H{"msg": "Invalid Params!"})
			return
		}
		res := model.FindTrack(data.Category)
		res1 := make([](map[string]interface{}),0)
		for i:=0;i<len(res);i++{
			tmp := make(map[string]interface{},0)
			tmp["id"] = res[i].Id
			tmp["events"] = res[i].Events
			res1 = append(res1,tmp)
		}
		util.RuntimeLog.Println(c.ClientIP()," request some track")
		c.JSON(200, gin.H{"news": res1})
	}
}
func GetNews() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data struct{
			Categories []string `json:"categories"`;
			Username string`json:"username"`;
			StartPublishDate string`json:"startPublishDate"`;
			EndPublishDate string`json:"endPublishDate"`
			NewsPerPage int`json:"newsPerPage"`
			PageNum int`json:"pageNum"`
		}
		if c.BindJSON(&data) != nil {
			c.JSON(400, gin.H{"msg": "Invalid Params!"})
			return
		}
		pages,res := model.FindNews(data.Categories,
						data.Username,
						data.StartPublishDate,
						data.EndPublishDate,
						data.NewsPerPage,
						data.PageNum-1,
					)
		util.RuntimeLog.Println(c.ClientIP()," request some news")
		c.JSON(200, gin.H{"totalPage":pages,"news": res})
	}
}

func GetOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		newsId := c.Param("id")
		int_id, _ := strconv.ParseInt(newsId, 0, 0)
		News,prev,next := model.News{Id: int_id}.GetWithPrevNext()
		if News == nil{
			c.JSON(400, gin.H{"msg": "No such news!"})
			return 
		}
		res := map[string]interface{}{
			"title":       News.Title,
			"name":        News.Name,
			"username":    News.CreateBy,
			"publishTime": News.PublishTime,
			"markdownText":News.MarkdownText,
			"text":        News.Text,
			"category":    News.Category,
			"events":      News.Events,
			"relationships":News.Relationships,
		}
		res["prevArticle"] = map[string]interface{}{
			"title": prev.Title,
			"id":    prev.Id,
		}
		res["nextArticle"] = map[string]interface{}{
			"title": next.Title,
			"id":    next.Id,
		}
	
		util.RuntimeLog.Println(c.ClientIP(),"request news",newsId)
		c.JSON(200, res)
	}
}

func PostOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data model.News
		if c.BindJSON(&data) != nil {
			c.JSON(400, gin.H{"msg": "Invalid Params!"})
			return
		}
		username:=c.MustGet("username").(string)
		user:=model.User{Username:username}.Get()
		data.CreateBy = username
		data.Name = user.Nickname
		data.PublishTime = time.Now().Format(time.RFC3339)
		data = data.Insert()
		util.RuntimeLog.Println("user",username,"post a news")
		c.JSON(200, gin.H{"msg": "add news success", "id": data.Id})
	}
}

func ModifyOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data model.News
		if c.BindJSON(&data) != nil {
			c.JSON(400, gin.H{"msg": "Invalid Params!"})
			return
		}
		newsId := c.Param("id")
		int_id, _ := strconv.ParseInt(newsId, 0, 0)
		data.Update(int_id)
		//username:=c.MustGet("username").(string)
		util.RuntimeLog.Println("user","modify news",int_id)
		c.JSON(200, gin.H{
			"msg": "change success",
		})
	}
}

func DeleteOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		newsId := c.Param("id")
		int_id, _ := strconv.ParseInt(newsId, 0, 0)
		model.News{Id: int_id}.Delete()
		//username:=c.MustGet("username").(string)
		util.RuntimeLog.Println("user","delete news", int_id)
		c.JSON(200, gin.H{"msg": "delete success"})
	}
}
