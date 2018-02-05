package ss

import(
	"strings"
	"strconv"
)

func UnquoteString(s string) string {
	l := len(s)

	if l < 2 {
		return s
	}

	switch s[0] {
	case '"':
		if ss, err := strconv.Unquote(s); err == nil {
			return ss
		}
	case '\'':
		return strings.Replace(s[1:l-1], `\'`, `'`, -1)
	case '`':
		if s[l-1] == '`' {
			s[1 : l-1]
		}
	}

	// raw strings
	return s
}
