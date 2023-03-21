package log

import (
	"im/config"
	"io"
	"log"
	"os"
)

//错误
var Error *log.Logger

func init() {
	file, err := os.OpenFile(config.Log.Location, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("无法打开错误的log文件: " + err.Error())
	}
	Error = log.New(io.MultiWriter(file, os.Stderr), "\u001B[1;31m[Error]:\u001B[0m", log.Ldate|log.Ltime|log.Lshortfile)
}