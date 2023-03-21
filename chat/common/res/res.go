package res

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type H struct {
	Code int `json:"code"`
	Msg  string `json:"msg"`
	Data any `json:"data,omitempty"`
	Rows any `json:"rows,omitempty"`
	Total any `json:"total,omitempty"`
}

func Err(c *gin.Context, msg string){
	Res(c,-1,msg,nil)
}

func Ok(c *gin.Context,msg string, data any){
	Res(c,0,msg,data)
}

func OkList(c *gin.Context,list,total any){
	//分页数目,
	ResList(c,0,list,total)
}

func Res(c *gin.Context, code int, msg string, data any)  {
	c.JSON(http.StatusOK, H {
		Code: code,
        Msg:  msg,
        Data: data,
	})
}

func ResList(c *gin.Context,code int,rows interface{},total interface{})  {
	c.JSON(http.StatusOK, H {
		Code: code,
        Rows: rows,
        Total:  total,
	})
}