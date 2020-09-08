package web

import (
	"backup-db/entity"
	"backup-db/util"
	"html/template"
	"log"
	"net/http"
)

// WritingConfig 填写配置信息
func WritingConfig(writer http.ResponseWriter, request *http.Request) {
	tmpl, err := template.ParseFiles("./static/pages/config.html")
	if err != nil {
		log.Println(err)
		return
	}

	conf, err := util.GetConfig()
	if err == nil {
		tmpl.Execute(writer, conf)
		return
	}

	// default config
	conf = &entity.Config{
		Server: entity.Server{
			Type: "client",
		},
		SaveDays: 3,
		SMTPConfig: entity.SMTPConfig{
			SMTPPort: 587,
		},
	}

	tmpl.Execute(writer, conf)
}
