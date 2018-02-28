package yaml

import (
	"io/ioutil"

	"github.com/go-yaml/yaml"
	"github.com/subchen/go-stack/data"
)

func NewQuery(data []byte) (*data.Query, error) {
	var result interface{}
	if err := yaml.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return data.NewQuery(result)
}

func NewStringQuery(data string) (*data.Query, error) {
	return NewQuery([]byte(data))
}

func NewFileQuery(filename string) (*data.Query, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return NewQuery(data)
}

func Marshal(data interface{}) ([]byte, error) {
	return yaml.Marshal(data)
}
