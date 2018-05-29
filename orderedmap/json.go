package orderedmap

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

var (
	errEndOfSlice = errors.New("End of slice")
)

var (
	_ json.Marshaler   = &OrderedMap{}
	_ json.Unmarshaler = &OrderedMap{}
)

// MarshalJSON implements the json.Marshaler interface.
func (m *OrderedMap) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0, 32))

	buf.WriteByte('{')
	for i, key := range m.keys {
		if i > 0 {
			buf.WriteByte(',')
		}

		// marshal key
		b, err := json.Marshal(key)
		if err != nil {
			return nil, err
		}
		buf.Write(b)
		buf.WriteByte(':')

		// marshal value
		b, err = json.Marshal(m.values[key])
		if err != nil {
			return nil, err
		}
		buf.Write(b)
	}
	buf.WriteByte('}')

	return buf.Bytes(), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (m *OrderedMap) UnmarshalJSON(b []byte) error {
	dec := json.NewDecoder(bytes.NewReader(b))

	dec.Token() // {
	m.unmarshalJSON(dec)
	dec.Token() // }

	return nil
}

func (m *OrderedMap) unmarshalJSON(dec *json.Decoder) error {
	for dec.More() {
		t, err := dec.Token()
		if err != nil {
			return err
		}

		key, isKey := t.(string)
		if !isKey {
			return fmt.Errorf("%t %s is not string(expected key)", t, t)
		}

		val, err := getVal(dec)
		if err != nil {
			return err
		}
		m.Set(key, val)
	}
	return nil
}

func decToSlice(dec *json.Decoder) ([]interface{}, error) {
	res := make([]interface{}, 0)

	for {
		v, err := getVal(dec)
		if err == errEndOfSlice {
			return res, nil
		}
		if err != nil {
			return nil, err
		}
		res = append(res, v)
	}
}

func getVal(dec *json.Decoder) (interface{}, error) {
	t, err := dec.Token()
	if err != nil {
		return nil, err
	}

	switch tok := t.(type) {
	case json.Delim:
		switch tok {
		case '[':
			return decToSlice(dec)
		case '{':
			m := New()
			err := m.unmarshalJSON(dec)
			if err != nil {
				return nil, err
			}
			_, err = dec.Token() // }
			return m, err
		case ']':
			return nil, errEndOfSlice
		case '}':
			return nil, errors.New("unexpected '}'")
		default:
			panic("unreachable code")
		}
	default:
		return tok, nil
	}
}
