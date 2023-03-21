package ws

var Dispatcher chan *Request

func StartDispatcher(workerNum, dispatcherLen uint) {
	//初始化请求队列
	Dispatcher = make(chan *Request, dispatcherLen)
	
	//启动dispatcher
	startWorker(workerNum)
}

func startWorker(workerNum uint) {
	for workerNum > 0 {
		go func() {
			//不断阻塞等待请求队列的请求
			for req := range Dispatcher {
				do(req)
			}
		}()
		workerNum--
	}
} 
