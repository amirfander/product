package search

import "context"

var repository RepositoryInterface

func SetRepository(ri RepositoryInterface) RepositoryInterface {
	repository = ri
	return ri
}

type RepositoryInterface interface {
	Create(ctx context.Context, id string, document interface{}, index string)
	Search(index string, search string, result interface{}, limit int, skip int)
}

func Create(ctx context.Context, id string, document interface{}, index string) {
	repository.Create(ctx, id, document, index)
}

func Search(index string, search string, result interface{}, limit int, skip int) {
	repository.Search(index, search, result, limit, skip)
}
