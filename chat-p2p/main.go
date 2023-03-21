package main

import (
	"net"
	"strconv"
	"strings"
	"syscall/js"
)

var (
	console = js.Global().Get("console")
	document = js.Global().Get("document")
	input = document.Call("getElementById", "input")
)

var conn *net.UDPConn

func log(args ...any) {
	console.Call("log", args...)
}

func main() {
	localAddr := &net.UDPAddr{Port: 8080}
	conn, err := net.DialUDP("udp", localAddr, &net.UDPAddr{Port: 8888})
	if err != nil {
		log("连接建立失败", err.Error())
		return
	}
	_, err = conn.Write([]byte{})
	if err != nil {
		log("向服务端发送请求失败", err.Error())
		return
	}
	buf := make([]byte, 256)
	n, _, err := conn.ReadFromUDP(buf)
	if err != nil {
		log("读取数据失败", err.Error())
		return
	}
	conn.Close()
	addr := strings.Split(string(buf[:n]), ":")
	if len(addr) != 2 {
		log("服务端错误")
		return
	}
	port, err := strconv.Atoi(addr[1])
	if err != nil {
		log("服务端参数错误")
		return
	}
	dstAddr := net.UDPAddr{IP: net.ParseIP(addr[0]), Port: port}
	p2p(localAddr, &dstAddr)
}

// p2p通信
func p2p(src *net.UDPAddr, dst *net.UDPAddr) {
	// 请求建立连接
	var err error
	conn, err = net.DialUDP("udp", src, dst)
	if err != nil {
		log(err.Error())		
	}
	log("连接建立")
	buf := make([]byte, 256)
	for {
		_, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			continue
		}
		log(addr.String() + ": " + string(buf))
	}
}

func Send() {
	if conn != nil {
		value := input.Get("value").String()
		log(value)
		conn.Write([]byte(value))
	}
}