package router

import (
	"chat/global"
	"chat/log"
	"chat/middleware"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func Init(r *gin.Engine) {
	r.Static("/asset", "./asset")
	page(r)
	r.Use(middleware.RateLimiter, middleware.Jwt)
	r.GET("/chat", Upgrade)
	user(r.Group("user"))
	contact(r.Group("contact"))
}

const (
	CMD_SINGLE_MSG = 10
	CMD_ROOM_MSG   = 11
	CMD_HEART      = 0
)

const (
    //文本样式
	MEDIA_TYPE_TEXT=1
	//新闻样式,类比图文消息
	MEDIA_TYPE_News=2
	//语音样式
	MEDIA_TYPE_VOICE=3
	//图片样式
	MEDIA_TYPE_IMG=4
	
	//红包样式
	MEDIA_TYPE_REDPACKAGR=5
	//emoj表情样式
	MEDIA_TYPE_EMOJ=6
	//超链接样式
	MEDIA_TYPE_LINK=7
	//视频样式
	MEDIA_TYPE_VIDEO=8
	//名片样式
	MEDIA_TYPE_CONCAT=9
	//其他自己定义,前端做相应解析即可
	MEDIA_TYPE_UDEF=100
)

type Message struct {
	Id      uint  `json:"id,omitempty" form:"id"` //消息ID
	Userid  uint  `json:"userid,omitempty" form:"userid"` //谁发的
	Cmd     int    `json:"cmd,omitempty" form:"cmd"` //群聊还是私聊
	Dstid   uint  `json:"dstid,omitempty" form:"dstid"`//对端用户ID/群ID
	Media   int    `json:"media,omitempty" form:"media"` //消息按照什么样式展示
	Content string `json:"content,omitempty" form:"content"` //消息的内容
	Pic     string `json:"pic,omitempty" form:"pic"` //预览图片
	Url     string `json:"url,omitempty" form:"url"` //服务的URL
	Memo    string `json:"memo,omitempty" form:"memo"` //简单描述
	Amount  int    `json:"amount,omitempty" form:"amount"` //其他和数字相关的
}

// websocket默认配置
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 升级协议
func Upgrade(c *gin.Context) {
	socket, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Warn.Println("websocket升级失败, err: ", err)
		c.JSON(http.StatusBadRequest, "websocket升级失败")
		c.Abort()
		return
	}
	defer socket.Close()
	select{}
	// //todo 获得conn
	// node := &global.Socket{
	// 	Conn: socket,
	// 	DataQueue:make(chan []byte,50),
	// 	Exit: make(chan struct{}),
	// }
	// //todo 获取用户全部群Id
	// id := c.GetUint("id")
	// contact := &model.Contact{
	// 	Id: id,
	// 	Cate: model.CONCAT_CATE_COMUNITY,
	// }
	// list, _ :=  dao.GetContactsByOwneridAndCate(contact)
	//todo userid和node形成绑定关系
	// global.Lock.Lock()
	// global.OnlineUsers[id]=node
	// global.Lock.Unlock()

	// log.Info.Println("----------------------------")
	// //todo 完成发送逻辑,con
	// go sendproc(node)
	// //todo 完成接收逻辑
	// go recvproc(node)
}

func sendproc(node *global.Socket) {
	for {
		select {
		case data := <- node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
                log.Warn.Println("websocket send err: ", err.Error())
                return
            }
		case <-node.Exit:
			return
		}
	}
}

func recvproc(node *global.Socket) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			log.Warn.Println("websocket recv err: ", err.Error())
		}
		log.Info.Println("recv msg: "+string(data))
		dispatch(data)
	}
}

func sendMsg(id uint, msg []byte) {
	global.Lock.RLock()
	user, ok := global.OnlineUsers[id]
	global.Lock.RUnlock()
	log.Info.Println("send msg: "+string(msg))
	if !ok {
		log.Info.Printf("id = %d的用户不在线\n", id)
		return
	}
	user.DataQueue <- msg
}

func dispatch(data []byte) {
	// 将data解析为message对象
	msg := &Message{}
	err := json.Unmarshal(data, msg)
	if err !=nil {
		log.Info.Println("data解析失败, err: ", err.Error())
		return
	}

	// 根据message的cmd属性做进一步的处理
	switch msg.Cmd {
	case CMD_SINGLE_MSG:
		// 单聊
		sendMsg(msg.Dstid, data)
	case CMD_ROOM_MSG:
		// 群聊

	case CMD_HEART:
	}
}