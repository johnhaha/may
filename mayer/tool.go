package mayer

import "strconv"

type Filter struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

func NewFilter(offset int, limit int) *Filter {
	return &Filter{Offset: offset, Limit: limit}
}

func (filter *Filter) GetQuery() map[string]string {
	ot := make(map[string]string)
	if filter.Offset != 0 {
		ot["offset"] = strconv.Itoa(filter.Offset)
	}
	if filter.Limit != 0 {
		ot["limit"] = strconv.Itoa(filter.Limit)
	}
	return ot
}
