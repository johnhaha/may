package may

import (
	"errors"

	"github.com/johnhaha/hakit/hadata"
	"github.com/meilisearch/meilisearch-go"
)

func UpdateSetting(data interface{}) error {
	if check, index := IndexExist(data); check {
		searchAbleAttr := GetSearchableAttribute(data)
		if len(searchAbleAttr) != 0 {
			attrs := GetRankedSearchableAttribute(searchAbleAttr)
			_, err := searchClient.Index(index).UpdateSettings(&meilisearch.Settings{
				SearchableAttributes: attrs,
			})
			if err != nil {
				return err
			}
		}
		return nil
	}
	return errors.New("index not exist")
}

type SearchableAttribute struct {
	Name string
	Rank int
}

func (a SearchableAttribute) GetValue() int {
	return a.Rank
}

func GetRankedSearchableAttribute(attr []SearchableAttribute) []string {
	sort := make([]hadata.Sort, len(attr))
	for i, a := range attr {
		sort[i] = a
	}
	res := hadata.QuickSort(sort, 0, len(attr)-1)
	var ot []string
	for _, r := range res {
		ot = append(ot, r.(SearchableAttribute).Name)
	}
	return ot
}
