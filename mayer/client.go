package mayer

import (
	"github.com/johnhaha/hakit/hareq"
)

type MeiliClient struct {
	Host   string
	APIKey string
}

func (c *MeiliClient) GetCaller(path string) *hareq.Caller {
	caller := hareq.NewCaller(c.Host + path)
	if c.APIKey != "" {
		caller.SetAuth("Bearer " + c.APIKey)
	}
	return caller
}

func (c *MeiliClient) GetIndex(index string) (*IndexRes, error) {
	var res IndexRes
	err := c.GetCaller("/indexes/" + index).Get().Decode(&res)
	if err != nil {
		return nil, err
	}
	if res.Error() != "" {
		return nil, res.ErrRes
	}
	return &res, nil
}

func (c *MeiliClient) CreateIndex(uid string, primaryKey string) (*ModifyIndexRes, error) {
	var res ModifyIndexRes
	err := c.GetCaller("/indexes").SetBody(map[string]string{
		"uid":        uid,
		"primaryKey": primaryKey,
	}).Post().Decode(&res)
	if err != nil {
		return nil, err
	}
	if res.Error() != "" {
		return nil, res.ErrRes
	}
	return &res, nil
}

func (c *MeiliClient) DeleteIndex(index string) error {
	caller := c.GetCaller("/indexes/" + index).Delete()
	if caller.Err != nil {
		return caller.Err
	}
	return nil
}

type MeiliIndex struct {
	*MeiliClient
	Index      string
	Setting    *IndexSetting
	Filter     *Filter
	SearchBody *SearchBody
}

func (c *MeiliClient) Index(index string) *MeiliIndex {
	return &MeiliIndex{
		MeiliClient: c,
		Index:       index,
		Setting:     &IndexSetting{},
		Filter:      &Filter{},
		SearchBody:  &SearchBody{},
	}
}

func (index *MeiliIndex) UpdateSetting(options ...SettingOption) error {
	for _, o := range options {
		o(index)
	}
	var res ModifyIndexRes
	caller := index.GetCaller("/indexes/" + index.Index + "/settings").SetBody(index.Setting).Post()
	if caller.Err != nil {
		return caller.Err
	}
	err := caller.Decode(&res)
	if err != nil {
		return err
	}
	if res.Error() != "" {
		return res.ErrRes
	}
	return nil
}

func (index *MeiliIndex) Add(data interface{}) error {
	var res ModifyIndexRes
	err := index.GetCaller("/indexes/" + index.Index + "/documents").SetBody(data).Post().Decode(&res)
	if err != nil {
		return err
	}
	if res.Error() != "" {
		return res.ErrRes
	}
	return nil
}

func (index *MeiliIndex) Update(data interface{}) error {
	var res ModifyIndexRes
	err := index.GetCaller("/indexes/" + index.Index + "/documents").SetBody(data).Put().Decode(&res)
	if err != nil {
		return err
	}
	if res.Error() != "" {
		return res.ErrRes
	}
	return nil
}
func (index *MeiliIndex) Doc(id string) (Doc, error) {
	doc := make(Doc)
	caller := index.GetCaller("/indexes/" + index.Index + "/documents/" + id).Get()
	err := caller.Decode(&doc)
	if err != nil {
		return nil, err
	}
	if doc.Error() != "" {
		return nil, doc
	}
	return doc, nil
}

func (index *MeiliIndex) Find(options ...FilterOption) ([]Doc, error) {
	for _, o := range options {
		o(index)
	}
	var doc []Doc
	caller := index.GetCaller("/indexes/" + index.Index + "/documents").SetQuery(index.Filter.GetQuery()).Get()
	err := caller.Decode(&doc)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func (index *MeiliIndex) DelDoc(id string) error {
	caller := index.GetCaller("/indexes/" + index.Index + "/documents/" + id).Delete()
	return caller.Err
}

func (index *MeiliIndex) DelManyDoc(id []string) error {
	caller := index.GetCaller("/indexes/" + index.Index + "/documents/delete-batch").SetBody(id).Post()
	return caller.Err
}

func (index *MeiliIndex) Search(query string, options ...SearchOption) ([]Doc, error) {
	index.SearchBody.Query = query
	for _, o := range options {
		o(index)
	}
	var res SearchRes
	err := index.GetCaller("/indexes/" + index.Index + "/search").SetBody(index.SearchBody).Post().Decode(&res)
	if err != nil {
		return nil, err
	}
	if res.Error() != "" {
		return nil, res
	}
	return res.Hits, nil
}
