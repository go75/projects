package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var addr = "127.0.0.1:8080"

func main() {
	url, err := url.Parse("http://www.bilibili.com")
	if err != nil {
		log.Println(err)
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(url)
	log.Println("Starting httpserver at "+ addr)
	log.Fatal(http.ListenAndServe(addr, proxy))
}