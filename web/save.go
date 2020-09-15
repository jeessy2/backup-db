package web

import (
	"backup-db/entity"
	"backup-db/util"
	"net/http"
	"strconv"
)

// Save 保存
func Save(writer http.ResponseWriter, request *http.Request) {
	conf := &entity.Config{}

	conf.Type = util.GetEnvType()
	conf.DBType = util.GetEnvDBType()

	// 覆盖以前的配置
	conf.UploadURL = request.FormValue("UploadURL")
	conf.Username = request.FormValue("Username")
	conf.Password = request.FormValue("Password")

	forms := request.PostForm
	for index, projectName := range forms["ProjectName"] {
		saveDays, _ := strconv.Atoi(forms["SaveDays"][index])
		conf.BackupConfig = append(
			conf.BackupConfig,
			entity.BackupConfig{
				ProjectName: projectName,
				Command:     forms["Command"][index],
				SaveDays:    saveDays,
			},
		)
	}

	conf.NoticeConfig.BackupSuccessNotice = request.FormValue("BackupSuccessNotice") == "on"

	// DingDing
	conf.DingDing.WebHook = request.FormValue("WebHook")
	conf.DingDing.Secret = request.FormValue("Secret")

	// Email
	conf.NoticeEmail = request.FormValue("NoticeEmail")
	conf.SMTPHost = request.FormValue("SMTPHost")
	smtpPort, _ := strconv.Atoi(request.FormValue("SMTPPort"))
	conf.SMTPPort = smtpPort
	conf.SMTPUsername = request.FormValue("SMTPUsername")
	conf.SMTPPassword = request.FormValue("SMTPPassword")

	// 保存到用户目录
	err := conf.SaveConfig()

	// 没有错误，运行一次
	if err == nil {
		go RunOnce()
	}

	// 回写错误信息
	if err == nil {
		writer.Write([]byte("ok"))
	} else {
		writer.Write([]byte(err.Error()))
	}

}
