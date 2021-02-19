package setting

import (
	"fmt"

	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting(config string) (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	//如果有参数
	if config != "" {
		vp.AddConfigPath(config)
	}
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig() //读取配置中心内容
	if err != nil {
		fmt.Printf("配置中心初始化失败，err:%v\n",err)
		return nil, err
	}
	s := &Setting{vp}
	return s, nil
}
//读取文件配置
func (s *Setting) ReadSection(k string, v interface{}) error {
	//前面是key 后面是结构体 将key所对应的信息与结构进行绑定
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {

		fmt.Println("解析k:", k, "v:", v, "失败！")
		return err
	}
	return nil
}