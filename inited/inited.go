package inited

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"

	"Goco/global"
	"Goco/internal/cache"
	"Goco/internal/model"
	"Goco/pkg/logger"
	"Goco/pkg/setting"
	"Goco/pkg/tracer"
)

func InitGlobal() (err error) {
	err = initSetting()
	if err != nil {
		fmt.Println("配置中心配置失败")
		return err
	}
	err = initLogger()
	if err != nil {
		fmt.Println("配置全局日志失败")
		return err
	}
	err = initDBEngine()
	if err != nil {
		fmt.Println("配置全局数据库失败")
		return err
	}
	err = initRedis()
	if err != nil {
		fmt.Println("配置全局redis对象失败")
		return err
	}
	err = initTracer()
	if err != nil {
		fmt.Println("配置全局tracer对象失败")
		return err
	}
	return nil
}
//设置全局Tracer对象
func initTracer() (err error) {
	jaegerTracer, _, err := tracer.NewJaegerTracer(global.TracerSetting.Name, global.TracerSetting.Host)
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}


//设置全局redis对象
func initRedis() (err error) {
	global.RedisClient, err = cache.NewRedisClient(global.RedisSetting)
	if err != nil {
		return err
	}
	return nil
}

//设置全局DB对象
func initDBEngine() (err error) {
	global.DBEngine, err = model.NewDBEngine(global.DataBaseSetting)
	if err != nil {
		return err
	}
	migration()
	return nil
}

//设置全局日志对象
func initLogger() error {
	format := time.Now().Format("20060102")
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName +format+ global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,  //日志文件最大存储MB
		MaxAge:    10,   //日志文件生命周期
		LocalTime: true, //时间格式为本地时间格式
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}

//设置全局配置
func initSetting() (err error) {
	setting, err := setting.NewSetting("configs/")
	if err != nil {
		fmt.Println("获取配置中心setting失败")
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		fmt.Println("绑定ServerSetting数据失败")
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		fmt.Println("绑定AppSetting数据失败")
		return err
	}
	err = setting.ReadSection("DataBase", &global.DataBaseSetting)
	if err != nil {
		fmt.Println("绑定DataBaseSetting数据失败")
		return err
	}
	err = setting.ReadSection("Redis", &global.RedisSetting)
	if err != nil {
		fmt.Println("绑定RedisSetting数据失败")
		return err
	}
	err = setting.ReadSection("Tracer", &global.TracerSetting)
	if err != nil {
		fmt.Println("绑定TracerSetting数据失败")
		return err
	}

	global.ServerSetting.WriteTimeout *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.RedisSetting.IdleTimeout *= time.Second

	log.Printf("===全局变量ServerSetting:\n%#v\n", global.ServerSetting)
	log.Printf("===全局变量AppSetting:\n%#v\n", global.AppSetting)
	log.Printf("===全局变量DataBaseSetting:\n%#v\n", global.DataBaseSetting)
	log.Printf("===全局变量RedisSetting:\n%#v\n", global.RedisSetting)
	log.Printf("===全局变量TracerSetting):\n%#v\n", global.TracerSetting)
	return nil
}
