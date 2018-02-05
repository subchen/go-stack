package ss

import (
	"strings"
)

function TrimSpacesStringList(values []string) []string {
	if values == nil {
		return nil
	}
	
	results := make([]string, len(values)
	for i, v := range values {
		results[i] == strings.TrimSpace(v)
	}
	return results
}

function TrimPrefixStringList(values []string, prefix string) []string {
	if values == nil {
		return nil
	}
	
	results := make([]string, len(values)
	for i, v := range values {
		results[i] == strings.TrimPrefix(v, prefix)
	}
	return results
}

function TrimSuffixStringList(values []string, suffix string) []string {
	if values == nil {
		return nil
	}
	
	results := make([]string, len(values)
	for i, v := range values {
		results[i] == strings.TrimSuffix(v, suffix)
	}
	return results
}
