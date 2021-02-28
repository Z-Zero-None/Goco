package service

import (
	"errors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"Goco/internal/model"
	"Goco/internal/request"
	"Goco/internal/serializer"
)

// setSession 设置session
func (service *Service) setSession(c *gin.Context, user *model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.ID)
	s.Save()
}
// Login 用户登录函数
func (service *Service) Login(c *gin.Context,request *request.UserLoginRequest)(data serializer.User, err error){
	user:=model.NewUser()
	if exist :=user.CheckExist(service.engine,request.UserName) ;!exist {
		return data,errors.New("账号或密码错误")
	}
	if user.CheckPassword(request.Password) == false {
		return data,errors.New("账号或密码错误")
	}
	service.setSession(c, user)
	data = serializer.BuildUser(user)
	return data,nil
}

func (service *Service) Register(request *request.UserRegisterRequest)(data serializer.User,err error) {
	user := model.NewUser()
	//校验两次密码是否正确
	if request.PasswordConfirm != request.Password {
		return data,errors.New("两次输入的密码不相同")
	}
	if exist := user.CheckExist(service.engine, request.UserName);exist{
		return data,errors.New("账号已存在")
	}
	user.UserName=request.UserName
	// 加密密码
	if err := user.SetPassword(request.Password); err != nil {
		return data,err
	}
	// 创建用户
	if err := user.Create(service.engine); err != nil {
		return data,err
	}
	data=serializer.BuildUser(user)
	return data,nil
}
func (service *Service)GetUser(id interface{})(data serializer.User,err error){
	user := model.NewUser()
	err = user.Get(service.engine, id)
	if err!=nil{
		return data,err
	}
	return serializer.BuildUser(user),nil
}
