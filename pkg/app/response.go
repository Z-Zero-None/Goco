package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"Goco/pkg/errcode"
)


type Response struct {
	Ctx *gin.Context
}


func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

//正常响应
func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, gin.H{
		"data":data,
	})
}

//响应列表
func (r *Response) ToResponseList(list interface{}, totalRows int) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"list": list,
		"pager": Pager{
			Page:      GetPage(r.Ctx),
			PageSize:  GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}

//响应错误
func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(), response)
}
