package cache

import (
	"fmt"
	"reflect"
	"time"

	"gopkg.in/redis.v3"
)

type RedisCache struct {
	*redis.Client
}

func NewRedisCache(host, port string, db int64) *RedisCache {
	return &RedisCache{
		Client: redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%s", host, port),
			DB:   db,
		})}
}

func (this *RedisCache) Set(k string, v interface{}, ex time.Duration) error {
	return this.Client.Set(k, v, ex).Err()
}

func (this *RedisCache) Get(k string, def interface{}) interface{} {
	switch reflect.TypeOf(def).String() {
	case "int64":
		if v, err := this.Client.Get(k).Int64(); err == nil {
			return v
		} else {
			return def
		}
	}
	if v, err := this.Client.Get(k).Result(); err == nil {
		return v
	} else {
		return def
	}
}
