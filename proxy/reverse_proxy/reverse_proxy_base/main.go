package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

var (
	proxy_addr = "http://127.0.0.1:2023"
	port = "2022"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// 解析代理地址, 更改请求体的协议和主机
	proxy, err := url.Parse(proxy_addr)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	r.URL.Scheme = proxy.Scheme
	r.URL.Host = proxy.Host

	transport := http.DefaultTransport
	// 请求下游
	resp, err := transport.RoundTrip(r)
	if err != nil {
		fmt.Println(err)
		return
	}	
	// 返回下游响应数据给上游
	for key, value := range resp.Header {
		for _, v := range value {
			w.Header().Add(key, v)
		}
	}
	bufio.NewReader(resp.Body).WriteTo(w)
	resp.Body.Close()
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Start serving on "+port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}