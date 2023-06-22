// pkg/backup/backup.go
package backup

import (
	"fmt"
	"os"
	"time"
)

// BackupType 备份类型
type BackupType int

const (
	// FullBackup 全库备份
	FullBackup BackupType = iota
	// SingleDBBackup 单库备份
	SingleDBBackup
	// SingleTableBackup 单表备份
	SingleTableBackup
)

// BackupResult 备份结果
type BackupResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// Backup 备份数据库
func Backup(dbType string, backupType BackupType, dbName, tableName, backupDir string) *BackupResult {
	var (
		err   error
		cmd   string
		now   = time.Now().Format("20060102150405")
		fname string
	)

	switch dbType {
	case "mysql":
		switch backupType {
		case FullBackup:
			fname = fmt.Sprintf("%s/%s_%s.sql", backupDir, dbName, now)
			cmd = fmt.Sprintf("mysqldump -h%s -P%d -u%s -p%s --single-transaction --routines --triggers --events --all-databases > %s", mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.User, mysqlConfig.Password, fname)
		case SingleDBBackup:
			fname = fmt.Sprintf("%s/%s_%s_%s.sql", backupDir, dbName, tableName, now)
			cmd = fmt.Sprintf("mysqldump -h%s -P%d -u%s -p%s --single-transaction --routines --triggers --events %s %s > %s", mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.User, mysqlConfig.Password, dbName, tableName, fname)
		case SingleTableBackup:
			fname = fmt.Sprintf("%s/%s_%s_%s.sql", backupDir, dbName, tableName, now)
			cmd = fmt.Sprintf("mysqldump -h%s -P%d -u%s -p%s --single-transaction --routines --triggers --events %s %s > %s", mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.User, mysqlConfig.Password, dbName, tableName, fname)
		}
	case "mongodb":
		switch backupType {
		case FullBackup:
			fname = fmt.Sprintf("%s/%s_%s.tar.gz", backupDir, dbName, now)
			cmd = fmt.Sprintf("mongodump --host %s:%d --username %s --password %s --authenticationDatabase admin --gzip --archive=%s", mongodbConfig.Host, mongodbConfig.Port, mongodbConfig.User, mongodbConfig.Password, fname)
		case SingleDBBackup:
			fname = fmt.Sprintf("%s/%s_%s.tar.gz", backupDir, dbName, now)
			cmd = fmt.Sprintf("mongodump --host %s:%d --username %s --password %s --authenticationDatabase admin --gzip --db %s --archive=%s", mongodbConfig.Host, mongodbConfig.Port, mongodbConfig.User, mongodbConfig.Password, dbName, fname)
		case SingleTableBackup:
			fname = fmt.Sprintf("%s/%s_%s_%s.bson.gz", backupDir, dbName, tableName, now)
			cmd = fmt.Sprintf("mongodump --host %s:%d --username %s --password %s --authenticationDatabase admin --gzip --db %s --collection %s --archive=%s", mongodbConfig.Host, mongodbConfig.Port, mongodbConfig.User, mongodbConfig.Password, dbName, tableName, fname)
		}
	}

	err = os.MkdirAll(backupDir, os.ModePerm)
	if err != nil {
		return &BackupResult{Success: false, Message: err.Error()}
	}

	_, err = exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		return &BackupResult{Success: false, Message: err.Error()}
	}

	return &BackupResult{Success: true, Message: fname}
}
