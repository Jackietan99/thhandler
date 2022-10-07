package main

import (
	"net/http"
	"os"
	"thHandler/internal/config"
	"thHandler/internal/handler"
	"thHandler/internal/znet"
)

func main() {

	//需要一个全局变量存放 一个TCP服务所有已链接的客户端数据(最少需要2个数据 链接ID serverid)
	//如果参数等于1启动 一个WEBAPI服务端口8888(用于接收http get请求) 一个TCP服务端口1111(接受TCP请求客户端) WEB服务会与TCP服务产生互动
	args := os.Args[1]

	config.Init()

	znet.NewServer()

	if args == "2" {
		srv := &http.Server{
			Addr:    ":" + config.Conf.HttpPort,
			Handler: handler.Init(),
		}
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}
	select {}
}
