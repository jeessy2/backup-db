package client

import (
	"backup-db/entity"
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
		conf, err := entity.GetConfigCache()
		if err == nil {
			for _, backupConf := range conf.BackupConfig {
				// read from current path
				backupFiles, err := ioutil.ReadDir(backupConf.GetProjectPath())
				if err != nil {
					log.Println("Read dir with error :", err)
					continue
				}

				// delete client files
				util.DeleteOlderFiles(backupConf.GetProjectPath(), backupFiles)
			}
		}
		// sleep
		util.SleepForFileDelete()
	}

}
