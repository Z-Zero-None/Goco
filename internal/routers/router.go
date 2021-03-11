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
	r:=gin.New()
	//启动双写日志
	r.Use(middleware.AccessLog())
	//默认的报错处理
	r.Use(gin.Recovery())
	//启动Session用于存储信息
	r.Use(middleware.Session(global.AppSetting.SessionSecret))
	//语言中间件
	r.Use(middleware.Translations())
	//获取当前用户
	r.Use(middleware.CurrentUser())
	//启动链路追踪
	r.Use(middleware.Tracing())
	//启动swagger查看api文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//测试使用ping接口
	r.GET("/ping",v1.Ping)
	v:=r.Group("/api/v1")
	{
		// 用户登录
		v.POST("user/register", v1.UserRegister)
		// 用户登录
		v.POST("user/login", v1.UserLogin)
		v.Use(middleware.AuthRequired())
		{
			v.GET("user/me", v1.UserMe)
			v.DELETE("user/logout", v1.UserLogout)
		}
	}
	global.Logger.Info(context.Background(),"服务成功启动查看日志")
	return r
}