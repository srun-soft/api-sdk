package cache

import (
	"srunsoft-api-sdk/configs"
	"time"
)

type RedisCache struct {
}

func (c *RedisCache) GetCache(key string) string {
	cache := configs.Cache
	if cache == nil {
		configs.Log.WithField("Cache", "连接失败").Error()
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
