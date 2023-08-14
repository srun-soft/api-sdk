package cache

import "github.com/srun-soft/api-sdk/configs"

// cache

type Cache interface {
	GetCache(key string) string
	SetCache(key, value string, expireAt float64)
}

// Component 组件,返回cache组件接口
func Component() Cache {
	if configs.Cache != nil {
		return &RedisCache{}
	} else {
		return &SyncMapCache{}
	}
}
