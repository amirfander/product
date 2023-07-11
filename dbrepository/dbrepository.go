package dbrepository

import (
	"context"
	"time"

	"product/cacherepository"
)

type RepositoryInterface interface {
	InsertOne(ctx context.Context, document interface{}, collectionName string) (bool, error)
	FindById(ctx context.Context, id string, collectionName string, result interface{}) error
}

var repository RepositoryInterface

func InsertOne(ctx context.Context, document interface{}, collectionName string) (bool, error) {
	return repository.InsertOne(ctx, document, collectionName)
}

func FindById(ctx context.Context, id string, collectionName string, result interface{}) error {
	cacheKey := collectionName + "-" + "FindById-" + id

	if err := cacherepository.Get(cacheKey, result); err == nil {
		return nil
	}

	res := repository.FindById(ctx, id, collectionName, result)
	if res == nil {
		cacherepository.Set(cacheKey, result, 60*time.Millisecond)
	}
	return res
}

func SetRepository(ri RepositoryInterface) RepositoryInterface {
	repository = ri
	return ri
}
