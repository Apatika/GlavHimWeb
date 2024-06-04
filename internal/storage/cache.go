package storage

import (
	"glavhim-app/internal/storage/heapcache"
)

type ICache interface {
	Get(key string) interface{}
	Update(key string, data interface{})
	Delete(key string)
}

var Cache ICache

func CacheInit() {
	Cache = heapcache.New()
}
