package util
import (
    "github.com/sirupsen/logrus"
	"ecvbackend/pkg"
    "os"
)
var ErrorLog *logrus.Logger
var RuntimeLog *logrus.Logger
func init () {
    initErrorLog()
    initRuntimeLog()
}

func initErrorLog() {
    ErrorLog = logrus.New()
    ErrorLog.SetFormatter(&logrus.JSONFormatter{})
	sec,_ := setting.Cfg.GetSection("log")
	fileName := sec.Key("ERROR_LOG").String()
    src, err := os.OpenFile(fileName, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0755)
    if err != nil {
        panic(err)
    }
    ErrorLog.SetOutput(src)
}

func initRuntimeLog() {
    RuntimeLog = logrus.New()
    RuntimeLog.SetFormatter(&logrus.JSONFormatter{})
	sec,_ := setting.Cfg.GetSection("log")
	fileName := sec.Key("RUNTIME_LOG").String()
    src, err := os.OpenFile(fileName, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0755)
    if err != nil {
        panic(err)
    }
    RuntimeLog.SetOutput(src)
}