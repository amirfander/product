package search

type SearchConfiger interface {
	ConnectEngine(uri string)
}

func ConnectEngine(searchConfig SearchConfiger, uri string) {
	searchConfig.ConnectEngine(uri)
}
