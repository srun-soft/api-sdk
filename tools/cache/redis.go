package cache

import (
	"fmt"
	"github.com/srun-soft/api-sdk/configs"
	"time"
)

type RedisCache struct {
}

func (c *RedisCache) GetCache(key string) string {
	cache := configs.Cache
	if cache == nil {
		fmt.Println("cache connection failed")
		return ""
	}
	value, err := cache.Get(cache.Context(), key).Result()
	if err != nil || value == "" {
		return ""
	}
	return value
}

func (c *RedisCache) SetCache(key, value string, expireAt float64) {
	cache := configs.Cache
	if cache != nil {
		_ = cache.Set(cache.Context(), key, value, time.Duration(expireAt)).Err()
	}
}
