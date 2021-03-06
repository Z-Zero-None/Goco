package model

import (
	"fmt"

	otgorm "github.com/eddycjy/opentracing-gorm"

	"Goco/global"
	"Goco/pkg/setting"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func NewDBEngine(s *setting.DataBaseSetting) (*gorm.DB, error) {
	conStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local", s.UserName, s.Password, s.Host, s.DBName, s.Charset, s.ParseTime)
	db, err := gorm.Open(s.DBType, conStr)
	if err != nil {
		fmt.Println("连接数据库失败,conStr:", conStr)
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	//单表模式
	db.SingularTable(true)

	//数据库属性设置
	db.DB().SetMaxIdleConns(s.MaxIdleConnNum)
	db.DB().SetMaxOpenConns(s.MaxOpenConnNum)
	otgorm.AddGormCallbacks(db) //添加回调函数获取sql链路追踪
	return db, nil
}

