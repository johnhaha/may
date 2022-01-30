package mayer

import (
	"encoding/json"
)

func DecodeManyDoc(doc []Doc, data interface{}) error {
	if doc == nil {
		return nil
	}
	d, err := json.Marshal(doc)
	if err != nil {
		return err
	}
	err = json.Unmarshal(d, data)
	return err
}
