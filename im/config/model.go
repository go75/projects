package config

import (
	"encoding/json"
	"io"
	"os"
)

type DataBaseModel struct {
	DSN string
	SlowThreshold int // 慢SQL阈值
	LogLevel int // 级别
	Colorful bool // 彩色
}

type RedisModel struct { 
	Addr string
    DB int
    PoolSize int
    MinIdleConns int
	ExpiredTime int
}

type JwtModel struct {
	ExpiredTime int
}

type LogModel struct {
	Location string
}

type DispatcherModel struct {
    DispatchSize int
	TextChannel string
}

var DataBase = new(DataBaseModel)
var Redis = new(RedisModel)
var Jwt = new(JwtModel)
var Log = new(LogModel)
var Dispatcher = new(DispatcherModel)

func inject(kvs map[string]any) error {
	for k, v := range kvs {
		file, err := os.Open(k + ".json")
		if err != nil {
			return err
		}
		defer file.Close()
		data, err := io.ReadAll(file)
		if err != nil {
			return err
		}
		err = json.Unmarshal(data, v)
		if err != nil {
			return err
		}
	}
	return nil
}