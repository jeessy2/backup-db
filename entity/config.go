package entity

// Config yml格式的配置文件
// go的实体需大写对应config.yml的key, key全部小写
type Config struct {
	Server struct {
		IP              string
		ServerPort      int
		SecretKey       []byte
	}
	ProjectName string
	// 命令
	Command     string
	MaxSaveDays int
	BackupPath  string
	NoticeEmail string
	SMTPConfig  struct {
		Host     string
		Port     int
		Username string
		Password string
	}
}
