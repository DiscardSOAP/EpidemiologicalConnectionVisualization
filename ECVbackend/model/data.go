package model

import (
	"ecvbackend/pkg/util"
)

type Event struct{
	StartDate string `xorm:"startDate" json:"startDate"`
	StartTime string `xorm:"startTime" json:"startTime"`
	EndDate string `xorm:"endDate" json:"endDate"`
	EndTime string `xorm:"endTime" json:"endTime"`
	Addr string `xorm:"addr" json:"addr"`
	Description string `xorm:"description" json:"description"`
}
type Relationship struct{
	Relationship string `xorm:"relationship" json:"relationship"`
	NewsID int64 `xorm:"newsid" json:"id"`
}

type News struct {
	Id int64 `xorm:"pk autoincr" json:"id"`
	Title string `xorm:"title" json:"title"`
	Text string `xorm:"text"  json:"text"`
	Category string `xorm:"category"  json:"category"`
	Events []Event `xorm:"events"  json:"events" `
	Name string `xorm:"name" json:"name"`
	CreateBy string `xorm:"createBy" json:"username"`
	PublishTime string `xorm:"publish_time" json:"publishTime"`
	MarkdownText string `xorm:"markdownText" json:"markdownText"`
	Relationships []Relationship `xorm:"relationships" json:"relationships"`
}

func FindNews(cats []string, username string,st string, ed string,num int,offset int) (int64,[]News){
	var ans []News
	qu := dbEngine.Where("id >?",-1)
	if len(st)>0{
		qu = qu.Where("publish_time>?",st)
	}
	if len(ed)>0{
		qu = qu.Where("publish_time<?",ed)
	}
	if username!=""{
		qu = qu.Where("createBy=?",username)
	}
	if len(cats)>0 {
		qu = qu.In("category",3)
	}
	ne := new(News)
	count,_:= qu.Count(ne)
	qu = dbEngine.Where("id >?",-1)
	if len(st)>0{
		qu = qu.Where("publish_time>?",st)
	}
	if len(ed)>0{
		qu = qu.Where("publish_time<?",ed)
	}
	if username!=""{
		qu = qu.Where("createBy=?",username)
	}
	if len(cats)>0 {
		qu = qu.In("category",cats)
	}
	if num!=0{
		qu.Limit(num,offset*num).Desc("publish_time").Find(&ans)
		return (count-1)/int64(num)+1, ans
	}else{
		return 0, ans
	}
}


func FindTrack(cats string) []News{
	var ans []News
	dbEngine.Where("id >?",-1).In("category",cats).Find(&ans)
	return ans
}

func (p News)Insert() News{
	_, err := dbEngine.Insert(&p)
	if err!=nil{
		return News{}
	}else {
		return p
	}
}

func (p News) Get() *News{
	exist, _ := dbEngine.Get(&p)
	if exist {
		return &p
	}
	return nil
}

func (p News) GetWithPrevNext() (*News,*News,*News){
	exists, _ := dbEngine.Get(&p)
	if exists {
		var next []*News
		var prev []*News
		dbEngine.Limit(1).Cols("id", "title").Asc("id").Where("id >?", p.Id).Find(&next)
		dbEngine.Limit(1).Cols("id", "title").Desc("id").Where("id <?", p.Id).Find(&prev)
		var nxt *News = &News{Id:-1,Title:"last"}
		var prv *News = &News{Id:-1,Title:"first"}
		if len(prev)>0{
			prv = prev[0]
		}
		if len(next)>0{
			nxt = next[0]
		}
		return &p,prv,nxt
	}
	return nil,nil,nil
}

func (p News) Delete() {
	dbEngine.Delete(&p)
}

func (p News) GetAll() []News{
	everyone := make([]News, 0)
	err := dbEngine.Find(&everyone)
	if err==nil {
		return everyone 
	}
	return nil
}

func (p News) Update(id int64) {
	_, err := dbEngine.Where("id = ?",id).AllCols().Update(&p)
	if err == nil {
		util.RuntimeLog.Println("update success")
	} else {
		util.RuntimeLog.Println("update error ",err)
	}
}

func FindRel(cat string) []News{
	ans := make([]News, 0)
	err := dbEngine.Cols("relationships","id").Where("category = ?",cat).Find(&ans)
	if err == nil {
		util.RuntimeLog.Println("update success")
	} else {
		util.RuntimeLog.Println("update error ",err)
	}
	return ans
}


func GetCatCount() map[string]int{
	res := make(map[string]int, 0)
	everyone := make([]News, 0)
	err := dbEngine.Find(&everyone)
	if err==nil {
		for i:=0;i<len(everyone);i++{
			res[everyone[i].Category]=res[everyone[i].Category]+1
		}
		return res
	}
	return nil
}