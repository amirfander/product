package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func SetClient(cl *redis.Client) {
	client = cl
}

type Redis struct{}

func (r Redis) Set(key string, document interface{}, expiration time.Duration) {
	ctx := context.Background()
	jsonDoc, err := json.Marshal(document)
	if err != nil {
		fmt.Println(err)
	}
	err = client.Set(ctx, key, jsonDoc, expiration).Err()
	if err != nil {
		panic(err)
	}
}

func (r Redis) Get(key string, result interface{}) error {
	ctx := context.Background()
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return err
	}
	json.Unmarshal([]byte(val), &result)
	return nil
}
