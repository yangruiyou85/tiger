// api/push.go
package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jdxj/backup/pkg/push"
)

// PushReq 推送请求
type PushReq struct {
	Message string `json:"message" binding:"required"`
}

// PushResp 推送响应
type PushResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// PushHandler 推送接口
func PushHandler(c *gin.Context) {
	var req PushReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := push.Push(req.Message)
	if result.ErrCode != 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.ErrMsg})
		return
	}

	c.JSON(http.StatusOK, PushResp{Success: true, Message: "push success"})
}
