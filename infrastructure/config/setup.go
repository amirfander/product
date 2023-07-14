package config

import (
	"product/infrastructure/config/broker"
	"product/infrastructure/config/broker/nats"
	"product/infrastructure/config/cache"
	"product/infrastructure/config/cache/redis"
	"product/infrastructure/config/db"
	"product/infrastructure/config/db/mongo"
	"product/infrastructure/config/env"
	"product/infrastructure/config/search"
	"product/infrastructure/config/search/elastic"
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

func ConnectSearch() {
	switch env.EnvSearchType() {
	case "elastic":
		search.ConnectEngine(elastic.ElasticConfig{}, env.EnvElasticURI())
	default:
		search.ConnectEngine(elastic.ElasticConfig{}, env.EnvElasticURI())
	}
}
