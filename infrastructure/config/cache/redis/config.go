package redis

import (
	"fmt"

	"github.com/redis/go-redis/v9"

	"product/repository/cache"
)

type RedisConfig struct {
}

func (rc RedisConfig) ConnectCache(uri string) {
	opt, err := redis.ParseURL(uri)
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)
	fmt.Println("Connected to Redis")
	SetClient(client)
	cache.SetRepository(Redis{})
}
