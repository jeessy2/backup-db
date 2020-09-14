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

	conf, err := entity.GetConfigCache()
	if err == nil {
		tmpl.Execute(writer, conf)
		return
	}

	// default config
	// 获得环境变量
	backupConf := []entity.BackupConfig{}
	for i := 0; i < 16; i++ {
		backupConf = append(backupConf, entity.BackupConfig{SaveDays: 30})
	}
	conf = entity.Config{
		Server: entity.Server{
			Type:   util.GetEnvType(),
			DBType: util.GetEnvDBType(),
		},
		BackupConfig: backupConf,
	}

	tmpl.Execute(writer, conf)
}
