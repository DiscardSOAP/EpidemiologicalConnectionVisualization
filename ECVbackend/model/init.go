package model

import (
	"log"
	"fmt"
	"ecvbackend/pkg"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var dbEngine *xorm.Engine

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
	dbEngine.CreateTables(&User{})
}
