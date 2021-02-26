package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	extrapolations "github.com/go-playground/validator/v10/translations/en"
	zillions "github.com/go-playground/validator/v10/translations/zh"
)

//国际化处理
func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		//New返回一个新的UniversalTranslator实例集，该实例集具有后备语言环境和应支持的语言环境
		uni := ut.New(en.New(), zh.New(), zh_Hant_TW.New())
		//获取当前context请求头
		locale := c.GetHeader("locale")
		//获取转译语言类型
		trans, _ := uni.GetTranslator(locale)
		//转译器
		v, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			switch locale {
			case "zh":
				//将验证器和对应的语言类型的Translator注册进来，实现支持多语言
				_ = zillions.RegisterDefaultTranslations(v, trans)
				break
			case "en":
				_ = extrapolations.RegisterDefaultTranslations(v, trans)
				break
			default:
				_ = zillions.RegisterDefaultTranslations(v, trans)
				break
			}
			c.Set("trans", trans)
		}
		c.Next()
	}
}
