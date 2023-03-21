package middleware

import (
	"chat/common/imerr"
	"chat/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 检权中间件
func Jwt(c *gin.Context) {

	// 允许直接通过的路径
	if strings.HasPrefix(c.Request.URL.Path, "/user/login.html") || strings.HasPrefix(c.Request.URL.Path, "/user/regist.html") {
		return
	}

	// 通过http cookie中的token解析来认证
	token,err := c.Cookie("chat-token")
	if err != nil {
		c.HTML(http.StatusOK, "login.html", nil)
		c.Abort()
		return
	}

	// 未携带token直接返回
	if token == "" {
		c.HTML(http.StatusOK, "login.html", nil)
		c.Abort()
		return
	}

	// 初始化一个JWT对象实例，并根据结构体方法来解析token
	// 解析token中包含的相关信息(有效载荷)
	claims, err := utils.ParseToken(token)

	if err != nil {
		// token过期
		if err == imerr.ExpiredTokenErr {
			c.HTML(http.StatusOK, "login.html", "token授权已过期, 请重新申请授权")
			c.Abort()
			return
		}

		// 其他错误
		c.HTML(http.StatusOK, "login.html", "无效token")
		c.Abort()
		return
	}
	// 将解析后的有效载荷claims的数据写入gin.Context引用对象中
	c.Set("id", claims.ID)
	c.Set("mobile", claims.Mobile)
	c.Set("nickname", claims.Nickname)
	c.Next()
}