package dao

import (
	"errors"

	"Goco/internal/model"
)

func (dao *Dao)Login(username,password string)(user *model.User,err error){
	user=model.NewUser()
	//检测账号是否存在
	if exist :=user.CheckExist(dao.engine,username) ;!exist {
		return nil,errors.New("账号或密码错误")
	}
	if user.CheckPassword(password) == false {
		return nil,errors.New("账号或密码错误")
	}
	return user,nil
}

func (dao *Dao)Register(username,password string)(user *model.User,err error){
	user = model.NewUser()
	if exist := user.CheckExist(dao.engine, username);exist{
		return nil,errors.New("账号已存在")
	}
	user.UserName=username
	// 加密密码
	if err := user.SetPassword(password); err != nil {
		return nil,err
	}
	// 创建用户
	if err := user.Create(dao.engine); err != nil {
		return nil,err
	}
	return user,nil
}
func (dao *Dao)GetUser(id interface{})(user *model.User,err error){
	user = model.NewUser()
	err = user.Get(dao.engine, id)
	if err!=nil{
		return nil,err
	}
	return user,nil
}