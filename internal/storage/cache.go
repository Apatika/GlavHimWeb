package storage

import (
	"glavhim-app/internal/storage/heapcache"
)

type ICache interface {
	Get(string) map[string]interface{}
	Update(string, string, interface{})
	Delete(string, string)
}

var Cache ICache

func CacheInit() {
	Cache = heapcache.New()
}
