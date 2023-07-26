package configs

import (
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type APIConfig struct {
	Scheme      string
	InterfaceIP string
	AppId       string
	AppSecret   string
	LogDir      string
}

var (
	Config *APIConfig
	Log    *logrus.Logger
	Cache  *redis.Client
)
