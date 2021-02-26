package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"Goco/global"
	"Goco/inited"
	"Goco/internal/routers"
)

func init(){
	err := inited.InitGlobal()
	if err != nil {
		log.Fatalf("全局变量配置失败")
	}
	log.Println("你的服务配置已初始化完成！")
}

// @title go开发脚手架
// @version 1.0
// @description 如何使用swagger生成api文档
func main(){
	gin.SetMode(global.ServerSetting.RunMode)
	router:=routers.NewRouter()
	s:=&http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}