package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	code    int      `json:"code"`    //错误码
	msg     string   `json:"msg"`     //错误信息
	details []string `json:"details"` //错误描述
}

//错误码集合
var codes = map[int]string{}

//创建新的错误
func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码%d已经存在,请更换一个", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

//错误
func (e *Error) Error() string {
	return fmt.Sprintf("错误码:%d,错误信息::%s", e.Code(), e.Msg())
}
func (e *Error) Code() int {
	return e.code
}
func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) MsgF(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}
func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	for _, d := range details {
		newError.details = append(newError.details, d)
	}
	return &newError
}

//根据错误获取状态码
func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code(): //成功
		return http.StatusOK
	case ServerError.Code(): //服务错误
		return http.StatusInternalServerError
	case InvalidParams.Code(): //参数错误
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}
	return http.StatusInternalServerError
}
