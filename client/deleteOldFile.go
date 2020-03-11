package client

import (
	"backup-db/util"
	"io/ioutil"
	"log"
	"time"
)

// DeleteOldBackup for client
func DeleteOldBackup() {
	// sleep 30 minutes
	time.Sleep(30 * time.Minute)
	for {
		log.Println("Start deleting old backup files")
		// read from current path
		backupFiles, err := ioutil.ReadDir(".")
		if err != nil {
			log.Println("Read dir with error :", err)
			continue
		}

		// delete client files
		util.DeleteOlderFiles(".", backupFiles)
		// sleep
		util.SleepForFileDelete()
	}

}
