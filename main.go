package main

import (
	"backup-db/entity"
	"backup-db/util"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	for {
		for _, cmd := range util.GetConfig().Cmds {
			// backup
			go backup(cmd)
			// sleep to tomorrow night
			sleep()
		}
	}
}

func backup(cmd entity.Commands) {
	log.Println("Starting backup:", cmd.Name)

	// create shell file
	shellName := cmd.Name + "backup.sh"
	shellString := strings.ReplaceAll(cmd.Command, "$DATE$", time.Now().Format("2006-01-02"))
	file, err := os.Create(shellName)
	file.Chmod(0700)
	if err == nil {
		file.WriteString(shellString)
		file.Close()
	}

	// run shell file
	shell := exec.Command("bash", shellName)
	shell.Stderr = os.Stderr
	shell.Stdout = os.Stdout
	shell.Run()
	log.Println(cmd.Name, "Complete backup!")

}

func sleep() {
	sleepHours := 24 - time.Now().Hour()
	log.Println("Run again after", sleepHours, "hours")
	time.Sleep(time.Hour * time.Duration(sleepHours))
}
