package ss

import (
	"strings"
	"unicode"
)

func Capitalize(s string) string {
	if len(s) == 0 {
		return ""
	}
	rs := []rune(s)
	rs[0] = unicode.ToUpper(rs[0])
	return string(rs)
}

func Uncapitalize(s string) string {
	if len(s) == 0 {
		return ""
	}
	rs := []rune(s)
	rs[0] = unicode.ToLower(rs[0])
	return string(rs)
}

// ToCamelcase can convert all words in a string to camel format.
//
// Some samples.
//     "Hello world"  => "helloWorld"
//     "Hello-world"  => "helloWorld"
//     "Hello_world"  => "helloWorld"
func ToCamelcase(str string) string {
	words := SplitLowerWords(strings.ToLower(str))
	for i, w := range words {
		if i > 0 {
			words[i] = Capitalize(w)
		} else {
			words[i] = w
		}
	}
	return strings.Join(words, "")
}

// ToDashizeName can convert all words in a string to dashizer format.
//
// Some samples.
//     "HelloWorld"   => "hello-world"
//     "Hello World"  => "hello-world"
//     "Hello-World"  => "hello-world"
//     "Hello_World"  => "hello-world"
func ToDashizerName(str string) string {
	words := SplitLowerWords(strings.ToLower(str))
	return strings.Join(words, "-")
}

// ToUnderlineName can convert all words in a string to underscore format.
//
// Some samples.
//     "HelloWorld"   => "hello_world"
//     "Hello World"  => "hello_world"
//     "Hello-World"  => "hello_world"
//     "Hello_World"  => "hello_world"
func ToUnderlineName(str string) string {
	words := SplitLowerWords(strings.ToLower(str))
	return strings.Join(words, "_")
}

// ToPropertyName can convert all words in a string to point format.
//
// Some samples.
//     "HelloWorld"   => "hello.world"
//     "Hello World"  => "hello.world"
//     "Hello-World"  => "hello.world"
//     "Hello_World"  => "hello.world"
func ToPropertyName(str string) string {
	words := SplitLowerWords(strings.ToLower(str))
	return strings.Join(words, ".")
}

func SplitLowerWords(str string) []string {
	var words []string
	lastpos := 0

	for i, c := range str {
		if c == '-' || c == '_' || unicode.IsSpace(c) {
			s := strings.ToLower(str[lastpos:i])
			if len(s) > 0 {
				words = append(words, s)
			}
			lastpos = i + 1
		} else if unicode.IsUpper(c) {
			s := strings.ToLower(str[lastpos:i])
			if len(s) > 0 {
				words = append(words, s)
			}
			lastpos = i
		}
	}

	// remained word
	if lastpos < len(str) {
		s := strings.ToLower(str[lastpos:len(str)])
		words = append(words, s)
	}

	return words
}
