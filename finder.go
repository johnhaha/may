package may

import "github.com/johnhaha/may/mayer"

type Finder struct {
	mayer.Filter
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
	doc, err := searchClient.Index(index).Find(mayer.SetLimit(f.Limit), mayer.SetOffset(f.Offset))
	if err != nil {
		return err
	}
	err = mayer.DecodeManyDoc(doc, data)
	return err
}

func Search(query string, data interface{}, options ...mayer.SearchOption) error {
	index := GetIndex(data)
	res, err := searchClient.Index(index).Search(query, options...)
	if err != nil {
		return err
	}
	err = mayer.DecodeManyDoc(res, data)
	return err
}
