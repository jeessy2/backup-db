package util

import (
	"backup-db/entity"
	"encoding/hex"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"
	"sync"

	"gopkg.in/yaml.v2"
)

var lock sync.Mutex
var conf *entity.Config

// GetConfig 获得配置
func GetConfig() (config *entity.Config, err error) {

	if conf != nil {
		return conf, nil
	}

	//获取锁
	lock.Lock()
	//释放锁
	defer lock.Unlock()

	//业务逻辑操作
	configFilePath := GetConfigFilePath()
	_, err = os.Stat(configFilePath)
	if err != nil {
		log.Println("没有找到配置文件！请在网页中输入")
		return conf, err
	}
	byt, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Println("config.yaml读取失败")
		return conf, err
	}

	conf = &entity.Config{}
	err = yaml.Unmarshal(byt, conf)
	if err == nil {
		conf.PasswordBytes = execPassword(conf.Password)
	}

	return conf, err
}

// ClearConfigCache 清空配置cache
func ClearConfigCache() {
	conf = nil
}

// GetConfigFilePath 获得配置文件路径
func GetConfigFilePath() string {
	u, err := user.Current()
	if err != nil {
		log.Println("Geting current user failed!")
		return "../.backup_db_docker_config.yaml"
	}
	return u.HomeDir + string(os.PathSeparator) + ".backup_db_docker_config.yaml"
}

func execPassword(secretKey string) []byte {
	if secretKey == "" {
		nonce, _ := hex.DecodeString("68af433ace5112d34fad3e24")
		return nonce
	}

	// replace others to 1~9, a~f
	for _, key := range secretKey {
		if key >= 48 && key <= 57 {
			continue
		}
		if key >= 97 && key <= 102 {
			continue
		}
		secretKey = strings.ReplaceAll(secretKey, string(key), "b")
	}
	// must be 24
	oriLen := len(secretKey)
	if oriLen < 24 {
		for i := 0; i < 24-oriLen; i++ {
			secretKey += "a"
		}
	}
	if oriLen > 24 {
		secretKey = secretKey[:24]
	}
	// decode
	decode, _ := hex.DecodeString(secretKey)
	return decode
}
