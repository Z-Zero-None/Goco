package setting

import (
	"time"
)

type ServerSetting struct {
	RunMode      string        //运行模式
	HttpPort     string        //http端口
	ReadTimeout  time.Duration //读超时
	WriteTimeout time.Duration //写超时
}
type AppSetting struct {
	DefaultPageSize       int           //默认路径大小
	MaxPageSize           int           //最大页数限制
	LogSavePath           string        //日志保存路径
	LogFileName           string        //日志文件名称
	LogFileExt            string        //日志文件后缀
}
type DataBaseSetting struct {
	DBType       string //数据库类型
	UserName     string //用户名
	Password     string //用户密码
	Host         string //host
	DBName       string //数据库名字
	TablePrefix  string //表名前缀
	Charset      string //字符类型
	ParseTime    bool   //时间
	MaxIdleConnNum int    //最大空闲连接
	MaxOpenConnNum int    //最大活跃连接
}
type RedisSetting struct {
	Address string //redis的地址
	Password string //redis密码
	DB int	//选择数据库
	IdleTimeout time.Duration	//超时时间
}
