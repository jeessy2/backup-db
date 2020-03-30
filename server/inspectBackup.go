package server

import (
	"backup-db/client"
	"backup-db/util"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

// InspectBackup Inspect Backup
func InspectBackup() {
	for {
		sleep()
		inspectInner()
	}
}

func inspectInner() {
	log.Println("Start inspect backup files")
	// get all projects
	projects, err := ioutil.ReadDir(client.ParentSavePath)
	if err != nil {
		log.Println("Read dir with error :", err)
	}

	todayString := time.Now().Format("2006-01-02")
	// delete
	for _, project := range projects {
		backupFiles, err := ioutil.ReadDir(client.ParentSavePath + "/" + project.Name())
		if err != nil {
			log.Println("Read dir with error :", err)
			break
		}

		find := false
		for _, backupFile := range backupFiles {
			if strings.Contains(backupFile.Name(), todayString) {
				find = true
				break
			}
		}

		// not find, send email
		if !find {
			util.SendEmail("The \""+project.Name()+"\" is not uploaded today", "Please check your project \""+project.Name()+"\"")
		}

	}

}

func sleep() {
	sleepHours := 10 - time.Now().Hour()
	if sleepHours < 0 {
		sleepHours = 24 + 10 - time.Now().Hour()
	}
	log.Println("Run inspect backup files again after", sleepHours, "hours")
	time.Sleep(time.Hour * time.Duration(sleepHours))
}
