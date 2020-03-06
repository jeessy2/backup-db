package util

import (
	"io/ioutil"
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

// DeleteOlderBackup delete older backup
func DeleteOlderBackup() {
	// sleep 30 minutes
	time.Sleep(1 * time.Minute)
	for {
		// get all projects
		projects, err := ioutil.ReadDir(parentSavePath)
		if err != nil {
			log.Println("Read dir with errr :", err)
			continue
		}

		// delete
		for _, project := range projects {
			backupFiles, err := ioutil.ReadDir(parentSavePath + "/" + project.Name())
			if err != nil {
				log.Println("Read dir with errr :", err)
				break
			}

			delete(project, backupFiles)
		}
		// sleep
		sleep()
	}
}

// delete one by one
func delete(project os.FileInfo, backupFiles []os.FileInfo) {

	ago := time.Now()
	lastDay, _ := time.ParseDuration("-" + strconv.Itoa(GetConfig().MaxSaveDays*24) + "h")
	ago = ago.Add(lastDay)

	// delete older file when file numbers gt MaxSaveDays
	for _, backupFile := range backupFiles {
		if backupFile.ModTime().Before(ago) {
			filepath := parentSavePath + "/" + project.Name() + "/" + backupFile.Name()
			err := os.Remove(filepath)
			if err != nil {
				log.Printf("Delete older file %s failed", filepath)
			} else {
				log.Printf("Delete older file %s successed", filepath)
			}
		}
	}
}

func sleep() {
	sleepHours := 24 - time.Now().Hour() 
	log.Println("Run delete older file after", sleepHours, "hours again")
	time.Sleep(time.Hour * time.Duration(sleepHours))
}
