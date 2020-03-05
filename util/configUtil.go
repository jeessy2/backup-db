package util

import (
	"backup-db/entity"
	"log"
	"os"
	"strconv"
)

// 缓存已经读取到的配置
var cachedConfig *entity.Config

// GetConfig 获得配置文件
func GetConfig() *entity.Config {
	if cachedConfig != nil {
		return cachedConfig
	}
	wd, err := os.Getwd()
	if err == nil {
		config, err := getConfigInner(wd)
		if err == nil {
			return config
		}
		panic(err)
	}
	panic(err)
}

func getConfigInner(workDir string) (*entity.Config, error) {
	config := entity.Config{}
	config.Server.IP = os.Getenv("backup_server_ip")

	serverPort, err := strconv.Atoi(os.Getenv("backup_server_port"))
	if err != nil {
		serverPort = 9977
		log.Println("backup_server_port default: ", serverPort)
	}
	config.Server.ServerPort = serverPort

	config.ProjectName = os.Getenv("backup_project_name")
	config.Command = os.Getenv("backup_command")

	maxSaveDays, err := strconv.Atoi(os.Getenv("max_save_days"))
	if err != nil {
		maxSaveDays = 3
		log.Println("max_save_days default: ", maxSaveDays)
	}
	if maxSaveDays < 3 {
		maxSaveDays = 3
	}
	config.MaxSaveDays = maxSaveDays

	config.NoticeEmail = os.Getenv("notice_email")

	return &config, nil

}
