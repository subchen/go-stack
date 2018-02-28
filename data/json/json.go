package json

import (
	"encoding/json"

	"github.com/subchen/go-stack/data"
	"github.com/subchen/go-stack/fs"
)

func NewQuery(data []byte) (*data.Query, error) {
	var result interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return data.NewQuery(result)
}

func NewStringQuery(data string) (*data.Query, error) {
	return NewQuery([]byte(data))
}

func NewFileQuery(filename string) (*data.Query, error) {
	data, err := fs.FileGetBytes(filename)
	if err != nil {
		return nil, err
	}
	return NewQuery(data)
}

func Marshal(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func MarshalIndent(data interface{}) ([]byte, error) {
	return json.MarshalIndent(data, "", "  ")
}
