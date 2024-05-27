package storage

import (
	"glavhim-app/internal/storage/heapcache"
)

type ICache interface {
	Get(field string) map[string]interface{}
	Update(field, key string, data interface{})
	Delete(field, key string)
}

var Cache ICache

func CacheInit() {
	Cache = heapcache.New()
}
