package ss

import(
	"strings"
)

func Unquote(s string) string {
	if len(s) < 2 {
		return s
	}

	c1 := s[0]
	c2 := s[len(s)-1]

	if c1 == '"' && c2 == '"' {
		return s[1 : len(s)-1]
	} else if c1 == '\"' && c2 == '\"' {
		return s[1 : len(s)-1]
	} else if c1 == '`' && c2 == '`' {
		return s[1 : len(s)-1]
	}

	return s
}
