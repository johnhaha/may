package may

import "github.com/johnhaha/may/mayer"

//search with query, will parse search result to data
func Search(query string, data interface{}, options ...mayer.SearchOption) error {
	index := GetIndex(data)
	res, err := searchClient.Index(index).Search(query, options...)
	if err != nil {
		return err
	}
	err = mayer.DecodeManyDoc(res, data)
	return err
}
