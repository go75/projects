package global

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Socket struct {
	Conn *websocket.Conn
	DataQueue chan[]byte
	Exit chan struct{}
}

var OnlineUsers = make(map[uint]*Socket)
var Lock = new(sync.RWMutex)