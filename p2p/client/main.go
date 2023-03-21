package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("参数错误")
		return
	}
	localAddr := &net.UDPAddr{Port: port}
	conn, err := net.DialUDP("udp", localAddr, &net.UDPAddr{Port: 8888})
	if err != nil {
		fmt.Println("连接建立失败", err.Error())
		return
	}
	_, err = conn.Write([]byte{})
	if err != nil {
		fmt.Println("向服务端发送请求失败")
		return
	}
	buf := make([]byte, 256)
	n, _, err := conn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println("读取数据失败")
		return
	}
	conn.Close()
	addr := strings.Split(string(buf[:n]), ":")
	if len(addr) != 2 {
		fmt.Println("服务端错误")
		return
	}
	port, err = strconv.Atoi(addr[1])
	if err != nil {
		fmt.Println("服务端参数错误")
		return
	}
	dstAddr := net.UDPAddr{IP: net.ParseIP(addr[0]), Port: port}
	p2p(localAddr, &dstAddr)
}

// p2p通信
func p2p(src *net.UDPAddr, dst *net.UDPAddr) {
	// 请求建立连接
	conn, err := net.DialUDP("udp", src, dst)
	if err != nil {
		fmt.Println(err)		
	}

	buf := make([]byte, 256)
	// 启动goroutinue
	go func() {	
		for {
			_, addr, err := conn.ReadFromUDP(buf)
			if err != nil {
				continue
			}
			fmt.Print(addr.String() + ": " + string(buf))
		}
	}()
	reader := bufio.NewReader(os.Stdin)
	for {
		data, err := reader.ReadString('\n')		
		if err != nil {
			continue
		}
		conn.Write([]byte(data))
	}
}