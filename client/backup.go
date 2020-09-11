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
	conf, err := entity.GetConfigCache()
	if err == nil {
		// 迭代所有项目
		for _, backupConf := range conf.BackupConfig {
			if backupConf.NotEmptyProject() {
				err := prepare(backupConf)
				// backup
				outFileName, err := backup(backupConf)
				if err == nil {
					// send file to server
					err = SendFile(&conf, backupConf, outFileName.Name())
					if err != nil {
						unSendFiles = append(unSendFiles, outFileName.Name())
					} else {
						unSendFiles = sendFileAgain(&conf, backupConf, unSendFiles)
					}
				} else {
					// send email
					util.SendEmail(fmt.Sprintf("The %s is backup failed!", backupConf.ProjectName), err.Error())
				}
			}
		}
	}
	return
}

// prepare
func prepare(backupConf entity.BackupConfig) (err error) {
	// create floder
	os.MkdirAll(backupConf.GetProjectPath(), 0755)

	if !strings.Contains(backupConf.Command, "#{DATE}") {
		err = errors.New("备份脚本必须包含#{DATE}")
	}

	return
}

func backup(backupConf entity.BackupConfig) (outFileName os.FileInfo, err error) {
	projectName := backupConf.ProjectName
	log.Printf("正在备份项目: %s ...", projectName)

	todayString := time.Now().Format("2006-01-02")
	shellString := strings.ReplaceAll(backupConf.Command, "#{DATE}", todayString)

	// create shell file
	shellName := time.Now().Format("2006_01_02_") + "backup.sh"

	file, err := os.Create(backupConf.GetProjectPath() + "/" + shellName)
	file.Chmod(0700)
	if err == nil {
		file.WriteString(shellString)
		file.Close()
	} else {
		log.Println("Create file with error: ", err)
	}

	// run shell file
	shell := exec.Command("bash", shellName)
	shell.Dir = backupConf.GetProjectPath()
	outputBytes, err := shell.CombinedOutput()
	log.Printf("<span title=\"%s\">执行shell的输出：鼠标移动此处查看</span>", util.EscapeShell(string(outputBytes)))
	// execute shell success
	if err == nil {
		// find backup file by todayString
		outFileName, err = findBackupFile(backupConf, todayString)

		// check file size
		if err != nil {
			log.Println(err)
		} else if outFileName.Size() >= 100 {
			log.Printf("成功备份项目: %s, 文件名: %s\n", projectName, outFileName.Name())
		} else {
			err = errors.New(projectName + " 备份后的文件大小小于100字节, 当前大小：" + strconv.Itoa(int(outFileName.Size())))
			log.Println(err)
		}
	} else {
		log.Println("执行备份shell失败: ", util.EscapeShell(string(outputBytes)))
	}

	return
}

// find backup file by todayString
func findBackupFile(backupConf entity.BackupConfig, todayString string) (backupFile os.FileInfo, err error) {
	files, err := ioutil.ReadDir(backupConf.GetProjectPath())
	for _, file := range files {
		if strings.Contains(file.Name(), todayString) {
			backupFile = file
			return
		}
	}
	err = errors.New("不能找到备份后的文件，没有找到包含 " + todayString + " 的文件名")
	return
}

// send file again
func sendFileAgain(conf *entity.Config, backupConf entity.BackupConfig, unSendFiles []string) []string {
	newUnSendFils := []string{}
	for _, file := range unSendFiles {
		if nil != SendFile(conf, backupConf, file) {
			newUnSendFils = append(newUnSendFils, file)
		}
	}
	return newUnSendFils
}

func sleep() {
	sleepHours := 24 - time.Now().Hour()
	log.Println("下次运行时间：", sleepHours, "hours")
	time.Sleep(time.Hour * time.Duration(sleepHours))
	// time.Sleep(time.Second * 10)
}
