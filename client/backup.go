package client

import (
	"backup-db/entity"
	"backup-db/util"
	"log"
	"os"
	"os/exec"
	"time"
)

// StartBackup start backup db
func StartBackup() {
	for {
		for _, cmd := range util.GetConfig().Cmds {
			// backup
			outFileName := backup(cmd)
			// send file to server
			SendFile(outFileName)
			// sleep to tomorrow night
			sleep()
		}
	}
}

func backup(cmd entity.Commands) (outFileName string) {
	log.Println("Starting backup:", cmd.Name)

	// create shell file
	shellName := cmd.Name + "backup.sh"
	outFileName = cmd.Name + "/" + cmd.Name + time.Now().Format("2006-01-02") + ".sql"
	file, err := os.Create(shellName)
	file.Chmod(0700)
	if err == nil {
		file.WriteString(cmd.Command + " > " + outFileName)
		file.Close()
	}

	// run shell file
	shell := exec.Command("bash", shellName)
	shell.Stderr = os.Stderr
	shell.Stdout = os.Stdout
	shell.Run()
	log.Println(cmd.Name, "Complete backup!")

	return
}

func sleep() {
	sleepHours := 24 - time.Now().Hour()
	log.Println("Run again after", sleepHours, "hours")
	time.Sleep(time.Hour * time.Duration(sleepHours))
	// time.Sleep(time.Second * 2)
}
