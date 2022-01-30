package may

import "github.com/johnhaha/may/mayer"

var searchClient *mayer.MeiliClient

func InitClient(host string, apiKey string) {
	client := &mayer.MeiliClient{
		Host:   host,
		APIKey: apiKey,
	}
	searchClient = client
}

func GetSearchClient() *mayer.MeiliClient {
	return searchClient
}
