package util

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
	"log"
	"strings"
)

// SendEmail send email
func SendEmail(subject string, content string) bool {
	if strings.Contains(GetConfig().NoticeEmail, "@") {
		config := GetConfig().SMTPConfig
		m := gomail.NewMessage()
		m.SetHeader("From", config.Username)
		m.SetHeader("To", strings.Split(GetConfig().NoticeEmail, ",")...)
		m.SetHeader("Subject", subject)
		m.SetBody("text/html", content)

		d := gomail.NewDialer(config.Host, config.Port, config.Username, config.Password)
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
		// Send the email to Bob, Cora and Dan.
		if err := d.DialAndSend(m); err != nil {
			log.Println(err)
		}else {
			log.Printf("Sending email success! To: %s, Subject: %s", GetConfig().NoticeEmail, subject)
			return true
		}
	}

	return false
}