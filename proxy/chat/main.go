package main

import (
	// "proxy/chat/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	
	r.LoadHTMLFiles("view/*")

	r.GET("/start", func(ctx *gin.Context) {
		ctx.HTML(200, "start.html", nil)
	})
	// router.Router(r)

	r.Run(":9999")
}