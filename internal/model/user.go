package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserName string
	Password string
}
func (u User)TableName() string{
	return "goco_user"
}
