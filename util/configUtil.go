package util

import (
	"backup-db/entity"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
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

// GetConfig 获得配置文件
func getConfigInner(workDir string) (*entity.Config, error) {
	configPath := workDir + string(os.PathSeparator) + "config.yml"
	_, err := os.Stat(configPath)

	if err == nil {
		// 找到配置文件
		configFile, err := os.Open(configPath)
		if err != nil {
			log.Fatal("打开配件文件失败")
			return nil, err
		}

		config := entity.Config{}
		b, err := ioutil.ReadAll(configFile)
		if err != nil {
			log.Fatal("读取配件文件失败")
			return nil, err
		}

		err = yaml.Unmarshal(b, &config)
		if err != nil {
			log.Fatal("读取config.yml配件文件失败, 注意冒号后面必须跟一个空格。\n", err)
			return nil, err
		}
		cachedConfig = &config
		return cachedConfig, nil
	}

	log.Fatal(err)
	return nil, err
}
