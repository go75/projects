package fuse

import "time"

type Fuse struct {
	// 单次请求的超时时间
	Timeout time.Duration
	// 最大并发量
	MaxConcurrentRequests uint32
	// 熔断后多久去尝试服务是否可用
	SleepWindow time.Duration
	// 验证熔断时的请求数量, 10秒内取样
	RequestNum uint16
	// 验证熔断的错误百分比
	ErrorPercent uint16
}

type Fuses struct {
	list map[string]*Fuse
}

func (f *Fuses) Do(name string, fn func() error, fallback func ()) error {
	
	return nil
}

