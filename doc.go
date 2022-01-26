package may

func Add(data interface{}) error {
	val, err := cleanPointerData(data)
	if err != nil {
		return err
	}
	index, docs := getManyDocumentAndIndexFromData(val)
	_, err = searchClient.Index(index).AddDocuments(docs)
	return err
}

//update data in search, receive struct or slice of struct data
func Update(data interface{}) error {
	val, err := cleanPointerData(data)
	if err != nil {
		return err
	}
	index, docs := getManyDocumentAndIndexFromData(val)
	_, err = searchClient.Index(index).UpdateDocuments(docs)
	return err
}

//get data from search, pass pointer of struct or slice of struct data
func Get(id string, data interface{}) error {
	index := GetIndex(data)
	err := searchClient.Index(index).GetDocument(id, data)
	return err
}

func Del(data interface{}) error {
	val, err := cleanPointerData(data)
	if err != nil {
		return err
	}
	err = delData(val)
	return err
}
