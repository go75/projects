package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{Port: 8888})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	addr0 := new(net.UDPAddr)
	addr1 := new(net.UDPAddr)
	buf := make([]byte, 256)
	_, addr0, err = conn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println("读取数据失败")
		return
	}
	fmt.Println("--------")
	_, addr1, err = conn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println("读取数据失败")
		return	
	}
	fmt.Println("--------")
	conn.WriteToUDP([]byte(addr0.String()), addr1)
	fmt.Println("...")
	conn.WriteToUDP([]byte(addr1.String()), addr0)
	fmt.Println("...")
	time.Sleep(time.Second << 1)
}