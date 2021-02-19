package app

import (
	"github.com/gin-gonic/gin"

	"Goco/global"
	"Goco/pkg/convert"
)

//页面信息
type Pager struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

//获取页面
func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page < 0 {
		return 1
	}
	return page
}

//获取页面大小
func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()
	//限制获取页数大小
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}
	return pageSize
}

//获取当前页面下表
func GetPageOffset(page, pageSize int) (result int) {
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return result
}
