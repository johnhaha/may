package may

func Add(data interface{}) error {
	val, err := cleanPointerData(data)
	if err != nil {
		return err
	}
	index, docs := getManyDocumentAndIndexFromData(val)
	err = searchClient.Index(index).Add(docs)
	return err
}

//update data in search, receive struct or slice of struct data
func Update(data interface{}, includeField ...string) error {
	val, err := cleanPointerData(data)
	if err != nil {
		return err
	}
	index, docs := getManyDocumentAndIndexFromData(val, includeField...)
	err = searchClient.Index(index).Update(docs)
	return err
}

//get data from search, pass pointer of struct or slice of struct data
func Get(id string, data interface{}) error {
	index := GetIndex(data)
	doc, err := searchClient.Index(index).Doc(id)
	if err != nil {
		return err
	}
	return doc.Decode(data)
}

func Del(data interface{}) error {
	val, err := cleanPointerData(data)
	if err != nil {
		return err
	}
	err = delData(val)
	return err
}
