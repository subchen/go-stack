package ss

import (
	"bytes"
	"unicode"
)

// DeleteWhitespaces deletes all whitespaces (unicode.IsSpace) in string.
func DeleteWhitespaces(s string) string {
	if len(s) == 0 {
		return s
	}

	var changes bool
	var buf bytes.Buffer
	for _, r := range s {
		if unicode.IsSpace(r) {
			changes = true
		} else {
			buf.WriteRune(r)
		}
	}

	if !changes {
		return s
	}
	return buf.String()
}

// IsBlank returns whether a string is whitespace or empty.
func IsBlank(s string) bool {
	if len(s) == 0 {
		return true
	}
	for _, r := range s {
		if unicode.IsSpace(r) == false {
			return false
		}
	}
	return true
}
