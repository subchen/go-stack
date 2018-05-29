package json

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

// Marshal returns the JSON encoding of v
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v)
	if err != nil {
		return nil, err
	}

	// trim last "\n"
	b := buf.Bytes()
	if n := len(b); n > 0 && b[n-1] == '\n' {
		b = b[:n-1]
	}
	return b, nil
}

// MarshalIndent is like Marshal but applies Indent to format the output
func MarshalIndent(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")
	err := encoder.Encode(v)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v.
func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// ReadFile reads json from file and unmarshal into result
func ReadFile(filename string, result interface{}) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return Unmarshal(data, result)
}

// WriteFile marshals data and writes to file
func WriteFile(filename string, data interface{}) error {
	bytes, err := MarshalIndent(data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, bytes, 0755)
}
