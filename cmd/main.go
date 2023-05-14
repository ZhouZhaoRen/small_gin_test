package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println()
	router := InitRouter()
	//pprof.Register(router)

	s := &http.Server{
		//Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort), // 监听的TCP地址
		Addr:    fmt.Sprintf(":%d", 8080), // 监听的TCP地址
		Handler: router,                   // http句柄，用于处理程序响应http请求
		//ReadTimeout:    setting.ServerSetting.ReadTimeout,                  //允许读取请求头的最大数时间
		//WriteTimeout:   setting.ServerSetting.WriteTimeout,                 // 允许写入的最大时间
		MaxHeaderBytes: 1 << 20, // 请求头的最大字节数
	}

	// 这里和r.Run()没有本质上的区别
	s.ListenAndServe() // 开始监听服务，监听 TCP 网络地址，Addr 和调用应用程序处理连接上的请求。
}
