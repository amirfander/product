package db

type DBConfiger interface {
	ConnectDB(uri string)
}

func ConnectDB(dbConfig DBConfiger, uri string) {
	dbConfig.ConnectDB(uri)
}
