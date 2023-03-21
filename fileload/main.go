package main

import (
	"embed"
	"io"
	"io/fs"
	"net"
	"os"
	"strings"
	"time"

	// "io/fs"
	"net/http"
	"github.com/skip2/go-qrcode"
	"github.com/gin-gonic/gin"
)

//go:embed front/dist/assets/*
var FS embed.FS

func main() {
	r := gin.Default()
	staticFS, _ := fs.Sub(FS, "front/dist/assets")
	r.StaticFS("/assets", http.FS(staticFS))
	r.LoadHTMLFiles("./front/dist/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/addresses", AddressesHandler)
	r.POST("/file/:path", FileHandler)
	r.GET("/upload/:path", UploadHandler)
	r.Run(":8888")
}

func uploadFile(content, postfix string) (string, error) {
	// 获取当前可执行文件的目录
	// exe, err := os.Executable()
	// if err != nil {
	// 	return "", err
	// }
	// dir :=  filepath.Dir(exe)
	filepath := "uploads/" + time.Now().String() + postfix
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0666)
	if err!= nil {
        return "", err
    }
	_, err = io.WriteString(file, content)
	if err != nil {
		return "", err
	}
	return filepath, nil
}

func AddressesHandler(c *gin.Context) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": err.Error(),
		})
	}
	var res []string
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				res = append(res, ipnet.IP.String())
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"addresses": res,
	})
}

func FileHandler(c *gin.Context) {
	// path := c.Param("path")

	// png, err := qrcode.Encode("http://127.0.0.1:8888/upload/"+path, qrcode.Medium, 256)
	png, err := qrcode.Encode("http://www.baidu.com", qrcode.Medium, 256)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	c.Data(http.StatusOK, "image/png", png)
}

func UploadHandler(c *gin.Context) {
	path := c.Param("path")
	parts := strings.Split(path, "/")
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Desposition", "attachment; filename="+parts[len(parts) - 1])
	c.Header("Content-Type", "application/octet-stream")
	c.File(path)
}