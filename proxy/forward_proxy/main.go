package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
)

type Pxy struct{}

func (p *Pxy)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request %s %s %s\n", r.Method, r.Host, r.RemoteAddr)
	transport := http.DefaultTransport
	// 拷贝http请求
	req := new(http.Request)
	*req = *r
	if clientIP, _ , err := net.SplitHostPort(r.RemoteAddr); err!=nil {
		if prior, ok := req.Header["X-Forwarded-For"]; ok {
			clientIP = strings.Join(prior, ", ") + ", " + clientIP 
		}
		req.Header.Set("X-Forwarded-For", clientIP)
	}

	// 请求下游
	resp, err := transport.RoundTrip(req)
	if err!=nil {
		w.WriteHeader(http.StatusBadGateway)
		log.Println(err)
		return
	}
	// 返回下游响应数据给上游
	for key, value := range req.Header {
		for _, v := range value {
			w.Header().Add(key, v)
		}
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
	resp.Body.Close()
}

func main() {
	fmt.Println("Serve on :8080")
	http.Handle("/", &Pxy{})
	http.ListenAndServe("0.0.0.0:8080", nil)
}