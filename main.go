package main

import (
	"context"
	"fmt"
	format "github.com/antonfisher/nested-logrus-formatter"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"srunsoft-api-sdk/configs"
	"srunsoft-api-sdk/sdk"
	"sync"
	"time"
)

var (
	wg  sync.WaitGroup
	ctx = context.Background()
)

func main() {
	// 初始化日志
	log := logrus.New()

	log.SetFormatter(&format.Formatter{
		HideKeys:        false,
		TimestampFormat: time.RFC3339,
		FieldsOrder:     []string{"component", "category"},
	})

	log.WithField("Log组件", "加载成功").Info()

	configs.Config = &configs.APIConfig{
		Scheme:      "http://",
		InterfaceIP: "127.0.0.1:8001",
		AppId:       "srunsoft",
		AppSecret:   "ba62f23e6212790052f387ee7af2943e4cfcece0",
		LogDir:      "runtime/",
	}
	configs.Log = log

	// redis-cache
	wg.Add(1)
	go func() {
		defer wg.Done()
		configs.Cache = createClient("192.168.0.100", "16384")
	}()
	wg.Wait()

	temp := &sdk.APIClient{}
	token, err := temp.GetToken()
	if err != nil {
		configs.Log.WithField("token failed", err).Error()
	}
	configs.Log.WithField("Token is ", token).Info()
}

// 创建 Redis 连接
func createClient(host, port string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", host, port),
		Password:     "srun_3000@redis",
		DB:           0,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     100,
		PoolTimeout:  30 * time.Second,
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		configs.Log.WithField(fmt.Sprintf("Redis[%s] init err", port), err).Warn()
		return nil
	}
	configs.Log.WithField(fmt.Sprintf("Redis[:%s] init", port), "Successful").Info()
	return rdb
}
