package inited

import (
	"Goco/global"
	"Goco/internal/model"
)

func migration(){
	global.DBEngine.AutoMigrate(&model.User{})
}