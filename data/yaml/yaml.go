package yaml

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
	"github.com/subchen/go-stack/data"
)

func NewQuery(v []byte) (*data.Query, error) {
	var result interface{}
	if err := yaml.Unmarshal(v, &result); err != nil {
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
	return yaml.Marshal(v)
}
