package errcode
//user相关错误
var (
	ErrLoginFail  = NewError(200100, "用户登录失败")
	ErrUnLogin  = NewError(200101, "用户未登录")
)

