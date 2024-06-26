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
func (service *Service) Login(c *gin.Context, request *request.UserLoginRequest) (data serializer.User, err error) {
	user, err := service.dao.Login(request.UserName, request.Password)
	if err != nil {
		return data, err
	}
	service.setSession(c, user)
	data = serializer.BuildUser(user)
	return data, nil
}

func (service *Service) Register(request *request.UserRegisterRequest) (data serializer.User, err error) {

	//校验两次密码是否正确
	if request.PasswordConfirm != request.Password {
		return data, errors.New("两次输入的密码不相同")
	}
	user, err := service.dao.Register(request.UserName, request.Password)
	if err != nil {
		return data, err
	}
	data = serializer.BuildUser(user)
	return data, nil
}
func (service *Service) GetUser(id interface{}) (data serializer.User, err error) {
	user, err := service.dao.GetUser(id)
	if err != nil {
		return data, err
	}
	return serializer.BuildUser(user), nil
}
