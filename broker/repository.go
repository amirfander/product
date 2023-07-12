package broker

var repository RepositoryInterface

func SetRepository(ri RepositoryInterface) RepositoryInterface {
	repository = ri
	return ri
}

type RepositoryInterface interface {
	Publish(subject string, data []byte)
}

func Publish(subject string, data []byte) {
	repository.Publish(subject, data)
}
