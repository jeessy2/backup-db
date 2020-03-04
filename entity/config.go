package entity

// Config yml格式的配置文件
// go的实体需大写对应config.yml的key, key全部小写
type Config struct {
	Server struct {
		IP         string
		ServerPort int `yaml:"server_port"`
	}
	MaxSaveDays int    `yaml:"max_save_days"`
	BackupPath  string `yaml:"backup_path"`
	Cmds        []Commands
	NoticeEmail []string `yaml:"notice_email"`
}

// Commands 命令实体
type Commands struct {
	// 名称
	Name string
	// 命令
	Command  string
}
