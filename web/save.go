package web

import (
	"backup-db/entity"
	"backup-db/util"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

// Save 保存
func Save(writer http.ResponseWriter, request *http.Request) {
	conf := &entity.Config{}

	typ := os.Getenv("BACKUP_TYPE")
	if typ == "" {
		conf.Type = "client"
	} else {
		conf.Type = typ
	}
	conf.DBType = os.Getenv("BACKUP_DB_TYPE")

	// 覆盖以前的配置
	conf.UploadURL = request.FormValue("UploadURL")
	conf.Username = request.FormValue("Username")
	conf.Password = request.FormValue("Password")

	conf.ProjectName = request.FormValue("ProjectName")
	conf.Command = request.FormValue("Command")
	saveDays, _ := strconv.Atoi(request.FormValue("SaveDays"))
	conf.SaveDays = saveDays

	conf.NoticeEmail = request.FormValue("NoticeEmail")
	conf.SMTPHost = request.FormValue("SMTPHost")
	smtpPort, _ := strconv.Atoi(request.FormValue("SMTPPort"))
	conf.SMTPPort = smtpPort
	conf.SMTPUsername = request.FormValue("SMTPUsername")
	conf.SMTPPassword = request.FormValue("SMTPPassword")

	// 保存到用户目录
	byt, err := yaml.Marshal(conf)
	if err != nil {
		log.Println(err)
	}

	err = ioutil.WriteFile(util.GetConfigFilePath(), byt, 0600)

	// clear cache
	if err == nil {
		util.ClearConfigCache()
		go RunOnce()
	} else {
		log.Println(err)
	}

	// 跳转
	http.Redirect(writer, request, "/?saveOk=true", http.StatusFound)

}
