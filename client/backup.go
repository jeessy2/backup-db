package client

import (
	"backup-db/util"
	"log"
	"os"
	"os/exec"
	"time"
)

// StartBackup start backup db
func StartBackup() {
	for {
		// backup
		outFileName := backup()
		// send file to server
		SendFile(outFileName)
		// sleep to tomorrow night
		sleep()
	}
}

func backup() (outFileName string) {
	projectName := util.GetConfig().ProjectName
	command := util.GetConfig().Command
	log.Println("Starting backup:", projectName)

	// create shell file
	shellName := projectName + "backup.sh"
	outFileName = projectName + "/" + projectName + time.Now().Format("2006-01-02") + ".sql"
	// create floder
	os.Mkdir(projectName, 0700)
	file, err := os.Create(shellName)
	file.Chmod(0700)
	if err == nil {
		file.WriteString(command + " > " + outFileName)
		file.Close()
	}

	// run shell file
	shell := exec.Command("bash", shellName)
	shell.Stderr = os.Stderr
	shell.Stdout = os.Stdout
	shell.Run()
	log.Println(projectName, "Complete backup!")

	return
}

func sleep() {
	sleepHours := 24 - time.Now().Hour()
	log.Println("Run again after", sleepHours, "hours")
	time.Sleep(time.Hour * time.Duration(sleepHours))
	// time.Sleep(time.Second * 10)
}
