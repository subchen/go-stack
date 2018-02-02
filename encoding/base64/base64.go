package base64

import (
	"encoding/base64"
)

func EncodeBytes(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func EncodeString(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func DecodeToBytes(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

func DecodeToString(data string) (string, error) {
	s, e := base64.StdEncoding.DecodeString(data)
	return string(s), e
}
