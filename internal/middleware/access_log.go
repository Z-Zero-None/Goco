package middleware

import (
	"bytes"
	"time"

	"github.com/gin-gonic/gin"

	"Goco/global"
	"Goco/pkg/logger"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

//写入操作
func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

//中间件日志操作
func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		//中间件前
		bodyWriter := &AccessLogWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = bodyWriter
		now := time.Now().Unix()
		c.Next()
		//中间件后
		end := time.Now().Unix()
		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}
		s := "access log:method:%s,status_code:%d" + ",begin_time:%d,end_time:%d"
		global.Logger.WithFields(fields).InfoF(s, c, c.Request.Method, bodyWriter.Status(), now, end)
	}
}
