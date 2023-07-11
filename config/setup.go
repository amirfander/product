package config

import (
	"product/config/cacheconfig"
	"product/config/cacheconfig/redis"
	"product/config/dbconfig"
	"product/config/dbconfig/mongo"
	"product/config/env"
)

func ConnectDB() {
	switch env.EnvDBType() {
	case "mongodb":
		dbconfig.ConnectDB(mongo.MongoDBConfig{}, env.EnvMongoURI())
	default:
		dbconfig.ConnectDB(mongo.MongoDBConfig{}, env.EnvMongoURI())
	}
}

func ConnectCache() {
	switch env.EnvCacheType() {
	case "redis":
		cacheconfig.ConnectCache(redis.RedisConfig{}, env.EnvRedisURI())
	default:
		cacheconfig.ConnectCache(redis.RedisConfig{}, env.EnvRedisURI())
	}
}
