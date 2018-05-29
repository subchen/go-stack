package splits

import (
	"strings"

	"github.com/subchen/go-stack/scanner"
)

func AttrSplit(input string) ([]string, error) {
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
			attr, ok = s.ScanUntil(']')
			if !ok {
				return nil, s.Errorf("no `]` end")
			}

			_, _ = s.ScanChar() // skip ']'
			attr += "]"
		} else {
			attr, ok = s.ScanUntil('.')
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
