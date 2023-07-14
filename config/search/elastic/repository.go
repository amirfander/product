package elastic

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
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

func (e Elastic) Search(index string, search string, result interface{}, limit int, skip int) {
	req := esapi.SearchRequest{
		Index: []string{index},
		From:  &skip,
		Size:  &limit,
	}
	res, _ := req.Do(context.Background(), client)
	defer res.Body.Close()
	var results []interface{} = make([]interface{}, 0)
	var response map[string]interface{}
	json.NewDecoder(res.Body).Decode(&response)
	for _, hit := range response["hits"].(map[string]interface{})["hits"].([]interface{}) {
		craft := hit.(map[string]interface{})["_source"].(map[string]interface{})
		results = append(results, craft)
	}
	jsonResult, _ := json.Marshal(results)
	json.Unmarshal(jsonResult, result)
	fmt.Println(results...)
}
