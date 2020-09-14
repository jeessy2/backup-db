package server

import (
	"backup-db/entity"
	"fmt"
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
	conf, err := entity.GetConfigCache()
	if err == nil {
		log.Println("开始检测备份文件")
		// get all projects
		projects, err := ioutil.ReadDir(entity.ParentSavePath)
		if err != nil {
			log.Println("Read dir with error :", err)
		}

		todayString := time.Now().Format("2006-01-02")
		// delete
		for _, project := range projects {
			backupFiles, err := ioutil.ReadDir(entity.ParentSavePath + "/" + project.Name())
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
				conf.SendMessage(
					fmt.Sprintf("%s 今日没有上传备份文件", project),
					fmt.Sprintf("%s 今日没有上传备份文件, 请检测!", project),
				)
			}
		}
	}
}

func sleep() {
	sleepHours := 10 - time.Now().Hour()
	if sleepHours <= 0 {
		sleepHours = 24 + 10 - time.Now().Hour()
	}
	log.Printf("%d小时后再次运行：检测备份文件", sleepHours)
	time.Sleep(time.Hour * time.Duration(sleepHours))
}
