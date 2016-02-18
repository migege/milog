package cache

import (
	"time"

	"github.com/astaxie/beego"
)

var (
	cache Cache
)

func CacheInst() Cache {
	return cache
}

type Cache interface {
	Set(k string, v interface{}, ex time.Duration) error
	Get(k string, def interface{}) interface{}
}

func init() {
	cacheengine := beego.AppConfig.String("cacheengine")
	cachehost := beego.AppConfig.String("cachehost")
	cacheport := beego.AppConfig.String("cacheport")
	cachedb, _ := beego.AppConfig.Int64("cachedb")

	switch cacheengine {
	case "redis":
		cache = NewRedisCache(cachehost, cacheport, cachedb)
	}
}
