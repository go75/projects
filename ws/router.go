package ws

import "sync"

type routerTable struct {
	table map[uint]func(r *Request)
	lock *sync.RWMutex
}

var table = routerTable {
	table: make(map[uint]func(r *Request)),
	lock: new(sync.RWMutex),
}

func Register(id uint, fn func(r *Request)) {
	table.lock.Lock()
	defer table.lock.Unlock()
	table.table[id] = fn
}

func do(r *Request) {
	table.lock.RLock()
	fn := table.table[r.ID]
	table.lock.RUnlock()
	fn(r)
}