package service

import (
	"errors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"Goco/internal/model"
	"Goco/internal/serializer"
)

type UserLogin struct {
	UserName string `form:"username" json:"username" binding:"required,min=5,max=20"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=20"`
}
func NewUserLogin() *UserLogin {
	return &UserLogin{}
}
// setSession 设置session
func (service *UserLogin) setSession(c *gin.Context, user *model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.ID)
	s.Save()
}
// Login 用户登录函数
func (service *UserLogin) Login(c *gin.Context)(data serializer.User, err error){
	user:=model.NewUser()
	if err = model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return data,errors.New("账号或密码错误")
	}
	if user.CheckPassword(service.Password) == false {
		return data,errors.New("账号或密码错误")
	}
	service.setSession(c, user)
	data = serializer.BuildUser(*user)
	return data,nil
}

type UserRegister struct {
	UserName        string `form:"username" json:"username" binding:"required,min=5,max=20"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=20"`
	PasswordConfirm string `form:"confirm" json:"confirm" binding:"required,min=8,max=20"`
}
type UserChangePWD struct {
	OldPwd string `json:"old_pwd"`
	NewPwd string `json:"new_pwd"`
	ComPwd string `json:"com_pwd"`
}