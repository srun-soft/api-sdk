package cache

import (
	"sync"
	"time"
)

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
