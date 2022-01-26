package may

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"

	"github.com/johnhaha/hakit/hadata"
	"github.com/meilisearch/meilisearch-go"
)

func getIndexName(name string) string {
	return name + "_idx"
}

func getJsonFieldName(f reflect.StructField) (name string, omitempty bool) {
	if n, ok := f.Tag.Lookup("json"); ok {
		ns := strings.Split(n, ",")
		if len(ns) == 1 {
			return n, false
		}
		return ns[0], true
	}
	return f.Name, false
}

func getManyDocumentAndIndexFromData(data interface{}, includeField ...string) (index string, docs []map[string]interface{}) {
	if reflect.TypeOf(data).Kind() == reflect.Struct {
		index = GetIndex(data)
		docs = append(docs, getDocumentFromData(data, includeField...))

	} else if reflect.TypeOf(data).Kind() == reflect.Slice {
		s := reflect.ValueOf(data)
		for i := 0; i < s.Len(); i++ {
			sData := s.Index(i).Interface()
			if index == "" {
				index = GetIndex(sData)
			}
			docs = append(docs, getDocumentFromData(sData, includeField...))
		}
	}
	return index, docs
}

//get doc map from struct
func getDocumentFromData(data interface{}, includeField ...string) map[string]interface{} {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	doc := make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		d := v.Field(i)
		jsonName, _ := getJsonFieldName(f)
		if d.IsZero() && !hadata.IsInStringSlice(includeField, jsonName) {
			continue
		}
		if _, ok := f.Tag.Lookup("may"); ok {

			doc[jsonName] = d.Interface()
		}
	}
	return doc
}

//data can be struct or slice of struct
func delData(data interface{}) error {
	if reflect.TypeOf(data).Kind() == reflect.Struct {
		err := delDoc(data)
		return err
	}
	if reflect.TypeOf(data).Kind() == reflect.Slice {
		err := delManyDoc(data)
		return err
	}
	return errors.New("can not del data")
}

//data can only be struct
func delManyDoc(data interface{}) error {
	t := reflect.ValueOf(data)
	sData := make([]interface{}, t.Len())
	for i := 0; i < t.Len(); i++ {
		sData[i] = t.Index(i).Interface()
	}
	uid := make([]string, len(sData))
	for i, d := range sData {
		k, err := GetPrimaryKeyFieldValue(d)
		if err != nil {
			return err
		}
		uid[i] = k
	}
	index := GetIndex(data)
	_, err := searchClient.Index(index).DeleteDocuments(uid)
	if err != nil {
		return err
	}
	return nil
}

//data can only be struct
func delDoc(data interface{}) error {
	key, err := GetPrimaryKeyFieldValue(data)
	if err != nil {
		return err
	}
	index := GetIndex(data)
	_, err = searchClient.Index(index).DeleteDocument(key)
	if err != nil {
		return err
	}
	return nil
}

func cleanPointerData(data interface{}) (interface{}, error) {
	t := reflect.ValueOf(data)
	if t.Kind() == reflect.Ptr {
		res, err := hadata.GetPointerData(data)
		return res, err
	}
	return data, nil
}

func decodeSearchRes(res *meilisearch.SearchResponse, data interface{}) error {
	if res.Hits == nil {
		return nil
	}
	d, err := json.Marshal(res.Hits)
	if err != nil {
		return err
	}
	err = json.Unmarshal(d, data)
	return err
}

func getSearchTag(f reflect.StructField) []string {
	if n, ok := f.Tag.Lookup("may"); ok {
		s := strings.Split(n, ",")
		return s
	}
	return nil
}
