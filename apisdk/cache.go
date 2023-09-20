package apisdk

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

type Cache interface {
	GetCache(key string) string
	SetCache(key, value string, expireAt float64)
}

// Component 组件,返回cache组件接口
func Component(ac *ApiConfig) Cache {
	if ac.Cache != nil {
		return &RedisCache{
			Cache: ac.Cache,
		}
	} else {
		return &SyncMapCache{}
	}
}

type SyncMapCache struct {
	sync.Map
	sync.RWMutex
}

// GetCache 从 sync.Map 缓存中获取值
func (c *SyncMapCache) GetCache(key string) string {
	c.RLock()
	defer c.RUnlock()

	value, found := c.Load(key)
	if !found {
		return ""
	}

	if strValue, ok := value.(string); ok {
		return strValue
	}
	return ""
}

// SetCache 向 sync.Map 缓存中设置值
func (c *SyncMapCache) SetCache(key, value string, expireAt float64) {
	c.Lock()
	defer c.Unlock()

	c.Store(key, value)

	if expireAt > 0 {
		time.AfterFunc(time.Duration(expireAt)*time.Second, func() {
			c.Lock()
			defer c.Unlock()
			c.Delete(key)
		})
	}
}

type RedisCache struct {
	Cache *redis.Client
}

func (c *RedisCache) GetCache(key string) string {
	cache := c.Cache
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
	cache := c.Cache
	if cache != nil {
		cache.Set(cache.Context(), key, value, time.Duration(expireAt)*time.Second)
	}
}
