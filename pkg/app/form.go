package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

//接口校验
//参数错误结构体
type ValidError struct {
	Key     string
	Message string
}

//参数错误集合
type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

//错误集合的信息集合
func (v ValidErrors) Errors() (errs []string) {
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

//绑定参数
func BindAndValid(c *gin.Context, v interface{}) (b bool, errs ValidErrors) {
	err := c.ShouldBind(v)
	fmt.Printf("%#v\n", v)
	if err != nil {
		v := c.Value("trans")
		trans, _ := v.(ut.Translator)
		verrs, ok := err.(validator.ValidationErrors)
		if !ok {

			return false, errs
		}
		for key, value := range verrs.Translate(trans) {
			errs = append(errs, &ValidError{
				Key:     key,
				Message: value,
			})
		}
		return false, errs
	}
	return true, nil
}
