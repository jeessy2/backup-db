package client

import (
	"backup-db/entity"
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

// SendFile send file to server
func SendFile(conf *entity.Config, fileName string) (err error) {
	log.Printf("Starting send file to server: %s:%d", conf.Server.IP, conf.Server.Port)

	// 创建表单文件
	// CreateFormFile 用来创建表单，第一个参数是字段名，第二个参数是文件名
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	writer.WriteField("path", conf.GetProjectPath())
	formFile, err := writer.CreateFormFile("uploadfile", fileName)
	if err != nil {
		log.Fatalf("Create form file failed: %s\n", err)
	}

	// 从文件读取数据，写入表单
	srcFile, err := os.Open(conf.GetProjectPath() + "/" + fileName)
	if err != nil {
		log.Fatalf("%Open source file failed: s\n", err)
	}
	defer srcFile.Close()
	_, err = io.Copy(formFile, srcFile)
	if err != nil {
		log.Fatalf("Write to form file falied: %s\n", err)
	}

	// 发送表单
	contentType := writer.FormDataContentType()
	// 发送之前必须调用Close()以写入结尾行
	writer.Close()
	_, err = http.Post("http://"+conf.IP+"/upload", contentType, buf)
	if err != nil {
		log.Fatalf("Post failed: %s\n", err)
	} else {
		log.Printf("Send file to server: %s:%d success!", conf.Server.IP, conf.Server.Port)
	}

	return err

}
