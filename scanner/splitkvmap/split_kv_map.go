package splitkvmap

import (
	"strings"

	"github.com/subchen/go-stack/scanner"
)

// Split splits "key:value ..." delimited by whitespaces.
//
// The value can be quoted by '"", '\'', '`' if value has spaces
//
// Example input: `exe:/bin/ls args:"-l -a --color"`
func Split(input string) (map[string]string, error) {
	m := make(map[string]string)

	s := scanner.New(strings.TrimSpace(input))
	lastDelim := '\u0000'
	for !s.Eof() {
		key, ok := s.ScanIdentifier()
		if !ok {
			return nil, s.Errorf("missing key")
		}

		delim, _ := s.ScanChar()
		if delim == ':' {
			if lastDelim == '=' {
				return nil, s.Errorf("missing '=', but got ':'")
			}
			lastDelim = delim
		} else if delim == '=' {
			if lastDelim == ':' {
				return nil, s.Errorf("missing ':', but got '='")
			}
			lastDelim = delim
		} else {
			return nil, s.Errorf("missing delim ':' or '='")
		}

		value, ok := s.ScanString()
		if !ok {
			return nil, s.Errorf("missing value")
		}

		m[key] = value

		if !s.Eof() {
			ok := s.SkipWhitespaces()
			if !ok {
				return nil, s.Errorf("missing whitespaces after value")
			}
		}
	}

	return m, nil
}
