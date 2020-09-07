package entity

// SMTPConfig 邮件配置
type SMTPConfig struct {
	NoticeEmail  string
	SMTPHost     string
	SMTPPort     int
	SMTPUsername string
	SMTPPassword string
}
