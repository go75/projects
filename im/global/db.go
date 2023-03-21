package global

import (
	"im/config"
	"log"
	"os"
	"time"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func initDB() {
	// sql日志
	logger := logger.New (
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config {
			SlowThreshold: time.Duration(config.DataBase.SlowThreshold) * time.Millisecond, // 慢SQL阈值
			//LogLevel: logger.Info, // 级别
			LogLevel: logger.LogLevel(config.DataBase.LogLevel),
			Colorful: config.DataBase.Colorful, // 彩色
		},
	)
	var err error
	DB, err = gorm.Open(mysql.Open(config.DataBase.DSN), &gorm.Config {
		Logger: logger,
	})
	if err!=nil {
		panic(err)
	}
}
