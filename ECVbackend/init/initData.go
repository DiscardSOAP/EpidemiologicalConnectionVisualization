package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"strings"
	"log"
	"ecvbackend/pkg"
	"ecvbackend/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
	"strconv"
)

func init() {
	var (
	err error
	dbType, dbName, user, password, host string
    )

	sec, err := setting.Cfg.GetSection("database")
	
    if err != nil {
	log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	//tablePrefix = sec.Key("TABLE_PREFIX").String()
	
	dbEngine, err = xorm.NewEngine(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", 
	user, 
	password, 
	host, 
	dbName))
	if err != nil {
		log.Println(err)
    }
	//dbEngine.ShowSQL = true
	dbEngine.ShowSQL(true)
	dbEngine.SetMaxIdleConns(5)
	dbEngine.SetMaxOpenConns(5)
	err = dbEngine.Ping()
	if err != nil {
		log.Println("ping sql wrong")
	} else {
		log.Println("ping sql ok")
	}
	dbEngine.CreateTables(&model.User{})
	dbEngine.CreateTables(&model.News{})
}

var dbEngine *xorm.Engine

func GetMarkDownText(s []string)string{
	return "# 病例Id"+s[0]+"\n"+
	"## 基本信息\n"+
	"姓名："+s[12]+"\n"+
	"性别："+s[9]+"\n"+
	"年龄："+s[10]+"\n"+
	"## 病情相关\n"+
	"确诊时间："+s[3]+"\n"+
	"病情程度："+s[11]+"\n"+
	"描述："+s[13]+"\n"
}
// 游戏读取数据，读取游戏配置数据
func ReadCsv_ConfigFile_Fun(fileName string) bool {
	// 获取数据，按照文件
	cntb, err := ioutil.ReadFile(fileName)
	if err != nil {
		return false
	}
	// 读取文件数据
	r2 := csv.NewReader(strings.NewReader(string(cntb)))
	ss, _ := r2.ReadAll()
	//sz := len(ss)
	// 循环取数据
	data:=model.News{}
	for i := 1; i < 23; i++ {
		fmt.Println(len(ss[i]))
		data = model.News{}
		data.Category="天津"
		tim,_:=time.Parse("2006-01-02 15:04:05",ss[i][3]+":00")
		fmt.Println(ss[i][3]+":00",tim)
		data.PublishTime=tim.Format(time.RFC3339)
		data.Title="病例"+strconv.Itoa(i)
		data.Name="Official"
		data.Text=ss[i][13]
		data.MarkdownText=GetMarkDownText(ss[i])
		data.CreateBy="jkxing1234"
		x := strings.Split(ss[i][4],",")
		data.Relationships=make([]model.Relationship,0)
		for i:=0;i<len(x);i++{
			nid,_:=strconv.ParseInt(x[i], 10, 64)
			data.Relationships = append(data.Relationships,model.Relationship{Relationship:"From",NewsID:nid})
		}
		x = strings.Split(ss[i][6],",")
		for i:=0;i<len(x);i++{
			nid,_:=strconv.ParseInt(x[i], 10, 64)
			data.Relationships = append(data.Relationships,model.Relationship{Relationship:"To",NewsID:nid})
		}
		data.Events=make([]model.Event,0)
		for j:=14;j<len(ss[i]);j+=4{
			if(ss[i][j]!=""){	
				data.Events=append(data.Events,model.Event{StartDate:ss[i][j],
															EndDate:ss[i][j+1],
															Addr:ss[i][j+2],
															Description:ss[i][j+3],
														})
			}
		}
		dbEngine.Insert(data)
		fmt.Println(data)
	}
	return true
}

/*type News struct {
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
}*/
func main(){
	ReadCsv_ConfigFile_Fun("data.csv")
}