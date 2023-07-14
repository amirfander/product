package db

import (
	"context"
	"time"

	"product/repository/cache"
)

type RepositoryInterface interface {
	InsertOne(ctx context.Context, document interface{}, collectionName string) (string, error)
	FindById(ctx context.Context, id string, collectionName string, result interface{}) error
	Find(ctx context.Context, collectionName string, filter interface{}, skip int, limit int, result interface{}) error
	UpdateById(ctx context.Context, collectionName string, id string, document interface{}) error
	DeleteById(ctx context.Context, collectionName string, id string) error
}

var repository RepositoryInterface

func InsertOne(ctx context.Context, document interface{}, collectionName string) (string, error) {
	return repository.InsertOne(ctx, document, collectionName)
}

func FindById(ctx context.Context, id string, collectionName string, result interface{}) error {
	cacheKey := collectionName + "-" + "FindById-" + id

	if err := cache.Get(cacheKey, result); err == nil {
		return nil
	}

	res := repository.FindById(ctx, id, collectionName, result)
	if res == nil {
		cache.Set(cacheKey, result, time.Minute)
	}
	return res
}

func Find(ctx context.Context, collectionName string, filter interface{}, skip int, limit int, result interface{}) error {
	return repository.Find(ctx, collectionName, filter, skip, limit, result)
}

func UpdateById(ctx context.Context, collectionName string, id string, document interface{}) error {
	return repository.UpdateById(ctx, collectionName, id, document)
}

func DeleteById(ctx context.Context, collectionName string, id string) error {
	return repository.DeleteById(ctx, collectionName, id)
}

func SetRepository(ri RepositoryInterface) RepositoryInterface {
	repository = ri
	return ri
}
