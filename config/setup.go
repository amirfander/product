package config

import (
	"product/config/broker"
	"product/config/broker/nats"
	"product/config/cache"
	"product/config/cache/redis"
	"product/config/db"
	"product/config/db/mongo"
	"product/config/env"
)

func ConnectDB() {
	switch env.EnvDBType() {
	case "mongodb":
		db.ConnectDB(mongo.MongoDBConfig{}, env.EnvMongoURI())
	default:
		db.ConnectDB(mongo.MongoDBConfig{}, env.EnvMongoURI())
	}
}

func ConnectCache() {
	switch env.EnvCacheType() {
	case "redis":
		cache.ConnectCache(redis.RedisConfig{}, env.EnvRedisURI())
	default:
		cache.ConnectCache(redis.RedisConfig{}, env.EnvRedisURI())
	}
}

func ConnectBroker() {
	switch env.EnvBrokerType() {
	case "nats":
		broker.ConnectBroker(nats.NatsConfig{}, env.EnvNatsURI())
	default:
		broker.ConnectBroker(nats.NatsConfig{}, env.EnvNatsURI())
	}
}
