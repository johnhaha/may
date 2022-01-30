package mayer

import (
	"encoding/json"
	"strings"
)

type IndexSetting struct {
	RankingRules         []string `json:"rankingRules,omitempty"`
	FilterableAttributes []string `json:"filterableAttributes,omitempty"`
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

func SetSearchFilter(filter []string) SettingOption {
	return func(mi *MeiliIndex) {
		mi.Setting.FilterableAttributes = filter
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
	Filter string   `json:"filter,omitempty"`
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

//set search filter, use '?' to be replaced by query
func SetFilter(filter string, query ...interface{}) SearchOption {
	s := strings.Split(filter, "?")
	l := len(s)
	f := s[0]
	if l > 1 {
		if len(query) != l-1 {
			panic("filter query not set properly")
		}
		for i := 1; i < l; i++ {
			j, _ := json.Marshal(query[i-1])
			f += (string(j) + s[i])
		}
	}
	return func(mi *MeiliIndex) {
		mi.SearchBody.Filter = f
	}
}
