package example

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"srunsoft-api-sdk/configs"
	"srunsoft-api-sdk/sdk"
	"sync"
	"time"
)

var (
	wg  sync.WaitGroup
	ctx = context.Background()
	API *sdk.APIClient
)

func main() {
	configs.Config = &configs.APIConfig{
		Scheme:      "https://",
		InterfaceIP: "127.0.0.1:8001",
		AppId:       "srunsoft",
		AppSecret:   "ba62f23e6212790052f387ee7af2943e4cfcece0",
	}

	// ... 初始化其他配置 ...

	// 配置 redis,如果需要将access_token缓存在redis,则将Cache改为你的redis连接,否则默认使用SyncMap
	// example:

	// ... redis connection start ...

	wg.Add(1)
	go func() {
		defer wg.Done()
		configs.Cache = createClient("192.168.0.100", "16384")
	}()
	wg.Wait()

	// ... redis connection end ...

	// 创建 API 客户端
	API = &sdk.APIClient{}

	// 调用示例方法
	createUserExample()
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
		_ = fmt.Errorf("redis[%s] init err", port)
		return nil
	}
	fmt.Printf("Redis[:%s] init Successful", port)
	return rdb
}

// 开户示例
func createUserExample() {
	data := make(map[string]interface{})
	data["user_name"] = "yt"
	data["group_id"] = 2
	data["user_password"] = "12345678"
	data["products_id"] = 1
	response, err := API.Users(data)
	if err != nil {
		return
	}
	fmt.Println(response)
}
