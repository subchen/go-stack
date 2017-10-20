package gstack

import (
	"strings"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func CharsetDecodeBytes(data []byte, charset string) []byte {
	charset = strings.ToUpper(charset)
	if charset == "" || charset == "UTF-8" {
		return data
	}

	var encoding encoding.Encoding
	if charset == "GBK" || charset == "GB2312" || charset == "GB18030" {
		encoding = simplifiedchinese.GB18030
	} else {
		panic("Unsupported charset: " + charset)
	}

	dst := make([]byte, len(data)*2)
	n, _, err := encoding.NewDecoder().Transform(dst, data, true)
	if err != nil {
		panic(err)
	}
	return dst[:n]
}

func CharsetEncodeBytes(data []byte, charset string) []byte {
	charset = strings.ToUpper(charset)
	if charset == "" || charset == "UTF-8" {
		return data
	}

	var encoding encoding.Encoding
	if charset == "GBK" || charset == "GB2312" || charset == "GB18030" {
		encoding = simplifiedchinese.GB18030
	} else {
		panic("Unsupported charset: " + charset)
	}

	dst := make([]byte, len(data)*2)
	n, _, err := encoding.NewEncoder().Transform(dst, data, true)
	if err != nil {
		panic(err)
	}
	return dst[:n]
}

func CharsetDecodeString(data string, charset string) string {
	return string(CharsetDecodeBytes([]byte(data), charset))
}

func CharsetEncodeString(data string, charset string) string {
	return string(CharsetEncodeBytes([]byte(data), charset))
}
