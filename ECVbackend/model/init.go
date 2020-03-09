package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var dbEngine *xorm.Engine

func init() {
	dbEngine, _ = xorm.NewEngine("mysql", "ecv_test:ecv_test@/ecv")
}
