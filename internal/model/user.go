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
//增删改查列表
func (u *User) Get(id string) error {
	return DB.Find(&u, id).Error
}
func (u *User) Create() error {
	return DB.Create(&u).Error
}
func (u *User) Update() error {
	return DB.Save(&u).Error
}
func (u *User) Delete(id string) error {
	return DB.Delete(&u, id).Error
}
func (u *User) List(offset, size int) (list []*User, err error) {
	db:=DB
	if offset >= 0 && size > 0 {
		db = db.Offset(offset).Limit(size)
	}
	if err = db.Where("deleted_at is null").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
func (u *User) Count() (count int, err error) {
	err = DB.Model(&u).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
