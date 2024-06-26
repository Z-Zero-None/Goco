package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	"Goco/global"
	"Goco/internal/model"
	"Goco/pkg/setting"
)

var DB *gorm.DB

func init() {
	setting, err := setting.NewSetting("configs/")
	if err != nil {
		log.Fatal("获取配置中心setting失败")
	}
	err = setting.ReadSection("DataBase", &global.DataBaseSetting)
	if err != nil {
		log.Fatal("绑定DataBaseSetting数据失败")

	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		log.Fatal("绑定ServerSetting数据失败")

	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		log.Fatal("绑定AppSetting数据失败")

	}
	fmt.Println(global.DataBaseSetting)
	db, err := model.NewDBEngine(global.DataBaseSetting)
	if err != nil {
		log.Fatal("初始化DB失败")
	}
	DB = db
}

func main() {
	var count int
	var zzn string
	//会出现报错，在获取值时出现问题
	zn := map[string]string{
		"user_name":  zzn,
		"created_at": "2021-03-08",
	}
	for k, v := range zn {
		fmt.Println(k, v)
	}

	//当使用struct查询时，GORM将只查询那些具有值的字段
	where := model.User{
		UserName: zzn,
	}
	err := DB.Table("goco_user").Where(where).Count(&count).Error
	if err != nil {
		fmt.Println("DB 查询出现错误", err)
		return
	}
	fmt.Println(count)
}
