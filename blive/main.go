package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

type BiliClient struct {
	RoomID uint32
	HTTPClient *http.Client
	Conn *websocket.Conn

	Ch chan PacketBody
}

type StartInfo struct {
	Code int `json:"code"`
	Message string `json:"message"`
	TTL int `json:"ttl"`
	Data Data `json:"data"`
}
type HostList struct {
	Host string `json:"host"`
	Port int `json:"port"`
	WssPort int `json:"wss_port"`
	WsPort int `json:"ws_port"`
}
type Data struct {
	Group string `json:"group"`
	BusinessID int `json:"business_id"`
	RefreshRowFactor float64 `json:"refresh_row_factor"`
	RefreshRate int `json:"refresh_rate"`
	MaxDelay int `json:"max_delay"`
	Token string `json:"token"`
	HostList []HostList `json:"host_list"`
}

var urlFormat = "https://api.live.bilibili.com/xlive/web-room/v1/index/getDanmuInfo?id=%d"

var start = new(StartInfo)

func NewBiliClient(roomID uint32) *BiliClient {
	return &BiliClient{
		RoomID: roomID,
		HTTPClient: new(http.Client),
		Ch: make(chan PacketBody, 1024),
	}
}

func (c *BiliClient) GetHostList() []url.URL {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(urlFormat, c.RoomID), nil)
	if err != nil {
		panic(err)
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, start)
	if err != nil {
		panic(err)
	}
	hostList := make([]url.URL, 0, 3)
	for _, host := range start.Data.HostList {
		u := url.URL{Scheme: "wss", Host: fmt.Sprintf("%s:%d", host.Host, host.WssPort), Path: "/sub"}
		hostList = append(hostList, u)
	}
	return hostList
}

func (c *BiliClient) Connect(urls []url.URL) {
	var err error
	for _, url := range urls {
		c.Conn, _, err = websocket.DefaultDialer.Dial(url.String(), nil)
		if err != nil {
			fmt.Printf("websocket dial %s failed\n", url.String())
			continue
		}
		fmt.Printf("websocket dial %s successfully\n", url.String())
		break
	}
	if err != nil {
		panic(err)
	}
	type handShakeInfo struct {
		UID uint8 `json:"uid"`
		Roomid uint32 `json:"roomid"`
		Protover uint8 `json:"protover"`
		Platform string `json:"platform"`
		Clientver string `json:"clientver"`
		Type uint8 `json:"type"`
		Key string `json:"key"`
	}

	hsi := &handShakeInfo {
		UID: 0,
		Roomid: c.RoomID,
		Protover: 2,
		Platform: "web",
		Clientver: "1.10.2",
		Type: 2,
		Key: start.Data.Token,
	}
	data, err := json.Marshal(hsi)
	if err != nil {
		panic(err)
	}

	err = c.SendPacket(0, 16, 0, 7, 0, data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("connect to room %d successfully\n", c.RoomID)
}

func (c *BiliClient) HeartBeat() {
	for {
		err := c.SendPacket(0, 16, 0, 2, 0, []byte(""))
		if err != nil {
			fmt.Println("heart beat err: ", err)
			time.Sleep(5 * time.Second)
			continue
		}
		time.Sleep(20 * time.Second)
	}
}

func (bc *BiliClient) Disconnect() {
	err := bc.Conn.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	if len(os.Args) != 2 {
		panic("argument number error")
	}

	id, err := strconv.ParseUint(os.Args[1], 10, 32)
	if err != nil {
		panic(err)
	}

	biliClient := NewBiliClient(uint32(id))
	hostList := biliClient.GetHostList()
	biliClient.Connect(hostList)
	defer biliClient.Disconnect()

	go biliClient.ReceiveMessages()
	go biliClient.ParseMessages()
	biliClient.HeartBeat()
}
