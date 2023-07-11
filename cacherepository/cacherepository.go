package cacherepository

import "time"

var repository RepositoryInterface

func SetRepository(ri RepositoryInterface) RepositoryInterface {
	repository = ri
	return ri
}

type RepositoryInterface interface {
	Set(key string, document interface{}, expiration time.Duration)
	Get(key string, result interface{}) error
}

func Set(key string, document interface{}, expiration time.Duration) {
	repository.Set(key, document, expiration)
}

func Get(key string, result interface{}) error {
	return repository.Get(key, result)
}
