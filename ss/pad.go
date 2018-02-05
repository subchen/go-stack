package ss

import (
	"bytes"
	"unicode/utf8"
)

// PadLeft returns a string with pad string at right side if str's rune length is smaller than length.
// If str's rune length is larger than length, str itself will be returned.
//
// If pad is an empty string, str will be returned.
//
// Samples:
//     PadLeft("hello", 4, " ")    => "hello"
//     PadLeft("hello", 10, " ")   => "hello     "
//     PadLeft("hello", 10, "123") => "hello12312"
func PadLeft(str string, length int, pad string) string {
	l := Len(str)

	if l >= length || pad == "" {
		return str
	}

	remains := length - l
	padLen := Len(pad)

	output := &bytes.Buffer{}
	output.Grow(len(str) + (remains/padLen+1)*len(pad))
	output.WriteString(str)
	writePadString(output, pad, padLen, remains)
	return output.String()
}

// PadRight returns a string with pad string at left side if str's rune length is smaller than length.
// If str's rune length is larger than length, str itself will be returned.
//
// If pad is an empty string, str will be returned.
//
// Samples:
//     PadRight("hello", 4, " ")    => "hello"
//     PadRight("hello", 10, " ")   => "     hello"
//     PadRight("hello", 10, "123") => "12312hello"
func PadRight(str string, length int, pad string) string {
	l := Len(str)

	if l >= length || pad == "" {
		return str
	}

	remains := length - l
	padLen := Len(pad)

	output := &bytes.Buffer{}
	output.Grow(len(str) + (remains/padLen+1)*len(pad))
	writePadString(output, pad, padLen, remains)
	output.WriteString(str)
	return output.String()
}

// PadCenter returns a string with pad string at both side if str's rune length is smaller than length.
// If str's rune length is larger than length, str itself will be returned.
//
// If pad is an empty string, str will be returned.
//
// Samples:
//     PadCenter("hello", 4, " ")    => "hello"
//     PadCenter("hello", 10, " ")   => "  hello   "
//     PadCenter("hello", 10, "123") => "12hello123"
func PadCenter(str string, length int, pad string) string {
	l := Len(str)

	if l >= length || pad == "" {
		return str
	}

	remains := length - l
	padLen := Len(pad)

	output := &bytes.Buffer{}
	output.Grow(len(str) + (remains/padLen+1)*len(pad))
	writePadString(output, pad, padLen, remains/2)
	output.WriteString(str)
	writePadString(output, pad, padLen, (remains+1)/2)
	return output.String()
}

func writePadString(output *bytes.Buffer, pad string, padLen, remains int) {
	var r rune
	var size int

	repeats := remains / padLen

	for i := 0; i < repeats; i++ {
		output.WriteString(pad)
	}

	remains = remains % padLen

	if remains != 0 {
		for i := 0; i < remains; i++ {
			r, size = utf8.DecodeRuneInString(pad)
			output.WriteRune(r)
			pad = pad[size:]
		}
	}
}
