package elastic

import (
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v8"

	"product/search"
)

type ElasticConfig struct {
}

func (ec ElasticConfig) ConnectEngine(uri string) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			uri,
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		fmt.Println(err)
		log.Fatalf("Error creating the client: %s", err)
	}
	fmt.Println("Connected to Elastic")
	SetClient(es)
	search.SetRepository(Elastic{})
}
