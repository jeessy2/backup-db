package entity

// Config yml格式的配置文件
// go的实体需大写对应config.yml的key, key全部小写
type Config struct {
	Server
	User
	BackupConfig
	SMTPConfig
}

// ParentSavePath Parent Save Path
const ParentSavePath = "backup-files"

// GetProjectPath 获得项目路径
func (conf *Config) GetProjectPath() string {
	return ParentSavePath + "/" + conf.ProjectName
}
