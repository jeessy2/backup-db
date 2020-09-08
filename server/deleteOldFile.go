package server

import (
	"backup-db/entity"
	"backup-db/util"
	"io/ioutil"
	"log"
	"time"
)

// DeleteOldBackup for server
func DeleteOldBackup() {
	// sleep 30 minutes
	time.Sleep(30 * time.Minute)
	for {
		log.Println("Start deleting old backup files")
		// get all projects
		projects, err := ioutil.ReadDir(entity.ParentSavePath)
		if err != nil {
			log.Println("Read dir with error :", err)
			continue
		}

		// delete
		for _, project := range projects {
			backupFiles, err := ioutil.ReadDir(entity.ParentSavePath + "/" + project.Name())
			if err != nil {
				log.Println("Read dir with error :", err)
				break
			}

			// delete server files
			util.DeleteOlderFiles(entity.ParentSavePath+"/"+project.Name(), backupFiles)
		}
		// sleep
		util.SleepForFileDelete()
	}

}
