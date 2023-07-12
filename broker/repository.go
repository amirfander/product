package broker

var repository RepositoryInterface

func SetRepository(ri RepositoryInterface) RepositoryInterface {
	repository = ri
	return ri
}

type RepositoryInterface interface {
}
