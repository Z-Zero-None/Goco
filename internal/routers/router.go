package routers

import (
	"github.com/gin-gonic/gin"

	v1 "Goco/internal/routers/api/v1"
)

func NewRouter()*gin.Engine{
	r:=gin.Default()
	r.GET("/ping",v1.Ping)
	return r
}