package cache

// cache

type Cache interface {
	GetCache(key string) string
	SetCache(key, value string, expireAt float64)
}
