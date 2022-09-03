package main

import (
	"frontend/controller"
	"net/http"
)

// 用于启动http服务，处理界面的请求

func main() {
	http.Handle("/search", controller.CreateSearchResultHandler("view/template.html"))

	// 建立静态文件的处理, 默认会把index.html当做首页
	http.Handle("/", http.FileServer(http.Dir("view/")))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
