package may

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/johnhaha/hakit/hadata"
	"github.com/johnhaha/hakit/hamsg"
)

//create index from given struct data
func CreateIndex(data ...interface{}) error {
	for _, d := range data {
		name := hadata.GetStructNameInLowerCase(d)
		index := getIndexName(name)
		if check, _ := IndexExist(d); check {
			fmt.Printf("index %v has already been created", hamsg.InRed(name))
			UpdateSetting(d)
			continue
		}
		primaryKey, err := GetPrimaryKeyField(d)
		if err != nil {
			fmt.Println("can not get primary key for data", name)
			return err
		}
		_, err = searchClient.CreateIndex(index, primaryKey)
		if err != nil {
			fmt.Printf("created index %v failed", hamsg.InRed(name))
			return err
		}
		UpdateSetting(d)
		fmt.Printf("created index %v success", hamsg.InGreen(name))
	}
	return nil
}

func IndexExist(data interface{}) (bool, string) {
	name := hadata.GetStructNameInLowerCase(data)
	_, err := searchClient.GetIndex(getIndexName(name))
	log.Println("index is", getIndexName(name))
	return err == nil, getIndexName(name)
}

func GetPrimaryKeyField(data interface{}) (string, error) {
	t := reflect.TypeOf(data)
	// v := reflect.ValueOf(data)
	var inferIndexField []string
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		isIndex, isInferIndex := IsIndexFiled(f)
		jsonName, _ := getJsonFieldName(f)
		if isIndex {
			return jsonName, nil
		}
		if isInferIndex {
			inferIndexField = append(inferIndexField, jsonName)
		}

	}
	if len(inferIndexField) > 0 {
		return inferIndexField[0], nil
	}
	return "", errors.New("no index found")

}

func GetIndex(data interface{}) string {
	s := reflect.ValueOf(data)
	if s.Kind() == reflect.Ptr {
		s = s.Elem()
	}
	var name string
	if s.Kind() == reflect.Struct {
		name = hadata.GetStructNameInLowerCase(data)
	}
	if s.Kind() == reflect.Slice {
		//get slice elemene type and make new slice to get index
		typ := reflect.TypeOf(s.Interface()).Elem()
		t := reflect.MakeSlice(reflect.SliceOf(typ), 1, 1)
		name = hadata.GetStructNameInLowerCase(t.Index(0).Interface())
	}
	return getIndexName(name)
}

func IsIndexFiled(f reflect.StructField) (index bool, inferIndex bool) {
	tag := getSearchTag(f)
	if len(tag) == 0 {
		return false, false
	}
	for _, t := range tag {
		if t == "index" {
			return true, false
		}
	}
	fName := strings.ToLower(f.Name)
	return false, strings.Contains(fName, "id")
}

func CheckSearchFiled(f reflect.StructField) (searchable bool, filterable bool, name string, rank int64) {
	tag := getSearchTag(f)
	if len(tag) == 0 {
		return
	}
	name, _ = getJsonFieldName(f)
	for _, t := range tag {
		if i, err := strconv.ParseInt(t, 10, 32); err == nil {
			searchable = true
			rank = i
		}
		if t == "filter" {
			filterable = true
		}
	}
	return
}

func GetPrimaryKeyFieldValue(data interface{}) (string, error) {
	v := reflect.ValueOf(data)
	t := reflect.TypeOf(data)

	var inferIndexValue []interface{}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		isIndex, isInferIndex := IsIndexFiled(f)
		if isIndex {
			return v.Field(i).String(), nil
		}
		if isInferIndex {
			inferIndexValue = append(inferIndexValue, v.Field(i).Interface())
		}
	}
	if inferIndexValue == nil {
		return "", errors.New("no index exist")
	}
	return inferIndexValue[0].(string), nil

}

func GetSearchAttribute(data interface{}) (searchableAttr []string, filterAttr []string) {
	var searchableAttrRank []SearchableAttribute
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		searchable, filterable, name, rank := CheckSearchFiled(f)
		if searchable {
			searchableAttrRank = append(searchableAttrRank, SearchableAttribute{
				Name: name,
				Rank: int(rank),
			})
		}
		if filterable {
			filterAttr = append(filterAttr, name)
		}
	}
	searchableAttr = GetRankedSearchableAttribute(searchableAttrRank)
	return
}
