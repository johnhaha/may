package may

import "github.com/meilisearch/meilisearch-go"

var searchClient *meilisearch.Client

func InitClient(host string, apiKey string) {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   host,
		APIKey: apiKey,
	})
	searchClient = client
}

func Client() *meilisearch.Client {
	if searchClient == nil {
		panic("search client not init")
	}
	return searchClient
}
