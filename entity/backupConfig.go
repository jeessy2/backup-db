package entity

// BackupConfig 备份配置
type BackupConfig struct {
	// 项目名称
	ProjectName string
	// 命令
	Command  string
	SaveDays int
}
