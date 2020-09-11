package util

import (
	"backup-db/entity"
	"log"
	"os"
	"strconv"
	"time"
)

// PathExists Get path exist
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// DeleteOlderFiles delete older files
func DeleteOlderFiles(path string, backupFiles []os.FileInfo) {

	conf, err := entity.GetConfigCache()
	if err == nil {
		ago := time.Now()
		for _, conf := range conf.BackupConfig {
			lastDay, _ := time.ParseDuration("-" + strconv.Itoa(conf.SaveDays*24) + "h")
			ago = ago.Add(lastDay)

			// delete older file when file numbers gt MaxSaveDays
			for _, backupFile := range backupFiles {
				if backupFile.ModTime().Before(ago) {
					filepath := path + "/" + backupFile.Name()
					err := os.Remove(filepath)
					if err != nil {
						log.Printf("删除过期的文件 %s 失败", filepath)
					} else {
						log.Printf("删除过期的文件 %s 成功", filepath)
					}
				}
			}
		}
	}

}

// SleepForFileDelete Sleep For File Delete
func SleepForFileDelete() {
	sleepHours := 24 - time.Now().Hour()
	log.Printf("%d小时后再次运行：删除过期的备份文件", sleepHours)
	time.Sleep(time.Hour * time.Duration(sleepHours))
}
