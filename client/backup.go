package client

import (
	"backup-db/entity"
	"backup-db/util"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var unSendFiles = []string{}

// StartBackup start backup db
func StartBackup() {
	for {
		RunOnce()
		// sleep to tomorrow night
		sleep()
	}
}

// RunOnce 运行一次
func RunOnce() (unSendFiles []string) {
	conf, err := util.GetConfig()
	if err == nil {
		err := prepare(conf)
		// backup
		outFileName, err := backup(conf)
		if err == nil {
			// send file to server
			err = SendFile(conf, outFileName.Name())
			if err != nil {
				unSendFiles = append(unSendFiles, outFileName.Name())
			} else {
				unSendFiles = sendFileAgain(conf, unSendFiles)
			}
		} else {
			// send email
			util.SendEmail(fmt.Sprintf("The %s is backup failed!", conf.ProjectName), err.Error())
		}
	}
	return
}

// prepare
func prepare(conf *entity.Config) (err error) {
	// create floder
	os.MkdirAll(conf.GetProjectPath(), 0755)

	if !strings.Contains(conf.Command, "#{DATE}") {
		err = errors.New("Command must containes #{DATE}")
	}

	return
}

func backup(conf *entity.Config) (outFileName os.FileInfo, err error) {
	projectName := conf.ProjectName
	log.Println("Starting backup:", projectName)

	todayString := time.Now().Format("2006-01-02")
	shellString := strings.ReplaceAll(conf.Command, "#{DATE}", todayString)

	// create shell file
	shellName := time.Now().Format("2006_01_02_") + "backup.sh"

	file, err := os.Create(conf.GetProjectPath() + "/" + shellName)
	file.Chmod(0700)
	if err == nil {
		file.WriteString(shellString)
		file.Close()
	} else {
		log.Println("Create file with error: ", err)
	}

	// run shell file
	shell := exec.Command("bash", shellName)
	shell.Dir = conf.GetProjectPath()
	shell.Stderr = os.Stderr
	shell.Stdout = os.Stdout
	err = shell.Run()
	// execute shell success
	if err == nil {
		// find backup file by todayString
		outFileName, err = findBackupFile(conf, todayString)

		// check file size
		if err != nil {
			log.Println(err)
		} else if outFileName.Size() >= 100 {
			log.Printf("Success backup: %s, file: %s", projectName, outFileName.Name())
		} else {
			err = errors.New("The \"" + projectName + "\" backup file size less than 100 bytes, Current size is " + strconv.Itoa(int(outFileName.Size())))
			log.Println(err)
		}
	} else {
		log.Println("Execute shell with error:", err)
	}

	return
}

// find backup file by todayString
func findBackupFile(conf *entity.Config, todayString string) (backupFile os.FileInfo, err error) {
	files, err := ioutil.ReadDir(conf.GetProjectPath())
	for _, file := range files {
		if strings.Contains(file.Name(), todayString) {
			backupFile = file
			return
		}
	}
	err = errors.New("Can't find the backup file which containes " + todayString)
	return
}

// send file again
func sendFileAgain(conf *entity.Config, unSendFiles []string) []string {
	newUnSendFils := []string{}
	for _, file := range unSendFiles {
		if nil != SendFile(conf, file) {
			newUnSendFils = append(newUnSendFils, file)
		}
	}
	return newUnSendFils
}

func sleep() {
	sleepHours := 24 - time.Now().Hour()
	log.Println("Run again after", sleepHours, "hours")
	time.Sleep(time.Hour * time.Duration(sleepHours))
	// time.Sleep(time.Second * 10)
}
