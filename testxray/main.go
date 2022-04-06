package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/opensearch-project/opensearch-go"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
	"log"
	"net/http"
	"strings"
	"time"
)

type aliasResponse struct {
	Actions []map[string]interface{} `json:"actions"`
}

func main() {
	client, err := connect()
	if err != nil {
		fmt.Println("cannot initialize", err)
		panic(err)
	}
	create(client)
	//currentIndex := getAlias("pschatbot-current", client)
	//updateAlias(currentIndex, "newindex", "pschatbot-current", client)
}

func connect() (*opensearch.Client, error) {
	// 这个方法会自动用OPENSEARCH_URL环境变量中的值，所以Addresses: []string{}就行
	client, err := opensearch.NewClient(opensearch.Config{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Addresses: []string{"https://search-supportchatbot-dev-xyxtx5r23imc3xpnb3dx5eevly.cn-north-1.es.amazonaws.com.cn"},
		Username:  "supportchatbot",
		Password:  "eNnkXYH+T8KgKVKAwLYS9xPzw9v0TuPfkm/mz7xWE0c=",
	})
	return client, err
}

func getAlias(alias string, client *opensearch.Client) string {
	var index string
	a := opensearchapi.IndicesGetAliasRequest{Index: []string{alias}}
	b, err := a.Do(context.Background(), client)
	if err != nil {
		fmt.Println("报错啦", err)
		panic(err)
	}
	defer b.Body.Close()
	var r map[string]interface{}
	if err := json.NewDecoder(b.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	for key, _ := range r {
		index = key
	}
	return index
}

func updateAlias(ri string, ai string, alias string, client *opensearch.Client) {
	removeIndex := map[string]string{
		"index": ri,
		"alias": alias,
	}
	addIndex := map[string]string{
		"index": ai,
		"alias": alias,
	}
	actionRemove := map[string]interface{}{"remove": removeIndex}
	actionAdd := map[string]interface{}{"add": addIndex}
	actions := []interface{}{actionRemove, actionAdd}
	bodyMap := map[string]interface{}{"actions": actions}
	jsonStr, err := json.Marshal(bodyMap)
	if err != nil {
		fmt.Println("报错啦", err)
		panic(err)
	}
	body := strings.NewReader(string(jsonStr))
	a := opensearchapi.IndicesUpdateAliasesRequest{Body: body}
	b, err := a.Do(context.Background(), client)
	if err != nil {
		fmt.Println("报错啦", err)
		panic(err)
	}
	defer b.Body.Close()
	fmt.Println(b)
}

func create(client *opensearch.Client) string {
	indexName := "knowledge-center-" + time.Now().Format("2006-01-02-15-04-05")
	a := opensearchapi.IndicesCreateRequest{Index: indexName}
	b, err := a.Do(context.Background(), client)
	if err != nil {
		fmt.Println("报错啦", err)
		panic(err)
	}
	defer b.Body.Close()
	var r map[string]interface{}
	if err := json.NewDecoder(b.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	return r["index"].(string)
}

func createIndex(client *opensearch.Client) {

	//	body := strings.NewReader(`{
	//  "actions": [
	//    {
	//      "remove": {
	//        "index": "newindex",
	//        "alias": "pschatbot-current"
	//      }
	//    },
	//    {
	//      "add": {
	//        "index": "testindex",
	//        "alias": "pschatbot-current"
	//      }
	//    }
	//  ]
	//}`)
	//	a := opensearchapi.IndicesUpdateAliasesRequest{Body: body}
	//	b, err := a.Do(context.Background(), client)
	//	if err != nil {
	//		fmt.Println("报错啦", err)
	//		panic(err)
	//	}
	//	fmt.Println(b)
}
