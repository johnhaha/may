package may

import (
	"errors"

	"github.com/johnhaha/hakit/hadata"
	"github.com/johnhaha/may/mayer"
)

func UpdateSetting(data interface{}) error {
	if check, index := IndexExist(data); check {
		searchAttr, filterAttr := GetSearchAttribute(data)
		err := searchClient.Index(index).UpdateSetting(mayer.SetSearchableAttribute(searchAttr), mayer.SetSearchFilter(filterAttr))
		if err != nil {
			return err
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
