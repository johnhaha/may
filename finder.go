package may

type Finder struct {
	Filter
}

func NewFinder() *Finder {
	return &Finder{}
}

func (f *Finder) SetLimit(l int) *Finder {
	f.Limit = l
	return f
}

func (f *Finder) SetOffset(o int) *Finder {
	f.Offset = o
	return f
}

func (f *Finder) GetDocuments(data interface{}) error {
	index := GetIndex(data)
	doc, err := searchClient.Index(index).Find(SetLimit(f.Limit), SetOffset(f.Offset))
	if err != nil {
		return err
	}
	err = decodeManyDoc(doc, data)
	return err
}

func (f *Finder) Search(query string, data interface{}) error {
	index := GetIndex(data)
	res, err := searchClient.Index(index).Search(query, SetSearchLimit(f.Limit), SetSearchOffset(f.Offset))
	if err != nil {
		return err
	}
	err = decodeManyDoc(res, data)
	return err
}
