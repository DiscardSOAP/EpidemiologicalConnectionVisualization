package model

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var dbEngine *xorm.Engine

func init() {
	var err error
	dbEngine, _ = xorm.NewEngine("mysql", "user:123456@tcp(127.0.0.1:3306)/ecvsql?charset=utf8")
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
	_, err = dbEngine.Insert(&User{Username: "jkxing", Password: "123456"})
	if err != nil {
		log.Println(err)
	} else {
		log.Println("init sql ok")
	}
}
