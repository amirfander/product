package cacheconfig

type CacheConfiger interface {
	ConnectCache(uri string)
}

func ConnectCache(cacheConfig CacheConfiger, uri string) {
	cacheConfig.ConnectCache(uri)
}
