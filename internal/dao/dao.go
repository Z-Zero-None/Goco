package dao

import "github.com/jinzhu/gorm"

//dao模块
type Dao struct {
	engine *gorm.DB
}

func NewDao(engine *gorm.DB) *Dao {
	return &Dao{
		engine: engine,
	}
}
