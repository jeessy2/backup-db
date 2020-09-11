package entity

// BackupConfig 备份配置
type BackupConfig struct {
	// 项目名称
	ProjectName string
	// 命令
	Command  string
	SaveDays int
}

// ParentSavePath Parent Save Path
const ParentSavePath = "backup-files"

// GetProjectPath 获得项目路径
func (backupConfig *BackupConfig) GetProjectPath() string {
	return ParentSavePath + "/" + backupConfig.ProjectName
}

// NotEmptyProject 是不是空的项目
func (backupConfig *BackupConfig) NotEmptyProject() bool {
	return backupConfig.Command != "" && backupConfig.ProjectName != ""
}
