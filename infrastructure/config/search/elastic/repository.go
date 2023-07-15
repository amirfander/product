package elastic

import (
	"bytes"
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
	var buffer bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"query_string": map[string]interface{}{
				"query": "*" + search + "*",
			},
		},
	}
	json.NewEncoder(&buffer).Encode(query)
	req := esapi.SearchRequest{
		Index: []string{index},
		Body:  &buffer,
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
}

func (e Elastic) UpdateById(index string, id string, document interface{}) {
	jsonData, _ := json.Marshal(document)
	if _, err := client.Update(index, id, bytes.NewReader([]byte(fmt.Sprintf(`{"doc":%s}`, jsonData)))); err != nil {
		fmt.Println(err)
	}
}

func (e Elastic) DeleteById(index string, id string) {
	if _, err := client.Delete(index, id); err != nil {
		fmt.Println(err)
	}
}
