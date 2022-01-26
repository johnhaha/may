package may

import (
	"github.com/meilisearch/meilisearch-go"
)

type Finder struct {
	Limit  int64
	Offset int64
}

func NewFinder() *Finder {
	return &Finder{}
}

func (f *Finder) SetLimit(l int64) *Finder {
	f.Limit = l
	return f
}

func (f *Finder) SetOffset(o int64) *Finder {
	f.Offset = o
	return f
}

func (f *Finder) DocumentsRequest() *meilisearch.DocumentsRequest {
	return &meilisearch.DocumentsRequest{
		Offset: f.Offset,
		Limit:  f.Limit,
	}
}

func (f *Finder) SearchRequest() *meilisearch.SearchRequest {
	return &meilisearch.SearchRequest{
		Offset: f.Offset,
		Limit:  f.Limit,
	}
}

func (f *Finder) GetDocuments(data interface{}) error {
	index := GetIndex(data)
	err := searchClient.Index(index).GetDocuments(f.DocumentsRequest(), data)
	return err
}

func (f *Finder) Search(query string, data interface{}) error {
	index := GetIndex(data)
	res, err := searchClient.Index(index).Search(query, f.SearchRequest())
	if err != nil {
		return err
	}
	err = decodeSearchRes(res, data)
	return err
}
