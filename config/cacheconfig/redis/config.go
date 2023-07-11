package redis

import (
	"fmt"

	"github.com/redis/go-redis/v9"

	cacherepository "product/cacherepository"
	"product/config/env"
)

type RedisConfig struct {
}

func (rc RedisConfig) ConnectCache(uri string) {
	opt, err := redis.ParseURL(env.EnvRedisURI())
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)
	fmt.Println("Connected to Redis")
	SetClient(client)
	cacherepository.SetRepository(Redis{})
}
