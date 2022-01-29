package may

import (
	"errors"

	"github.com/johnhaha/hakit/hadata"
)

func UpdateSetting(data interface{}) error {
	if check, index := IndexExist(data); check {
		searchAbleAttr := GetSearchableAttribute(data)
		if len(searchAbleAttr) != 0 {
			attrs := GetRankedSearchableAttribute(searchAbleAttr)
			err := searchClient.Index(index).UpdateSetting(SetSearchableAttribute(attrs))
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

type IndexSetting struct {
	RankingRules         []string `json:"rankingRules,omitempty"`
	DistinctAttribute    string   `json:"distinctAttribute,omitempty"`
	SearchableAttributes []string `json:"searchableAttributes,omitempty"`
	DisplayedAttributes  []string `json:"displayedAttributes,omitempty"`
	StopWords            []string `json:"stopWords,omitempty"`
	SortableAttributes   []string `json:"sortableAttributes,omitempty"`
	Synonyms             Synonyms `json:"synonyms,omitempty"`
}

type Synonyms struct {
	Wolverine []string `json:"wolverine,omitempty"`
	Logan     []string `json:"logan,omitempty"`
}

type SettingOption func(*MeiliIndex)
type FilterOption func(*MeiliIndex)
type SearchOption func(*MeiliIndex)

func SetSearchableAttribute(attr []string) SettingOption {
	return func(mi *MeiliIndex) {
		mi.Setting.SearchableAttributes = attr
	}
}

func SetRankingRules(attr []string) SettingOption {
	return func(mi *MeiliIndex) {
		mi.Setting.RankingRules = attr
	}
}

func SetOffset(offset int) FilterOption {
	return func(mi *MeiliIndex) {
		mi.Filter.Offset = offset
	}
}

func SetLimit(limit int) FilterOption {
	return func(mi *MeiliIndex) {
		mi.Filter.Limit = limit
	}
}

type SearchBody struct {
	Query  string   `json:"q"`
	Offset int      `json:"offset,omitempty"`
	Limit  int      `json:"limit,omitempty"`
	Sort   []string `json:"sort,omitempty"`
}

func SetSearchOffset(offset int) SearchOption {
	return func(mi *MeiliIndex) {
		mi.SearchBody.Offset = offset
	}
}

func SetSearchLimit(limit int) SearchOption {
	return func(mi *MeiliIndex) {
		mi.SearchBody.Limit = limit
	}
}

func SetSearchSort(sort []string) SearchOption {
	return func(mi *MeiliIndex) {
		mi.SearchBody.Sort = sort
	}
}
