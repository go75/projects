package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.GET("/start", func(c *gin.Context) {
		c.HTML(http.StatusOK, "start.html", nil)
	})
}