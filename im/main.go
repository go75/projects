package main

import (
	"im/docs"
	"im/router"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title im系统接口文档
// @version 1.0
// @description  API文档
// @BasePath /
func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
    r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Static("/assets", "./front/dist/assets")
	r.Static("/image", "./assets/static/img")
	r.Static("/head", "./head")
	r.LoadHTMLFiles("./front/dist/index.html", "./assets/view/login.html", "./assets/view/regist.html")
	router.Router(r)
	r.Run(":9999")
}