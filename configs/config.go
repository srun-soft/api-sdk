package configs

import (
	"github.com/go-redis/redis/v8"
)

// APIConfig 接口配置
type APIConfig struct {
	Scheme      string `comment:"接口协议"`
	InterfaceIP string `comment:"接口地址"`
	AppId       string `comment:"appid"`
	AppSecret   string `comment:"appSecret"`
}

var (
	Config *APIConfig
	Cache  *redis.Client
)
