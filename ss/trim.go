package ss

import (
	"strings"
)

func TrimSpaceStringList(values []string) []string {
	if values == nil {
		return nil
	}

	results := make([]string, len(values))
	for i, v := range values {
		results[i] = strings.TrimSpace(v)
	}
	return results
}

func TrimPrefixStringList(values []string, prefix string) []string {
	if values == nil {
		return nil
	}

	results := make([]string, len(values))
	for i, v := range values {
		results[i] = strings.TrimPrefix(v, prefix)
	}
	return results
}

func TrimSuffixStringList(values []string, suffix string) []string {
	if values == nil {
		return nil
	}

	results := make([]string, len(values))
	for i, v := range values {
		results[i] = strings.TrimSuffix(v, suffix)
	}
	return results
}

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
