package request

type UserLoginRequest struct {
	UserName string `form:"username" json:"username" binding:"required,min=5,max=20"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=20"`
}
type UserRegisterRequest struct {
	UserName        string `form:"username" json:"username" binding:"required,min=5,max=20"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=20"`
	PasswordConfirm string `form:"confirm" json:"confirm" binding:"required,min=8,max=20"`
}
type UserChangePWDRequest struct {
	OldPwd string `json:"old_pwd" binding:"required,min=5,max=20"`
	NewPwd string `json:"new_pwd" binding:"required,min=8,max=20"`
	ComPwd string `json:"com_pwd" binding:"required,min=8,max=20"`
}