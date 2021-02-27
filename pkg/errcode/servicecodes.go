package errcode
//user相关错误
var (
	ErrLoginFail  = NewError(200100, "用户登录失败")
)

