package splits

import (
	"fmt"
	"strings"

	"github.com/subchen/go-stack/scanner"
)

// KVSplit splits "key:value ..." delimited by whitespaces.
// The value can be quoted by '"", '\'', '`' if value has spaces
// Example input: `exe:/bin/ls args:"-l -a --color"`
func KVSplit(input string) (map[string]string, error) {
	m := make(map[string]string)

	s := scanner.New(strings.TrimSpace(input))
	for !s.Eof() {
		key, ok := s.ScanIdentifier()
		if !ok {
			return nil, s.Errorf("missing key")
		}

		if !s.ScanChar(':') {
			return nil, s.Errorf("missing ':'")
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
