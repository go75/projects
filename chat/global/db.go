package global

import (
	"chat/config"
	"log"
	"os"
	"time"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func initDB() {
	// sql日志
	logger := logger.New (
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config {
			SlowThreshold: time.Second, // 慢SQL阈值
			LogLevel: logger.Info, // 级别
			Colorful: true, // 彩色
		},
	)
	var err error
	DB, err = gorm.Open(mysql.Open(config.DbBsn), &gorm.Config {
		Logger: logger,
	})
	if err!=nil {
		panic(err)
	}
}
