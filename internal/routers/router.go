package routers

import (
	"context"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "Goco/docs"
	"Goco/global"
	"Goco/internal/middleware"
	v1 "Goco/internal/routers/api/v1"
)

func NewRouter()*gin.Engine{
	r:=gin.Default()
	//启动Session用于存储信息
	r.Use(middleware.Session(global.AppSetting.SessionSecret))
	//语言中间件
	r.Use(middleware.Translations())
	//启动swagger查看api文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//测试使用ping接口
	r.GET("/ping",v1.Ping)
	global.Logger.Info(context.Background(),"服务成功启动查看日志")
	return r
}