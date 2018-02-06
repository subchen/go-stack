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
