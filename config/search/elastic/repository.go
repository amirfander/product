package elastic

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

var client *elasticsearch.Client

func SetClient(cl *elasticsearch.Client) {
	client = cl
}

type Elastic struct{}

func (e Elastic) Create(ctx context.Context, id string, document interface{}, index string) {
	jsonData, _ := json.Marshal(document)
	if _, error := client.Create(index, id, strings.NewReader(string(jsonData))); error != nil {
		fmt.Println(error)
	}
}

func (e Elastic) Search(index string, search string, fields []string, result []interface{}) {
	query := `{ query:{query_string:{query: "` + search + `" }} }`
	response, _ := client.Search(client.Search.WithIndex(index), client.Search.WithBody(strings.NewReader(query)))
	json.NewDecoder(response.Body).Decode(&result)
}
