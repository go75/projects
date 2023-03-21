package ws

import (
	"im/config"
	"im/log"
	"im/ws/api"
	"im/ws/entity"
)

var channel = make(chan *entity.Request, 1024)

func init() {
	for i := 0; i < config.Dispatcher.DispatchSize; i++ {
		log.Info.Printf("dispatcher%d, start...\n", i)
		go func() {
			var req *entity.Request
			//订阅消息, 不断阻塞等待请求队列的请求
			for req = range channel {
				log.Info.Printf("请求: %v", *req)
				api.Do(req)
			}
		}()
	}
}