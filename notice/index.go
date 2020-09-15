package notice

// NoticeConfig 通知设置
type NoticeConfig struct {
	BackupSuccessNotice bool
}

// SendMessage interface
type SendMessage interface {
	CanBeSend() bool
	SendMessage(title, message string) error
}
