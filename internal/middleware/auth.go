package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"Goco/internal/serializer"
	"Goco/internal/service"
	"Goco/pkg/app"
	"Goco/pkg/errcode"
)

// CurrentUser 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		//从cookie获取user_id	查看是否登录
		uid := session.Get("user_id")

		if uid != nil {
			//获取user信息
			service := service.NewService(c.Request.Context())
			user, err := service.GetUser(uid)
			//存在即放入设置全局user
			if err == nil {
				c.Set("user", &user)
			}
		}
		c.Next()
	}
}

// AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if _, ok := user.(*serializer.User); ok {
				c.Next()
				return
			}
		}
		response := app.NewResponse(c)
		response.ToErrorResponse(errcode.ErrUnLogin)
		c.Abort()
	}
}
