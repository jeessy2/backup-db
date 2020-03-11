package server

import (
	"backup-db/client"
	"backup-db/util"
	"io/ioutil"
	"log"
)

// DeleteOlderBackup for server
func DeleteOlderBackup() {
	for {
		// get all projects
		projects, err := ioutil.ReadDir(client.ParentSavePath)
		if err != nil {
			log.Println("Read dir with error :", err)
			continue
		}

		// delete
		for _, project := range projects {
			backupFiles, err := ioutil.ReadDir(client.ParentSavePath + "/" + project.Name())
			if err != nil {
				log.Println("Read dir with error :", err)
				break
			}

			// delete server files
			util.DeleteOlderFiles(client.ParentSavePath + "/" + project.Name(), backupFiles)
		}
		// sleep
		util.SleepForFileDelete()
	}

}
