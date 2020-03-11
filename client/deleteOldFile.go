package client

import (
	"backup-db/util"
	"io/ioutil"
	"log"
)

// DeleteOlderBackup for client
func DeleteOlderBackup() {
	for {
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
