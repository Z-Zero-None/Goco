package v1

import (
	"github.com/gin-gonic/gin"

	"Goco/global"
	"Goco/internal/service"
	"Goco/pkg/app"
	"Goco/pkg/errcode"
)

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	service:= service.NewUserLogin()
	//创建对应的响应对象
	response := app.NewResponse(c)
	valid,errs := app.BindAndValid(c,service)
	//绑定不成功处理
	if !valid {
		global.Logger.ErrorF("app.BindAndValid errs:%v", c, errs)
		//封装错误集合
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		//响应错误
		response.ToErrorResponse(errRsp)
		return
	}
	data,err := service.Login(c)
	if err != nil {
		global.Logger.ErrorF("svc.UpdateTag err:%v", c, err)
		response.ToErrorResponse(errcode.ErrLoginFail)
		return
	}
	response.ToResponse(data)
	return
}
