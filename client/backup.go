package client

import (
	"backup-db/util"
	"log"
	"os"
	"os/exec"
	"time"
)

const parentSavePath = "backup-files"

// StartBackup start backup db
func StartBackup() {
	for {
		// backup
		outFileName, err := backup()
		if err == nil {
			// send file to server
			SendFile(outFileName)
		}
		// sleep to tomorrow night
		sleep()
	}
}

func backup() (outFileName string, err error) {
	projectName := util.GetConfig().ProjectName
	command := util.GetConfig().Command
	log.Println("Starting backup:", projectName)

	// create shell file
	shellName := projectName + "backup.sh"
	// create floder
	os.MkdirAll(parentSavePath+"/"+projectName, 0755)

	outFileName = parentSavePath + "/" + projectName + "/" + projectName + time.Now().Format("2006-01-02") + ".sql"

	file, err := os.Create(shellName)
	file.Chmod(744)
	if err == nil {
		file.WriteString(command + " > " + outFileName)
		file.Close()
	} else {
		log.Println("Create file with error: ", err)
	}

	// run shell file
	shell := exec.Command("bash", shellName)
	shell.Stderr = os.Stderr
	shell.Stdout = os.Stdout
	shell.Run()

	fileInfo, err := os.Stat(outFileName)
	if err != nil {
		log.Println("Backup failed:", projectName)
	} else if fileInfo.Size() >= 100 {
		log.Println("Success backup:", projectName)
	} else {
		log.Println("Backup file size less than 100 bytes")
	}

	return
}

func sleep() {
	sleepHours := 24 - time.Now().Hour()
	log.Println("Run again after", sleepHours, "hours")
	time.Sleep(time.Hour * time.Duration(sleepHours))
	// time.Sleep(time.Second * 10)
}
