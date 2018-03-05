package json

import (
	"encoding/json"
	"io/ioutil"

	"github.com/subchen/go-stack/data"
)

func NewQuery(v []byte) (*data.Query, error) {
	var result interface{}
	if err := json.Unmarshal(v, &result); err != nil {
		return nil, err
	}
	return data.NewQuery(result), nil
}

func NewStringQuery(v string) (*data.Query, error) {
	return NewQuery([]byte(v))
}

func NewFileQuery(filename string) (*data.Query, error) {
	v, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return NewQuery(v)
}

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func MarshalIndent(v interface{}) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}
