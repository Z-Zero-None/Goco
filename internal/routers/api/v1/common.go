package v1

import (
	"github.com/gin-gonic/gin"

	"Goco/pkg/app"
)

// @Summary	用于测试服务是否正常启动
// @Produce	json
// @Param	name query string false "案例" maxlength(100)
// @Success 200 {string} pong
// @Failure	400 {string} none
// @Router /ping	[get]
func Ping(c *gin.Context) {
	resp := app.NewResponse(c)
	resp.ToResponse(gin.H{
		"data": "pong",
	})
}
