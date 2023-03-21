package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func UserLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func UserRegistPage(c *gin.Context) {
	c.HTML(http.StatusOK, "regist.html", nil)
}