package main

import (
	"backup-db/web"

	"log"
	"net/http"
	"time"
)

const port = "9978"

func main() {

	// 启动静态文件服务
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/favicon.ico", http.StripPrefix("/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", web.WritingConfig)
	http.HandleFunc("/save", web.Save)
	http.HandleFunc("/logs", web.Logs)

	// 运行
	go web.Run()

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Println("启动端口发生异常, 1分钟后自动关闭此窗口", err)
		time.Sleep(time.Minute)
	}

}
