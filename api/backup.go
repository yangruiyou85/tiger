// api/backup.go
package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yangruiyou85/tiger/backup/pkg/backup"
)

// BackupReq 备份请求
type BackupReq struct {
	DBType     string `json:"db_type" binding:"required"`
	BackupType int    `json:"backup_type" binding:"required"`
	DBName     string `json:"db_name" binding:"required"`
	TableName  string `json:"table_name"`
}

// BackupResp 备份响应
type BackupResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// BackupHandler 备份接口
func BackupHandler(c *gin.Context) {
	var req BackupReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var backupType backup.BackupType
	switch req.BackupType {
	case 0:
		backupType = backup.FullBackup
	case 1:
		backupType = backup.SingleDBBackup
	case 2:
		backupType = backup.SingleTableBackup
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid backup type"})
		return
	}

	result := backup.Backup(req.DBType, backupType, req.DBName, req.TableName, "backup")
	if !result.Success {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Message})
		return
	}

	c.JSON(http.StatusOK, BackupResp{Success: true, Message: result.Message})
}
