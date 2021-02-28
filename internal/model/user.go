package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	UserName string
	Password string
	AuthKey  string
}
const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
)

//创建用户
func NewUser() *User {
	return &User{}
}
//通过该方法变化生成表的表名
func (u User) TableName() string {
	return "goco_user"
}
//设置密码
func (u *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}
//校验密码
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
//检测账号是否存在
func (u *User)CheckExist(db *gorm.DB,username string)bool{
	err:=db.Where("user_name = ?", username).First(&u).Error
	if err!=nil{
		return false
	}
	return true
}
//增删改查列表
func (u *User) Get(db *gorm.DB,id interface{}) error {
	return db.Find(&u, id).Error
}
func (u *User) Create(db *gorm.DB) error {
	return db.Create(&u).Error
}
func (u *User) Update(db *gorm.DB) error {
	return db.Save(&u).Error
}
func (u *User) Delete(db *gorm.DB,id string) error {
	return db.Delete(&u, id).Error
}
func (u *User) List(db *gorm.DB,offset, size int) (list []*User, err error) {
	if offset >= 0 && size > 0 {
		db = db.Offset(offset).Limit(size)
	}
	if err = db.Where("deleted_at is null").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
func (u *User) Count(db *gorm.DB,) (count int, err error) {
	err = db.Model(&u).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
