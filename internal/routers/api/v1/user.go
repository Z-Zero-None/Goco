package v1

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"Goco/global"
	"Goco/internal/model"
	"Goco/internal/request"
	"Goco/internal/service"
	"Goco/pkg/app"
	"Goco/pkg/errcode"
)
//校验用户详情
func UserMe(c *gin.Context){
	//创建对应的响应对象
	response := app.NewResponse(c)
	if user, _ := c.Get("user"); user != nil {
		if u, ok := user.(*model.User); ok {
			response.ToResponse(u)
			return
		}
	}
	response.ToResponse(gin.H{})
	return
}
// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	request:=request.UserLoginRequest{}
	service:= service.NewService(c.Request.Context())
	//创建对应的响应对象
	response := app.NewResponse(c)
	//指针绑定参数
	valid,errs := app.BindAndValid(c,&request)
	//绑定不成功处理
	if !valid {
		global.Logger.ErrorF("app.BindAndValid errs:%v", c, errs)
		//封装错误集合
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		//响应错误
		response.ToErrorResponse(errRsp)
		return
	}
	data,err := service.Login(c,&request)
	if err != nil {
		global.Logger.ErrorF("service.Login err:%v", c, err)
		response.ToErrorResponse(errcode.ErrLoginFail)
		return
	}
	response.ToResponse(data)
	return
}
// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	request:=request.UserRegisterRequest{}
	service:= service.NewService(c.Request.Context())
	//创建对应的响应对象
	response := app.NewResponse(c)
	valid,errs := app.BindAndValid(c,&request)
	//绑定不成功处理
	if !valid {
		global.Logger.ErrorF("app.BindAndValid errs:%v", c, errs)
		//封装错误集合
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		//响应错误
		response.ToErrorResponse(errRsp)
		return
	}
	data, err := service.Register(&request)
	if err != nil {
		global.Logger.ErrorF("service.Register err:%v", c, err)
		response.ToErrorResponse(errcode.ErrLoginFail)
		return
	}
	response.ToResponse(data)
	return
}
// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	response := app.NewResponse(c)
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	response.ToResponse(gin.H{
		"msg":"退出成功",
	})
}