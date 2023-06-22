// api/router.go
package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() http.Handler {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.POST("/backup", BackupHandler)
	r.POST("/push", PushHandler)

	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	return r
}
