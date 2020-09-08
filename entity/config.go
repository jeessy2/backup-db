package entity

// Config yml格式的配置文件
// go的实体需大写对应config.yml的key, key全部小写
type Config struct {
	Server
	ProjectName string
	// 命令
	Command    string
	SaveDays   int
	BackupPath string
	SMTPConfig
}

// ParentSavePath Parent Save Path
const ParentSavePath = "backup-files"

func (conf *Config) GetProjectPath() string {
	return ParentSavePath + "/" + conf.ProjectName
}
