package main

import (
	"chat/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.Init(r)
	r.Run(":8080")
}