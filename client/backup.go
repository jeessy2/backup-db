package client

import (
	"backup-db/util"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// ParentSavePath Parent Save Path
const ParentSavePath = "backup-files"

// StartBackup start backup db
func StartBackup() {
	var unSendFiles = []string{}
	err := prepare()
	if err == nil {
		for {
			// backup
			outFileName, err := backup()
			if err == nil {
				// send file to server
				err = SendFile(outFileName.Name())
				if err != nil {
					unSendFiles = append(unSendFiles, outFileName.Name())
				} else {
					unSendFiles = sendFileAgain(unSendFiles)
				}
			} else {
				// send email
				util.SendEmail("The \""+util.GetConfig().ProjectName+"\" is backup failed!", err.Error())
			}
			// sleep to tomorrow night
			sleep()
		}
	} else {
		log.Println(err)
		// sleep to tomorrow night
		sleep()
	}
}

// prepare
func prepare() (err error) {
	projectPath := ParentSavePath + "/" + util.GetConfig().ProjectName
	// create floder
	os.MkdirAll(projectPath, 0755)
	os.Chdir(projectPath)

	if !strings.Contains(util.GetConfig().Command, "#{DATE}") {
		err = errors.New("backup_command must contains #{DATE}")
	}

	return
}

func backup() (outFileName os.FileInfo, err error) {
	projectName := util.GetConfig().ProjectName
	log.Println("Starting backup:", projectName)

	todayString := time.Now().Format("2006-01-02")
	shellString := strings.ReplaceAll(util.GetConfig().Command, "#{DATE}", todayString)

	// create shell file
	shellName := time.Now().Format("2006_01_02_") + "backup.sh"

	file, err := os.Create(shellName)
	file.Chmod(0700)
	if err == nil {
		file.WriteString(shellString)
		file.Close()
	} else {
		log.Println("Create file with error: ", err)
	}

	// run shell file
	shell := exec.Command("bash", shellName)
	shell.Stderr = os.Stderr
	shell.Stdout = os.Stdout
	err = shell.Run()
	// execute shell success
	if err == nil {
		// find backup file by todayString
		outFileName, err = findBackupFile(todayString)

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
func findBackupFile(todayString string) (backupFile os.FileInfo, err error) {
	files, err := ioutil.ReadDir(".")
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
func sendFileAgain(unSendFiles []string) []string {
	newUnSendFils := []string{}
	for _, file := range unSendFiles {
		if nil != SendFile(file) {
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
