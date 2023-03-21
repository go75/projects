package manager

import (
	"im/common/imerr"
	"sync"

	"github.com/gorilla/websocket"
)

type connectionManager struct {
	// key:在线用户id, value:websocket连接
	sockets map[uint]*websocket.Conn
	lock *sync.RWMutex
}

var connManager = &connectionManager {
	sockets: make(map[uint]*websocket.Conn, 0),
	lock: new(sync.RWMutex),
}

func Put(id uint, conn *websocket.Conn) error {
	if _, ok := connManager.sockets[id]; ok {
		return imerr.AlreadyExistConnErr
	}

	connManager.lock.Lock()
	defer connManager.lock.Unlock()
	connManager.sockets[id] = conn
	return nil
}

func Get(id uint) *websocket.Conn {
	connManager.lock.RLock()
	defer connManager.lock.RUnlock()
	return connManager.sockets[id]
}

func Remove(id uint) {
	connManager.lock.Lock()
    defer connManager.lock.Unlock()
    delete(connManager.sockets, id)
}