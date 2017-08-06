package loader

import (
	"time"

	cache "github.com/patrickmn/go-cache"
)

//MainCache to stores database records
var MainCache *cache.Cache

//ReadFromCache reads from cachestore using the key
func ReadFromCache(key string) (interface{}, bool) {
	return MainCache.Get(key)
}

//WriteToCache writes to cachestore (file) based on key.
func WriteToCache(key string, obj interface{}, expiration time.Duration) {
	MainCache.Set(key, obj, expiration)
}
