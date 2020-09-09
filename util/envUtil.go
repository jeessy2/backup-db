package util

import "os"

// GetEnvType 获得类型
func GetEnvType() string {
	typ := os.Getenv("BACKUP_TYPE")
	if typ == "" {
		typ = "client"
	}
	return typ
}

// GetEnvDBType 获得数据库类型
func GetEnvDBType() string {
	return os.Getenv("BACKUP_DB_TYPE")
}
