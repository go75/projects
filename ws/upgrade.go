package ws

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// use default options
var upgrader = websocket.Upgrader {
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Upgrade(c *gin.Context) (ws *websocket.Conn, err error) {
	return upgrader.Upgrade(c.Writer, c.Request, nil)
}