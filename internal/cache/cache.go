package cache

import (
	"fmt"

	"github.com/go-redis/redis"

	"Goco/pkg/setting"
)

func NewRedisClient(s *setting.RedisSetting)( *redis.Client,error){
	options:=&redis.Options{
		Addr: s.Address,
		Password: s.Password,
		DB: s.DB,
		IdleTimeout: s.IdleTimeout,
	}
	client := redis.NewClient(options)
	_, err := client.Ping().Result()
	if err !=nil{
		fmt.Println("连接redis失败")
		return nil,err
	}
	return client,nil
}