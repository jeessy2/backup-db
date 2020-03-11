package util

import (
	"log"
	"os"
	"strconv"
	"time"
)

const parentSavePath = "backup-files"

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

	ago := time.Now()
	lastDay, _ := time.ParseDuration("-" + strconv.Itoa(GetConfig().MaxSaveDays*24) + "h")
	ago = ago.Add(lastDay)

	// delete older file when file numbers gt MaxSaveDays
	for _, backupFile := range backupFiles {
		if backupFile.ModTime().Before(ago) {
			filepath := path + "/" + backupFile.Name()
			err := os.Remove(filepath)
			if err != nil {
				log.Printf("Delete older file %s failed", filepath)
			} else {
				log.Printf("Delete older file %s successed", filepath)
			}
		}
	}
}

// SleepForFileDelete Sleep For File Delete
func SleepForFileDelete() {
	sleepHours := 24 - time.Now().Hour()
	log.Println("Run delete older file after", sleepHours, "hours again")
	time.Sleep(time.Hour * time.Duration(sleepHours))
}
