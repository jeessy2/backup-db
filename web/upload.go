package web

import (
	"io"
	"log"
	"net/http"
	"os"
)

// Upload 上传接口
func Upload(w http.ResponseWriter, r *http.Request) {

	// 根据字段名获取表单文件
	formFile, header, err := r.FormFile("uploadfile")
	path := r.FormValue("path")
	if err != nil {
		log.Printf("Get form file failed: %s\n", err)
		return
	}
	defer formFile.Close()

	// 创建目录
	if path != "" {
		_, err := os.Stat(path)
		if err != nil {
			os.MkdirAll(path, 0755)
		}
	}

	// 创建保存文件
	destFile, err := os.Create(path + "/" + header.Filename)
	if err != nil {
		log.Printf("Create failed: %s\n", err)
		return
	}
	defer destFile.Close()

	// 读取表单文件，写入保存文件
	_, err = io.Copy(destFile, formFile)
	if err != nil {
		log.Printf("Write file failed: %s\n", err)
		return
	}

	log.Printf("Client %s upload file \"%s\" success! ", r.RemoteAddr, destFile.Name())
}
