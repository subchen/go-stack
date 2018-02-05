package ss

import (
	"regexp"
	"strings"
	"unicode"
)

func Capitalize(s string) string {
	if len(s) == 0 {
		return ""
	}
	return strings.ToUpper(s[0]) + s[1:]
}

func Uncapitalize(s string) string {
	if len(s) == 0 {
		return ""
	}
	return strings.ToLower(s[0]) + s[1:]
}

func ToCamelcase(str string) string {
	words := SplitLowerWords(strings.ToLower(str))
	for i, w := range words {
		if i>0 {
			words[i] = ToCapitalized(w)
		} else {
			words[i] = w
		}
	}
	return strings.Join(words, "")
}

func ToDashizeName(str string) string {
	return toSpecialNameWithSep(str, "-")
}

func ToUnderlineName(str string) string {
	return toSpecialNameWithSep(str, "_")
}

func ToPropertyName(str string) string {
	return toSpecialNameWithSep(str, ".")
}

func toSpecialNameWithSep(str string, sep string) string {
	words := SplitLowerWords(strings.ToLower(str))
	for i, w := range words {
		if i>0 || title {
			words[i] = ToCapitalized(w)
		} else {
			words[i] = w
		}
	}
	return strings.Join(words, sep)
}

func SplitLowerWords(str string) []string {
	var words []string
	ibegin := 0

	for i, c := range str {
		if c == '-' || c == '_' || unicode.IsSpace(c) {
			s := strings.ToLower(str[ibegin:i])
			if len(s) > 0 {
				words = append(words, s)
			}
			ibegin = i+1
		} else if unicode.IsUpper(c) {
			s := strings.ToLower(str[ibegin:i])
			 if len(s) > 0 {
				words = append(words, s)
			 }
			ibegin = i
		}
	}

	// remained word
	if ibegin < len(str) {
		s := strings.ToLower(str[ibegin:len(str)])
		words = append(words, s)
	}

	return words
}
