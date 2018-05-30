package splitattrs

import (
	"strings"

	"github.com/subchen/go-stack/scanner"
)

// Split splits "attr1.attr2.[index].[key=value].[a:b]..." delimited by dot.
//
// The attr can be quoted by '"", '\'', '`' if value has spaces
//
// Example input: `nodes."availables".[0].status.[cpu=2.0]`
func Split(input string) ([]string, error) {
	attrs := make([]string, 0, 4)

	s := scanner.New(strings.TrimSpace(input))
	for !s.Eof() {
		var attr string
		var ok bool

		peek := s.Peek()
		if peek == '"' || peek == '\'' || peek == '`' {
			attr, ok = s.ScanQuoteString()
			if !ok {
				return nil, s.Errorf("invalid quote string")
			}
		} else if peek == '[' {
			attr, ok = s.ScanUntil(']', true)
			if !ok {
				return nil, s.Errorf("no `]` end")
			}
		} else {
			attr, ok = s.ScanUntil('.', false)
			if !ok {
				attr = s.ScanToEnd()
			}
		}

		attrs = append(attrs, attr)

		if !s.Eof() {
			delim, _ := s.ScanChar()
			if delim != '.' {
				return nil, s.Errorf("required `.` as delim")
			}
		}
	}

	return attrs, nil
}
