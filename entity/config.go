package entity

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"sync"

	"gopkg.in/yaml.v2"
)

// Config yml格式的配置文件
// go的实体需大写对应config.yml的key, key全部小写
type Config struct {
	Server
	User
	BackupConfig []BackupConfig
	SMTPConfig
}

// ConfigCache ConfigCache
type cacheType struct {
	ConfigSingle *Config
	Err          error
	Lock         sync.Mutex
}

var cache = &cacheType{}

// GetConfigCache 获得配置
func GetConfigCache() (conf Config, err error) {

	if cache.ConfigSingle != nil {
		return *cache.ConfigSingle, cache.Err
	}

	cache.Lock.Lock()
	defer cache.Lock.Unlock()

	// init config
	cache.ConfigSingle = &Config{}

	configFilePath := GetConfigFilePath()
	_, err = os.Stat(configFilePath)
	if err != nil {
		log.Println("没有找到配置文件！请在网页中输入")
		cache.Err = err
		return *cache.ConfigSingle, err
	}

	byt, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Println("config.yaml读取失败")
		cache.Err = err
		return *cache.ConfigSingle, err
	}

	err = yaml.Unmarshal(byt, cache.ConfigSingle)
	if err != nil {
		log.Println("反序列化配置文件失败", err)
		cache.Err = err
		return *cache.ConfigSingle, err
	}
	// remove err
	cache.Err = nil
	return *cache.ConfigSingle, err
}

// SaveConfig 保存配置
func (conf *Config) SaveConfig() (err error) {
	byt, err := yaml.Marshal(conf)
	if err != nil {
		log.Println(err)
		return err
	}

	err = ioutil.WriteFile(GetConfigFilePath(), byt, 0600)
	if err != nil {
		log.Println(err)
		return
	}

	// 清空配置缓存
	cache.ConfigSingle = nil

	return
}

// GetConfigFilePath 获得配置文件路径
func GetConfigFilePath() string {
	u, err := user.Current()
	if err != nil {
		log.Println("Geting current user failed!")
		return "../.backup_db_docker_config.yaml"
	}

	// 自定义path
	if os.Getenv("configPath") != "" {
		return u.HomeDir + string(os.PathSeparator) + os.Getenv("configPath")
	}

	return u.HomeDir + string(os.PathSeparator) + ".backup_db_docker_config.yaml"
}
