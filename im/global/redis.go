package global

// import (
// 	"im/config"
// 	"im/log"
// 	"github.com/go-redis/redis"
// )

// var Rd *redis.Client

// func initRedis() {
// 	Rd = redis.NewClient(&redis.Options {
// 		Addr: config.Redis.Addr,
// 		DB: config.Redis.DB,
// 		PoolSize: config.Redis.PoolSize,
// 		MinIdleConns: config.Redis.MinIdleConns,
// 	})
// 	pong, err := Rd.Ping().Result()
// 	if err != nil {
// 		panic(err)
// 	} else {
// 		log.Info.Println("redis inited, ", pong)
// 	}
// }
