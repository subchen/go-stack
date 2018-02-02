package zlib

import (
	"bytes"
	"compress/zlib"
	"io"
)

func Encode(data []byte) []byte {
	if data == nil || len(data) < 13 {
		return data
	}
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	w.Write(data)
	w.Close()
	return in.Bytes()
}

func Decode(data []byte) []byte {
	if data == nil || len(data) < 13 {
		return data
	}
	b := bytes.NewReader(data)
	var out bytes.Buffer
	r, err := zlib.NewReader(b)
	if err != nil {
		panic(err)
	}
	io.Copy(&out, r)
	return out.Bytes()
}
