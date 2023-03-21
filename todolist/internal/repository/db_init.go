package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/driver/mysql"
	"todolist/config"
	"os"
	"log"
	"time"
)
var DB *gorm.DB

func InitDB()  {
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
	DB, err = gorm.Open(mysql.Open(config.DbDsn), &gorm.Config {
		Logger: logger,
	})
	if err!=nil {
		panic(err)
	}
	tmp, err := DB.DB()
	if err!=nil {
		panic(err)
	}
	tmp.SetMaxIdleConns(10)
	tmp.SetMaxOpenConns(50)
	tmp.SetConnMaxLifetime(30*time.Second)
	migration()
}