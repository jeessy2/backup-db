package util

import (
	"crypto/tls"
	"log"
	"strings"

	"gopkg.in/gomail.v2"
)

// SendEmail send email
func SendEmail(subject string, content string) bool {
	conf, err := GetConfig()
	if err == nil {
		if strings.Contains(conf.NoticeEmail, "@") {
			m := gomail.NewMessage()
			m.SetHeader("From", conf.SMTPUsername)
			m.SetHeader("To", strings.Split(conf.NoticeEmail, ",")...)
			m.SetHeader("Subject", subject)
			m.SetBody("text/html", content)

			d := gomail.NewDialer(conf.SMTPHost, conf.SMTPPort, conf.SMTPUsername, conf.SMTPPassword)
			d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
			// Send the email to Bob, Cora and Dan.
			if err := d.DialAndSend(m); err != nil {
				log.Println(err)
			} else {
				log.Printf("Sending email success! To: %s, Subject: %s", conf.NoticeEmail, subject)
				return true
			}
		} else {
			log.Printf("邮箱不正确%s", conf.NoticeEmail)
		}
	}

	return false
}
