package notice

import (
	"crypto/tls"
	"fmt"
	"log"
	"strings"

	"gopkg.in/gomail.v2"
)

// EmailConfig smtp
type EmailConfig struct {
	NoticeEmail  string
	SMTPHost     string
	SMTPPort     int
	SMTPUsername string
	SMTPPassword string
}

// CanBeSend 能否发送
func (email *EmailConfig) CanBeSend() bool {
	return email.NoticeEmail != "" && strings.Contains(email.NoticeEmail, "@") &&
		email.SMTPHost != "" && email.SMTPUsername != "" && email.SMTPPassword != ""
}

// SendMessage 发送钉钉消息
func (email *EmailConfig) SendMessage(title, message string) (err error) {
	if strings.Contains(email.NoticeEmail, "@") {
		m := gomail.NewMessage()
		m.SetHeader("From", email.SMTPUsername)
		m.SetHeader("To", strings.Split(email.NoticeEmail, ",")...)
		m.SetHeader("Subject", title)
		m.SetBody("text/html", message)

		d := gomail.NewDialer(email.SMTPHost, email.SMTPPort, email.SMTPUsername, email.SMTPPassword)
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
		// Send the email to Bob, Cora and Dan.
		if err := d.DialAndSend(m); err != nil {
			log.Printf("发送邮件失败! 主题: %s。错误信息：%s\n", title, err.Error())
		} else {
			log.Printf("发送邮件成功! 主题: %s\n", title)
			return nil
		}
	} else {
		return fmt.Errorf("邮箱不正确%s", email.NoticeEmail)
	}

	return
}
