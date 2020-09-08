package main

import (
	"backup-db/util"
	"backup-db/web"
	"os"

	"log"
	"net/http"
	"time"
)

var defaultPort = "9977"

func main() {

	// 启动静态文件服务
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/favicon.ico", http.StripPrefix("/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", util.BasicAuth(web.WritingConfig))
	http.HandleFunc("/save", util.BasicAuth(web.Save))
	http.HandleFunc("/logs", util.BasicAuth(web.Logs))
	http.HandleFunc("/upload", util.BasicAuth(web.Upload))

	// 运行
	go web.Run()

	if os.Getenv("port") != "" {
		defaultPort = os.Getenv("port")
	}

	err := http.ListenAndServe(":"+defaultPort, nil)

	if err != nil {
		log.Println("启动端口发生异常, 1分钟后自动关闭此窗口", err)
		time.Sleep(time.Minute)
	}

}
