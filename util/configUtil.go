package util

import (
	"backup-db/entity"
)

// 缓存已经读取到的配置
var cachedConfig *entity.Config

// GetConfig 获得配置文件
// func GetConfig() *entity.Config {
// 	if cachedConfig != nil {
// 		return cachedConfig
// 	}
// 	wd, err := os.Getwd()
// 	if err == nil {
// 		config, err := getConfigInner(wd)
// 		if err == nil {
// 			return config
// 		}
// 		panic(err)
// 	}
// 	panic(err)
// }

// func getConfigInner(workDir string) (*entity.Config, error) {
// 	config := entity.Config{}
// 	config.Server.IP = os.Getenv("backup_server_ip")

// 	serverPort, err := strconv.Atoi(os.Getenv("backup_server_port"))
// 	if err != nil {
// 		serverPort = 9977
// 		log.Println("backup_server_port default: ", serverPort)
// 	}
// 	config.Server.ServerPort = serverPort

// 	secretKey := os.Getenv("server_secret_key")
// 	if secretKey == "" {
// 		nonce, _ := hex.DecodeString("68af433ace5112d34fad3e24")
// 		config.Server.SecretKey = nonce
// 	} else {
// 		// replace others to 1~9, a~f
// 		for _, key := range secretKey {
// 			if key >= 48 && key <= 57 {
// 				continue
// 			}
// 			if key >= 97 && key <= 102 {
// 				continue
// 			}
// 			secretKey = strings.ReplaceAll(secretKey, string(key), "b")
// 		}
// 		// must be 24
// 		oriLen := len(secretKey)
// 		if oriLen < 24 {
// 			for i := 0; i < 24-oriLen; i++ {
// 				secretKey += "a"
// 			}
// 		}
// 		if oriLen > 24 {
// 			secretKey = secretKey[:24]
// 		}
// 		// decode
// 		decode, _ := hex.DecodeString(secretKey)
// 		config.Server.SecretKey = decode
// 	}

// 	config.ProjectName = os.Getenv("backup_project_name")
// 	config.Command = os.Getenv("backup_command")

// 	maxSaveDays, err := strconv.Atoi(os.Getenv("max_save_days"))
// 	if err != nil {
// 		maxSaveDays = 3
// 		log.Println("max_save_days default: ", maxSaveDays)
// 	}
// 	if maxSaveDays < 3 {
// 		maxSaveDays = 3
// 	}
// 	config.MaxSaveDays = maxSaveDays

// 	config.NoticeEmail = os.Getenv("notice_email")
// 	config.SMTPConfig.Host = os.Getenv("smtp_host")
// 	port, err := strconv.Atoi(os.Getenv("smtp_port"))
// 	if err != nil {
// 		port = 587
// 		log.Println("smtp_port default: ", port)
// 	}
// 	config.SMTPConfig.Port = port
// 	config.SMTPConfig.Username = os.Getenv("smtp_username")
// 	config.SMTPConfig.Password = os.Getenv("smtp_password")

// 	return &config, nil

// }
